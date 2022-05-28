<div align="center">
  <h1><code>notify-con</code></h1>
  <p>
    <strong>⏲ notify-con notify when a internet connection is lost or back.</strong>
  </p>
</div>

## Usage

To create a Windows Service from an executable, you can use sc.exe:

```
sc create InternetCheckService binPath="<path_to_the_service_executable>"
sc start create InternetCheckService
```

```
nssm install InternetCheckService "<path_to_the_service_executable>"
```

As for me, I use this path `C:\bin\notify-con.exe`.
