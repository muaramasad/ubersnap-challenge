**Run the Service:**
<code>docker-compose up</code>.

**Run the test:**
<code>docker-compose run --rm web go test ./test/  -coverpkg=./... -coverprofile ./coverage.out</code>.

**Postman CURL:**
***Convert:***
<code>curl --location --request POST 'http://127.0.0.1:8080/api/v1/convert' \
--form 'image=@"IMAGE_PATH"'</code>.

***Compress:***
<code>curl --location --request POST 'http://127.0.0.1:8080/api/v1/compress' \
--form 'image=@"IMAGE_PATH"'</code>.

***Resize:***
<code>curl --location --request POST 'http://127.0.0.1:8080/api/v1/resize' \
--form 'image=@"IMAGE_PATH"' \
--form 'width="100"' \
--form 'height="50"'</code>.

***View:***
<code>curl --location --request GET 'http://localhost:8080/api/v1/view/7f6bd20ebb8d4601b22b6d74ad38e0c3_compressed.png'</code>.