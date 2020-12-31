/*
 * 使用Go生成的动态库的命令：g++ main.cpp mq.so -o hello
 */
#include <iostream>
#include <thread>
#include "mq.h"

using namespace std;


char * GoStringToCString( GoString gstr ) {
    char *str = new char[gstr.n + 1];
    memcpy(str, gstr.p, gstr.n);
    str[gstr.n] = '\0';
    return str;
}

GoString CStringToGoString( char *str) {
    GoString gstr;
    gstr.p = str;
    gstr.n = strlen(str);
    return gstr;
}

void onMessage( char *msg ) {
    std::cout << "new message: " << msg << std::endl;
}

int main()
{
    char *str = HelloGolang( CStringToGoString( "China" ) );
    std::cout << str << std::endl;

    GoSetCallback( onMessage );

    while ( true ) {
        std::this_thread::sleep_for( std::chrono::seconds( 1 ) );
    }
    return 0;
}
