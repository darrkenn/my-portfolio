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

#[derive(Serialize, Deserialize, Debug)]
pub struct Blog {
    title: String,
    html: String,
    date: String,
}

#[derive(Serialize, Deserialize)]
pub struct Blogs {
    blogs: Vec<Blog>,
}

pub async fn load_projects() -> Vec<Project> {
    let mut projects: Vec<Project> = Vec::new();

    if let Ok(entries) = fs::read_dir("/var/lib/portfolio/projects") {
        for entry in entries.flatten() {
            if entry.path().extension().and_then(|e| e.to_str()) == Some("json") {
                if let Ok(file) = fs::read_to_string(entry.path()) {
                    if let Ok(json) = serde_json::from_str::<Project>(&file) {
                        projects.push(Project {
                            title: json.title,
                            info: json.info,
                            github: json.github,
                            website: json.website,
                            blog: json.blog,
                        })
                    }
                }
            }
        }
    };

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
