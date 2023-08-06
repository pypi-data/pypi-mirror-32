@echo off
Rem Intel(R) MPI Library Build Environment

if "%1" == "quiet" (
    shift
    goto EXPORTS
) else if "%2" == "quiet" (
    goto EXPORTS
)

title <productname> Build Environment for Intel(R) 64 applications

echo.
echo <productname> Build Environment for Intel(R) 64 applications
echo Copyright 2007-2018 Intel Corporation.
echo.

:EXPORTS
SET I_MPI_ROOT=C:\TCAgent2\work\11562aa05e25f578\conda-build\impi_rt_1525422267555\_b_env
SET PATH=%I_MPI_ROOT%\Library\bin;%PATH%
SET LIB=%I_MPI_ROOT%\Library\bin;%LIB%
SET INCLUDE=%I_MPI_ROOT%\Library\include;%INCLUDE%

if /i "%1"=="debug" (
    goto EXTRA_EXPORTS
)
if /i "%1"=="release" (
    goto EXTRA_EXPORTS
)
if /i "%1"=="debug_mt" (
    goto EXTRA_EXPORTS
)
if /i "%1"=="release_mt" (
    goto EXTRA_EXPORTS
)
goto END
                       
:EXTRA_EXPORTS
set PATH=%I_MPI_ROOT%\Library\bin\%1;%PATH%

:END
