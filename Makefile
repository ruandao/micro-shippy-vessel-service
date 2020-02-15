build:
	protoc -I=. --go_out=plugin=micro:. proto/vessel/vessel.proto