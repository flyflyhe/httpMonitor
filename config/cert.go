package config

func GetRoot() []byte {
	return resourceRootPem.StaticContent
}

func GetServerPrivateKey() []byte {
	return resourcePrivateKey.StaticContent
}

func GetServerCertChain() []byte {
	return resourceCertPem.StaticContent
}

func GetClientPrivateKey() []byte {
	return resourceClientprivatePem.StaticContent
}

func GetClientCertChain() []byte {
	return resourceClientcertPem.StaticContent
}
