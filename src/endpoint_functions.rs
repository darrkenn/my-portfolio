use std::fs;

use crate::Project;

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
