package main

import "os"

import "testing"

func TestGetAPIKey(t *testing.T) {

	// 環境変数が設定されている時
	expected := "thisistestapikey1234567890"
	os.Setenv("BING_API_KEY", expected)
	apikey, err := getAPIKey()
	if expected != apikey {
		t.Error("環境変数が正しく取得できていません")
	}

	if err != nil {
		t.Error("環境変数が空でない時にもエラーが発生しています")
	}

	// 環境変数が設定されていない時
	expected = ""
	os.Clearenv()
	apikey, err = getAPIKey()
	if apikey != expected {
		t.Error("環境変数が正しく取得できていません")
	}

	if err == nil {
		t.Error("環境変数が空だった時にエラーが正しく発生していません")
	}
}

func TestGetImageType(t *testing.T) {
	var contentType string
	var expected string

	// JPEG のとき
	contentType = "image/jpeg"
	expected = "jpeg"
	if imageType, err := getImageType(contentType); imageType != expected {
		t.Error("正しくimage/jpegのファイルタイプを識別できません")
	} else {
		// 正しく判定できているのにエラーがでている場合
		if nil != err {
			t.Error("ContentType が image/jpeg のときにエラーがでています")
		}
	}

	// PNG のとき
	contentType = "image/png"
	expected = "png"
	if imageType, err := getImageType(contentType); imageType != expected {
		t.Error("正しくimage/pngのファイルタイプを識別できません")
	} else {
		// 正しく判定できているのにエラーがでている場合
		if nil != err {
			t.Error("ContentType が image/png のときにエラーがでています")
		}
	}

	// GIF のとき
	contentType = "image/gif"
	expected = "gif"
	if imageType, err := getImageType(contentType); imageType != expected {
		t.Error("正しくimage/gifのファイルタイプを識別できません")
	} else {
		// 正しく判定できているのにエラーがでている場合
		if nil != err {
			t.Error("ContentType が image/gif のときにエラーがでています")
		}
	}

	// それ以外の ContentType のとき
	contentType = "image/hoge"
	expected = ""
	if imageType, err := getImageType(contentType); imageType != expected {
		t.Error("正しく未定義の ContentType の処理ができていません")
	} else {
		// 判定できていない場合にエラーがでていない場合
		if err == nil {
			t.Error("未定義の ContentType の時にエラーが発生していません")
		}
	}

}
