package middleware

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"example-nunu/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
)

// 从配置中加载公钥
func loadPublicKeyFromConfig(conf *viper.Viper) (*rsa.PublicKey, error) {
	publicKeyString := conf.GetString("security.api_encrypt.public_key")
	publicKeyPEM := []byte(publicKeyString)

	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("公钥文件格式错误")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("公钥类型错误")
	}

	return rsaPublicKey, nil
}

// 从配置中加载私钥
func loadPrivateKeyFromConfig(conf *viper.Viper) (*rsa.PrivateKey, error) {
	privateKeyPEM := []byte(conf.GetString("security.api_encrypt.private_key"))

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("私钥文件格式错误")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// 加密处理函数
func ResponseEncryptHandlerMiddleware(logger *log.Logger, conf *viper.Viper) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 设置一个缓冲区来暂存响应
		writer := &responseWriter{ResponseWriter: ctx.Writer, body: bytes.NewBufferString("")}
		ctx.Writer = writer

		// 执行请求
		ctx.Next()

		// 加载 RSA 公钥
		publicKey, err := loadPublicKeyFromConfig(conf)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "加载公钥失败"})
			return
		}

		// 获取响应数据
		responseBody := writer.body.String()

		// 将响应解析为 map 结构
		var responseData map[string]interface{}
		err = json.Unmarshal([]byte(responseBody), &responseData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "解析响应数据失败"})
			return
		}

		dataField, exists := responseData["data"]
		logger.Debug("responseData", zap.Any("dataField", dataField))
		if exists {
			encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(""))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "加密失败"})
				return
			}

			// 将加密后的数据替换 `data` 字段
			encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)
			responseData["data"] = encryptedBase64

			// 将加密后的响应重新编码为 JSON 并发送
			encryptedResponse, err := json.Marshal(responseData)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "生成加密响应失败"})
				return
			}

			// 写入加密后的响应
			ctx.Writer.Header().Set("Content-Type", "application/json")
			ctx.Writer.Write(encryptedResponse)
		}
	}

}

// responseWriter 用于捕获响应的内容
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	// 将响应内容写入缓冲区
	w.body.Write(b)
	// 继续写入原始响应
	return w.ResponseWriter.Write(b)
}

// 解密处理函数
func RequestDecryptMiddleware(conf *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData map[string]string
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求"})
			return
		}

		encryptedData := requestData["encrypted_data"]
		encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效加密数据"})
			return
		}

		privateKey, err := loadPrivateKeyFromConfig(conf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "加载私钥失败"})
			return
		}

		decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "解密失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"decrypted_data": string(decryptedBytes)})
	}
}
