package main

import(
  "fmt"

)

func infometion()string{
  return "http://web-int.u-aizu.ac.jp/labs/istc/ipc/index.html"
}
func literacy1()string{
  return "http://web-int.u-aizu.ac.jp/course/lit1/"
}
func literacy2()string{
  return "http://web-int.u-aizu.ac.jp/course/lit2/"
}
func sisugaiOUT()string{
  return "http://www.u-aizu.ac.jp/course/csI/"
}
func sisugaiIN()string{
  return "file:///home/course/csI/public_html/2016/welcome.html"
}
func rikougaku()string{
  return "http://web-int.u-aizu.ac.jp/course/cselab/"
}
func InformationSecurity()string{
  return "http://web-int.u-aizu.ac.jp/~yodai/course/IS/welcome.html"
}
func Multimedia3()string{
  return "http://web-int.u-aizu.ac.jp/~asada/class/MMS.html"
}
func Multimedia2()string{
  return "http://www.u-aizu.ac.jp/~shigeo/course/mms/"
}
func ComputerNet2()string{
  return "http://web-int.u-aizu.ac.jp/~pham/cn1/"
}
func ComputerNet1()string{
  return " http://web-int.u-aizu.ac.jp/~pham/cn1/"
}
func prog0OUT()string{
  return "http://www.u-aizu.ac.jp/course/prog0/"
}
func prog0IN()string{
  return "file:///home/course/prog0/public_html/2016/welcome.html"
}
func prog1OUT()string{
  return "http://www.u-aizu.ac.jp/course/prog1/"
}
func prog1IN()string{
  return "file:///home/course/prog1/public_html/2016/welcome.html"
}
func java1()string{
  return "http://web-int.u-aizu.ac.jp/~vkluev/courses/javaone/"
}
func CPP()string{
  return "授業中に指示されます"
}
func ComputerGengo()string{
  return "http://provit.u-aizu.ac.jp/clang/"
}
func algo()string{
  return "AOJでググれ"
}
func algotoku()string{
  return "http://hare.u-aizu.ac.jp/classaa/2013"
}
func gengoshori()string{
  return "http://elvis.rowan.edu/~bergmann/books.html."
}
func jyouhouassyuku()string{
  return "http://web-int.u-aizu.ac.jp/~asada/class/DC.html"
}
func keisannkikagaku()string{
  return "http://i-health.u-aizu.ac.jp/CompuGeo/index.html"
}
func Softwarekougaku()string{
  return "http://borealis.u-aizu.ac.jp/classes/se1"
}
func AI()string{
  return "http://web-ext.u-aizu.ac.jp/~qf-zhao/TEACHING/AI/AI.html"
}
func ComputerGraphics()string{
  return "http://web-int.u-aizu.ac.jp/~fayolle/teaching/2015/cg/index.html"
}
func gashori()string{
  return "http://iplab.u-aizu.ac.jp/moodle/"
}
func gashoriHand()string{
  return "http://iplab.u-aizu.ac.jp/moodle/"
}
func baio()string{
  return "http://web-ext.u-aizu.ac.jp/course/bmclass/"
}
func robot()string{
  return "http://iplab.u-aizu.ac.jp/moodle"
}
func hyuumann()string{
  return "http://web-int.u-aizu.ac.jp/~mcohen/welcome/courses/AizuDai/undergraduate/HI&VR"
}
func dezitaru()string{
  return " http://www.dspguide.com/pdfbook.htm"
}
func Webengin()string{
  return "http://web-int.u-aizu.ac.jp/~vkluev/courses/webengineering/"
}
func Softwarekougaku2()string{
  return "http://sealpv0.u-aizu.ac.jp/moodle/"
}
func SofteareDtdio()string{
  return "http://borealis.u-aizu.ac.jp/classes/studio/"
}
func bunsann()string{
  return " www.tinyurl.com/aizu-dc"
}
func messageReturn(Message string){

  switch Message {
  case "リテラシー1":
    fmt.Println(literacy1());
    break;
  case "リテラシー2":
    fmt.Println(literacy2());
    break;
  case "情報センター":
      fmt.Println(infometion());
      break;
  case "シス外":
  case "システム概論":
    fmt.Println(sisugaiOUT());
    fmt.Println(sisugaiIN());
    break;
  case "理工学実験":
    fmt.Println(rikougaku());
    break;
  case "情報セキュリティ":
    fmt.Println(InformationSecurity());
    break;
  case "マルチメディアシステム概論":
    fmt.Println(Multimedia3());
    fmt.Println(Multimedia2());
    break;
  case "コンピュータネットワーク概論":
    fmt.Println(ComputerNet2());
    fmt.Println(ComputerNet1());
    break;
  case "プログラミング入門":
  case "プログ入門":
  case "prog0":
    fmt.Println(prog0OUT());
    fmt.Println(prog0IN());
    break;
  case "プログラミングC":
  case "プログC":
  case "plog1":
    fmt.Println(prog1OUT());
    fmt.Println(prog1IN());
    break;
  case "プログラミングjava":
  case "java":
    fmt.Println(java1());
    break;
  case "C++":
  case "プログラミングC++":
    fmt.Println(CPP());
    break;
  case "コンピュータ言語論":
    fmt.Println(ComputerGengo());
    break;
  case "アルゴ":
  case "アルゴリズムとデータ構造":
    fmt.Println(algo());
    break;
  case "アルゴリズム特論":
    fmt.Println(algotoku());
    break;
  case "言語処理系論":
    fmt.Println(gengoshori());
    break;
  case "情報圧縮":
    fmt.Println(jyouhouassyuku());
    break;
  case "計算幾何学":
    fmt.Println(keisannkikagaku());
    break;
  case "ソフトウェア工学概論":
    fmt.Println(Softwarekougaku());
    fmt.Println(Softwarekougaku2());
    break;
  case "人工知能":
  case "AI":
    fmt.Println(AI());
    break;
  case "コンピュータグラフィックス論":
    fmt.Println(ComputerGraphics());
    break;
  case "画像処理":
    fmt.Println(gashori());
    fmt.Println(gashoriHand());
    break;
  case "バイオメディカル情報工学":
    fmt.Println(baio());
    break;
  case "ロボット工学と自動制御":
    fmt.Println(robot());
    break;
  case "ヒューマインインターフェイスと仮想現実":
    fmt.Println(hyuumann());
    break;
  case "デジタル信号処理":
    fmt.Println(dezitaru());
    break;
  case "ウェブエンジニアリング":
    fmt.Println(Webengin());
    break;
  case "ソフトウェアスタジオ":
    fmt.Println(SofteareDtdio());
    break;
  case "分散コンピューティング":
    fmt.Println(bunsann());
    break;
   default:
     fmt.Print("該当ページがないか、存在しません");
     fmt.Println(Message)
     break;




  }
}
func main(){
  messageReturn("java")
}
