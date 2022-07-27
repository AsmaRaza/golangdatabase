## Database Operations in Golang
## Dependencies 
[Golang](https://go.dev/learn/)
[Postman](https://www.postman.com/)
[Postgresql](https://www.pgadmin.org/)
### 
first of all we create sechema in pgadmin and create a table in database like this format
create table "myscehma".arthimetic{
id SERAIL PRIMARY KEY,
Num1 int not null,
Num2 int not null,
result float not null,
operations text not null
}
After than run program 
## command
go run .\main.go
## output be like
![database](https://user-images.githubusercontent.com/93153939/181231743-d44b62ff-149f-496c-a7b8-dea654200d72.PNG)
![del](https://user-images.githubusercontent.com/93153939/181231771-dac9704a-e64e-4e48-a9b6-a911a49104a8.PNG)
![gett](https://user-images.githubusercontent.com/93153939/181231850-f9bfc7ba-38b6-4fc7-a869-61b062d5c204.PNG)
![post](https://user-images.githubusercontent.com/93153939/181231875-27eb554f-9a3f-4e4e-88aa-3e7cad8e894e.PNG)
