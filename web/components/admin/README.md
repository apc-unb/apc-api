# Admin

## Create Admins
* HTTP Request : ```POST http://api.com/admin```
* Send admins's data in the request body in the following format 
* *PS* : Only admin with `Professor` : `True` can make this request

	``` 
		[
			{
				"classid"	:	ObjectId,
				"firstname"	:	String,
				"lastname"	: 	String,
				"matricula"	: 	String
			},...
		]
	```
* http StatusCreated (201) will be sent if the admin has been created correctly

## Create Admins by CSV file
* HTTP Request : ```POST http://api.com/admin/file```
* Send data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request

	|    ANO/SEMESTE/TURMA   |             2019/2/A 
	|------------------------|-------------------------------
	|       160140000        | 	Thiago Veras Machado    
	|       160140000        | 	Giovanni Guidini       
	|       160140000        | 	Vitor Dullens     

	``` 
		{  
			"file"	: file.csv,
		},...
	```

* http StatusCreated (201) will be sent if the admin has been created correctly

## Get all Admins
* HTTP Request : ```GET http://api.com/admin```
* Return a list of object in json format as follow

    ``` 
		[
			{
                "id"        :   ObjectId,
                "classid"   :   ObjectId,
                "firstname" :   String,
                "lastname"  :   String,
                "matricula" :   String,
                "photourl"  :   String,
                "email"     :   String,
                "projects"  :   Integer,
                "professor" :   Bool
			},...
		]
    ```

## Get all Admins from a specif class
* HTTP Request : ```GET http://api.com/admin/{classid}```
* Return a list of object in json format as follow

    ``` 
		[
			{
                "id"        :   ObjectId,
                "classid"   :   ObjectId,
                "firstname" :   String,
                "lastname"  :   String,
                "matricula" :   String,
                "photourl"  :   String,
                "email"     :   String,
                "projects"  :   Integer,
                "professor" :   Bool
			},...
		]
    ```

## Update Admins
* HTTP Request : ```PUT http://api.com/admin```
* Send data in the request body in the following format.

**PS:** (`id` and `password` required)

	``` 
    [
        {  
            "id"            :   ObjectId,
            "classid"       :   ObjectId,
            "email"         :   String,
            "password"      :   String,
            "newpassword"   :   String,
            "photourl"      :   String
        },...
    ]
	```
	
* http StatusCreated (201) will be sent if the admin has been updated correctly

## Update Students from Admin request
* HTTP Request : ```PUT http://api.com/admin/student```
* Send data in the request body in the following format (```stundentid```, ```adminid``` and ```adminpassword``` is required)

	``` 
        {  
            "adminid"           :   ObjectId,                
            "studentid"         :   ObjectId,
            "classid"           :   ObjectId,
            "adminpassword"     :   String,
            "firstname"         :   String,
            "lastname"          :   String,
            "matricula"         :   String,
            "handles"           :	{
                "codeforces"        :	String,
                "uri"               :	String
            },
            "photourl"          :   String,
            "email"             :   String,
            "grades"            :   StudentGrades {
                "exams"             :   []float64
                "lists"             :   []float64
            }
        }
	```
* http StatusCreated (201) will be sent if the student has been updated correctly by an admin

## Delete Admin
* HTTP Request : ```DELETE http://api.com/admin```
* Send data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request

	``` 
        {  
            "id"	:	ObjectId
        }
	```
* http StatusOK (200) will be sent if the Admin have been deleted correctly

## Delete Admins
* HTTP Request : ```DELETE http://api.com/admins```
* Send data in the request body in the following format

	``` 
		[
			{  
				"id"	:	ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the Admins have been deleted correctly


## Log in Admin
* HTTP Request : ```POST http://api.com/admin/login```

    ``` 
		{
            "matricula"     :   String,
            "password"      :   String,
		}
    ```
* Return a json format as follow with http StatusOK (200) if login was succeeded

	```
	{
        "admin"         :   Admin,
        "class"         :   SchoolClass,
        "news"          :   []News 
	}
	```
