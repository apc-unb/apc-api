
# Exam

## Get all News
* HTTP Request : ```GET http://api.com/news```
* Return a list of object in json format as follow

    ``` 
    [
        {
            "id"            :   ObjectId,
            "classid"       :   ObjectId,
            "authorID"      :   ObjectId,
            "authorName"    :   String,
            "admin"         :   Bool,
            "title"         :   String,
            "description"   :   String,
            "tags"          :   []String,
            "createdAT"     :   Time,
            "updatedAT"     :   Time
        }...,
    ]
    ```

## Get all News from a class
* HTTP Request : ```GET http://api.com/news/{classID}```
* Return a list of object in json format as follow

    ``` 
        [
			{
                "id"            :   ObjectId,
                "classid"       :   ObjectId,
                "authorID"      :   ObjectId,
                "authorName"    :   String,
                "admin"         :   Bool,
                "title"         :   String,
                "description"   :   String,
                "tags"          :   []String,
                "createdAT"     :   Time,
                "updatedAT"     :   Time
            }...,
		]
    ```


## Create News
* HTTP Request : ```POST http://api.com/news```
* Send News's data in the request body in the following format 

	``` 
        [
			{
                "classid"       :   ObjectId,
                "authorID"      :   ObjectId,
                "authorName"    :   String,
                "admin"         :   Bool,
                "title"         :   String,
                "description"   :   String,
                "tags"          :   []String,
			}...
		]
    ```

* http StatusCreated (201) will be sent if the student has been created correctly


## Update News
* HTTP Request : ```PUT http://api.com/news```
* Send data in the request body in the following format

    ``` 
        [
			{
				"id"           :	ObjectId,
				"classid"       :	ObjectId,
				"title"         :	String,
				"description"   :	String
                "tags"          :   []String
			}...
		]
    ```

* http StatusCreated (201) will be sent if the student has been updated correctly


## Delete News
* HTTP Request : ```DELETE http://api.com/news```
* Send data in the request body in the following format

	``` 
		[
			{  
				"id" : ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly
