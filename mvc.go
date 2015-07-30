package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"github.com/mattn/go-colorable"
	"os"
)

const (
	Red    = "\x1b[31;1m"
	Green  = "\x1b[32;1m"
	Yellow = "\x1b[33;1m"
	White  = "\x1b[37;1m"
	End    = "\x1b[0m"
)

var (
	stdout = colorable.NewColorableStdout()
)

type CoreValue struct {
	title string
	what  string
}

var coreValues []CoreValue = []CoreValue{
	{"1. User-First",
		"ユーザのことを一番に考え、期待を超える"},
	// <<趣旨>>
	// 自分ファーストではなく、外向きに。そして我々の提供するサービスを使ってくれる人
	// （顧客の先にいるエンドユーザー）の考えを超えて「あっ、これが欲しかった」と思うものを提供する。

	{"2. Think Big",
		"大胆に構想し、出来ない理由ではなく可能にする手段を考える"},
	// <<趣旨>>
	// 物事を発想するときに、できそうなことを小さく考えるのではなく、
	// でっかく考え、そしてできない理由は探さずどうやったらできるかを考える

	{"3. Challenge",
		"チャンスを逃さず、挑戦することをためらわない"},
	// <<趣旨>>
	// 「あの時やっていれば」という後悔・言い訳をせずに、
	// やれると思ったら躊躇せずチャレンジする

	{"4. Commitment",
		"プロのビジネスパーソンとして妥協せず、継続して最後までやりきる"},
	// <<趣旨>>
	// （給料という対価を受け取る）プロとしてビジネスを行うからには、
	// 本気で（妥協せず）そしてすぐに成功しなくてもあきらめずにやりきる

	{"5. Professionalism",
		"得意分野を作り出し、専門性で尖り続ける"},
	// <<趣旨>>
	// ACCESSならではの得意分野（差別化できる分野）を作り、
	// そこの第一人者（ニッチトップ）となるため、
	// 広く浅く取り組むのではなく、専門性を発掘していく

	{"6. Collaboration",
		"ひとりひとりに敬意を払い、互いの英知を結集する"},
	// <<趣旨>>
	// 相手の立場や専門性を理解・尊敬して、英知を集結させることで
	// 一人一人そして全体でのパフォーマンスを最大化する協業を行う

	{"7. Deliver WOW",
		"ワクワクを作ってワクワクを提供する"},
	// <<趣旨>>
	// せっかく仕事をするのであれば相手がワクワクするようなものを提供する方がよい。
	// そのためには自分がワクワクすると思えるものを作る必要がある

	{"8. Simplicity",
		"合理性を重んじ、ものごとをシンプルに進める"},
	// <<趣旨>>
	// プロセス・開発・デザイン等全ての業務について
	// 無駄をなくし、合理的に、そして簡単に行う

	{"9. Evolution",
		"現状に満足することなく変革し続ける"},
	// <<趣旨>>
	// 「現状に満足せず、過去の失敗から学び、過去の成功も超える」
	// という意識で常に進化し続ける

	{"10. Frugality",
		"コストを意識し、常に創意工夫を心がける"},
	// <<趣旨>>
	// 企業の利益を最大化するために、コスト意識をしっかり持ち、
	// 限られた予算の中で創意工夫をしていく。
	// （ただし必要だと思うことには惜しまず予算は投資する）

}

type Options struct {
	Random []bool `short:"r" long:"random" description:"Show one of MVC in random"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	parser.Name = "acsmvc"
	parser.Usage = "[OPTIONS]"
	_, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}

	if opts.Random != nil && opts.Random[0] {
		fmt.Println("random mode")
		os.Exit(0)
	} else {
		fmt.Println("all mode")
		os.Exit(0)
	}

	fmt.Fprintln(stdout, "")
	fmt.Fprintln(stdout, Red+"Mission"+End)
	fmt.Fprintln(stdout, "")
	fmt.Fprintln(stdout, Green+"我々は、ソフトウェアを人々の身近な存在にし、世界に新しい日常を提供し続けます。"+End)
	fmt.Fprintln(stdout, "")
	// <<趣旨>>
	// あらゆる人々（ITリテラシーが低い人でも、気づかないところでも）が
	// 当たり前のように使うソフトウェアを通じ、人々に新しい体験や価値を
	// 提供し暮らしをより豊かにする。
	// そのようなソフトウェアを提供し続ける。

	fmt.Fprintln(stdout, "")
	fmt.Fprintln(stdout, Red+"Vision"+End)
	fmt.Fprintln(stdout, "")
	fmt.Fprintln(stdout, Green+"自らの技術力と人間力で 「あらゆる機器・サービス・人」をつなげ、"+End)
	fmt.Fprintln(stdout, Green+"ユーザ・業界・地域の課題を解き続けるイノベーションソフトウェア会社となる。"+End)
	fmt.Fprintln(stdout, "")
	// <<趣旨>>
	// 創業以来培ってきた強みである「つなぐ」を軸に、 エンドユーザ（最終使用者）、
	// お客様を含めた業界全体、 そして地域の今ある課題やまだ見ぬ課題に気づき、
	// 技術力だけでなく、社員が持つ熱意や魅力・能力などの総合的な人間力を発揮し、
	// ソフトウェアを通じ革新的な解決策を提供し続ける。
	// そのようなイノベーションソフト企業になる。

	fmt.Fprintln(stdout, "")
	fmt.Fprintln(stdout, Red+"Core Value"+End)
	fmt.Fprintln(stdout, "")

	for i := 0; i < len(coreValues); i++ {
		fmt.Fprintln(stdout, Green+coreValues[i].title+End)
		fmt.Fprintln(stdout, Yellow+coreValues[i].what+End)
		fmt.Fprintln(stdout, "")
	}
}
