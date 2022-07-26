# sap-api-integrations-memo-activity-reads  
sap-api-integrations-memo-activity-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API メモアクティビティデータを取得するマイクロサービスです。  
sap-api-integrations-memo-activity-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-memo-activity-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/memoactivity/overview  

## 動作環境
sap-api-integrations-memo-activity-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-memo-activity-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-memo-activity-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/memoactivity/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-memo-activity-reads には、次の API をコールするためのリソースが含まれています。  

* ActivityCollection（メモアクティビティ - アクティビティ）

## API への 値入力条件 の 初期値
sap-api-integrations-memo-activity-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ActivityCollection.ID（ID）


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ActivityCollection" が指定されています。    
  
```
	"api_schema": "MemoActivityActivityCollection",
	"accepter": ["ActivityCollection"],
	"memo_activity_code": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "MemoActivityActivityCollection",
	"accepter": ["All"],
	"memo_activity_code": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetMemoActivity(iD string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ActivityCollection":
			func() {
				c.ActivityCollection(iD)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP メモアクティビティ の アクティビティデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "EntityLastChangedOn" は、/SAP_API_Output_Formatter/type.go 内 の Type ActivityCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-memo-activity-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-memo-activity-reads/SAP_API_Caller.(*SAPAPICaller).ActivityCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E04B6021ED2B6E362FB24C3595E",
			"ETag": "2013-06-22T18:07:02+09:00",
			"ID": "1",
			"ProcessingTypeCode": "0004",
			"PriorityCode": "3",
			"ReportedDate": "2013-06-22T09:00:00+09:00",
			"ReportedDateTime": "2013-06-22T18:06:58+09:00",
			"SubjectName": "please help the latch does not open",
			"UUID": "00163E04-B602-1ED2-B6E3-62FB24C3595E",
			"TypeCode": "39",
			"LifeCycleStatusCode": "3",
			"InitiatorCode": "2",
			"CreationDate": "2013-06-22T09:00:00+09:00",
			"EntityLastChangedOn": "2013-06-22T18:07:02+09:00"
		},
		{
			"ObjectID": "00163E063FDC1EE489B25090E9DBE3F3",
			"ETag": "2014-08-17T07:13:15+09:00",
			"ID": "1",
			"ProcessingTypeCode": "0007",
			"PriorityCode": "3",
			"ReportedDate": "2014-08-16T09:00:00+09:00",
			"ReportedDateTime": "2014-08-17T07:13:14+09:00",
			"SubjectName": "Chat From:terence.chesire@sap.com",
			"UUID": "00163E06-3FDC-1EE4-89B2-5090E9DBE3F3",
			"TypeCode": "1976",
			"LifeCycleStatusCode": "1",
			"InitiatorCode": "2",
			"CreationDate": "2014-08-16T09:00:00+09:00",
			"EntityLastChangedOn": "2014-08-17T07:13:15+09:00"
		},
		{
			"ObjectID": "00163E7842081ED9BEDBB1AA319238F9",
			"ETag": "2019-10-30T15:21:48+09:00",
			"ID": "1",
			"ProcessingTypeCode": "0011",
			"PriorityCode": "3",
			"ReportedDate": "2019-10-30T09:00:00+09:00",
			"ReportedDateTime": "2019-10-30T15:21:47+09:00",
			"SubjectName": "Memo Activity",
			"UUID": "00163E78-4208-1ED9-BEDB-B1AA319238F9",
			"TypeCode": "2574",
			"LifeCycleStatusCode": "1",
			"InitiatorCode": "1",
			"CreationDate": "2019-10-30T09:00:00+09:00",
			"EntityLastChangedOn": "2019-10-30T15:21:48+09:00"
		}
	],
	"time": "2022-07-26T20:14:33+09:00"
}

```