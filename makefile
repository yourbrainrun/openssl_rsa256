init:
	openssl genrsa -out storage/initPri.pem 2048
	openssl pkcs8 -topk8 -inform PEM -in storage/initPri.pem -outform PEM -nocrypt -out storage/pri.pem
	openssl rsa -in storage/pri.pem -pubout -out storage/pub.pem
