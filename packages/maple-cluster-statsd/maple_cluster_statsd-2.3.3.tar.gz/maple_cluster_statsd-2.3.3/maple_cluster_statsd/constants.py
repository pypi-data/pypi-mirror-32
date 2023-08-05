# -*- coding: utf-8 -*-

# gateway
# 直接gauge的
GATEWAY_GAUGE_STAT_LIST = [
    'CLIENTS',
    'AUTHED_CLIENTS',

    'TRIGGERS',

    'WORKERS',
    'IDLE_WORKERS',

    'STORES',
    'CONNECTED_STORES',

    'CLIENT_PENDING_SEND_BUF',

    'INNER_PENDING_TASKS',
    'INNER_PENDING_SEND_BUF',

    'STORE_PENDING_TASKS',
]

# 间隔一段时间，然后调用incr的
GATEWAY_INCR_STAT_LIST = [
    'CLIENT_REQ',
    'CLIENT_CREATED',
    'CLIENT_CLOSED',
    'CLIENT_TIMEOUT',
    'CLIENT_CONN_OVERFLOW',
    'CLIENT_SEND_BUF_OVERFLOW',
    'CLIENT_RECV_BUF_OVERFLOW',
    'CLIENT_CREATE_RATE_LIMITED',
    'CLIENT_UNAUTHED_REQ_RATE_LIMITED',
    'CLIENT_AUTHED_REQ_RATE_LIMITED',

    'INNER_REQ',
    'INNER_CREATED',
    'INNER_CLOSED',
    'INNER_TIMEOUT',
    'INNER_CONN_OVERFLOW',
    'INNER_TASK_OVERFLOW',
    'INNER_SEND_BUF_OVERFLOW',
    'INNER_RECV_BUF_OVERFLOW',
    'INNER_TASK_ALLOC_RATE_LIMITED',

    'STORE_REQ',
    'STORE_RSP',
    'STORE_RSP_FAIL',
    'STORE_CONNECT',
    'STORE_CONNECTED',
    'STORE_CLOSED',
    'STORE_TIMEOUT',
    'STORE_TASK_OVERFLOW',

    'CMD_WORKER_ASK_FOR_TASK',
    'CMD_WRITE_TO_CLIENT',
    'CMD_WRITE_TO_USERS',
    'CMD_CLOSE_CLIENT',
    'CMD_CLOSE_USERS',
    'CMD_LOGIN_CLIENT',
    'CMD_LOGOUT_CLIENT',
    'CMD_WRITE_TO_WORKER',
    'CMD_CLEAR_CLIENT_TASKS',

    'TASK_TIME_1_MS',
    'TASK_TIME_3_MS',
    'TASK_TIME_5_MS',
    'TASK_TIME_10_MS',
    'TASK_TIME_30_MS',
    'TASK_TIME_50_MS',
    'TASK_TIME_100_MS',
    'TASK_TIME_300_MS',
    'TASK_TIME_500_MS',
    'TASK_TIME_1_S',
    'TASK_TIME_3_S',
    'TASK_TIME_5_S',
    'TASK_TIME_10_S',
    'TASK_TIME_MORE',

    'STORE_TIME_1_MS',
    'STORE_TIME_3_MS',
    'STORE_TIME_5_MS',
    'STORE_TIME_10_MS',
    'STORE_TIME_30_MS',
    'STORE_TIME_50_MS',
    'STORE_TIME_100_MS',
    'STORE_TIME_300_MS',
    'STORE_TIME_500_MS',
    'STORE_TIME_1_S',
    'STORE_TIME_3_S',
    'STORE_TIME_5_S',
    'STORE_TIME_10_S',
    'STORE_TIME_MORE',
]

# forwarder
FORWARDER_GAUGE_STAT_LIST = [
    'TRIGGERS',

    'STORES',
    'CONNECTED_STORES',

    'NODES',
    'CONNECTED_NODES',

    'STORE_PENDING_TASKS',

    'NODE_PENDING_SEND_BUF',
]

FORWARDER_INCR_STAT_LIST = [
    'TRIGGER_REQ',
    'TRIGGER_CREATED',
    'TRIGGER_CLOSED',
    'TRIGGER_TIMEOUT',
    'TRIGGER_CONN_OVERFLOW',

    'STORE_REQ',
    'STORE_RSP',
    'STORE_RSP_FAIL',
    'STORE_CONNECT',
    'STORE_CONNECTED',
    'STORE_CLOSED',
    'STORE_TIMEOUT',
    'STORE_TASK_OVERFLOW',

    'NODE_REQ',
    'NODE_CONNECT',
    'NODE_CONNECTED',
    'NODE_CLOSED',
    'NODE_TIMEOUT',
    'NODE_SEND_BUF_OVERFLOW',

    'CMD_WRITE_TO_CLIENT',
    'CMD_WRITE_TO_USERS',
    'CMD_CLOSE_CLIENT',
    'CMD_CLOSE_USERS',
    'CMD_LOGIN_CLIENT',
    'CMD_LOGOUT_CLIENT',
    'CMD_WRITE_TO_WORKER',
    'CMD_CLEAR_CLIENT_TASKS',

    'STORE_TIME_1_MS',
    'STORE_TIME_3_MS',
    'STORE_TIME_5_MS',
    'STORE_TIME_10_MS',
    'STORE_TIME_30_MS',
    'STORE_TIME_50_MS',
    'STORE_TIME_100_MS',
    'STORE_TIME_300_MS',
    'STORE_TIME_500_MS',
    'STORE_TIME_1_S',
    'STORE_TIME_3_S',
    'STORE_TIME_5_S',
    'STORE_TIME_10_S',
    'STORE_TIME_MORE',
]
