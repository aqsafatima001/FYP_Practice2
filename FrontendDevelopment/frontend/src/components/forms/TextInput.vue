<template>
    <div class="mb-3">
        <label :for="name" class="form-label">{{ label }}</label>
        <input 
            :type="type" 
            :name="name" 
            :placeholder="placeholder" 
            :required="required" 
            :min="min" 
            :max="max"
            :value="value" 
            :autocomplete="name + '-new'" 
            class="form-control"
            @input="handleInput" 
        />
    </div>
</template>
  
<script>
export default {
    name: 'TextInput',
    props: {
        name: String,
        type: String,
        label: String,
        placeholder: String,
        required: Boolean,
        min: String,
        max: String,
        value: String,
    },
    methods: {
        handleInput(event) {
            console.log('Input event emitted:', event.target.value);
            this.$emit('input', event.target.value);
        },
    },
};
</script>
  







<!-- <template>
    <div class="mb-3">
        <label :for="name" class="form-label">{{label}}</label>
        <input
        :type="type"
        :name="name"
        :placeholder="placeholder"
        :required="required"
        :min="min"
        :max="max"
        :value="value"
        :autocomplete="name + '-new'"
        class="form-control"
        @input="$emit('update:value', $event.target.value)"
        @input="$emit('input', $event.target.value)" 
      />
    </div>
</template>

<script>
export default {
    name: 'TextInput',
    props: {
        name: String,
        type: String,
        label: String,
        placeholder: String,
        required: String,
        min: String,
        max: String,
        value: String,
    },
}
</script> -->