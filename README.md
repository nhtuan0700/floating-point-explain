Dấu phẩỷ động được biểu diễn ở binary ở dạng 32bit hoặc 64 bit

32 bit:
1. sign (1 bit): bit đầu tiên
   - dùng để chỉ dấu + -
2. biased exponent (8 bit): 2 -> 9
   - dùng để chỉ vị trí dấu chấm di chuyển 
3. mantisa (23 bit): 10 -> 32
   - phần thập phân

64 bit:
1. sign (1 bit): bit đầu tiên
   - dùng để chỉ dấu + -
2. biased exponent (11 bit): 2 -> 12
   - dùng để chỉ vị trí dấu chấm di chuyển 
3. mantisa (52 bit): 13 -> 64
   - phần thập phân

=====
Tiếng việt: https://viblo.asia/p/01-02-03000004-lap-trinh-vien-khong-nen-bo-qua-khai-niem-nay-3RlL5GBP4bB


Tiếng anh: https://www.geeksforgeeks.org/ieee-standard-754-floating-point-numbers/

Tool convert:
https://baseconvert.com/ieee-754-floating-point
# Lưu ý

Bởi vì tính chất của dấu chấm động nên 1 số phép toán sẽ không đúng bởi vì nó sẽ làm tròn

## Ví dụ 1: `0.1`
- `0.1` được lưu trong máy tính ở dạng binary: 
  - `0 01111011 10011001100110011001101`

- => nhưng khi chuyển từ binary sang giá trị thập phân chính xác thì nó:
  - `0.10000000149011612`


Giải thích
```
Tính toán cụ thể, mình từng note lại kiểu này hi vọng bạn hiểu sau khi đọc link trên.

` không dùng chuẩn

0.1 ->

bin: 0.00011001100110011001100110011001100110011001100110011010.....

dec: 0.09999..

chuẩn IEEE754

0.1 ->

binary: 0 01111011 10011001100110011001101 [32bit]

dec: 0.100000001490116119384765625

==================================================================

tính 0.1 theo IEEE 754 - 32 bit,

chọn faction (mantissa) [chỉ chọn 23 hoặc 52 bit], mình chọn 23 bit (52 bit cho IEEE754 - 64bit)

=> sign = 0 [0 cho không dấu, 1 cho có dấu]

=> phần nguyên là 0, không cần đổi sang binary

phần thập phân nhân 2 cho tới khi bằng 1.0

=> 0.0001 10011001100110011001100110011001... [vô hạn]

=> exponent = 2^(-4) // dịch chuyển dấu chấm động sang phải, dịch [trái là + /phải là -] cho đến khi có dạng 1.xyz => 1.10011001100110011001100110011001......

=> biased exponent = 127 + (-4) = 123 [binary: 01111011]

bias cho 32bit = 127, 64bit = 1023

faction: '100110011001100110011001'.length // 24, nhưng bit thứ 24 có số '1' nên bit thứ 23 phải làm tròn lên 1 (Có thuật toán làm tròn mà mình cũng không rõ cơ chế, quá phức tạp, nếu bạn check ở dạng IEEE 754 - 64 bit bạn sẽ thấy khác biệt - chuẩn xác hơn)

=> mantissa: '10011001100110011001101'

<sign> <biased exponent [dạng binary]> <mantissa>

=> 0 01111011 10011001100110011001101

decimal: 0.100000001490116119384765625

nhưng default precision của javascript thể hiện tới 19 kí tự, default precision của PHP chỉ tới 17 kí tự (muốn đúng như js phải sửa trong php.ini)

=> làm tròn dẫn tới decimal thể hiện chỉ là : 0.10000000149011612

```

## Ví dụ 1: 0.2

- `0.2` được lưu trong máy tính ở dạng binary: 
  - `0 01111100 10011001100110011001101`

- => nhưng khi chuyển từ binary sang giá trị thập phân chính xác thì nó:
  - `0.20000000298023223876953125`


=> Từ đó dẫn đến kết quả 0.1 + 0.2 = 0.30000...04
