#!/bin/bash
echo "Creating a new project"

RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
CLEAR=$(tput sgr0)

readonly urlRegex="https?://.+"

echo "${GREEN}Enter the project title: ${CLEAR} "
while :; do
  read -r projectTitle
  if [ ${#projectTitle} -gt 1 ]; then
    break
  else
    echo "${RED}Title must be greater than 1${CLEAR}"
  fi
done

echo "${GREEN}Enter the project description (Minimum 20 words):${CLEAR}"
while :; do
  read -r projectDescription
  if [ ${#projectDescription} -gt 10 ]; then
    break
  else
    echo "${RED}Description must be greater than 10${CLEAR}"
  fi
done

echo "${GREEN}Enter the technologies (split by a |):${CLEAR}"
while :; do
  read -r projectTechnologies
  if [ ${#projectTechnologies} -gt 1 ]; then
    break
  else
    echo "${RED}Must have atleast one technology${CLEAR}"
  fi
done

echo "${GREEN}Enter the Github Link (Press Enter for null):${CLEAR}"
while :; do
  read -r projectGitLink
  if [ -n "${projectGitLink}" ]; then
    if [[ "${projectGitLink}" =~ ${urlRegex} ]]; then
      projectGitLink="\"${projectGitLink}\""
      break
    else
      echo "Url must match the regex https?://.+"
    fi
  else
    projectGitLink=null
    break
  fi
done

echo "${GREEN}Enter the Website Link (Press enter for null):${CLEAR}"
while :; do
  read -r projectWebLink
  if [ -n "${projectWebLink}" ]; then
    if [[ "${projectWebLink}" =~ ${urlRegex} ]]; then
      projectWebLink="\"${projectWebLink}\""
      break
    else
      echo "Url must match the regex https?://.+"
    fi
  else
    projectWebLink=null
    break
  fi
done

echo "${GREEN}Enter the Blog Id associated (Press enter for null):${CLEAR}"
read projectBlogId
if [ -n "${projectBlogId}" ]; then
  echo ""
else
  projectBlogId=null
fi

echo -e "${GREEN}Are you happy with these values?${CLEAR}\n"
echo -e "${GREEN}Project Title:${CLEAR}$projectTitle"
echo -e "${GREEN}Project Description:${CLEAR}\n$projectDescription"
echo -e "${GREEN}Project Technologies:${CLEAR} $projectTechnologies"
echo -e "${GREEN}Project Github Link:${CLEAR} $projectGitLink"
echo -e "${GREEN}Project Website Link:${CLEAR} $projectWebLink"
echo -e "${GREEN}Project Blog ID:${CLEAR} $projectBlogId"

echo "${GREEN}y${CLEAR}/${RED}n${CLEAR}"
while :; do
  read satisifed
  if [ "${satisifed}" == "y" ]; then
    break
  elif [ "${satisifed}" == "n" ]; then
    exit
  else
    echo -e "${RED} y or n${CLEAR}"
  fi
done

touch newProjectReq.json
echo "{
  \"title\":\"${projectTitle}\",
  \"desc\":\"${projectDescription}\",
  \"tech\":\"${projectTechnologies}\",
  \"git_link\":${projectGitLink},
  \"web_link\":${projectWebLink},
  \"blog_id\":${projectBlogId}
}" >newProjectReq.json


curl -X POST http://localhost:5731/addProject -H "Content-Type: application/json" \
  -d @newProjectReq.json | jq

rm newProjectReq.json