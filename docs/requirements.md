# System requirement specification

## Introduction

### Purpose

Tài liệu này trình bày chi tiết hệ thống blockchainx. Nó giới thiệu tổng quan
về hệ thống, đối tượng sử dụng hệ thống, liệt kê các chức năng mà nó cần đáp ứng, đồng thời đưa ra các yêu cầu về giao diện sử dụng.

Tài liệu sẽ được cập nhập liên tục để đáp ứng với yêu cầu mới hay thay đổi từ bên ngoài.

### Indended Audiences:

Tài liệu là cơ sơ để thiết kế kiến trúc hệ thống (xem [system design document](design-document.md)), được sử dụng bởi developers để tham khảo khi phát triển hệ thống. Nó cũng được dùng làm cơ sở để
đánh giá tính hoàn thiện của hệ thống. Ngoài ra nó cũng giúp nhà đầu tư, hay
người dùng có thể hiểu cách hệ thống vận hành.

### Terms and Definitions

- *Nhà sản xuất*: là cá nhân, hộ kinh doanh hay doanh nghiệp nhỏ có nhu cầu minh bạch hóa thông tin về quy trình sản xuất sản phẩm với mục đích tăng doanh thu bán hàng. 
- *Nhà tiêu thụ*: là các chủ siêu thị thu mua sản phẩm từ nhà sản xuất để bán trên trang thương mại điện tử của họ

- *Quản trị viên*: là người điều hành hệ thống blockchainx

- *Hỗ trợ viên*: là người có thể sử dụng các tính năng nội bộ của hệ thống để hỗ trợ người dùng (nhà tiêu thụ)

- *Tài nguyên sản xuất*: sở hữu của nhà sản xuất, là phương tiện để tạo ra sản phẩm đem bán (ví dụ ao tôm, vườn chuối ...). Chúng được số hóa trên hệ thống với định danh riêng và thông tin mô tả chi tiết. 

- *Sản phẩm thô*: Là gói sản phẩm chuyển từ nhà sản xuất sang nhà tiêu thị. Có
  thể tách ra thành nhiều sản phẩm trên kệ.
  
- *Sản phẩm trên kệ*: là sản phẩm bày bán trên kệ siêu thị, có gắn mã vạch và có
  thể truy xuất nguồn gốc


## Overall description

BlockchainX là một hệ thống nền tảng cho phép lưu trữ và truy xuất thông tin
nguồn gốc sản phẩm. Đối tượng chủ yếu là các sản phẩm có bán ở các siêu thị
thương mại điện tử. Nó cho phép nhà sản xuất có thể chủ động công khai minh
bạch thông tin về quy trình sản xuất và liên kết chúng với từng sản phầm đầu ra
 trưng bán trên các trang thương mại điện tử được liên kết với hệ thống.
 
 [Figure: tổng quan hệ thống - sẽ cập nhập]
 
- Người dùng (hộ sản xuất) kết nối với hệ thống thông qua mobile app (chính) và giao diện web để cập nhập quy trình sản xuất, đăng ký sản phẩm
- Siêu thị kết nối với hệ thống qua giao diện lập trình (restful API) để truy xuất thông tin
- Quản trị viên/hỗ trợ viên truy cập hệ thống qua giao diện web để quản lý thông tin và điều khiển hệ thống.
- Backend - front end, cơ sở dữ liệu, và mạng blockchain 

....

 - Hộ sản xuất: sử dụng hệ thống để quản lý tài nguyên sản xuất và cập nhập các sự kiện
   liên quan đến quá trình sản xuất sản phẩm. Kết nối thông tin này với sản
   phẩm đóng gói. Trên cơ sở minh bạch thông tin về quy trình sản xuất để tạo uy tín và thương hiệu từ đó gia tăng doanh thu hiện tại hay tìm kiếm kênh bán hàng mới.
 
 - Chủ siêu thị: sử dụng hệ thống với hai mục đích:  1) truy xuất thông tin nhà sản xuất để đánh giá mức độ tin cậy từ đó hỗ trợ việc ra quyết định hợp tác kinh doanh. 2) kết nối với hệ thống để lấy thông tin nguồn gốc sản phẩm kéo về trang thương mại điện tử của họ.
 
 - Quản trị viên:  giám sát và điều hành hệ thống blockchainx.

 - Hỗ trợ viên: truy cập thông tin người dùng hệ thống để hỗ trợ kịp thời
   các khó khăn trong việc sử dụng hệ thống. Giải đáp các khúc mắc của người
   dùng. Hỗ trợ viết content quảng cáo cho khách hàng.

## Functional requirements

  - Quản lý hộ sản xuất (đang kí, chỉnh sửa thông tin, khóa tài khoản etc.)
  - Quản lý tài nguyên và quy trình sản xuất (tạo chỉnh sửa tài nguyên, cập nhập sự kiện trong quá trình sản xuất)
  - Tạo mã định danh cho sản phẩm, liên kết với thông tin quy trình sản xuất
  - Tạo catalog cho hộ sản xuất
  - 
   ....

## External Interface requirements

### Đối với nhà sản xuất:

Giao diện web hoặc mobile
 - Đăng ký tài khoản
 - Đăng nhâp tài khoản
 - Tạo tài nguyên sản xuất
 - Ghi chép sự kiện
 - Đóng gói sản phẩm đầu ra, tạo mã định danh cho sản phẩm
 - Tạo, chỉnh sửa nội dung cho trang cá nhân, catalog sản phẩm
 ...
 
### Đối với nhà tiêu thụ:

Giao diện APIs
 - truy xuất thông tin nhà sản xuất
 - truy xuất thông tin sản phẩm
...

### Đối với quản trị viên

Giao diện quản trị web/console
 - Chỉnh sửa thông tin hộ cá nhân theo yêu cầu
 - Hiển thị thông tin thống kê chung toàn hệ thống
 - Truy xuất thông tin (sản phẩm, sự kiện, người dùng)
 ..
 
### Đối với hỗ trợ viên
 - Tạo/chỉnh sửa catalog cho hộ cá nhân
...


## Non-functional requirements:

### Performance requirements

 - Hệ thống cân được thiết kế cho phép mở rộng khi lượng người dùng gia tăng 
 ...
 
### Security requirements:

 - Hỗ trợ đăng nhập 2 bước
 - Bảo mật thông tin người dùng
 ...
