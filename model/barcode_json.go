package model

type BarcodeJson struct {
	Code    string `json:"code"`
	Product struct {
		Id                       string        `json:"_id"`
		Keywords                 []string      `json:"_keywords"`
		AddedCountriesTags       []interface{} `json:"added_countries_tags"`
		AdditivesOldTags         []interface{} `json:"additives_old_tags"`
		AdditivesOriginalTags    []interface{} `json:"additives_original_tags"`
		AdditivesTags            []interface{} `json:"additives_tags"`
		Allergens                string        `json:"allergens"`
		AllergensFromIngredients string        `json:"allergens_from_ingredients"`
		AllergensFromUser        string        `json:"allergens_from_user"`
		AllergensHierarchy       []interface{} `json:"allergens_hierarchy"`
		AllergensTags            []interface{} `json:"allergens_tags"`
		AminoAcidsTags           []interface{} `json:"amino_acids_tags"`
		CategoriesProperties     struct {
		} `json:"categories_properties"`
		CategoriesPropertiesTags []string `json:"categories_properties_tags"`
		CategoryProperties       struct {
		} `json:"category_properties"`
		CheckersTags            []interface{} `json:"checkers_tags"`
		CiqualFoodNameTags      []string      `json:"ciqual_food_name_tags"`
		Code                    string        `json:"code"`
		CodesTags               []string      `json:"codes_tags"`
		Complete                int           `json:"complete"`
		Completeness            float64       `json:"completeness"`
		CorrectorsTags          []string      `json:"correctors_tags"`
		Countries               string        `json:"countries"`
		CountriesHierarchy      []string      `json:"countries_hierarchy"`
		CountriesTags           []string      `json:"countries_tags"`
		CreatedT                int           `json:"created_t"`
		Creator                 string        `json:"creator"`
		DataQualityBugsTags     []interface{} `json:"data_quality_bugs_tags"`
		DataQualityErrorsTags   []interface{} `json:"data_quality_errors_tags"`
		DataQualityInfoTags     []string      `json:"data_quality_info_tags"`
		DataQualityTags         []string      `json:"data_quality_tags"`
		DataQualityWarningsTags []string      `json:"data_quality_warnings_tags"`
		DataSources             string        `json:"data_sources"`
		DataSourcesTags         []string      `json:"data_sources_tags"`
		EcoscoreData            struct {
			Adjustments struct {
				OriginsOfIngredients struct {
					AggregatedOrigins []struct {
						EpiScore            string      `json:"epi_score"`
						Origin              string      `json:"origin"`
						Percent             int         `json:"percent"`
						TransportationScore interface{} `json:"transportation_score"`
					} `json:"aggregated_origins"`
					EpiScore                int      `json:"epi_score"`
					EpiValue                int      `json:"epi_value"`
					OriginsFromCategories   []string `json:"origins_from_categories"`
					OriginsFromOriginsField []string `json:"origins_from_origins_field"`
					TransportationScore     int      `json:"transportation_score"`
					TransportationScores    struct {
						Ad    int `json:"ad"`
						Al    int `json:"al"`
						At    int `json:"at"`
						Ax    int `json:"ax"`
						Ba    int `json:"ba"`
						Be    int `json:"be"`
						Bg    int `json:"bg"`
						Ch    int `json:"ch"`
						Cy    int `json:"cy"`
						Cz    int `json:"cz"`
						De    int `json:"de"`
						Dk    int `json:"dk"`
						Dz    int `json:"dz"`
						Ee    int `json:"ee"`
						Eg    int `json:"eg"`
						Es    int `json:"es"`
						Fi    int `json:"fi"`
						Fo    int `json:"fo"`
						Fr    int `json:"fr"`
						Gg    int `json:"gg"`
						Gi    int `json:"gi"`
						Gr    int `json:"gr"`
						Hr    int `json:"hr"`
						Hu    int `json:"hu"`
						Ie    int `json:"ie"`
						Il    int `json:"il"`
						Im    int `json:"im"`
						Is    int `json:"is"`
						It    int `json:"it"`
						Je    int `json:"je"`
						Lb    int `json:"lb"`
						Li    int `json:"li"`
						Lt    int `json:"lt"`
						Lu    int `json:"lu"`
						Lv    int `json:"lv"`
						Ly    int `json:"ly"`
						Ma    int `json:"ma"`
						Mc    int `json:"mc"`
						Md    int `json:"md"`
						Me    int `json:"me"`
						Mk    int `json:"mk"`
						Mt    int `json:"mt"`
						Nl    int `json:"nl"`
						No    int `json:"no"`
						Pl    int `json:"pl"`
						Ps    int `json:"ps"`
						Pt    int `json:"pt"`
						Ro    int `json:"ro"`
						Rs    int `json:"rs"`
						Se    int `json:"se"`
						Si    int `json:"si"`
						Sj    int `json:"sj"`
						Sk    int `json:"sk"`
						Sm    int `json:"sm"`
						Sy    int `json:"sy"`
						Tn    int `json:"tn"`
						Tr    int `json:"tr"`
						Ua    int `json:"ua"`
						Uk    int `json:"uk"`
						Us    int `json:"us"`
						Va    int `json:"va"`
						World int `json:"world"`
						Xk    int `json:"xk"`
					} `json:"transportation_scores"`
					TransportationValue  int `json:"transportation_value"`
					TransportationValues struct {
						Ad    int `json:"ad"`
						Al    int `json:"al"`
						At    int `json:"at"`
						Ax    int `json:"ax"`
						Ba    int `json:"ba"`
						Be    int `json:"be"`
						Bg    int `json:"bg"`
						Ch    int `json:"ch"`
						Cy    int `json:"cy"`
						Cz    int `json:"cz"`
						De    int `json:"de"`
						Dk    int `json:"dk"`
						Dz    int `json:"dz"`
						Ee    int `json:"ee"`
						Eg    int `json:"eg"`
						Es    int `json:"es"`
						Fi    int `json:"fi"`
						Fo    int `json:"fo"`
						Fr    int `json:"fr"`
						Gg    int `json:"gg"`
						Gi    int `json:"gi"`
						Gr    int `json:"gr"`
						Hr    int `json:"hr"`
						Hu    int `json:"hu"`
						Ie    int `json:"ie"`
						Il    int `json:"il"`
						Im    int `json:"im"`
						Is    int `json:"is"`
						It    int `json:"it"`
						Je    int `json:"je"`
						Lb    int `json:"lb"`
						Li    int `json:"li"`
						Lt    int `json:"lt"`
						Lu    int `json:"lu"`
						Lv    int `json:"lv"`
						Ly    int `json:"ly"`
						Ma    int `json:"ma"`
						Mc    int `json:"mc"`
						Md    int `json:"md"`
						Me    int `json:"me"`
						Mk    int `json:"mk"`
						Mt    int `json:"mt"`
						Nl    int `json:"nl"`
						No    int `json:"no"`
						Pl    int `json:"pl"`
						Ps    int `json:"ps"`
						Pt    int `json:"pt"`
						Ro    int `json:"ro"`
						Rs    int `json:"rs"`
						Se    int `json:"se"`
						Si    int `json:"si"`
						Sj    int `json:"sj"`
						Sk    int `json:"sk"`
						Sm    int `json:"sm"`
						Sy    int `json:"sy"`
						Tn    int `json:"tn"`
						Tr    int `json:"tr"`
						Ua    int `json:"ua"`
						Uk    int `json:"uk"`
						Us    int `json:"us"`
						Va    int `json:"va"`
						World int `json:"world"`
						Xk    int `json:"xk"`
					} `json:"transportation_values"`
					Value  int `json:"value"`
					Values struct {
						Ad    int `json:"ad"`
						Al    int `json:"al"`
						At    int `json:"at"`
						Ax    int `json:"ax"`
						Ba    int `json:"ba"`
						Be    int `json:"be"`
						Bg    int `json:"bg"`
						Ch    int `json:"ch"`
						Cy    int `json:"cy"`
						Cz    int `json:"cz"`
						De    int `json:"de"`
						Dk    int `json:"dk"`
						Dz    int `json:"dz"`
						Ee    int `json:"ee"`
						Eg    int `json:"eg"`
						Es    int `json:"es"`
						Fi    int `json:"fi"`
						Fo    int `json:"fo"`
						Fr    int `json:"fr"`
						Gg    int `json:"gg"`
						Gi    int `json:"gi"`
						Gr    int `json:"gr"`
						Hr    int `json:"hr"`
						Hu    int `json:"hu"`
						Ie    int `json:"ie"`
						Il    int `json:"il"`
						Im    int `json:"im"`
						Is    int `json:"is"`
						It    int `json:"it"`
						Je    int `json:"je"`
						Lb    int `json:"lb"`
						Li    int `json:"li"`
						Lt    int `json:"lt"`
						Lu    int `json:"lu"`
						Lv    int `json:"lv"`
						Ly    int `json:"ly"`
						Ma    int `json:"ma"`
						Mc    int `json:"mc"`
						Md    int `json:"md"`
						Me    int `json:"me"`
						Mk    int `json:"mk"`
						Mt    int `json:"mt"`
						Nl    int `json:"nl"`
						No    int `json:"no"`
						Pl    int `json:"pl"`
						Ps    int `json:"ps"`
						Pt    int `json:"pt"`
						Ro    int `json:"ro"`
						Rs    int `json:"rs"`
						Se    int `json:"se"`
						Si    int `json:"si"`
						Sj    int `json:"sj"`
						Sk    int `json:"sk"`
						Sm    int `json:"sm"`
						Sy    int `json:"sy"`
						Tn    int `json:"tn"`
						Tr    int `json:"tr"`
						Ua    int `json:"ua"`
						Uk    int `json:"uk"`
						Us    int `json:"us"`
						Va    int `json:"va"`
						World int `json:"world"`
						Xk    int `json:"xk"`
					} `json:"values"`
					Warning string `json:"warning"`
				} `json:"origins_of_ingredients"`
				Packaging struct {
					NonRecyclableAndNonBiodegradableMaterials int    `json:"non_recyclable_and_non_biodegradable_materials"`
					Value                                     int    `json:"value"`
					Warning                                   string `json:"warning"`
				} `json:"packaging"`
				ProductionSystem struct {
					Labels  []interface{} `json:"labels"`
					Value   int           `json:"value"`
					Warning string        `json:"warning"`
				} `json:"production_system"`
				ThreatenedSpecies struct {
					Warning string `json:"warning"`
				} `json:"threatened_species"`
			} `json:"adjustments"`
			Agribalyse struct {
				Warning string `json:"warning"`
			} `json:"agribalyse"`
			Missing struct {
				Categories  int `json:"categories"`
				Ingredients int `json:"ingredients"`
				Labels      int `json:"labels"`
				Origins     int `json:"origins"`
				Packagings  int `json:"packagings"`
			} `json:"missing"`
			MissingAgribalyseMatchWarning int `json:"missing_agribalyse_match_warning"`
			MissingKeyData                int `json:"missing_key_data"`
			Scores                        struct {
			} `json:"scores"`
			Status string `json:"status"`
		} `json:"ecoscore_data"`
		EcoscoreGrade          string        `json:"ecoscore_grade"`
		EcoscoreTags           []string      `json:"ecoscore_tags"`
		EditorsTags            []string      `json:"editors_tags"`
		EntryDatesTags         []string      `json:"entry_dates_tags"`
		FoodGroupsTags         []interface{} `json:"food_groups_tags"`
		ID                     string        `json:"id"`
		ImageFrontSmallURL     string        `json:"image_front_small_url"`
		ImageFrontThumbURL     string        `json:"image_front_thumb_url"`
		ImageFrontURL          string        `json:"image_front_url"`
		ImageNutritionSmallURL string        `json:"image_nutrition_small_url"`
		ImageNutritionThumbURL string        `json:"image_nutrition_thumb_url"`
		ImageNutritionURL      string        `json:"image_nutrition_url"`
		ImageSmallURL          string        `json:"image_small_url"`
		ImageThumbURL          string        `json:"image_thumb_url"`
		ImageURL               string        `json:"image_url"`
		Images                 struct {
			Num1 struct {
				Sizes struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				UploadedT int    `json:"uploaded_t"`
				Uploader  string `json:"uploader"`
			} `json:"1"`
			Num2 struct {
				Sizes struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				UploadedT int    `json:"uploaded_t"`
				Uploader  string `json:"uploader"`
			} `json:"2"`
			FrontFr struct {
				Angle     int         `json:"angle"`
				Geometry  string      `json:"geometry"`
				Imgid     string      `json:"imgid"`
				Normalize interface{} `json:"normalize"`
				Rev       string      `json:"rev"`
				Sizes     struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num200 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"200"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				WhiteMagic interface{} `json:"white_magic"`
				X1         string      `json:"x1"`
				X2         string      `json:"x2"`
				Y1         string      `json:"y1"`
				Y2         string      `json:"y2"`
			} `json:"front_fr"`
			NutritionFr struct {
				Angle       int         `json:"angle"`
				Geometry    string      `json:"geometry"`
				Imgid       string      `json:"imgid"`
				Normalize   interface{} `json:"normalize"`
				Ocr         int         `json:"ocr"`
				Orientation string      `json:"orientation"`
				Rev         string      `json:"rev"`
				Sizes       struct {
					Num100 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"100"`
					Num200 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"200"`
					Num400 struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"400"`
					Full struct {
						H int `json:"h"`
						W int `json:"w"`
					} `json:"full"`
				} `json:"sizes"`
				WhiteMagic interface{} `json:"white_magic"`
				X1         string      `json:"x1"`
				X2         string      `json:"x2"`
				Y1         string      `json:"y1"`
				Y2         string      `json:"y2"`
			} `json:"nutrition_fr"`
		} `json:"images"`
		InformersTags                       []string      `json:"informers_tags"`
		IngredientsFromPalmOilTags          []interface{} `json:"ingredients_from_palm_oil_tags"`
		IngredientsThatMayBeFromPalmOilTags []interface{} `json:"ingredients_that_may_be_from_palm_oil_tags"`
		InterfaceVersionCreated             string        `json:"interface_version_created"`
		InterfaceVersionModified            string        `json:"interface_version_modified"`
		Lang                                string        `json:"lang"`
		Languages                           struct {
			EnFrench int `json:"en:french"`
		} `json:"languages"`
		LanguagesCodes struct {
			Fr int `json:"fr"`
		} `json:"languages_codes"`
		LanguagesHierarchy []string      `json:"languages_hierarchy"`
		LanguagesTags      []string      `json:"languages_tags"`
		LastEditDatesTags  []string      `json:"last_edit_dates_tags"`
		LastEditor         string        `json:"last_editor"`
		LastImageDatesTags []string      `json:"last_image_dates_tags"`
		LastImageT         int           `json:"last_image_t"`
		LastModifiedBy     string        `json:"last_modified_by"`
		LastModifiedT      int           `json:"last_modified_t"`
		LastUpdatedT       int           `json:"last_updated_t"`
		Lc                 string        `json:"lc"`
		MainCountriesTags  []interface{} `json:"main_countries_tags"`
		MaxImgid           string        `json:"max_imgid"`
		MineralsTags       []interface{} `json:"minerals_tags"`
		MiscTags           []string      `json:"misc_tags"`
		NovaGroupDebug     string        `json:"nova_group_debug"`
		NovaGroupError     string        `json:"nova_group_error"`
		NovaGroupsTags     []string      `json:"nova_groups_tags"`
		NucleotidesTags    []interface{} `json:"nucleotides_tags"`
		NutrientLevels     struct {
		} `json:"nutrient_levels"`
		NutrientLevelsTags []interface{} `json:"nutrient_levels_tags"`
		Nutriments         struct {
			Carbohydrates           float64 `json:"carbohydrates"`
			Carbohydrates100G       float64 `json:"carbohydrates_100g"`
			CarbohydratesUnit       string  `json:"carbohydrates_unit"`
			CarbohydratesValue      float64 `json:"carbohydrates_value"`
			Energy                  int     `json:"energy"`
			EnergyKcal              int     `json:"energy-kcal"`
			EnergyKcal100G          int     `json:"energy-kcal_100g"`
			EnergyKcalUnit          string  `json:"energy-kcal_unit"`
			EnergyKcalValue         int     `json:"energy-kcal_value"`
			EnergyKcalValueComputed float64 `json:"energy-kcal_value_computed"`
			Energy100G              int     `json:"energy_100g"`
			EnergyUnit              string  `json:"energy_unit"`
			EnergyValue             int     `json:"energy_value"`
			Fat                     float64 `json:"fat"`
			Fat100G                 float64 `json:"fat_100g"`
			FatUnit                 string  `json:"fat_unit"`
			FatValue                float64 `json:"fat_value"`
			Proteins                float64 `json:"proteins"`
			Proteins100G            float64 `json:"proteins_100g"`
			ProteinsUnit            string  `json:"proteins_unit"`
			ProteinsValue           float64 `json:"proteins_value"`
			Salt                    float64 `json:"salt"`
			Salt100G                float64 `json:"salt_100g"`
			SaltUnit                string  `json:"salt_unit"`
			SaltValue               float64 `json:"salt_value"`
			SaturatedFat            float64 `json:"saturated-fat"`
			SaturatedFat100G        float64 `json:"saturated-fat_100g"`
			SaturatedFatUnit        string  `json:"saturated-fat_unit"`
			SaturatedFatValue       float64 `json:"saturated-fat_value"`
			Sodium                  float64 `json:"sodium"`
			Sodium100G              float64 `json:"sodium_100g"`
			SodiumUnit              string  `json:"sodium_unit"`
			SodiumValue             float64 `json:"sodium_value"`
			Sugars                  float64 `json:"sugars"`
			Sugars100G              float64 `json:"sugars_100g"`
			SugarsUnit              string  `json:"sugars_unit"`
			SugarsValue             float64 `json:"sugars_value"`
		} `json:"nutriments"`
		Nutriscore struct {
			Num2021 struct {
				CategoryAvailable int `json:"category_available"`
				Data              struct {
					Energy                                   int     `json:"energy"`
					Fiber                                    int     `json:"fiber"`
					FruitsVegetablesNutsColzaWalnutOliveOils int     `json:"fruits_vegetables_nuts_colza_walnut_olive_oils"`
					IsBeverage                               int     `json:"is_beverage"`
					IsCheese                                 int     `json:"is_cheese"`
					IsFat                                    int     `json:"is_fat"`
					IsWater                                  int     `json:"is_water"`
					Proteins                                 float64 `json:"proteins"`
					SaturatedFat                             float64 `json:"saturated_fat"`
					Sodium                                   int     `json:"sodium"`
					Sugars                                   float64 `json:"sugars"`
				} `json:"data"`
				Grade                string `json:"grade"`
				NutrientsAvailable   int    `json:"nutrients_available"`
				NutriscoreApplicable int    `json:"nutriscore_applicable"`
				NutriscoreComputed   int    `json:"nutriscore_computed"`
			} `json:"2021"`
			Num2023 struct {
				CategoryAvailable int `json:"category_available"`
				Data              struct {
					Energy                  int     `json:"energy"`
					Fiber                   int     `json:"fiber"`
					FruitsVegetablesLegumes int     `json:"fruits_vegetables_legumes"`
					IsBeverage              int     `json:"is_beverage"`
					IsCheese                int     `json:"is_cheese"`
					IsFatOilNutsSeeds       int     `json:"is_fat_oil_nuts_seeds"`
					IsRedMeatProduct        int     `json:"is_red_meat_product"`
					IsWater                 int     `json:"is_water"`
					Proteins                float64 `json:"proteins"`
					Salt                    float64 `json:"salt"`
					SaturatedFat            float64 `json:"saturated_fat"`
					Sugars                  float64 `json:"sugars"`
				} `json:"data"`
				Grade                string `json:"grade"`
				NutrientsAvailable   int    `json:"nutrients_available"`
				NutriscoreApplicable int    `json:"nutriscore_applicable"`
				NutriscoreComputed   int    `json:"nutriscore_computed"`
			} `json:"2023"`
		} `json:"nutriscore"`
		Nutriscore2021Tags                          []string      `json:"nutriscore_2021_tags"`
		Nutriscore2023Tags                          []string      `json:"nutriscore_2023_tags"`
		NutriscoreGrade                             string        `json:"nutriscore_grade"`
		NutriscoreTags                              []string      `json:"nutriscore_tags"`
		NutriscoreVersion                           string        `json:"nutriscore_version"`
		NutritionDataPer                            string        `json:"nutrition_data_per"`
		NutritionDataPreparedPer                    string        `json:"nutrition_data_prepared_per"`
		NutritionGradeFr                            string        `json:"nutrition_grade_fr"`
		NutritionGrades                             string        `json:"nutrition_grades"`
		NutritionGradesTags                         []string      `json:"nutrition_grades_tags"`
		NutritionScoreBeverage                      int           `json:"nutrition_score_beverage"`
		NutritionScoreDebug                         string        `json:"nutrition_score_debug"`
		NutritionScoreWarningNoFiber                int           `json:"nutrition_score_warning_no_fiber"`
		NutritionScoreWarningNoFruitsVegetablesNuts int           `json:"nutrition_score_warning_no_fruits_vegetables_nuts"`
		OtherNutritionalSubstancesTags              []interface{} `json:"other_nutritional_substances_tags"`
		PackagingMaterialsTags                      []interface{} `json:"packaging_materials_tags"`
		PackagingRecyclingTags                      []interface{} `json:"packaging_recycling_tags"`
		PackagingShapesTags                         []interface{} `json:"packaging_shapes_tags"`
		Packagings                                  []interface{} `json:"packagings"`
		PackagingsMaterials                         struct {
		} `json:"packagings_materials"`
		PhotographersTags    []string      `json:"photographers_tags"`
		PnnsGroups1          string        `json:"pnns_groups_1"`
		PnnsGroups1Tags      []string      `json:"pnns_groups_1_tags"`
		PnnsGroups2          string        `json:"pnns_groups_2"`
		PnnsGroups2Tags      []string      `json:"pnns_groups_2_tags"`
		PopularityKey        int           `json:"popularity_key"`
		ProductName          string        `json:"product_name"`
		ProductNameFr        string        `json:"product_name_fr"`
		RemovedCountriesTags []interface{} `json:"removed_countries_tags"`
		Rev                  int           `json:"rev"`
		SelectedImages       struct {
			Front struct {
				Display struct {
					Fr string `json:"fr"`
				} `json:"display"`
				Small struct {
					Fr string `json:"fr"`
				} `json:"small"`
				Thumb struct {
					Fr string `json:"fr"`
				} `json:"thumb"`
			} `json:"front"`
			Nutrition struct {
				Display struct {
					Fr string `json:"fr"`
				} `json:"display"`
				Small struct {
					Fr string `json:"fr"`
				} `json:"small"`
				Thumb struct {
					Fr string `json:"fr"`
				} `json:"thumb"`
			} `json:"nutrition"`
		} `json:"selected_images"`
		Sortkey               int           `json:"sortkey"`
		States                string        `json:"states"`
		StatesHierarchy       []string      `json:"states_hierarchy"`
		StatesTags            []string      `json:"states_tags"`
		Traces                string        `json:"traces"`
		TracesFromIngredients string        `json:"traces_from_ingredients"`
		TracesFromUser        string        `json:"traces_from_user"`
		TracesHierarchy       []interface{} `json:"traces_hierarchy"`
		TracesTags            []interface{} `json:"traces_tags"`
		UnknownNutrientsTags  []interface{} `json:"unknown_nutrients_tags"`
		UpdateKey             string        `json:"update_key"`
		VitaminsTags          []interface{} `json:"vitamins_tags"`
	} `json:"product"`
	Status        int    `json:"status"`
	StatusVerbose string `json:"status_verbose"`
}
