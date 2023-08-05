# -*- coding: utf-8 -*-
# Copyright (c) 2016, 2017, 2018 Sqreen. All rights reserved.
# Please refer to our terms for more information:
#
#     https://www.sqreen.io/terms.html
#
""" Record request context
"""

from logging import getLogger

from ..frameworks.wsgi import WSGIRequest
from ..rules import RuleCallback
from ..runtime_storage import runtime

LOGGER = getLogger(__name__)


class RecordRequestContext(RuleCallback):

    def pre(self, *args, **kwargs):
        self._store_request(WSGIRequest(args[-2]))

    def _store_request(self, request):
        runtime.store_request(request)
        if self.runner is None or not self.runner.whitelisted_metric:
            return
        whitelist_match = runtime.get_whitelist_match(self.runner.settings)
        if whitelist_match is not None:
            self.record_observation('whitelisted', whitelist_match, 1)

    def _clear_request(self):
        if self.runner is None:
            runtime.clear_request(None, None)
        else:
            runtime.clear_request(self.runner.queue,
                                  self.runner.observation_queue)

    def post(self, *args, **kwargs):
        self._clear_request()

    def failing(self, *args, **kwargs):
        self._clear_request()

    @property
    def whitelisted(self):
        return False
