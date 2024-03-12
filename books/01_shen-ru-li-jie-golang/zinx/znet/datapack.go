package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"playground/books/01_shen-ru-li-jie-golang/zinx/utils"
	"playground/books/01_shen-ru-li-jie-golang/zinx/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeaderLen() uint32 {
	// Id + DataLen === uint32(4 byte) + uint32 (4 byte)
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建 byte buffer
	buf := bytes.NewBuffer([]byte{})
	//	写 data len
	if err := binary.Write(buf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	//	写 data id
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	//	写 data
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	// 创建 io reader
	buf := bytes.NewReader(data)
	msg := &Message{}
	//	读 data len
	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//	读 data id
	if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//	判断 data len 是否超出服务器允许的最大长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("数据包过大")
	}
	//	只需要讲 head 的数据拆包即可, 后续可以通过 head 中 data len, 再从 conn 中读取一次数据即可获取 data
	return msg, nil
}
