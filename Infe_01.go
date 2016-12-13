package main

type IFile interface {
	Read(buf [] byte) (len int,err error)
	Write(buf []byte)(len int,err error)
	Seek(off int,whence int)(pos int64, err error)
	Close()(err error)
}

type IRead interface {
	Read(buf []byte)(len int,err error)
}
type IWrite interface {
	Write(buf []byte)(len int,err error)
}

type IClose interface {
	Close()(err error)
}

type File struct {

}

func (file *File) Read(buf []byte) (len int,err error){
	return
}
func (file *File) Write(buf []byte)(len int,err error)  {
	return
}
func (file *File) Close()(err error)  {
	return
}
func (file *File) Seek(off,whence int)(pos int64,err error)  {
	return
}



func main() {
	var file IFile = new(File)
	byte := []byte{}

	by := make([]byte,20)
	file.Read(byte)
	file.Write(by[2:4])

}

