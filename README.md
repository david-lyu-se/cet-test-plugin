# We want:
# - to check if file exists in ~/.cet-wp-plugins ✔️
#  - If not create file environment.json ✔️
#  - Get users input environment name (key) : path (value)
# - Menu ask for plugin
#  - Add plugin
#   - plugin name(key) : path (value)
#   - save in json file
#  - Use plugin

# Using plugin
# - composer install root
# - block library recusively go through folders and npm install
# - rsync plugin into wpvip/plugins file

# Extra credit: Seperate models and menu (Generic filepicker, Generic UserInput)

**Menus:**
1) Main menu to:
 - Add application (Submenu 1)
 - Add mono repo directory (Submenu 2)
 - Enter application (Submenu 3)
 - Edit application (Tell user to edit in .json file)
 - Delete Application (Tell user to edit in .json file)

**SubMenu:**
1) Add application
 - FilePicker
 - Input name
  - Do I want verification?
 - Update in File and return UpdateApplicationMsg
2) Mono repo directory
 - FilePicker (again) <- should move logic out to DRY
 - Input name
 - Edit/Delete ? (Tell user to edit in .json file)
3) Application Selected
 - Go to 2nd Level Menu
**2nd Level Menu**
- Update Plugins Dir
- Select Plugin
  - FilePicker (Not sure what file to use composer.json should be good)
  - Looking for vendor
    - No vendor composer install
  - Options:
    - Install all
      - composer, npm install, npm run watch
    - Install block
      - Get userinput
    - Enter to start sync and leave menu on
    - Edit install block choice
    - Exit
      - npm run build
      - rysnc
  - rsync to <application>/plugin
