#!/usr/bin/env python
# -*- encoding: utf-8 -*-
import os
import sys
import time
import atexit
import signal
import util

class Daemon(object):
    def __init__(self, pidfile='/tmp/daemon.pid', stdin='/dev/null', stdout='/dev/null', stderr='/dev/null',input_data={}):
        self.stdin = stdin
        self.stdout = stdout
        self.stderr = stderr
        self.pidfile = pidfile
        self.input_data = input_data
        self.now_path = os.getcwd()

    def check_exist(self):
        i_cmd = self.input_data["cmd"]
        check_pid,check_cmd,pro_cmd = util.get_proc(i_cmd)

        if check_pid is not None and check_cmd is not None:
            raise RuntimeError('Already running.')
        else:
            return True

    def daemonize(self):
        if os.path.exists(self.pidfile):
            #raise RuntimeError('Already running.')
            pass

        if self.check_exist():
            pass

        # First fork (detaches from parent)
        try:
            if os.fork() > 0:
                raise SystemExit(0)
        except OSError as e:
            raise RuntimeError('fork #1 faild: {0} ({1})\n'.format(e.errno, e.strerror))

        os.chdir('/')
        os.setsid()
        os.umask(0o22)

        # Second fork (relinquish session leadership)
        try:
            if os.fork() > 0:
                raise SystemExit(0)
        except OSError as e:
            raise RuntimeError('fork #2 faild: {0} ({1})\n'.format(e.errno, e.strerror))

        # Flush I/O buffers
        sys.stdout.flush()
        sys.stderr.flush()

        # Replace file descriptors for stdin, stdout, and stderr
        with open(self.stdin, 'rb', 0) as f:
            os.dup2(f.fileno(), sys.stdin.fileno())
        with open(self.stdout, 'ab', 0) as f:
            os.dup2(f.fileno(), sys.stdout.fileno())
        with open(self.stderr, 'ab', 0) as f:
            os.dup2(f.fileno(), sys.stderr.fileno())

        # Write the PID file
        with open(self.pidfile, 'w') as f:
            #print os.getpid() >> file=f
            print>>f,os.getpid()

        # Arrange to have the PID file removed on exit/signal
        atexit.register(lambda: os.remove(self.pidfile))

        signal.signal(signal.SIGTERM, self.__sigterm_handler)

    # Signal handler for termination (required)
    @staticmethod
    def __sigterm_handler(signo, frame):
        raise SystemExit(1)

    def start(self):
        try:
            self.daemonize()
        except RuntimeError as e:
            #print(e, file=sys.stderr)
            print>>sys.stderr,e
            raise SystemExit(1)

        self.run()

    def stop(self):
        try:
            if os.path.exists(self.pidfile):
                with open(self.pidfile) as f:
                    os.kill(int(f.read()), signal.SIGTERM)
            else:
                #print('Not running.', file=sys.stderr)
                print>>sys.stderr,'Not running.'
                raise SystemExit(1)
        except OSError as e:
            if 'No such process' in str(e) and os.path.exists(self.pidfile): 
                os.remove(self.pidfile)

    def restart(self):
        self.stop()
        self.start()

    def run(self,input_data):
        pass
        #while 1:
            #sys.stdout.write("Daemon Alive! {}\n".format(time.ctime()))
            #sys.stdout.flush()
            #time.sleep(1)








