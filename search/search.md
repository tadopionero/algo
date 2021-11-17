1. Binary search

chia arr thành 2 mảng trai phải rồi so sánh giá trị phần tử đó với phần tử giữa mảng

nếu chúng không bằng nhau thì dĩ nhiên một nửa không chứa mục tiêu sẽ bị bỏ đi.

Và việc tìm kiếm tiếp tục ở nửa còn lại, một lần nữa lấy phần tử ở giữa được chọn để so sánh với giá trị đích và lặp lại điều này cho đến khi tìm thấy giá trị đích.

Nếu tìm kiếm kết thúc với nửa còn lại trống, mục tiêu không nằm trong mảng.

2. linear search

scan toàn bộ array, tìm xem thằng nào bằng nó thì là ra