#!/bin/bash
echo "Creating a new project"

RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
BLUE=$(tput setaf 4)
PURPLE=$(tput setaf 5)
CYAN=$(tput setaf 6)
WHITE=$(tput setaf 7)
BOLD=$(tput bold)
CLEAR=$(tput sgr0)

echo "${GREEN}Enter the project title: ${CLEAR} "
while :; do
  read projectTitle
  if [ ${#projectTitle} -gt 1 ]; then
    break
  else
    echo "${RED}Title must be greater than 1${CLEAR}"
  fi
done

echo "${GREEN}Enter the project description (Minimum 20 words):${CLEAR}"
while :; do
  read projectDescription
  if [ ${#projectDescription} -gt 10 ]; then
    break
  else
    echo "${RED}Description must be greater than 10${CLEAR}"
  fi
done

echo "${GREEN}Enter the technologies (split by a |):${CLEAR}"
while :; do
  read projectTechnologies
  if [ ${#projectTechnologies} -gt 1 ]; then
    break
  else
    echo "${RED}Must have atleast one technology${CLEAR}"
  fi
done

echo "${GREEN}Enter the Github Link (Press Enter for null):${CLEAR}"
read projectGitLink
echo "${GREEN}Enter the Website Link (Press enter for null):${CLEAR}"
read projectWebLink
echo "${GREEN}Enter the Blog Id associated (Press enter for null):${CLEAR}"
read projectBlogId

clear

echo -e "${PURPLE}Are you happy with these values?${CLEAR}\n"
echo -e "${PURPLE}Project Title: $projectTitle${CLEAR}\n"
echo -e "${PURPLE}Project Description:\n$projectDescription${CLEAR}\n"
echo -e "${PURPLE}Project Technologies:\n$projectTechnologies${CLEAR}\n"
echo -e "${PURPLE}Project Github Link: $projectGitLink${CLEAR}\n"
echo -e "${PURPLE}Project Website Link: $projectWebLink${CLEAR}\n"
echo -e "${PURPLE}Project Blog ID: $projectBlogId${CLEAR}\n"

echo "${GREEN}y${CLEAR}/${RED}n${CLEAR}"
while :; do
  read satisifed
  if [ ${satisifed} == "y" ]; then
    break
  elif [ ${satisifed} == "n" ]; then
    exit
  else
    echo -e "${RED} y or n${CLEAR}"
  fi
done
