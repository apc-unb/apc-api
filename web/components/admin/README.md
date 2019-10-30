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
				"_id"       :	ObjectId,
				"classid"   :	ObjectId,
				"firstname" :	String,
				"lastname"  :	String,
				"matricula" :	String,
				"photourl"  :	String,
				"email"     :	String
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
				"_id"      	:   ObjectId,
				"classid"   	:   ObjectId,
				"email" 	:   String,
				"password"	:   String,
				"newpassword"	:   String,
				"photourl"  	:   String
			},...
		]
	```
* http StatusCreated (201) will be sent if the admin has been updated correctly

## Update Students from Admin request
* HTTP Request : ```PUT http://api.com/admin/student```
* Send data in the request body in the following format (```stundentid```, ```adminid``` and ```adminpassword``` is required)

	``` 
        {  
            "adminid"      	:   ObjectId,                
            "studentid"          :   ObjectId,
            "classid"   	        :   ObjectId,
            "adminpassword"      :   String,
            "firstname" 	        :   String,
            "lastname"           :   String,
            "matricula" 	        :   String,
            "handles"   	        :	{
                "codeforces"	    :	String,
                "uri"		    :	String
            },
            "photourl"  	        :   String,
            "email"  	        :   String,
            "grades"    	        :	{
                "exams"		    :	[]float64,
                "projects" 	    :	[]float64,
                "lists"    	    :	[]float64
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
            "_id"	:	ObjectId
        }
	```
* http StatusOK (200) will be sent if the Admin have been deleted correctly

## Delete Admins
* HTTP Request : ```DELETE http://api.com/admins```
* Send data in the request body in the following format

	``` 
		[
			{  
				"_id"	:	ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the Admins have been deleted correctly


## Log in Admin
* HTTP Request : ```POST http://api.com/admin/login```

    ``` 
		{
			"matricula" :	String,
			"password"  :	String,
		}
    ```
* Return a json format as follow

	```
	{
	    "userexist"	:	Boolean,
	    "admin"	:	AdminInfo,
	    "class"	:	SchoolClass,
	    "news"	:	[]News 
	}
	```
