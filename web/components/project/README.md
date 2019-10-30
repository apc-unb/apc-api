# Projects

## Get all Projects from a student
* HTTP Request : ```GET http://api.com/project/{studentid}```
* Return a list of object in json format as follow

    ``` 
    [
        {
            "_id"           :	ObjectId,
            "studentid"     :       ObjectID,
            "projecttypeid" :       ObjectID,
            "monitorid"     :	ObjectId,
            "classid"       :	ObjectId,
            "createdat"     :	time.Time,
            "filename"      :	String,
            "status"        :       String
        }...
    ]
    ```


## Update a project status
* HTTP Request : ```PUT http://api.com//project/status```
* Send Project's data in the request body in the following format 

    ``` 
        {
            "_id"           :   ObjectId,
            "status"        :   String
        }
    ```


## Create Project
* HTTP Request : ```POST http://api.com/project```
* Send Project's data in the request body in the following format 

	``` 
        {
            "studentid"     :    ObjectID,
            "projecttypeid" :    ObjectID,
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
  
# Project Type

## Create Project Type
* HTTP Request : ```POST http://api.com/project/type```
* Send ProjectType's data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request 

    ```
        {
            "name"          :   String,
            "description"   :   String,
            "ClassID"       :   ObjectID,
            "start"         :   Time,
            "end"           :   Time,
            "score"         :   Float
        }
    ```
  
## Update Project Type
* HTTP Request : ```PUT http://api.com/project/type```
* Send data in the request body in the following format (`_id` is required)
* *PS* : Only admin with `Professor` : `True` can make this request

    ```
        {
            "_id"           :   ObjectID,
            "name"          :   String,
            "description"   :   String,
            "ClassID"       :   ObjectID,
            "start"         :   Time,
            "end"           :   Time,
            "score"         :   Float
        }
    ```
 
## Delete Project Type
* HTTP Request : ```DELETE http://api.com/project/type```
* Send data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request

	``` 
        {  
            "_id"	:	ObjectId
        }
	```
* http StatusOK (200) will be sent if the Project Type have been deleted correctly
