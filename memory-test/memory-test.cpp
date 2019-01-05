#include <cstdio>

#include <unistd.h>
#include <signal.h>

#define LENGTH 2000000

void handler(int signum)
{
    printf("siginthandler: %d\n",signum);
}

int main() {

    char *p = new char[LENGTH];
    for(int i = 0; i < LENGTH; i ++)
        p[i] = i % 37;

    signal(SIGHUP, handler);
    signal(SIGINT, handler);
    signal(SIGQUIT, handler);
    signal(SIGILL, handler);
    signal(SIGTRAP, handler);
    signal(SIGABRT, handler);
    signal(SIGBUS, handler);
    signal(SIGFPE, handler);
    signal(SIGKILL, handler);
    signal(SIGUSR1, handler);
    signal(SIGSEGV, handler);
    signal(SIGUSR2, handler);
    signal(SIGPIPE, handler);
    signal(SIGALRM, handler);
    signal(SIGTERM, handler);

    signal(SIGCHLD, handler);
    signal(SIGSTOP, handler);

    printf("init done.\n");
    while(1) {
        sleep(3000);
    }
    return 0;

}
