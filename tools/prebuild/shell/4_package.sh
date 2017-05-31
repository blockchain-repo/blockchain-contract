#!/bin/bash

# "----------------------------------------------"
current_path=`pwd`
#默认的打包名称
package_name="package.tar.gz"
# "----------------------------------------------"

# "----------------------------------------------"
cd $current_path
if [ -d "package" ]
then
    rm -rf "package"
fi
mkdir "package"
mkdir "package/log"
mkdir "package/data"
cp -R "bin" "package"
cp -R "conf" "package"

cd $current_path
if [ -f $package_name ]
then
    rm -f $package_name
fi
tar -zcvf $package_name "package"
rm -rf "package"
# "----------------------------------------------"

# "----------------------------------------------"
echo "ok"
# "----------------------------------------------"
