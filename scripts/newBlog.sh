RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
CLEAR=$(tput sgr0)

echo "${GREEN}Enter the blog title:${CLEAR}"
while :; do
  read -r blogTitle
  if [ ${#blogTitle} -gt 0 ]; then
    break
  else
    echo "${RED}Must have a title${CLEAR}"
  fi
done

blogDate=$(date '+%d-%m-%Y')

touch newBlogReq.json
echo "{
  \"title\":\"${blogTitle}\",
  \"date\":\"${blogDate}\"
}" >newBlogReq.json

curl -X POST http://localhost:5731/addBlog -H "Content-Type: application/json" \
  -d @newBlogReq.json | jq

rm newBlogReq.json
