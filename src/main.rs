mod endpoint_functions;
use axum::{
    Router,
    extract::State,
    response::{Html, IntoResponse},
    routing::get,
};
use serde::{Deserialize, Serialize};
use std::{fs, sync::Arc};
use tera::Tera;
use tower_http::services::ServeDir;

use crate::endpoint_functions::load_projects;

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
    let html = fs::read_to_string("html/pages/projects.html")
        .unwrap_or_else(|_| "<h1>File deleted or im working on it</h1>".to_string());
    Html(html)
}
async fn notes() -> impl IntoResponse {
    let html = fs::read_to_string("html/pages/notes.html")
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
        .route("/notes", get(notes))
        //Endpoints
        .route("/api/list_projects", get(list_projects))
        .nest_service("/static", ServeDir::new("static"))
        .with_state(tera_state);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await.unwrap();

    println!("Started server!");
    axum::serve(listener, app).await.unwrap();
}
