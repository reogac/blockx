# System requirements document

## Introduction

### Purpose

Tài liệu này trình bày chi tiết hệ thống blockchainx. Nó giới thiệu tổng quan
về hệ thống, chức năng của nó, đối tượng sử dụng hệ thống, các tính năng mà nó
cần phải đáp ứng.

### Indended Audiences:

Tài liệu là cơ sơ để thiết kế hệ thống. Nó cũng được dùng làm cơ sở để
đánh giá tính hoàn thiện của hệ thống. Ngoài ra nó cũng giúp nhà đầu tư, hay
người dùng có thể hiểu cách hệ thống sẽ hệ thống.

### Terms and Definitions

- Nhà sản xuất: 
- Nhà tiêu thụ:

- Nhà quản trị hệ thống:
- Hỗ trợ viên:

- Tài nguyên sản xuất: sở hữu của nhà sản xuất là phương tiện để tạo ra sản
  phẩm đem bán. 
- Sản phẩm thô: Là gói sản phẩm chuyển từ nhà sản xuất sang nhà tiêu thị. Có
  thể tách ra thành nhiều sản phẩm trên kệ.
- Sản phẩm trên kệ: là sản phẩm bày bán trên kệ siêu thị, có gắn mã vạch và có
  thể truy xuất nguồn gốc

## Overall description

BlockchainX là một hệ thống nền tảng cho phép lưu trữ và truy xuất thông tin
nguồn gốc sản phẩm. Đối tượng chủ yếu là các sản phẩm có bán ở các siêu thị
thương mại điện tử. Nó cho phép nhà sản xuất có thể chủ động công khai minh
bạch thông tin về quy trình sản xuất và liên kết chúng với từng sản phầm đầu ra
 trưng bán trên các trang thương mại điện tử được liên kết với hệ thống.

### Người dùng:

 - Cá nhân, hộ sản xuất: quản lý tài nguyên sản xuất và cập nhập các sự kiện
   liên quan đến quá trình sản xuất sản phẩm. Kết nối thông tin này với sản
   phẩm đóng gói.
 
 - Chủ siêu thị: Là đối tượng thu ua sản phẩm từ nhà sản xuất và bán lại (chỉa
   lẻ hoặc nguyên khối) trên trang thương mại. Họ có thể truy xuất thông tin cá
   nhân/hộ sản xuất để đánh giá mức độ uy tín. Có thể truy xuất thông tin người
   gốc của sản phẩm để kéo về trang thương mại điện tử của họ.
 
 - Quản trị hệ thống blockchainx:  giám sát và điều hành hệ thống blockchainx.

 - Nhân viên hỗ trợ: truy cập thông tin người dùng hệ thống để hỗ trợ kịp thời
   các khó khăn trong việc sử dụng hệ thống. Giải đáp các khúc mắc của người
   dùng. Hỗ trợ viết content quảng cáo cho khách hàng.

 (các use case)


### Tính năng của hệ thống:
  - Quản lý người dùng hệ thống
  - Quản lý tài nguyên củaA nhà sản xuất
  - Quản lý chuỗi sự kiện trong quy trình sản xuất và liên kết tới sản phẩm
  - Quản lý sản phẩm mà nhà sản xuất tạo ra
	

## External Interface requirements

### Đối với người dùng là hộ sản xuất:

 - Đăng ký tài khoản
 - Đăng nhâp tài khoản
 - Tạo tài nguyên sản xuất
 - 
### Đối với siêu thị:

 - truy xuất thông tin nhà sản xuất
 - truy xuất thông tin sản phẩm

### Đối với quản trị hệ thống

### Đối với nhân viên hỗ trợ


## Non-functional requirements:

### Performance requirements

 - Hệ thống cân được thiết kế cho phép mở rộng khi lượng người dùng gia tăng 

### Security requirements:

 - Hỗ trợ đăng nhập 2-stéps 

