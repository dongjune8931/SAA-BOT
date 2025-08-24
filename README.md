# SAA-BOT: AI 기반 AWS SAA 문제 자동 정리 CLI

AI를 이용해 AWS SAA(Solutions Architect - Associate) 문제와 해설, 오답 노트를 Notion에 자동으로 정리해주는 CLI 툴입니다.

---

## ✨ 주요 기능 (Features)

* **🤖 AI 자동 분석**: 문제 내용을 입력하면 AI가 문제의 핵심, 정답/오답 해설, 관련 AWS 서비스 개념을 자동으로 분석하고 정리합니다.
* **✍️ Notion 자동화**: 분석된 모든 내용을 지정된 Notion 데이터베이스에 깔끔한 서식의 페이지로 자동 생성합니다.
* **💬 대화형 CLI**: 터미널에서 질문에 답하는 방식으로 쉽고 편리하게 문제를 등록할 수 있습니다.

---

## 🛠️ 사전 준비 (Prerequisites)

이 프로그램을 사용하기 전에 다음을 준비해야 합니다.

1.  **Google AI (Gemini) API 키**:
    * [Google AI for Developers](https://aistudio.google.com/app/apikey)에서 API 키를 발급받으세요.

2.  **Notion Integration 생성 및 설정**:
    * [Notion Integrations](https://www.notion.so/my-integrations)에서 새 Integration을 만들고 **Secret Key**를 복사하세요.
    * 문제를 저장할 Notion **데이터베이스**를 생성하고, 우측 상단 `···` 메뉴에서 생성한 Integration을 연결해주세요.
    * 데이터베이스의 **ID**를 확인하세요. (URL `notion.so/your-workspace/DATABASE_ID?v=...` 에서 `DATABASE_ID` 부분)

---

## 🚀 설치 및 사용법 (Installation and Usage)

1.  **프로젝트 복제 (Clone)**:
    ```bash
    git clone https://github.com/dongjune8931/SAA-BOT.git
    cd SAA-BOT
    ```

2.  **의존성 설치**:
    ```bash
    go mod tidy
    ```

3.  **환경 변수 설정**:
    프로젝트 루트 디렉터리에 `.env` 파일을 생성하고 아래 내용을 채워주세요.
    ```.env
    GEMINI_API_KEY="발급받은 Gemini API 키"
    NOTION_API_KEY="발급받은 Notion Integration 비밀 키"
    NOTION_DATABASE_ID="사용할 Notion 데이터베이스 ID"
    ```

4.  **프로그램 빌드**:
    ```bash
    go build -o saa-note .
    ```

5.  **프로그램 실행**:
    
    ```bash
    ./saa-note 
    ```

---

## 📝 사용 예시 (Example Usage)

`./saa-note`를 실행하면 환영 메시지와 함께 전용 프롬프트가 나타납니다.

1.  **프로그램 시작**:
    ```
      ██████╗  █████╗  █████╗    ██████╗  ██████╗ ████████╗
     ██╔════╝ ██╔══██╗██╔══██╗   ██╔══██╗██╔═══██╗╚══██╔══╝
     ███████╗ ███████║███████║   ██████╔╝██║   ██║   ██║
     ╚════██║ ██╔══██║██╔══██║   ██╔══██╗██║   ██║   ██║
     ███████╔╝ ██║  ██║██║  ██║   ██████╔╝╚██████╔╝   ██║
     ╚══════╝  ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═════╝  ╚═════╝    ╚═╝

    AI를 이용해 SAA 문제 노트를 Notion에 자동으로 정리하는 CLI 툴입니다.
    'add'를 입력해 새 노트를 추가하거나 'help'로 명령어 목록을 확인하세요. 종료는 'exit'입니다.
    --------------------------------------------------------------------
    SAA-BOT >
    ```

2.  **문제 추가**:
    프롬프트에 `add`를 입력하고 Enter를 누르면 문제 등록 절차가 시작됩니다.
    ```
    SAA-BOT > add
    ? 문제 번호(제목)를 입력하세요: 6번 문제
    ? 문제 내용을 모두 붙여넣은 후, Ctrl+D를 눌러 입력을 완료하세요.
    > (여기에 문제 내용 붙여넣기)
    ```

3.  **종료**:
    프로그램을 마치려면 `exit`를 입력합니다.
    ```
    SAA-BOT > exit
    SAA-BOT을 종료합니다.
    ```
    ```
