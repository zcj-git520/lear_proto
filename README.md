## proto3 语法学习

#### 基本规范

- 文件以.proto为文件后缀，除了结构体外，其他语句以分号结束

- 结构定义可以包含：message、service、enum

- rpc方法定义结尾的分号可有可无

- Message命名采用驼峰命名方式，字段命名采用小写字母加下划线分隔方式；举例：

  ```protobuf
  message ServerMessage{
	required string server_name = server1;
  }
  ```

- Enums类型名采用驼峰命名方式，字段名采用带大写字母加下划线分隔方式；举例：

  ```protobuf
  enum Sought{
	FIRST_VALUE = 100;
	SECONED_VALUE = 1200;
  }
  ```

- Service类型采用驼峰命名方式；举例：

  ```protobuf
  service UserService
  {
   rpc Login(LoginRequest)return(LoginResponse);
  }
  ```

  #### 字段规则

  - 字段格式：限定修饰符|数据类型|字段名称|=|字段编码值|[字段默认值]

  - 限定修饰符号 required\optional\repeated

    - required:表示是一个必须的字段，必须要发送给对方，在发送之前必须设置该字段的值，对于接收方，必须能够识别该字段的意思，在发送之前没有设置required字段或者无法识别该字段都会引起编码异常，导致消息被丢弃
    - optional:表示为可选字段，对于发送方，在发送消息时，可以选择性设置或者不设置该字段值，对于接收方，如果能够识别可选字段就进行相应的处理，如果无法识别，则就忽略该字段，消息中的其他字段正常处理。因为optional字段的特性，很多接口在升级版本中把后面添加大的字段都统一设置为optional字段，这样在老版本无法升级的时，也能正常于新的软件进行通信，只不过新的字段无法识别而已，因为并不是每个节点都需要新的功能，因此可以做到按需升级和平滑过渡
    - repeated:表示该字段可以包含0-n个元素，其特性和optional一样，但是每一次可以包含个多个值，可以看作是在传递一个数组的值

  - 数据类型：protobuf定义了一套基本的数据类型，几乎可以映射到c++\java等语言的基础数据类型

    ![image-20220818205947322](image-20220818205947322.png)

- 字段名称
  -  字段名称的命名与c、c++、java等语言的变量命名方式几乎相同
  - protobuf建议字段的命名采用采用下滑线分割的驼峰式
- 字段编码值
  - 有了字段编码值，通信双方才能互相识别对方的字段，相同的编码值，其限定修饰符和数据类型必须相同，编码值的范围为：1-2^32
  - 其中1-15的编码时间和空间效率都很高，编码值越大，其编码的是时间和空间效率就越低，所以建议把经常要传递的值把其字段编码设置为1-15之前的值
  - 1900-2000编码值为Google protobuf系统内部保留值，建议不要使用
- 字段默认值
  - 当在传递数据时，对于required数据类型，如果用户没有设置值，则使用默认值传递到对端

#### service定义

- 如果想要将消息类型用在rpc系统中，可以在.proto文件中定义一个RPC服务接口， protocol buffer 编译器会根据所选择的不同语言生成对应的服务器接口代码

- 例如，想要定义一个RPC服务并具有一个方法，该方法接收SearchRequest并返回一个SearchResponse，此时可以在.proto文件中进行如下定义：

  ```protobuf
   service SearchService {
          rpc Search (SearchRequest) returns (SearchResponse) {}
      }
  ```

- 生成的接口代码作为客户端与服务器的约定。服务端必须实现定义的所有接口方法，客户端直接调用同名方法向服务端发送请求，值得一提的是，即使业务上不需要参数也必须指定一个请求消息，一般会定义空的message

#### Message定义

- 一个message类型定义了一个请求或者响应的消息格式，可以包含多种类型的字段

- 字段名用小写，转为go文件时会自动转为大写，message就相当于结构体

- 首行声明使用的protobuf版本为proto3

  ```protobuf
  syntax = "proto3"
  massage SearchRequest{
   	string name = "zcj";  // 查询的名字
   	int32  pirce= 600;    //价格
   	int32 producet_num = 1; // 产品编号
  }
  ```

- 一个.proto文件中可以定义多个消息类型，一般用于同时定义的多个相关的信息

  ```protobuf
  syntax = "proto3"
  massage SearchRequest{
    string name = "zcj";  // 查询的名字
    int32  pirce= 600;    //价格
    int32 producet_num = 1; // 产品编号
  }
  // 响应
  message SearchResponse{
  // 响应信息
  }
  ```

- message 支持嵌套使用，作为另一个message中的字段类型
    ```protobuf
    message Result{
        string name = "zcj";
        string title = "proto3";
        int32 price = 1000;
    }
    message SearchResponse{
        repeated Result result = 1;
    }
    ```

- 支持嵌套信息，消息可以包含另一个消息作为字段，也可以在消息内定义一个新的消息

- 内部声明的message类型明显可以在内部直接使用
    ```protobuf
    message SearchReponse{
        message Resylt{
        string name = "zcj";
        string title = "proto3";
        int32 price = 1000;
        }
        repeated Resylt re = 1;
    }
    ```

#### proto3的Map类型

- proto3支持map类型声明

- 键、值类型可以是内置的类型，也可以是自定义的message类型

- 字段不支持repeated属性

    ```protobuf
  map<key_type, value_type>map_filed = N;
  //支持定义的类型
  message Produce{
	string name = "zcj";
  }
  map<string, Product> pro = 1;
  ```

  #### proto3文件编译

  - 通过定义好的.proto文件生成的Java，python，c/c++， GO等语言的代码，需要安装编译器protoc
  - 当使用protocol buffer 编译器允许.proto文件时，编译器生成所选择语言的代码，用于使用在.proto文件中定义的消息类型，服务接口约定等，不同的语言生成代码格式不同：
    - c++:每个.proto文件生成一个.h和一个.cc文件，每一个消息类型对应一个类
    - Java：生成一个.Java文件，同样每个消息对应一个类，同时还有一个特殊的Builder类用于创建信息接口
    - Go：生成一个.pb.go文件，每一个消息类型对应这一个结构体
    - Python: 姿势不太一样，每个.proto文件中的消息类型生成一个含有静态描述符的模块，该模块与一个元类metaclass在运行时创建需要的Python数据访问类

  #### import导入定义

  - 可以使用import 接口语句导入使用其他描述文件中的声明的类型
  - protobuf 接口文件可以像C语言的#include或者Java的import的行为大致一致
  - protocol buffer编译器会在-I /--proto_path参数指定的目录中查找导入的文件，如果没有指定该参数，默认使用当前的目录查找

  #### 包的使用

  - 在proto文件中使用package声明包，避免命名冲突
  - 在其他的消息格式定义中可以使用包名+消息名的方式使用类型
  - 在不同语言中，包含定义对编译后生成的代码影响不同：
    - C++：对应C++命令空间
    - Java 中：package会作为Java包名，除非指定了option jave_package选项
    - Python 中：package被忽略
    - Go 中：默认使用package名作为包名，除非指定了option go_package选项

  #### 测试
  - 测试将proto 转化为go/c++/python
  - 测试message与service