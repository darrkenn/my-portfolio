use std::fs;

use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Project {
    title: String,
    info: String,
    github: Option<String>,
    website: Option<String>,
    blog: Option<String>,
}

#[derive(Serialize, Deserialize)]
pub struct Blog {
    title: String,
    html: String,
    date: String,
}

#[derive(Serialize, Deserialize)]
pub struct Blogs {
    blogs: Vec<Blog>,
}

#[derive(Serialize, Deserialize)]
pub struct CurrentlyWorking {
    project: String,
    link: String,
}

pub async fn load_projects() -> Vec<Project> {
    let mut projects: Vec<Project> = Vec::new();

    let entries = match fs::read_dir("/var/lib/portfolio/projects") {
        Ok(entries) => entries.flatten(),
        Err(_) => return projects,
    };

    for entry in entries {
        if entry.path().extension().and_then(|e| e.to_str()) == Some("json") {
            if let Ok(file) = fs::read_to_string(entry.path()) {
                match serde_json::from_str::<Project>(&file) {
                    Ok(project) => projects.push(project),
                    Err(_) => continue,
                }
            }
        }
    }

    projects
}

pub async fn load_blogs() -> Vec<Blog> {
    match fs::read_to_string("/var/lib/portfolio/blogs.json") {
        Ok(file) => match serde_json::from_str::<Blogs>(&file) {
            Ok(data) => data.blogs,
            Err(_) => Vec::new(),
        },
        Err(_) => Vec::new(),
    }
}

pub async fn load_cw() -> CurrentlyWorking {
    match fs::read_to_string("/var/lib/portfolio/cw.json") {
        Ok(file) => match serde_json::from_str::<CurrentlyWorking>(&file) {
            Ok(data) => CurrentlyWorking {
                project: data.project,
                link: data.link,
            },
            Err(_) => CurrentlyWorking {
                project: String::from("Nothing"),
                link: String::from(""),
            },
        },
        Err(_) => CurrentlyWorking {
            project: String::from("ERROR"),
            link: String::from("https://doc.rust-lang.org/book/ch09-00-error-handling.html"),
        },
    }
}
