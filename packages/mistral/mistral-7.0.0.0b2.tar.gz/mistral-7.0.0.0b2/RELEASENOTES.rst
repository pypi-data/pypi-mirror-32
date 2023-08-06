=======
mistral
=======

.. _mistral_7.0.0.0b2:

7.0.0.0b2
=========

.. _mistral_7.0.0.0b2_New Features:

New Features
------------

.. releasenotes/notes/add-py-mini-racer-javascript-evaluator-9d8f9e0e36504d72.yaml @ 5f89e2e71fb95a35b549106fb3cc1792e80c89b7

- Added new JavaScript evaluator py_mini_racer. The py_mini_racer package
  allows us to get a JavaScript evaluator that doesn't require compilation.
  This is much lighter and easier to get started with.

.. releasenotes/notes/add_yaql_engine_options-200fdcfda04683ca.yaml @ 360fd8bd61dadf4738ea49ff0b15f7eb3ec4108b

- Added several config options that allow to tweak some aspects of the YAQL
  engine behavior.

.. releasenotes/notes/force-stop-executions-00cd67dbbc9b5483.yaml @ af84fa9181ecfce9386387300a564b63f37211f4

- Use of the parameter force to forcefully delete executions. Note using this parameter on unfinished executions might cause a cascade of errors.

.. releasenotes/notes/mistral-vitrage-actions-a205b8ea82b43cab.yaml @ f25fb43177767ef968d8ac6a7849ea3b028f3b5c

- Add Mistral actions for OpenStack Vitrage, the RCA service

.. releasenotes/notes/safe-rerun-in-task-defaults-87a4cbe12558bc6d.yaml @ a29c7d9f7f6cb3c22f3548171cdeb325d7a35d0f

- Added 'safe-rerun' policy to task-defaults section

.. releasenotes/notes/support-qinling-action-99cd323d4df36d48.yaml @ 4fe4198ac8590a3453f191ae8be5ab06a4264cb4

- Add Mistral actions for Openstack Qinling, the function management service.

.. releasenotes/notes/support-zun-action-3263350334d1d34f.yaml @ 11896ba7337fde1c3e806fcad6f5cde0c5050c57

- Add Mistral actions for Openstack Zun, the container service.


.. _mistral_7.0.0.0b2_Known Issues:

Known Issues
------------

.. releasenotes/notes/force-stop-executions-00cd67dbbc9b5483.yaml @ af84fa9181ecfce9386387300a564b63f37211f4

- Deleting unfinished executions might cause a cascade of errors, so the standard behaviour has been changed to delete only safe to delete executions and a new parameter force was added to forceful delete ignoring the state the execution is in.


.. _mistral_7.0.0.0b2_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/add_yaql_engine_options-200fdcfda04683ca.yaml @ 360fd8bd61dadf4738ea49ff0b15f7eb3ec4108b

- Fixed how Mistral initializes a child YAQL context before evaluating
  YAQL expressions. The given data context needs to go through a special
  filter that prepares the data properly, does conversion into internal
  types etc. Also, without this change YAQL engine options are not applied
  properly.

.. releasenotes/notes/fix-jinja-expression-handling-135451645d7a4e6f.yaml @ a70df9e99140413a76564cb7f0f0f8acfcf95b31

- Fixed jinja expression  error handling where invalid expression could prevent action or task status to be correctly updated.

.. releasenotes/notes/using_passive_deletes_in_sqlalchemy-4b3006b3aba55155.yaml @ d3c6bf7767fda544be79af94d3e0b26e1e3b2a3e

- Used "passive_deletes=True" in the configuration of relationships in
  SQLAlchemy models. This improves deletion of graphs of related objects
  stored in DB because dependent objects don't get loaded prior to
  deletion which also reduces the memory requirement on the system.
  More about using this flag can be found at:
  http://docs.sqlalchemy.org/en/latest/orm/collections.html#using-passive-deletes

.. releasenotes/notes/wf_final_context_evaluation_with_batches-6292ab64c131dfcc.yaml @ 5073274dd1f15a59a07da99f3aa387fb6e5d51dc

- Evaluation of final workflow context was very heavy in cases when the workflow had a lot of parallel tasks with large inbound contexts. Merging of those contexts in order to evaluate the workflow output consumed a lot of memory. Now this algorithm is rewritten with batched DB query and Python generators so that GS has a chance to destroy objects that have already been processed. Previously all task executions had to stay in memory until the end of the processing. The result is that now it consumes 3 times less memory on heavy cases.

.. releasenotes/notes/workflow_environment_optimizations-deb8868df3f0dc36.yaml @ cb1cabbfe652bcfdda4a116c290d8474c3f1f4e1

- Mistral was storing, in fact, two copies of a workflow environment, one in workflow parameters (the 'params' field) and another one in a context (the 'context' field). Now it's stored only in workflow parameters. It saves space in DB and increases performance in case of big workflow environments.

.. releasenotes/notes/workflow_environment_optimizations-deb8868df3f0dc36.yaml @ cb1cabbfe652bcfdda4a116c290d8474c3f1f4e1

- Mistral was copying a workflow environment into all of their sub workflows. In case of a big workflow environment and a big number of sub workflows it caused serious problems, used additional space in DB and used a lot of RAM (e.g. when the 'on-success' clause has a lot of tasks where each one of them is a subworkflow). Now it is fixed by evaluating a workflow environment through the root execution reference.


.. _mistral_7.0.0.0b1:

7.0.0.0b1
=========

.. _mistral_7.0.0.0b1_New Features:

New Features
------------

.. releasenotes/notes/add-execution-event-notifications-0f77c1c3eb1d6929.yaml @ 422c89a05aa7867509a6a989a536fc87b55bc792

- Introduce execution events and notification server and plugins for
  publishing these events for consumers. Event notification is defined per
  workflow execution and can be configured to notify on all the events or
  only for specific events.

.. releasenotes/notes/add_action_definition_caching-78d4446d61c6d739.yaml @ 6683e154bb5a840a5aa72c9b0f0abc8655816b02

- Enable caching of action definitions in local memory. Now, instead of
  downloading the definitions from the database every time, mistral engine
  will store them in a local cache. This should reduce the number of
  database requests and improve the whole performance of the system.
  Cache ttl can be configured with ``action_definition_cache_time`` option
  from [engine] group. The default value is 60 seconds.


.. _mistral_7.0.0.0b1_Bug Fixes:

Bug Fixes
---------

.. releasenotes/notes/add-task_execution_id-indexes-16edc58085e47663.yaml @ e7da5b6dc8381b78e572c903f3a6351697c823d8

- Added new indexes on the task_execution_id column of the
  action_executions_v2 and workflow_executions_v2 tables.

.. releasenotes/notes/fix-regression-when-logging-58faa35f02cefb34.yaml @ 15e95d9b3cc7ce8cbe8de22ef80e8230f22107ee

- A regression was introduced that caused an error when logging a specific
  message. The string formatting was broken, which caused the logging to
  fail.

.. releasenotes/notes/fix_pause_command-58294f613488511c.yaml @ 8b30743d0a5e44c7673f90b4509989f8f35a7cb3

- Fixed the logic of the 'pause' command. Before the fix Mistral wouldn't run any commands specified in 'on-success', 'on-error' and 'on-complete' clauses following after the 'pause' command when a workflow was resumed after it. Now it works as expected. If Mistral encounters 'pause' in the list of commands it saves all commands following after it to the special backlog storage and when/if the workflow is later resumed it checks that storage and runs commands from it first.

.. releasenotes/notes/remove_redundant_persistent_data_from_task_context-c5281a5f5ae688f1.yaml @ e42c515a5fcc459644c75b177a2a49faf066c414

- Mistral was storing some internal information in task execution inbound context ('task_executions_v2.in_contex' DB field) to DB. This information was needed only to correctly implement the YAQL function task() without arguments. A fix was made to not store this information in the persistent storage and rather include it into a context view right before evaluating expressions where needed. So it slightly optimizes spaces in DB.

