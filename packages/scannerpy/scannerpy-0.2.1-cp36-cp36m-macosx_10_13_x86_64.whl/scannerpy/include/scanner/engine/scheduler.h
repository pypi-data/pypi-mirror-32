/* Copyright 2018 Carnegie Mellon University
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

#include <grpc/grpc_posix.h>
#include <grpc/support/log.h>
#include <atomic>
#include <thread>

namespace scanner {
namespace internal {

using TaskID = i64;
using StreamID = i64;
using ElementIndex = i64;

class MachineSpecs {
  std::map<DeviceType, std::vector<i32>> compute_resources;
  i32 io_resources;
}

class Scheduler {
 public:
  Scheduler(const MachineSpecs& specs);

  struct SubmitTaskResult {
    TaskID task;
    std::vector<StreamID> output_streams;
  };
  SubmitTaskResult submit_task(
      std::function task,
      const std::vector<TaskID>& dependencies);

  TaskID submit_delete(
      const std::map<StreamID, std::vector<ElementIndex>>& deletions,
      const std::vector<TaskID>& dependencies);

};

}
}
