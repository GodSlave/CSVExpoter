set Client_Code_PATH="E:\Works\Cloud_Unreal\Source\Cloud\GameData\"

set Generated_Client_CODE_PATH="%~dp0%\out\client\CODE\*"
set Generated_Client_CSV_PATH="%~dp0%\out\client\CSV\*"

echo %Generated_Client_Code_PATH%

csvExporter.exe

xcopy %Generated_Client_CODE_PATH% %Client_Code_PATH% /s /i /y