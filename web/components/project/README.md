# Projects

## Get all Projects from a student
* HTTP Request : ```GET http://api.com/project/{studentid}```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"           :	ObjectId,
				"studentid"     :   ObjectID,
				"projecttypeid" :   ObjectID,
                "monitorid"     :	ObjectId,
                "classid"       :	ObjectId,
                "createdat"     :	time.Time,
				"filename"      :	String,
                "status"        :   String
			}...
		]
    ```


## Update a project status
* HTTP Request : ```PUT http://api.com//project/status```
* Return a list of object in json format as follow

    ``` 
        {
            "_id"           :   ObjectId,
            "status"        :   String
        }
    ```


## Create Project
* HTTP Request : ```POST http://api.com/project```
* Send News's data in the request body in the following format 

	``` 
        {
            "studentid"     :   ObjectID,
            "projecttypeid" :   ObjectID,
            "classid"       :	ObjectID,
            "filename"      :	String,
        }
    ```

* http StatusCreated (201) will be sent if the project has been created correctly and will return a JSON in the following format

    ```
       {
           "status": "success",
           "content": {
               "monitorEmail": "monitor.email@gmail.com",
               "monitorName": "Monitor Name"
           }
       } 
    ```
  
 ## Check Project
 
 * HTTP Request : ```POST localhost:8080/project/check``` 
    ```
        {
            "studentid": "ObjectID",
            "projecttypeid": "ObjectID"
        }
    ```
   
 * http StatusCreated (200) will be sent if the project exist  and will return a JSON in the following format
                                                              
    ```
        {
            "filename": String,
            "monitorEmail": String,
            "monitorName": String,
            "monitorid": ObjectID,
            "projectid": ObjectID,
            "status": String,
            "updatedat": Time
        }
    ```
   
  * In case that the project don't exist a http StatusNotFound (404) will be sent if the project don't exist