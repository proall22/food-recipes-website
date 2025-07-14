<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2">Create New Recipe</h1>
      <p class="text-gray-600">Share your delicious recipe with the RecipeHub community</p>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-8">
      <!-- Basic Information -->
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6">Basic Information</h2>
        
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div class="lg:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Recipe Title *
            </label>
            <Field
              name="title"
              type="text"
              v-model="form.title"
              :class="[
                'w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                errors.title ? 'border-red-500' : 'border-gray-300'
              ]"
              placeholder="Enter a catchy title for your recipe"
            />
            <ErrorMessage name="title" class="text-red-500 text-sm mt-1" />
          </div>

          <div class="lg:col-span-2">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Description *
            </label>
            <textarea
              v-model="form.description"
              rows="4"
              :class="[
                'w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors resize-none',
                errors.description ? 'border-red-500' : 'border-gray-300'
              ]"
              placeholder="Describe your recipe, what makes it special, and any tips for success"
            ></textarea>
            <ErrorMessage name="description" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Category *
            </label>
            <select
              v-model="form.categoryId"
              :class="[
                'w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                errors.categoryId ? 'border-red-500' : 'border-gray-300'
              ]"
            >
              <option value="">Select a category</option>
              <option v-for="category in categories" :key="category.id" :value="category.id">
                {{ category.name }}
              </option>
            </select>
            <ErrorMessage name="categoryId" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Difficulty Level *
            </label>
            <select
              v-model="form.difficulty"
              :class="[
                'w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                errors.difficulty ? 'border-red-500' : 'border-gray-300'
              ]"
            >
              <option value="">Select difficulty</option>
              <option value="Easy">Easy</option>
              <option value="Medium">Medium</option>
              <option value="Hard">Hard</option>
            </select>
            <ErrorMessage name="difficulty" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Preparation Time (minutes) *
            </label>
            <div class="relative">
              <Clock class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <input
                type="number"
                v-model="form.prepTime"
                min="1"
                :class="[
                  'w-full pl-10 pr-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.prepTime ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="30"
              />
            </div>
            <ErrorMessage name="prepTime" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Servings *
            </label>
            <div class="relative">
              <Users class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <input
                type="number"
                v-model="form.servings"
                min="1"
                :class="[
                  'w-full pl-10 pr-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.servings ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="4"
              />
            </div>
            <ErrorMessage name="servings" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Price (Optional)
            </label>
            <div class="relative">
              <DollarSign class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <input
                type="number"
                step="0.01"
                v-model="form.price"
                min="0"
                class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors"
                placeholder="4.99"
              />
            </div>
            <p class="text-sm text-gray-500 mt-1">Set a price if you want to sell this recipe</p>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Cuisine Type
            </label>
            <select
              v-model="form.cuisineType"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors"
            >
              <option value="">Select cuisine type</option>
              <option value="Italian">Italian</option>
              <option value="Mexican">Mexican</option>
              <option value="Asian">Asian</option>
              <option value="Mediterranean">Mediterranean</option>
              <option value="American">American</option>
              <option value="French">French</option>
              <option value="Indian">Indian</option>
              <option value="Other">Other</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Recipe Images -->
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6">Recipe Images</h2>
        
        <div class="space-y-4">
          <!-- Featured Image -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Featured Image * (This will be used as thumbnail)
            </label>
            <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center hover:border-orange-500 transition-colors">
              <div v-if="featuredImagePreview" class="relative">
                <img :src="featuredImagePreview" class="max-h-64 mx-auto rounded-lg" />
                <button
                  type="button"
                  @click="removeFeaturedImage"
                  class="absolute top-2 right-2 bg-red-500 text-white rounded-full p-1 hover:bg-red-600"
                >
                  <X class="h-4 w-4" />
                </button>
              </div>
              <div v-else>
                <ImageIcon class="mx-auto h-12 w-12 text-gray-400" />
                <div class="mt-4">
                  <button
                    type="button"
                    @click="$refs.featuredImageInput.click()"
                    class="bg-orange-500 text-white px-4 py-2 rounded-lg hover:bg-orange-600 transition-colors"
                  >
                    Upload Featured Image
                  </button>
                </div>
              </div>
              <input
                ref="featuredImageInput"
                type="file"
                accept="image/*"
                @change="handleFeaturedImageChange"
                class="hidden"
              />
            </div>
          </div>

          <!-- Additional Images -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Additional Images (Optional)
            </label>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div v-for="(image, index) in additionalImagePreviews" :key="index" 
                class="relative aspect-square">
                <img :src="image" class="w-full h-full object-cover rounded-lg" />
                <button
                  type="button"
                  @click="removeAdditionalImage(index)"
                  class="absolute top-1 right-1 bg-red-500 text-white rounded-full p-1 hover:bg-red-600"
                >
                  <X class="h-3 w-3" />
                </button>
              </div>
              
              <div v-if="additionalImagePreviews.length < 8" 
                class="aspect-square border-2 border-dashed border-gray-300 rounded-lg flex items-center justify-center hover:border-orange-500 transition-colors cursor-pointer"
                @click="$refs.additionalImagesInput.click()">
                <Plus class="h-8 w-8 text-gray-400" />
              </div>
            </div>
            <input
              ref="additionalImagesInput"
              type="file"
              accept="image/*"
              multiple
              @change="handleAdditionalImagesChange"
              class="hidden"
            />
          </div>
        </div>
      </div>

      <!-- Ingredients -->
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-xl font-semibold text-gray-900">Ingredients</h2>
          <button
            type="button"
            @click="addIngredient"
            class="bg-orange-500 text-white px-4 py-2 rounded-lg hover:bg-orange-600 transition-colors flex items-center"
          >
            <Plus class="h-4 w-4 mr-2" />
            Add Ingredient
          </button>
        </div>

        <div class="space-y-4">
          <div v-for="(ingredient, index) in form.ingredients" :key="ingredient.id" 
            class="grid grid-cols-12 gap-4 items-end">
            <div class="col-span-3">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Amount
              </label>
              <input
                type="text"
                v-model="ingredient.amount"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                placeholder="1"
              />
            </div>
            
            <div class="col-span-3">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Unit
              </label>
              <select
                v-model="ingredient.unit"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              >
                <option value="">Select unit</option>
                <option value="cup">Cup</option>
                <option value="tbsp">Tablespoon</option>
                <option value="tsp">Teaspoon</option>
                <option value="lb">Pound</option>
                <option value="oz">Ounce</option>
                <option value="g">Gram</option>
                <option value="kg">Kilogram</option>
                <option value="ml">Milliliter</option>
                <option value="l">Liter</option>
                <option value="piece">Piece</option>
                <option value="clove">Clove</option>
                <option value="slice">Slice</option>
              </select>
            </div>
            
            <div class="col-span-5">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Ingredient Name
              </label>
              <input
                type="text"
                v-model="ingredient.name"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                placeholder="e.g., All-purpose flour"
              />
            </div>
            
            <div class="col-span-1">
              <button
                type="button"
                @click="removeIngredient(index)"
                class="w-full p-2 text-red-500 hover:bg-red-50 rounded-lg transition-colors"
              >
                <Trash2 class="h-4 w-4" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Cooking Steps -->
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-xl font-semibold text-gray-900">Cooking Instructions</h2>
          <button
            type="button"
            @click="addStep"
            class="bg-orange-500 text-white px-4 py-2 rounded-lg hover:bg-orange-600 transition-colors flex items-center"
          >
            <Plus class="h-4 w-4 mr-2" />
            Add Step
          </button>
        </div>

        <div class="space-y-6">
          <div v-for="(step, index) in form.steps" :key="step.id" 
            class="border border-gray-200 rounded-lg p-4">
            <div class="flex items-center justify-between mb-4">
              <h3 class="font-medium text-gray-900">Step {{ index + 1 }}</h3>
              <button
                type="button"
                @click="removeStep(index)"
                class="text-red-500 hover:bg-red-50 p-1 rounded transition-colors"
              >
                <Trash2 class="h-4 w-4" />
              </button>
            </div>

            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Instruction
                </label>
                <textarea
                  v-model="step.instruction"
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent resize-none"
                  placeholder="Describe this step in detail..."
                ></textarea>
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    Timer (minutes, optional)
                  </label>
                  <input
                    type="number"
                    v-model="step.timer"
                    min="0"
                    class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                    placeholder="5"
                  />
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">
                    Step Image (optional)
                  </label>
                  <div class="flex items-center space-x-2">
                    <input
                      type="file"
                      accept="image/*"
                      @change="(e) => handleStepImageChange(e, index)"
                      class="hidden"
                      :ref="`stepImageInput${index}`"
                    />
                    <button
                      type="button"
                      @click="$refs[`stepImageInput${index}`][0].click()"
                      class="px-3 py-2 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors text-sm"
                    >
                      Choose Image
                    </button>
                    <span v-if="step.imagePreview" class="text-sm text-green-600">Image selected</span>
                  </div>
                </div>
              </div>

              <div v-if="step.imagePreview" class="mt-2">
                <img :src="step.imagePreview" class="max-h-32 rounded-lg" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Nutrition Information -->
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6">Nutrition Information (Optional)</h2>
        
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Calories
            </label>
            <input
              type="number"
              v-model="form.nutrition.calories"
              min="0"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              placeholder="350"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Protein (g)
            </label>
            <input
              type="number"
              v-model="form.nutrition.protein"
              min="0"
              step="0.1"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              placeholder="25"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Carbs (g)
            </label>
            <input
              type="number"
              v-model="form.nutrition.carbs"
              min="0"
              step="0.1"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              placeholder="45"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Fat (g)
            </label>
            <input
              type="number"
              v-model="form.nutrition.fat"
              min="0"
              step="0.1"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              placeholder="15"
            />
          </div>
        </div>
      </div>

      <!-- Submit Buttons -->
      <div class="flex justify-end space-x-4">
        <button
          type="button"
          @click="saveDraft"
          class="px-6 py-3 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors"
        >
          Save as Draft
        </button>
        <button
          type="submit"
          :disabled="isLoading"
          class="px-6 py-3 bg-orange-500 text-white rounded-lg hover:bg-orange-600 transition-colors disabled:opacity-50 flex items-center"
        >
          <span v-if="isLoading" class="mr-2">
            <Loader2 class="h-4 w-4 animate-spin" />
          </span>
          {{ isLoading ? 'Publishing...' : 'Publish Recipe' }}
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useForm, Field, ErrorMessage } from 'vee-validate'
import * as yup from 'yup'
import { 
  Clock, Users, DollarSign, Plus, X, Trash2, 
  ImageIcon, Loader2 
} from 'lucide-vue-next'
import { navigateTo } from '#app'
import { useAuthStore } from '~/stores/auth'
import { useNuxtApp } from '#app'
import { gql } from 'graphql-tag'
import { useHead } from '#head'

// Check authentication
const authStore = useAuthStore()
const { $apollo } = useNuxtApp()

const ingredients_name_key = 'ingredients_name_key' // Declare the variable here

const isLoading = ref(false)
const categories = ref([])
const featuredImagePreview = ref('')
const additionalImagePreviews = ref([])

const form = reactive({
  title: '',
  description: '',
  categoryId: '',
  difficulty: '',
  prepTime: null,
  servings: null,
  price: null,
  cuisineType: '',
  featuredImage: null,
  additionalImages: [],
  ingredients: [
    { id: Date.now(), amount: '', unit: '', name: '' }
  ],
  steps: [
    { id: Date.now(), instruction: '', timer: null, image: null, imagePreview: '' }
  ],
  nutrition: {
    calories: null,
    protein: null,
    carbs: null,
    fat: null
  }
})

// Validation schema
const schema = yup.object({
  title: yup.string().required('Recipe title is required'),
  description: yup.string().required('Recipe description is required'),
  categoryId: yup.string().required('Please select a category'),
  difficulty: yup.string().required('Please select difficulty level'),
  prepTime: yup.number().positive('Prep time must be positive').required('Prep time is required'),
  servings: yup.number().positive('Servings must be positive').required('Servings is required')
})

const { errors, validate } = useForm({
  validationSchema: schema
})

// Load categories
onMounted(async () => {
  try {
    const GET_CATEGORIES_QUERY = gql`
      query GetCategories {
        categories {
          id
          name
          slug
        }
      }
    `
    
    const { data } = await $apollo.query({
      query: GET_CATEGORIES_QUERY
    })
    
    categories.value = data.categories
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
})

// Image handling
const handleFeaturedImageChange = (event) => {
  const file = event.target.files[0]
  if (file) {
    form.featuredImage = file
    const reader = new FileReader()
    reader.onload = (e) => {
      featuredImagePreview.value = e.target.result
    }
    reader.readAsDataURL(file)
  }
}

const removeFeaturedImage = () => {
  form.featuredImage = null
  featuredImagePreview.value = ''
}

const handleAdditionalImagesChange = (event) => {
  const files = Array.from(event.target.files)
  files.forEach(file => {
    if (form.additionalImages.length < 8) {
      form.additionalImages.push(file)
      const reader = new FileReader()
      reader.onload = (e) => {
        additionalImagePreviews.value.push(e.target.result)
      }
      reader.readAsDataURL(file)
    }
  })
}

const removeAdditionalImage = (index) => {
  form.additionalImages.splice(index, 1)
  additionalImagePreviews.value.splice(index, 1)
}

const handleStepImageChange = (event, stepIndex) => {
  const file = event.target.files[0]
  if (file) {
    form.steps[stepIndex].image = file
    const reader = new FileReader()
    reader.onload = (e) => {
      form.steps[stepIndex].imagePreview = e.target.result
    }
    reader.readAsDataURL(file)
  }
}

// Ingredient management
const addIngredient = () => {
  form.ingredients.push({
    id: Date.now(),
    amount: '',
    unit: '',
    name: ''
  })
}

const removeIngredient = (index) => {
  if (form.ingredients.length > 1) {
    form.ingredients.splice(index, 1)
  }
}

// Step management
const addStep = () => {
  form.steps.push({
    id: Date.now(),
    instruction: '',
    timer: null,
    image: null,
    imagePreview: ''
  })
}

const removeStep = (index) => {
  if (form.steps.length > 1) {
    form.steps.splice(index, 1)
  }
}

// Form submission
const handleSubmit = async () => {
  const { valid } = await validate()
  if (!valid) return

  // Validate required fields
  if (!form.featuredImage) {
    alert('Please upload a featured image')
    return
  }

  if (form.ingredients.some(ing => !ing.name || !ing.amount)) {
    alert('Please fill in all ingredient fields')
    return
  }

  if (form.steps.some(step => !step.instruction)) {
    alert('Please fill in all step instructions')
    return
  }

  isLoading.value = true

  try {
    // Upload featured image
    const featuredImageUrl = await authStore.uploadFile(form.featuredImage, 'recipe')

    // Upload additional images
    const additionalImageUrls = await Promise.all(
      form.additionalImages.map(file => authStore.uploadFile(file, 'recipe'))
    )

    // Upload step images
    const stepsWithImages = await Promise.all(
      form.steps.map(async (step) => {
        if (step.image) {
          const imageUrl = await authStore.uploadFile(step.image, 'recipe')
          return { 
            instruction: step.instruction, 
            timer_minutes: step.timer,
            image_url: imageUrl,
            step_number: form.steps.indexOf(step) + 1
          }
        }
        return { 
          instruction: step.instruction, 
          timer_minutes: step.timer,
          step_number: form.steps.indexOf(step) + 1
        }
      })
    )

    // Create recipe using GraphQL mutation
    const CREATE_RECIPE_MUTATION = gql`
      mutation CreateRecipe($recipe: recipes_insert_input!) {
        insert_recipes_one(object: $recipe) {
          id
          slug
          title
        }
      }
    `

    const { data } = await $apollo.mutate({
      mutation: CREATE_RECIPE_MUTATION,
      variables: {
        recipe: {
          title: form.title,
          description: form.description,
          category_id: form.categoryId,
          difficulty: form.difficulty,
          prep_time: form.prepTime,
          cook_time: 0,
          servings: form.servings,
          price: form.price || 0,
          cuisine_type: form.cuisineType,
          featured_image: featuredImageUrl,
          status: 'published',
          recipe_images: {
            data: additionalImageUrls.map((url, index) => ({
              image_url: url,
              sort_order: index
            }))
          },
          recipe_ingredients: {
            data: form.ingredients
              .filter(ing => ing.name && ing.amount)
              .map((ing, index) => ({
                ingredient: {
                  data: { name: ing.name },
                  on_conflict: {
                    constraint: ingredients_name_key,
                    update_columns: []
                  }
                },
                amount: parseFloat(ing.amount),
                unit: ing.unit,
                sort_order: index
              }))
          },
          recipe_steps: {
            data: stepsWithImages
          },
          recipe_nutrition: form.nutrition.calories ? {
            data: {
              calories: form.nutrition.calories,
              protein: form.nutrition.protein,
              carbohydrates: form.nutrition.carbs,
              fat: form.nutrition.fat
            }
          } : undefined
        }
      }
    })

    if (data.insert_recipes_one) {
      await navigateTo(`/recipes/${data.insert_recipes_one.id}`)
    }
  } catch (error) {
    console.error('Failed to create recipe:', error)
    alert('Failed to create recipe. Please try again.')
  } finally {
    isLoading.value = false
  }
}

const saveDraft = async () => {
  // Similar to handleSubmit but with status: 'draft'
  isLoading.value = true
  
  try {
    // Implementation for saving draft
    console.log('Saving draft...')
  } catch (error) {
    console.error('Failed to save draft:', error)
  } finally {
    isLoading.value = false
  }
}

// SEO
useHead({
  title: 'Create Recipe - RecipeHub',
  meta: [
    { name: 'description', content: 'Share your delicious recipe with the RecipeHub community' }
  ]
})
</script>
