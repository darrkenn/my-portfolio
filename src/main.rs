mod endpoint_functions;
use axum::{
    Router,
    extract::{Path, State},
    response::{Html, IntoResponse},
    routing::get,
};
use comrak::{ComrakOptions, markdown_to_html};
use serde::{Deserialize, Serialize};
use std::{fs, sync::Arc};
use tera::Tera;
use tower_http::services::{ServeDir, ServeFile};

use crate::endpoint_functions::{load_blogs, load_projects};

#[derive(Clone)]
struct TeraState {
    tera: Arc<Tera>,
}

#[derive(Serialize, Deserialize)]
struct Project {
    title: String,
    info: String,
    github: Option<String>,
    website: Option<String>,
    blog: Option<String>,
}

//Pages
async fn about() -> impl IntoResponse {
    let html = fs::read_to_string("html/pages/about.html")
        .unwrap_or_else(|_| "<h1>File deleted or im working on it</h1>".to_string());
    Html(html)
}
async fn projects() -> impl IntoResponse {
    let html = fs::read_to_string("html/pages/projects.html")
        .unwrap_or_else(|_| "<h1>File deleted or im working on it</h1>".to_string());
    Html(html)
}
async fn blogs() -> impl IntoResponse {
    let html = fs::read_to_string("html/pages/blogs.html")
        .unwrap_or_else(|_| "<h1>File deleted or im working on it</h1>".to_string());
    Html(html)
}
async fn blog(Path(title): Path<String>, State(state): State<TeraState>) -> impl IntoResponse {
    let mut context = tera::Context::new();
    context.insert("title", &title);

    let html = state.tera.render("blog.html", &context).unwrap();
    Html(html)
}
async fn contact() -> impl IntoResponse {
    let html = fs::read_to_string("html/pages/contact.html")
        .unwrap_or_else(|_| "<h1>File deleted or im working on it</h1>".to_string());
    Html(html)
}

//Endpoints
async fn list_projects(State(state): State<TeraState>) -> impl IntoResponse {
    let projects = load_projects().await;

    let mut context = tera::Context::new();
    context.insert("projects", &projects);

    let html = state.tera.render("project_list.html", &context).unwrap();
    Html(html)
}
async fn list_blogs(State(state): State<TeraState>) -> impl IntoResponse {
    let blogs = load_blogs().await;

    let mut context = tera::Context::new();
    context.insert("blogs", &blogs);
    let html = state.tera.render("blog_list.html", &context).unwrap();
    Html(html)
}
async fn get_blog(Path(title): Path<String>) -> String {
    let md = fs::read_to_string(format!("/var/lib/portfolio/blogs/{}.md", title))
        .unwrap_or_else(|_| "<h1>Cant get markdown</h1>".to_string());

    markdown_to_html(&md, &ComrakOptions::default())
}

#[tokio::main]
async fn main() {
    let tera = Tera::new("html/templates/*").expect("Couldnt load templates folder");

    let tera_state = TeraState {
        tera: Arc::new(tera),
    };

    let app = Router::new()
        //Pages
        .route("/", get(about))
        .route("/projects", get(projects))
        .route("/blogs", get(blogs))
        .route("/blogs/{title}", get(blog))
        .route("/contact", get(contact))
        //Endpoints
        .route("/api/list_projects", get(list_projects))
        .route("/api/list_blogs", get(list_blogs))
        .route("/api/blog/{title}", get(get_blog))
        //Services
        .nest_service("/static", ServeDir::new("static"))
        .nest_service("/favicon.ico", ServeFile::new("static/favicon.ico"))
        .with_state(tera_state);
    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await.unwrap();

    println!("Started server!");
    axum::serve(listener, app).await.unwrap();
}
