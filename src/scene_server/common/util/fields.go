/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"fmt"
	"strings"

	"configcenter/src/common"
	gutil "configcenter/src/common/util"
	sourceAPI "configcenter/src/source_controller/api/object"
)

// ConvByPropertytype convert str to property type
func ConvByPropertytype(fields *sourceAPI.ObjAttDes, val string) (interface{}, error) {
	switch fields.PropertyType {
	case common.FieldTypeInt:
		return gutil.GetIntByInterface(val)
	case common.FieldTypeBool:
		val := strings.ToLower(val)
		switch val {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			return false, fmt.Errorf("%s not bool", val)
		}
	}
	return val, nil
}
