package logger

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func ReqGroup() logrus.Fields {
	return logrus.Fields{
		"request": map[string]string{
			"method": "GET",
		},
	}
}

func PostGroup() logrus.Fields {
	return logrus.Fields{
		"request": map[string]string{
			"method": "POST",
		},
	}
}

func PutGroup() logrus.Fields {
	return logrus.Fields{
		"request": map[string]string{
			"method": "PUT",
		},
	}
}

func DeleteGroup() logrus.Fields {
	return logrus.Fields{
		"request": map[string]string{
			"method": "DELETE",
		},
	}
}
