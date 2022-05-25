1. Nội dung : xây dựng CRUD API sử dụng GORM để thao tác với database

2. Sử dụng gin để tạo service, gorm để thao tác với database, database mysql build trên docker

3. Ý tưởng : dựa theo 1 phần ý tưởng của các app giao đồ ăn hiện nay, người dùng sẽ khởi tạo và chọn nhà hàng mong muốn, sau đó có thể sẽ hiện các món ăn hoặc việc like/dislike nhà hàng. Hiện tại mới làm xong API cho nhà hàng. (cần tìm hiểu việc các service/modules riêng biệt giao tiếp với nhau)

4. Các thức thực hiện : chia các thành phần thành các modules khác nhau và thực hiện trên từng module. Ví dụ module nhà hàng, module người dùng.

5. Kiến trúc : Clean Architecture, chia modules thành 3 layer khác nhau 
    Transport : http handler method pares & validate data from request
    Business : implement business logic
    storage: store & retrieve data

6. To-do : tiếp tục đọc và implement về error handling (lý thuyết về error, panic and recover, sử dụng gin middleware để handle panic recover)

7. Cấu trúc thư mục code 
    common : chứa các phương thức chung sẽ dùng nhiều lần
    component : tạo appcontext để quản lý các phương thức đầu vào của transport để dễ quản lý khi sau này source code thay đổi
    modules : là thư mục chứa các modules, ở đây ms làm đc modules restaurant(next : sẽ tạo module dành cho user)
        restauranbiz : business logic layer
        restaurant model : định dạng model của dữ liệu
        restaurantstorage: storage layer
        restauranttransport/ginrestaurant : transport layer