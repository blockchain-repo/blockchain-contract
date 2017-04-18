#!/bin/bash

# "----------------------------------------------"
current_path=`pwd`
#默认的打包名称
package_name="package.tar.gz"
# "----------------------------------------------"

# "----------------------------------------------"
cd $current_path
mkdir package
cp -R bin package
cp -R conf package
cp -R data package
cp -R log package

cd $current_path
tar -zcvf $package_name ./package
# "----------------------------------------------"

# "----------------------------------------------"
echo "ok"
# "----------------------------------------------"
