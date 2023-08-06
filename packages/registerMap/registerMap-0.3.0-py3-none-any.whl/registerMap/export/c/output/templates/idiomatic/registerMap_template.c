/*
 *
 * {{ registerMap.name }}
 *
 {%- if licenseText is not none %}
 {%- for line in licenseText %}
 * {{ line }}
 {%- endfor %}
 {%- endif %}
 *
 */

#include "{{ prefixPath }}/macro/extern.h"
#include "{{ prefixPath }}/memory/memory.h"
#include "{{ prefixPath }}/registerMap.h"


{{ registerMap.name|upper }}_OPEN_EXTERN_C

#ifdef OFF_TARGET_MEMORY

struct MemorySpace_t myRegisterMap_memory = {
  .allocated_memory_span = {{ memory.size }},
};

#else

{% set memoryPointerType = memory.sizeType~' volatile* const' -%}
struct MemorySpace_t myRegisterMap_memory = {
  .allocated_memory_span = {{ memory.size }},
  .base = ( {{ memoryPointerType }} ) {{ memory.baseAddress }},
};

#endif

struct {{ registerMap.name }}_t {{ registerMap.name }} = {
{%- for thisModule in registerMap.modules %}
{%- set pointerType = thisModule.name~'_t volatile* const' %}
  .{{ thisModule.name }} = ( {{ pointerType }} )( {{ registerMap.name }}_memory.base + {{ thisModule.offset }} ),
{% endfor -%}
};

{{ registerMap.name|upper }}_CLOSE_EXTERN_C
