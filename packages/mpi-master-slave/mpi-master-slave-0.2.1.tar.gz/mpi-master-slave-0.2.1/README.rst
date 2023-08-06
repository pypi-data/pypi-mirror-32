mpi-master-slave: Easy to use mpi master slave library with mpi4py
==================================================================

author: Luca Scarabello

Why MPI with Python?
--------------------

Python is such a nice and features rich language that writing the high level logic of your application with it require so little effort. Also, considering that you can wrap C or Fortran functions in Python, there is no need to write your whole application in a low level language that makes hard to change application logic, refactoring, maintenance, porting and adding new features.  Instead you can write the application in Python and possibly write in C/Fortran the functions that are computationally expensive. There is no need to write the whole application in C/Fortran if 90% of the execution time is spent in a bunch of functions. Just focus on optimizing those functions.

Also I just want to highligth that MPI allows to scale you application, that means it can handle indefinite amount of work given that enough hardware (nodes) is avalible. This shouldn't be confused with code optimization (multi threading, GPU code) that increases execution speed but don't allow you application to scale. 

MPI can also be used on your local machine spawing a process for each core, so you don't need a cluster. The resulting performance are comparable with multithreded code but usually is quicker to modify your application to support multithreding while MPI enforce a strinct design of your application as each process cannot share memory with others.


Writing you application
-----------------------

Here and at the end of this little tutorial, we'll cover meaningful example code that can serve as base for your projects.

`Example 1 <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example1.py>`__

Writing a master slave application is as simple as extenging Slave class, implementing the 'do_work' method with the specific task you need and creating a Master object that controls the slaves. In this example we use WorkQueue too, a convenient class that keeps running slaves until the work queue (the list of tasks your application has to do) is empty.


.. code:: python

    from mpi4py import MPI
    from mpi_master_slave import Master, Slave
    from mpi_master_slave import WorkQueue
    import time

    class MyApp(object):
        """
        This is my application that has a lot of work to do so it gives work to do
        to its slaves until all the work is done
        """

        def __init__(self, slaves):
            # when creating the Master we tell it what slaves it can handle
            self.master = Master(slaves)
            # WorkQueue is a convenient class that run slaves on a tasks queue
            self.work_queue = WorkQueue(self.master)

        def terminate_slaves(self):
            """
            Call this to make all slaves exit their run loop
            """
            self.master.terminate_slaves()

        def run(self, tasks=10):
            """
            This is the core of my application, keep starting slaves
            as long as there is work to do
            """
            #
            # let's prepare our work queue. This can be built at initialization time
            # but it can also be added later as more work become available
            #
            for i in range(tasks):
                # 'data' will be passed to the slave and can be anything
                self.work_queue.add_work(data=('Do task', i))

            #
            # Keeep starting slaves as long as there is work to do
            #
            while not self.work_queue.done():

                #
                # give more work to do to each idle slave (if any)
                #
                self.work_queue.do_work()

                #
                # reclaim returned data from completed slaves
                #
                for slave_return_data in self.work_queue.get_completed_work():
                    done, message = slave_return_data
                    if done:
                        print('Master: slave finished is task and says "%s"' % message)

                # sleep some time
                time.sleep(0.3)


    class MySlave(Slave):
        """
        A slave process extends Slave class, overrides the 'do_work' method
        and calls 'Slave.run'. The Master will do the rest
        """

        def __init__(self):
            super(MySlave, self).__init__()

        def do_work(self, data):
            rank = MPI.COMM_WORLD.Get_rank()
            name = MPI.Get_processor_name()
            task, task_arg = data
            print('  Slave %s rank %d executing "%s" task_id "%d"' % (name, rank, task, task_arg) )
            return (True, 'I completed my task (%d)' % task_arg)


    def main():

        name = MPI.Get_processor_name()
        rank = MPI.COMM_WORLD.Get_rank()
        size = MPI.COMM_WORLD.Get_size()

        print('I am  %s rank %d (total %d)' % (name, rank, size) )

        if rank == 0: # Master

            app = MyApp(slaves=range(1, size))
            app.run()
            app.terminate_slaves()

        else: # Any slave

            MySlave().run()

        print('Task completed (rank %d)' % (rank) )

    if __name__ == "__main__":
        main()


More advanced exaples are explained at the end of this tutorial, here is a summary:

`**Example 2** <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example2.py>`__ is the same code above without the WorkQueue class, this is helpful in case you like to know have the Master class works.

`**Example 3** <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example3.py>`__ shows how to assign specific tasks to specific slaves so that the latter can re-use part of previous work (resource already acquired or some computation already performed)

`**Example 4** <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example4.py>`__ shows how slaves can handle multiple type of tasks.

`**Example 5** <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example5.py>`__ shows how to  limit the number of slaves reserved to one or more type of tasks.



Running the application
-----------------------

::


    mpiexec -n 4 python example1.py

If you run your application on your machine, usually you get the bet performance creatin n processes, where n is the number of core of your machine. If your Master process doesn't do any computation and it is mostly sleeping, then make n = cores + 1. In our example the best performance is n=5

Output:

::

    I am  lucasca-desktop rank 3 (total 4)
    I am  lucasca-desktop rank 1 (total 4)
    I am  lucasca-desktop rank 2 (total 4)
    I am  lucasca-desktop rank 0 (total 4)
      Slave lucasca-desktop rank 2 executing "Do task" task_id "0"
      Slave lucasca-desktop rank 3 executing "Do task" task_id "1"
    Master: slave finished is task and says "I completed my task (0)"
      Slave lucasca-desktop rank 1 executing "Do task" task_id "2"
    Master: slave finished is task and says "I completed my task (1)"
    Master: slave finished is task and says "I completed my task (2)"
      Slave lucasca-desktop rank 2 executing "Do task" task_id "3"
      Slave lucasca-desktop rank 3 executing "Do task" task_id "4"
    Master: slave finished is task and says "I completed my task (3)"
    Master: slave finished is task and says "I completed my task (4)"
      Slave lucasca-desktop rank 1 executing "Do task" task_id "5"
      Slave lucasca-desktop rank 2 executing "Do task" task_id "6"
      Slave lucasca-desktop rank 3 executing "Do task" task_id "7"
    Master: slave finished is task and says "I completed my task (5)"
    Master: slave finished is task and says "I completed my task (7)"
    Master: slave finished is task and says "I completed my task (6)"
      Slave lucasca-desktop rank 1 executing "Do task" task_id "8"
      Slave lucasca-desktop rank 3 executing "Do task" task_id "9"
    Master: slave finished is task and says "I completed my task (9)"
    Master: slave finished is task and says "I completed my task (8)"
    Task completed (rank 2)
    Task completed (rank 1)
    Task completed (rank 3)
    Task completed (rank 0)




Debugging
---------

We'll open a xterm terminal for each mpi process so that we can debug each process independently:

::
 
    mpiexec -n 4 xterm -e "python example1.py ; bash"


"bash" is optional - it ensures that the xterm windows will stay open; even if finished

.. image:: https://github.com/luca-s/mpi-master-slave/raw/master/examples/debugging.png

Option 1: if you want the debugger to stop at a specific position in the code then add the following at the line where you want the debugger to stop:

::

    import ipdb; ipdb.set_trace()


Then run the application as above.


Option 2: start the debugger right after each process has started

::

    mpiexec -n 4 xterm -e "python -m pdb example1.py ; bash"


Profiling
---------

Eventually you'll probably like to profile your code to understand if there are bottlenecks. To do that you have to first include the profiling module and create one profiler object somewhere in the code


.. code:: python

    import cProfile

    pr = cProfile.Profile()


Then you have to start the profiler just before the part of the code you like to profile (you can also start/stop the profiler in different part of the code).
Once you want to see the results (or partial results) stop the profiler and print statistics.

.. code:: python

    pr.enable()

    [...code to be profiled here...]

    pr.disable()

    pr.print_stats(sort='tottime')
    pr.print_stats(sort='cumtime')


For example let's say we like to profile the Master process in the example above 

.. code:: python

    import cProfile

    [...]

        if rank == 0: # Master

            pr = cProfile.Profile()
            pr.enable()

            app = MyApp(slaves=range(1, size))
            app.run()
            app.terminate_slaves()

            pr.disable()
            pr.print_stats(sort='tottime')
            pr.print_stats(sort='cumtime')

        else: # Any slave
    [...]


Output:

::

   Ordered by: internal time

   ncalls  tottime  percall  cumtime  percall filename:lineno(function)
      100   30.030    0.300   30.030    0.300 {built-in method time.sleep}
      240    0.008    0.000    0.008    0.000 {built-in method builtins.print}
      221    0.003    0.000    0.004    0.000 master_slave.py:52(get_avaliable)
        1    0.002    0.002   30.049   30.049 example2.py:24(run)
      532    0.002    0.000    0.002    0.000 {method 'Iprobe' of 'mpi4py.MPI.Comm' objects}
      219    0.001    0.000    0.003    0.000 master_slave.py:74(get_completed)
      121    0.001    0.000    0.001    0.000 {method 'send' of 'mpi4py.MPI.Comm' objects}
      242    0.001    0.000    0.001    0.000 {method 'recv' of 'mpi4py.MPI.Comm' objects}
      121    0.001    0.000    0.003    0.000 master_slave.py:66(run)
      119    0.000    0.000    0.001    0.000 master_slave.py:87(get_data)
      440    0.000    0.000    0.000    0.000 {method 'keys' of 'dict' objects}
      243    0.000    0.000    0.000    0.000 {method 'add' of 'set' objects}
      241    0.000    0.000    0.000    0.000 {method 'remove' of 'set' objects}
      242    0.000    0.000    0.000    0.000 {method 'Get_source' of 'mpi4py.MPI.Status' objects}
        1    0.000    0.000    0.000    0.000 master_slave.py:12(__init__)
        1    0.000    0.000    0.000    0.000 example2.py:14(__init__)
        1    0.000    0.000    0.000    0.000 {method 'disable' of '_lsprof.Profiler' objects}


         3085 function calls in 30.049 seconds

   Ordered by: cumulative time

   ncalls  tottime  percall  cumtime  percall filename:lineno(function)
        1    0.002    0.002   30.049   30.049 example2.py:24(run)
      100   30.030    0.300   30.030    0.300 {built-in method time.sleep}
      240    0.008    0.000    0.008    0.000 {built-in method builtins.print}
      221    0.003    0.000    0.004    0.000 master_slave.py:52(get_avaliable)
      219    0.001    0.000    0.003    0.000 master_slave.py:74(get_completed)
      121    0.001    0.000    0.003    0.000 master_slave.py:66(run)
      532    0.002    0.000    0.002    0.000 {method 'Iprobe' of 'mpi4py.MPI.Comm' objects}
      121    0.001    0.000    0.001    0.000 {method 'send' of 'mpi4py.MPI.Comm' objects}
      242    0.001    0.000    0.001    0.000 {method 'recv' of 'mpi4py.MPI.Comm' objects}
      119    0.000    0.000    0.001    0.000 master_slave.py:87(get_data)
      440    0.000    0.000    0.000    0.000 {method 'keys' of 'dict' objects}
      243    0.000    0.000    0.000    0.000 {method 'add' of 'set' objects}
      241    0.000    0.000    0.000    0.000 {method 'remove' of 'set' objects}
      242    0.000    0.000    0.000    0.000 {method 'Get_source' of 'mpi4py.MPI.Status' objects}
        1    0.000    0.000    0.000    0.000 example2.py:14(__init__)
        1    0.000    0.000    0.000    0.000 master_slave.py:12(__init__)
        1    0.000    0.000    0.000    0.000 {method 'disable' of '_lsprof.Profiler' objects}


From the output above we can see most of the Master time is spent in time.sleep and this is good as the Master doesn't have to be busy as its role is to control the slaves. If the Master process become the bottleneck of your application, the slaves nodes will be idle due to the Master not being able to efficiently control them.


More examples covering common scenarios
---------------------------------------

Example 3
---------

In `Example 3 <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example3.py>`__ we'll see how to assign specific tasks to specific slaves so that the latter can re-use part of previous work.  This is a common scenario when a slave has to perform an initialization phase where it acquires resources (Database, network directory, network service, etc) or it has to pre-compute something, before starting its task. If the Master can assign the next task that deal with the same resources to the slave that has already loaded that resources, that would save much time becasue the slave has the resources in memory already.

This is the Slave code that simulate the time required to initialize the job for a specific resource.

.. code:: python

    class MySlave(Slave):

        def __init__(self):
            super(MySlave, self).__init__()
            self.resource = None

        def do_work(self, data):

            task, task_id, resource = data

            print('  Slave rank %d executing "%s" task id "%d" with resource "%s"' % 
                 (MPI.COMM_WORLD.Get_rank(), task, task_id, str(resource)) )

            #
            # The slave can check if it has already acquired the resource and save
            # time
            #
            if self.resource != resource:
                #
                # simulate the time required to acquire this resource
                #
                time.sleep(10)
                self.resource = resource

            # Make use of the resource in some way and then return
            return (True, 'I completed my task (%d)' % task_id)


On the Master code there is little to change from example 1. Both WorkQueue.add_work and MultiWorkQueue.add_work methods support an additional parameter **resource** that is a simple identifier (string, integer or any hashable object) that specify what resource the data is going to need. 

.. code:: python

    WorkQueue.add_work(data, resource_id=some_id)
    MultiWorkQueue.add_work(task_id, data, resource_id=some_id)


.. code:: python

    class MyApp(object):

        [...]

        def run(self, tasks=100):

            [...]

            for i in range(tasks):
                #
                # the slave will be working on one out of 3 resources
                #
                resource_id = random.randint(1, 3)
                data = ('Do something', i, resource_id)
                self.work_queue.add_work(data, resource_id)
           
            [...]



WorkQueue and  MultiWorkQueue will try their best to assign the same resource id to a slave that has previously worked with the same resource.

We can test the code and see that each slave keep processing the same resource until all the tasks associated with that resource are completed. At that point the slave starts processing another resource:

::

    mpiexec -n 4 xterm -e "python example3.py ; bash"

.. image:: https://github.com/luca-s/mpi-master-slave/raw/master/examples/example3.png


::

    mpiexec -n 6 xterm -e "python example3.py ; bash"


.. image:: https://github.com/luca-s/mpi-master-slave/raw/master/examples/example3bis.png



Example 4
---------

In `Example 4 <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example4.py>`__ we can see how to the **slaves can handle multiple type of tasks.** 

.. code:: python

    Tasks = IntEnum('Tasks', 'TASK1 TASK2 TASK3')


Instead of extending a Slave class for each type of task we have, we create only one class that can handle any type of work. This avoids having idle processes if, at certain times of the execution, there is only a particular type of work to do but the Master doesn't have the right slave for that task. If any slave can do any job, there is always a slave that can perform that task.

.. code:: python

    class MySlave(Slave):

        def __init__(self):
            super(MySlave, self).__init__()

        def do_work(self, args):
    
            # the data contains the task type
            task, data = args

            #
            # Every task type has its specific data input and return output
            #
            ret = None
            if task == Tasks.TASK1:

                arg1 = data
                [... do something...]
                ret = (True, arg1)

            elif task == Tasks.TASK2:

                arg1, arg2 = data
                [... do something...]
                ret = (True, 'All done')

            elif task == Tasks.TASK3:

                arg1, arg2, arg3 = data
                [... do something...]
                ret = (True, arg1+arg2, arg3)

            return (task, ret)


The master simply passes the task type to the slave together with the task specific data.

.. code:: python

    class MyApp(object):

        [...]

        def run(self, tasks=100):

            #
            # let's prepare our work queue. This can be built at initialization time
            # but it can also be added later as more work become available
            #
            for i in range(tasks):
                self.__add_next_task(i)
           
            #
            # Keeep starting slaves as long as there is work to do
            #
            while not self.work_queue.done():

                #
                # give more work to do to each idle slave (if any)
                #
                self.work_queue.do_work()

                #
                # reclaim returned data from completed slaves
                #
                for slave_return_data in self.work_queue.get_completed_work():
                    #
                    # each task type has its own return type
                    #
                    task, data = slave_return_data
                    if task == Tasks.TASK1:
                        done, arg1 = data
                    elif task == Tasks.TASK2:
                        done, arg1, arg2, arg3 = data
                    elif task == Tasks.TASK3:
                        done, arg1, arg2 = data    
                    if done:
                        print('Master: slave finished is task returning: %s)' % str(data))

                # sleep some time
                time.sleep(0.3)

    def __add_next_task(self, i, task=None):
        """
        we create random tasks 1-3 and add it to the work queue
        Every task has specific arguments
        """
        if task is None:
            task = random.randint(1,3)

        if task == 1:
            args = i
            data = (Tasks.TASK1, args)
        elif task == 2:
            args = (i, i*2)
            data = (Tasks.TASK2, args)
        elif task == 3:
            args = (i, 999, 'something')
            data = (Tasks.TASK3, args)

        self.work_queue.add_work(data)

Ourput

::

    $ mpiexec -n 16 python example4.py

    I am  lucasca-desktop rank 8 (total 16)
    I am  lucasca-desktop rank 15 (total 16)
    I am  lucasca-desktop rank 9 (total 16)
    I am  lucasca-desktop rank 0 (total 16)
    I am  lucasca-desktop rank 13 (total 16)
    I am  lucasca-desktop rank 6 (total 16)
      Slave lucasca-desktop rank 8 executing Tasks.TASK3 with task_id 0 arg2 999 arg3 something
      Slave lucasca-desktop rank 9 executing Tasks.TASK3 with task_id 1 arg2 999 arg3 something
      Slave lucasca-desktop rank 13 executing Tasks.TASK2 with task_id 2 arg2 4
      Slave lucasca-desktop rank 15 executing Tasks.TASK2 with task_id 3 arg2 6
    Master: slave finished is task returning: (True, 0, 'something'))
    Master: slave finished is task returning: (True, 1, 'something'))
    Master: slave finished is task returning: (True, 2, 'something', 'else'))
    I am  lucasca-desktop rank 1 (total 16)
    I am  lucasca-desktop rank 11 (total 16)
    I am  lucasca-desktop rank 5 (total 16)
    I am  lucasca-desktop rank 10 (total 16)
    I am  lucasca-desktop rank 2 (total 16)
    I am  lucasca-desktop rank 7 (total 16)
    I am  lucasca-desktop rank 14 (total 16)
    I am  lucasca-desktop rank 12 (total 16)
    I am  lucasca-desktop rank 4 (total 16)
    I am  lucasca-desktop rank 3 (total 16)
      Slave lucasca-desktop rank 3 executing Tasks.TASK2 with task_id 5 arg2 10
      Slave lucasca-desktop rank 2 executing Tasks.TASK1 with task_id 4
      Slave lucasca-desktop rank 6 executing Tasks.TASK3 with task_id 8 arg2 999 arg3 something
      Slave lucasca-desktop rank 5 executing Tasks.TASK3 with task_id 7 arg2 999 arg3 something
      Slave lucasca-desktop rank 4 executing Tasks.TASK2 with task_id 6 arg2 12
      Slave lucasca-desktop rank 9 executing Tasks.TASK1 with task_id 11
      Slave lucasca-desktop rank 7 executing Tasks.TASK3 with task_id 9 arg2 999 arg3 something
      Slave lucasca-desktop rank 10 executing Tasks.TASK2 with task_id 12 arg2 24
      Slave lucasca-desktop rank 12 executing Tasks.TASK3 with task_id 14 arg2 999 arg3 something
      Slave lucasca-desktop rank 11 executing Tasks.TASK1 with task_id 13
      Slave lucasca-desktop rank 13 executing Tasks.TASK2 with task_id 15 arg2 30
      Slave lucasca-desktop rank 14 executing Tasks.TASK3 with task_id 16 arg2 999 arg3 something
    Master: slave finished is task returning: (True, 5, 'something', 'else'))
    Master: slave finished is task returning: (True, 6, 'something', 'else'))
    Master: slave finished is task returning: (True, 7, 'something'))
    Master: slave finished is task returning: (True, 8, 'something'))
    Master: slave finished is task returning: (True, 9, 'something'))
    Master: slave finished is task returning: (True, 11))
    Master: slave finished is task returning: (True, 12, 'something', 'else'))
    Master: slave finished is task returning: (True, 13))
    Master: slave finished is task returning: (True, 14, 'something'))
    Master: slave finished is task returning: (True, 15, 'something', 'else'))
    Master: slave finished is task returning: (True, 16, 'something'))
    Master: slave finished is task returning: (True, 3, 'something', 'else'))
      Slave lucasca-desktop rank 8 executing Tasks.TASK1 with task_id 10
      Slave lucasca-desktop rank 1 executing Tasks.TASK2 with task_id 17 arg2 34
      Slave lucasca-desktop rank 3 executing Tasks.TASK3 with task_id 18 arg2 999 arg3 something
      Slave lucasca-desktop rank 4 executing Tasks.TASK1 with task_id 19
      Slave lucasca-desktop rank 5 executing Tasks.TASK2 with task_id 20 arg2 40
      Slave lucasca-desktop rank 10 executing Tasks.TASK1 with task_id 24
      Slave lucasca-desktop rank 11 executing Tasks.TASK1 with task_id 25
    Master: slave finished is task returning: (True, 10))
    Master: slave finished is task returning: (True, 4))
      Slave lucasca-desktop rank 7 executing Tasks.TASK3 with task_id 22 arg2 999 arg3 something
    Master: slave finished is task returning: (True, 18, 'something'))
    Master: slave finished is task returning: (True, 20, 'something', 'else'))
      Slave lucasca-desktop rank 12 executing Tasks.TASK2 with task_id 26 arg2 52
      Slave lucasca-desktop rank 14 executing Tasks.TASK1 with task_id 28
      Slave lucasca-desktop rank 15 executing Tasks.TASK2 with task_id 29 arg2 58
      Slave lucasca-desktop rank 13 executing Tasks.TASK2 with task_id 27 arg2 54
      Slave lucasca-desktop rank 6 executing Tasks.TASK2 with task_id 21 arg2 42
      Slave lucasca-desktop rank 9 executing Tasks.TASK2 with task_id 23 arg2 46
      Slave lucasca-desktop rank 2 executing Tasks.TASK1 with task_id 31
      Slave lucasca-desktop rank 5 executing Tasks.TASK2 with task_id 33 arg2 66




Example 5
---------

In `Example 5 <https://github.com/luca-s/mpi-master-slave/blob/master/examples/example5.py>`__ we still have that slaves handle multiple type of tasks but we also want to **limit the number of slaves reserved to one or more tasks**. This comes handy when, for example, one or more tasks deal with resources such as database conncetions, network services and so on, and you have to limit the number of concurrent accesses to those resources. 
In this example the Slave code is the same as the previous one but now each task has its own Master instead of letting a single Master handling all the tasks.

.. code:: python

    class MyMaster(Master):
        """
        This Master class handles a specific task
        """

        def __init__(self, task, slaves = None):
            super(MyMaster, self).__init__(slaves)
            self.task = task

        def run(self, slave, data):
            args = (self.task, data)
            super(MyMaster, self).run(slave, args)

        def get_data(self, completed_slave):
            task, data = super(MyMaster, self).get_data(completed_slave)
            return data


At this point one could create each Master with a specific number of slaves and a WorkQueue for each Master. Unfortunately this would produce bad performance as one or more Masters might not have tasks to do at certain times of the execution and their slaves would be idle while other Masters might have plenty of work to do.

What we want to achieve is to let Masters lend/borrow slaves with each others when they are idle so that they make the most out of their slaves. To do that we make use of the MultiWorkQueue class that handles multiple Masters and where each Master can have an optional limit on the number of slaves. MultiWorkQueue moves slaves between Masters when some of them are idles and gives slaves back when the Masters have work again.

.. code:: python

    from mpi.multi_work_queue import MultiWorkQueue

    class MyApp(object):

        def __init__(self, slaves,  task1_num_slave=None, task2_num_slave=None, task3_num_slave=None):
            """
            Each task/master can be limited on the number of slaves by the init
            arguments. Leave them None if you don't want to limit a specific Master
            """
            #
            # create a Master for each task
            #
            self.master1 = MyMaster(task=Tasks.TASK1)
            self.master2 = MyMaster(task=Tasks.TASK2)
            self.master3 = MyMaster(task=Tasks.TASK3)

            #
            # MultiWorkQueue is a convenient class that run multiple work queues
            # Each task needs a Tuple  with (someID, Master, None or max slaves)
            #
            masters_details = [(Tasks.TASK1, self.master1, task1_num_slave),
                               (Tasks.TASK2, self.master2, task2_num_slave),
                               (Tasks.TASK3, self.master3, task3_num_slave) ]
            self.work_queue = MultiWorkQueue(slaves, masters_details)


        def terminate_slaves(self):
            """
            Call this to make all slaves exit their run loop
            """
            self.master1.terminate_slaves()
            self.master2.terminate_slaves()
            self.master3.terminate_slaves()

        def __add_next_task(self, i, task=None):
            """
            Create random tasks 1-3 and add it to the right work queue
            """
            if task is None:
                task = random.randint(1,3)

            if task == 1:
                args = i
                self.work_queue.add_work(Tasks.TASK1, args)
            elif task == 2:
                args = (i, i*2)
                self.work_queue.add_work(Tasks.TASK2, args)
            elif task == 3:
                args = (i, 999, 'something')
                self.work_queue.add_work(Tasks.TASK3, args)

        def run(self, tasks=100):
            """
            This is the core of my application, keep starting slaves
            as long as there is work to do
            """

            #
            # let's prepare our work queue. This can be built at initialization time
            # but it can also be added later as more work become available
            #
            for i in range(tasks):
                self.__add_next_task(i)
           
            #
            # Keeep starting slaves as long as there is work to do
            #
            while not self.work_queue.done():

                #
                # give more work to do to each idle slave (if any)
                #
                self.work_queue.do_work()

                #
                # reclaim returned data from completed slaves
                #
                for data in self.work_queue.get_completed_work(Tasks.TASK1):
                    done, arg1 = data
                    if done:
                        print('Master: slave finished his task returning: %s)' % str(data))

                for data in self.work_queue.get_completed_work(Tasks.TASK2):
                    done, arg1, arg2, arg3 = data
                    if done:
                        print('Master: slave finished his task returning: %s)' % str(data))

                for data in self.work_queue.get_completed_work(Tasks.TASK3):
                    done, arg1, arg2 = data
                    if done:
                        print('Master: slave finished his task returning: %s)' % str(data))

                # sleep some time
                time.sleep(0.3)


For example, you can test the application like this:

.. code:: python

    app = MyApp(slaves=range(1, size), task1_num_slave=2, task2_num_slave=None, task3_num_slave=1)

::

    mpiexec -n 9 xterm -e "python example5.py ; bash"


You can see from the output the number of slaves for task1 is 2, task3 is 1 and task2 takes all the remaining slaves:


.. image:: https://github.com/luca-s/mpi-master-slave/raw/master/examples/example5.png




