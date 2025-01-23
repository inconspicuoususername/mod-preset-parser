pnpm run build
node --experimental-sea-config sea-config.json
node -e "require('fs').copyFileSync(process.execPath, './build/parser.exe')"
@REM signtool remove /s parser.exe
pnpx postject ./build/parser.exe NODE_SEA_BLOB ./build/sea-prep.blob ^
    --sentinel-fuse NODE_SEA_FUSE_fce680ab2cc467b6e072b8b5df1996b2