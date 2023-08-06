import pandas as pd
import matplotlib.pyplot as plt
import matplotlib as mpl
import seaborn as sns
import numpy as np
import string
import random
from scipy.ndimage.filters import gaussian_filter
import pytraj as pt

check = sns.__version__
check2 = check.split('.')
if int(check2[1]) <= 7:
    print('\x1b[4;30;41m'+'WARNING!!'+'\x1b[0m'+'\x1b[3;30;41m'+' The detected seaborn version installed is below 0.8.0! ('+str(check)+')\nSome utilities of the seaborn v0.8.0 are used in this module\n'+'\x1b[0m'+'\x1b[3;30;43m'+'UPGRADE SEABORN!'+'\x1b[0m')

names_default = ('Serie 1', 'Serie 2', 'Serie 3', 'Serie 4', 'Serie 5', 'Serie 6', 'Serie 7', 'Serie 8', 'Serie 9')
palette_default = ((sns.blend_palette(('#fbeaff','#fbeaff'), n_colors=50, as_cmap=True)),(sns.blend_palette(('#CCCCF5','#CCCCF5' ), n_colors=50, as_cmap=True)))
palette_default_r = ((sns.blend_palette(('#9444a8', '#be8eca','#be8eca','#be8eca','#be8eca','#be8eca'), n_colors=50, as_cmap=True)),(sns.blend_palette(('#0000A4','#7f7fd1','#7f7fd1','#7f7fd1','#7f7fd1','#7f7fd1'), n_colors=50, as_cmap=True)))



def get_color(index):
    '''
    Function that returns a selected color in HEX format.
    
    Parameters: ---------------------------------------------------------------------------------
    
    index: integer
    The index selects the color that will be returned. Index and order are selected to return nice colors, comparable with the previous and Osuna friendly.
    
    
    Returns: -------------------------------------------------------------------------------------------
    
    -String of the HEX code of the color selected
    
    '''
    
    colors=[]
    
    medium_orchid=(186,85,211)
    medium_blue=(0,0,205)
    dark_orange=(255,140,0)
    mediumseagreen=(60,179,113)
    slate_blue=(106,90,205)
    crimson=(220,20,60)
    gold=(255,215,0)
    medium_turquoise=(72,209,204)
    deep_sky_blue=(0,191,255)
    teal=(0,128,128)
    deep_pink=(255,20,147)
    midnight_blue=(25,25,112)
    dark_sea_green=(143,188,143)
    orange_red=(255,69,0)
    green_yellow=(173,255,47)
    dark_blue=(0,0,139)
    olive=(128,128,0)
    plum=(221,160,221)
    slate_gray=(112,128,144)
    salmon=(250,128,114)
    purple=(128,0,128)
    black=(0,0,0)
    dark_red=(139,0,0)
    dark_violet=(148,0,211)
    khaki=(240,230,140)
    hot_pink=(255,105,180)
    aqua=(0,255,255)
    silver=(192,192,192)
    limegreen=(50,205,50)
    magenta=(255,0,255)
    lightslategray=(119,136,153)
    khaki=(240,230,140)
    sea_green=(46,139, 87)
    
    
    colors.append(medium_orchid)
    colors.append(medium_blue)
    colors.append(medium_turquoise)
    colors.append(dark_orange)
    colors.append(crimson)
    colors.append(limegreen)
    colors.append(slate_blue)
    colors.append(gold)
    colors.append(hot_pink)
    colors.append(green_yellow)
    colors.append(deep_pink)
    colors.append(dark_violet)
    colors.append(teal)
    colors.append(midnight_blue)
    colors.append(deep_sky_blue)
    colors.append(purple)
    colors.append(olive)
    colors.append(orange_red)
    colors.append(slate_gray)
    colors.append(salmon)
    colors.append(dark_sea_green)
    colors.append(black)
    colors.append(dark_red)
    colors.append(khaki)
    colors.append(plum)
    colors.append(aqua)
    colors.append(silver)
    colors.append(lightslategray)
    
    if (index <0 or index > (len(colors)-1)):
        return '#%02x%02x%02x' % get_color_random()
    else:
        return '#%02x%02x%02x' % colors[index]


#---------------------------------------------------------------------------------------------------------------


def get_color_random():
    '''
    Function that returns a random color in RGB format
        
    '''
    a2=int(int(random.randrange(0,255)))
    b2=int(int(random.randrange(0,255)))
    c2=int(int(random.randrange(0,255)))
        
    color=(a2,b2,c2)
        
    return(color)

#---------------------------------------------------------------------------------------------------------------


def plot_dist(data, legend=True, names='default',read_every=1, color='default', window_size=50, alpha=0.2, shadow='sd', style='seaborn-white', font_size=20, figure_size=(10,8), title=False, title_fontsize=30, x_label='Simulation time (ns)', y_label='Distance ($\AA$)', x_limits='default', y_limits='default', vertical_lines=False, horizontal_lines=False,  line_width=1.8, x_label_fontsize=20, y_label_fontsize=20):
    '''
        
        Plot one or more timeseries with flexible and multiple options for the CompBioLab.
        
        This function is intended to be used with data extracted from MD silulations, such as: distances, angles, dihedrals, rmsd, etc.
        
        This function can display the row data (window_size=1) or the mean calculated every window_size, and the deviation is displayed as a 'unit_traces' (check seaborn tsplot).
        
        The default values are only prepared to plot a maximum numbers of 9 timeseries, if you want to plot more than 9 (not recomended, because the plot will become crowded), modify the default parameters.
        
        If you experience some problem or you have any comment with this function --> miqueleg@gmail.com :p
        
        Parameters: ---------------------------------------------------------------------------------
        
        data: list, nested list or ndarray
        Data for the plot. This data is the Y values. Be aware that the NAN values can make your plots weird.
        Headers has not to be included in the array. Remove all strings from the data
        
        legend: boolean(True or False)
        Set legend to True to display a legend in the plot
        
        names: list of strings
        Set the names of the series displayed on the legend.
        The 'default' will show Serie 1, Serie 2, etc.
        
        read_every: integer or floating point
        Value that changes the values of the X axis to fit correctly with the time unit used.
        The default value is calibrated for cpptraj outputs in MD simulations written every 10000 steps (GALATEA),
        and readed every 1
        
        color: list of strings
        List of the colors for the series. Seaborn, matplotlib and Hex formats are accepted.
        In general are not recomended to change the 'default' mode. The 'default' colors are cool!
        
        window_size: integer
        Width of the window of values that are used to calculate the mean in each position.
        If this value is higher, more noise are removed, but also more information.
        
        alpha: float
        Value that defines the transparency of the ci_areas(shadows)
        The default value is 0.2. 
        
        shadow: 'sd' or 'values'
        Type of the data plotted as a error ci_areas(shadows).
        If 'sd'(default) is selected, the ci_areas are the mean values +- StandardDeviation.
        If 'values' is selected, the ci_areas are the plot of the real values.
        
        style: string or None
        Matplotlib style used for the plot(nore information on the seaborn webpage).
        By default is used, the style will be 'seaborn-white'. Also, 'seaborn-darkgrid' is recomended too.
        
        font_size: integer
        Basic fontsize value.
        
        figure_size: list of two integers
        Dimensions of the resulting image.
        
        title: string or False
        Title that will be displayed.
        
        x_label: string of False
        Name of the X axis that will be displayed in the plot.
        
        y_label: string of False
        Name of the Y axis that will be displayed in the plot.
        
        x_limits: 'default' or list of two integers
        The 'default' mode will display the plot with tightened X axis limits that will change for each series.
        *This parameter cannot be 'default' if you use the horizontal_line parameter.
        
        y_limits: 'default' or list of two integers
        The 'default' mode will display the plot with tightened Y axis limits that will change for each timeseries.
        *This parameter cannot be 'default' if you use the vertical_line parameter.
        
        vertical_lines: False, integer or list of integers
        This paremeter will add a line (or multiple lines) parellel to the Y axis at the selected X position.
        *If this parameter is not False, then the y_limits parameter cannot be in 'default' mode.
        
        horizontal_lines: False, integer or list of integers
        This paremeter will add a line (or multiple lines) parellel to the X axis at the selected Y position.
        *If this parameter is not False, then the x_limits parameter cannot be in 'default' mode.
        
        line_width: integer or floating point
        Value that sets the width of the lines in the plot.
        
        x_label_fontsize: integer
        Value that sets the size of the font that makes the x_label parameter.
        
        y_label_fontsize: integer
        Value that sets the size of the font that makes the y_label parameter.
        
        Returns: -------------------------------------------------------------------------------------------
        
        -Nothing
        
        In jupyter-notebook (RECOMENDED!)
        put the command '%matplotlib inline' before this command to automatically display the output
        
        
        '''
    
    C = read_every*0.02

    if len(data) < 10 and type(data[0]).__module__ == np.__name__:
        data = [pd.Series(i) for i in data]
        print('list_numpy')
    elif len(data) >= 10 and type(data).__module__ == np.__name__:
        data = pd.Series(data)
        print('numpy')
    elif len(data) >= 10 and type(data) == list:
        print('list')
        data = pd.Series(data)
   
    
    
    #Selecting the variables depending if they are 'default', str, or list(in a multiple plot)
    if color == 'default':
        color=[]
        for data_set in range(len(data)):
            color.append(get_color(data_set))
    elif type(color) == str:
        color=[color, '0']
    
    if names == 'default':
        names = names_default
    elif type(names) == str:
        names=[names, '0']
        
    

    
    
    #Applying variables that will appear always(specs)
    mpl.style.use(style)
    mpl.rcParams['lines.linewidth'] = line_width
    mpl.rcParams['font.size'] = font_size
    plt.figure(figsize=figure_size)




    #The if checking legend, is here because the legend depends on the 'condition' in the tsplot command.
    #The plots in the true has this condition and in the else, not.
    if legend == True:
    
    #Then, the correct tsplots are selected in the try/except, depending on the number of 'datas'
        try:
            data = np.ravel(data.values)
            
            m = pd.Series(data).rolling(window_size).mean()
            sd = pd.Series(data).rolling(window_size).std()
            xs = np.array(range(len(data)))*C

            plt.plot(xs, m, color=color[0], label=names[0])
            
            if shadow == 'values':
                plt.plot(xs, data, color=color[0], alpha=alpha)
                
            elif shadow == 'sd':
                plt.fill_between(xs, m+2*sd,m-2*sd, color=color[0], alpha=alpha)
                
  
        except:
            for l in list(range(len(data))):
                ds = np.ravel(data[l].values)
                
                m = pd.Series(ds).rolling(window_size).mean()
                sd = pd.Series(ds).rolling(window_size).std()
                xs = np.array(range(len(ds)))*C

                plt.plot(xs, m, color=color[l], label=names[l])

                if shadow == 'values':
                    plt.plot(xs, ds, color=color[l], alpha=alpha)

                elif shadow == 'sd':
                    plt.fill_between(xs, m+2*sd,m-2*sd, color=color[l], alpha=alpha)
        plt.legend(loc='upper left', bbox_to_anchor=(1,1))



    elif legend == False:
    
        try:
            data = np.ravel(data.values)
            
            m = pd.Series(data).rolling(window_size).mean()
            sd = pd.Series(data).rolling(window_size).std()
            xs = np.array(range(len(data)))*C

            plt.plot(xs, m, color=color[0], label=names[0])
            
            if shadow == 'values':
                plt.plot(xs, data, color=color[0], alpha=alpha)
                
            elif shadow == 'sd':
                plt.fill_between(xs, m+2*sd,m-2*sd, color=color[0], alpha=alpha) 
            
        except:
            for l in list(range(len(data))):
                ds = np.ravel(data[l].values)
                
                m = pd.Series(ds).rolling(window_size).mean()
                sd = pd.Series(ds).rolling(window_size).std()
                xs = np.array(range(len(ds)))*C

                plt.plot(xs, m, color=color[l])

                if shadow == 'values':
                    plt.plot(xs, ds, color=color[l], alpha=alpha)

                elif shadow == 'sd':
                    plt.fill_between(xs, m+2*sd,m-2*sd, color=color[l], alpha=alpha)

    #Post-plotting modifications
    plt.tick_params(direction='out', length=6, width=2)
    
    ##Applying variables that will not appear if they are in 'default'/False
    if title:
        plt.title(title, fontsize=title_fontsize)

    if x_label:
        plt.xlabel(x_label, fontsize=x_label_fontsize, fontstyle='italic')
    
    if y_label:
        plt.ylabel(y_label, fontsize=y_label_fontsize, fontstyle='italic')

    if x_limits != 'default':
        plt.xlim(x_limits)
    else:
        plt.xlim(0)
    
    if y_limits != 'default':
        plt.ylim(y_limits)
    else:
        try:
            ymin = min([min(data),0])
            ymax = max(data)
        except:
            ymin = min([min([min(y.values) for y in data]),0])
            ymax = max([max(y.values) for y in data])
        plt.ylim(ymin,ymax)

    if vertical_lines:
        plt.vlines(vertical_lines, y_limits[0], y_limits[1])
    
    if horizontal_lines:
        plt.hlines(horizontal_lines,x_limits[0], x_limits[1])



#---------------------------------------------------------------------------------------------------------------

def plot_kde(data1, data2=None, mode=1,  names='default', x_limits='default', y_limits='default', style='ticks', alpha=0.1, n=30, font_scale=2, dpi=45, title=False, title_fontsize=45, font_style='italic', ticks_size=30, axis_label_fontsize=35, hist=False, kde=True):
    '''
        
        Plot one or two pair of variables in a KDE plot. For CompBioLab reserchers.
        
        This function is intended to be used with data extracted from MD silulations, such as: distances, angles, dihedrals, rmsd, etc.
        
        This function can display the row data (mode=1), the kernel density function (mode=1) or both(mode=2). In the sides are displayed the distributions, in KDE formar to histogram (or both).
        
        The default values are only prepared to plot the KDE density plot and the KDE distributions.
        
        If you experience some problem or you have any comment with this function --> miqueleg@gmail.com :p
        
        Parameters: ---------------------------------------------------------------------------------
        
        data1: two variables list. This variables can be: lists, ndarray or pandas.Series
        Be aware that the NAN values can make your plots weird.
        Headers has not to be included in the array. Remove all strings from the data.
        The Format of the data has to be: nested list with len() == 2, np.array with 2 columns or a list with two pandas.Series. If you have the data in a pandas.DataFrame, pass this data to Series(eg: [data['column1'], data['column2'])
        
        data2: two variables list. This variables can be: lists, ndarray or pandas.Series
        Be aware that the NAN values can make your plots weird.
        Headers has not to be included in the array. Remove all strings from the data.
        The Format of the data has to be: nested list with len() == 2, np.array with 2 columns or a list with two pandas.Series. If you have the data in a pandas.DataFrame, pass this data to Series(eg: [data['column1'], data['column2'])
        
        mode: integer (0,1 or 2)
        Set the mode of the function.
        This function can display the row data (mode=1), the kernel density function (mode=1) or both(mode=2). In the sides are displayed the distributions, in KDE formar to histogram (or both).
        
        n: integer
        Number of lines or levels in the bivariate KDE plot.

        names: list of strings
        Set the names of the series displayed on the legend.
        The 'default' will show Serie 1, Serie 2, etc.
        
        style: string
        Seaborn style used for the plot(more information on the seaborn webpage).
        By default is used 'ticks'
        
        alpha: float
        Alpha value (transparency) for the scatter plot in mode 0 and 2.
        
        font_scale: integer
        Multiplying factor of fontsize of all letters on the plot.
        Later, this value is modifyed by th other values of fontsize. This only affects to fonts that are no normaly by default(annotations and text)
        
        dpi: integer
        Size of the resulting image.
        
        title: string or False
        Title that will be displayed.
        
        title_fontsize: integer.
        Value that sets the size of the font that makes the title parameter.
        
        font_style: string
        Sets the style of the fonts. See rcParams to get the list of possibilities.
        
        ticks_size: 
        Value that sets the size of the font of the ticks(numbers).
        
        x_limits: 'default' or list of two integers
        The 'default' mode will display the plot with tightened X axis limits that will change for each series.
        *This parameter cannot be 'default' if you use the horizontal_line parameter.
        
        y_limits: 'default' or list of two integers
        The 'default' mode will display the plot with tightened Y axis limits that will change for each timeseries.
        *This parameter cannot be 'default' if you use the vertical_line parameter.
        
        axis_label_fontsize: integer
        Value that sets the size of the font that makes the x_label and y_label parameters.
        
        kde: boolean
        If trre, displays the KDE distribution in the side plots.
        
        hist: boolean
        If true, displays the histogram distribution in the side plots.
        
        Returns: -------------------------------------------------------------------------------------------
        
        -Nothing
        
        In jupyter-notebook (RECOMENDED!)
        put the command '%matplotlib inline' before this command to automatically display the output
        
        
        
        '''
    assert (mode < 3),'Mode {} is selected. Only mode 0,1 and 2 are valid ones'.format(mode)
    
    # Selecting Mode. If data2 exists, multiple is setted to True, and the data is converted into pandas pandas.Series
    multiple = False
    try:
        data2 = [pd.Series(d) for d in data2]
        data1 = [pd.Series(d) for d in data1]
        multiple = True
    
    except:
        try:
            if len(data1) == 1:
                data1 = [pd.Series(d)for d in data1]
        except:
            try:
                data2 = [d for d in data2]
                data1 = [d for d in data1]
                multiple = True
            
            except:
                try:
                    data1 = [d for d in data1]
                except:
                    raise FormatErrror('data1 is not in the correct format.\nThe Format of data1 has to be: nested list vith len() == 2, np.array with 2 columns or a list with two pandas.Series. If you have the data in a pandas.DataFrame, pass this data to Series(eg: [data["column1"], data["column2"])')

    #Setting some configuration to the figure
    plt.figure(figsize=(6,5))
    sns.set(font_scale=font_scale, rc={"font.style":font_style, 'figure.dpi':dpi,'xtick.labelsize':ticks_size,'ytick.labelsize':ticks_size, 'axes.labelsize':axis_label_fontsize})
    sns.set_style(style)
    
    if names == 'default':
        names = ['Serie 1','Serie 2']

    #If the plot is multiple, the data with he higher histogram value has to be plotted first in the distplot, if not the plot will be cutted
    if multiple:
        data1 = pd.DataFrame({names[0]:data1[0],names[1]:data1[1]}).astype(float)
        data2 = pd.DataFrame({names[0]:data2[0],names[1]:data2[1]}).astype(float)
        
        a1,b1,c1 = plt.hist(data1[names[0]].astype(float),bins=100)
        perint(a1)
        a2,b2,c2 = plt.hist(data2[names[0]].astype(float),bins=100)
        plt.close()
        
        if max(a1) > max(a2):
            X0 = data1[names[0]]
            X1 = data2[names[0]]
            X_color_default = [get_color(0),get_color(1)]
            print('The max value of the x_jointgrid is selected from data1')
        else:
            X0 = data2[names[0]]
            X1 = data1[names[0]]
            X_color_default = [get_color(1),get_color(0)]
            print('The max value of the x_jointgrid is selected from data2')

        a1,b1,c1 = plt.hist(data1[names[1]].astype(float),bins=100)
        a2,b2,c2 = plt.hist(data2[names[1]].astype(float),bins=100)
        plt.close()

        if max(a1) > max(a2):
            Y0 = data1[names[1]]
            Y1 = data2[names[1]]
            Y_color_default = [get_color(0),get_color(1)]
            print('The max value of the y_jointgrid is selected from data1')
        else:
            Y0 = data2[names[1]]
            Y1 = data1[names[1]]
            Y_color_default = [get_color(1),get_color(0)]
            print('The max value of the y_jointgrid is selected from data2')

        # Creates the grid
        g = sns.JointGrid(data1[names[0]], data1[names[1]], size=15, ratio=5, space=0, xlim=None, ylim=None)

        #Plots depending on the mode selected. The order of the plot lines are important
        if mode == 0:
            ax = g.ax_joint.plot(X1, Y0, "o", mew=0.2, ms=3, c=get_color(0), alpha=alpha)
            ax = g.ax_joint.plot(X0, Y1, "o", mew=0.2, ms=3, c=get_color(1), alpha=alpha)
            sns.distplot(X0, hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=X_color_default[0])
            sns.distplot(Y0, hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=Y_color_default[0])
            sns.distplot(X1, hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=X_color_default[1])
            sns.distplot(Y1, hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=Y_color_default[1])


        elif mode == 1:
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, shade=True, shade_lowest=False, cmap=palette_default[0])
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, cmap=palette_default_r[0])
            ax = sns.kdeplot(data2[names[0]], data2[names[1]], ax=g.ax_joint, n_levels=n, shade=True, shade_lowest=False, cmap=palette_default[1])
            ax = sns.kdeplot(data2[names[0]], data2[names[1]], ax=g.ax_joint, n_levels=n, cmap=palette_default_r[1])
            
            sns.distplot(X0, hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=X_color_default[0])
            sns.distplot(Y0, hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=Y_color_default[0])
            sns.distplot(X1, hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=X_color_default[1])
            sns.distplot(Y1, hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=Y_color_default[1])

        elif mode == 2:
            ax = g.ax_joint.plot(X1, Y0, "o", mew=0.2, ms=3, c=get_color(0), alpha=alpha)
            ax = g.ax_joint.plot(X0, Y1, "o", mew=0.2, ms=3, c=get_color(1), alpha=alpha)
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, shade=True, shade_lowest=False, cmap=palette_default[0])
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, cmap=palette_default_r[0])
            ax = sns.kdeplot(data2[names[0]], data2[names[1]], ax=g.ax_joint, n_levels=n, shade=True, shade_lowest=False, cmap=palette_default[1])
            ax = sns.kdeplot(data2[names[0]], data2[names[1]], ax=g.ax_joint, n_levels=n, cmap=palette_default_r[1])
            
            sns.distplot(X0, hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=X_color_default[0])
            sns.distplot(Y0, hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=Y_color_default[0])
            sns.distplot(X1, hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=X_color_default[1])
            sns.distplot(Y1, hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=Y_color_default[1])

    else:
        data1 = pd.DataFrame({names[0]:data1[0],names[1]:data1[1]}).astype(float)
        
        # Creates the grid, Because is not multiple, the order of the data plotted in the histogram is not important
        g = sns.JointGrid(data1[names[0]], data1[names[1]], size=15, ratio=5, space=0, xlim=None, ylim=None)
        
        if mode == 0:
            ax = g.ax_joint.plot(data1[names[0]], data1[names[1]], "o", mew=0.2, ms=3, c=get_color(0), alpha=alpha)
            sns.distplot(data1[names[0]], hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=get_color(0))
            sns.distplot(data1[names[1]], hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=get_color(0))
        
        elif mode == 1:
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, shade=True, shade_lowest=False, cmap=palette_default[0])
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, cmap=palette_default_r[0])
            
            sns.distplot(data1[names[0]], hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=get_color(0))
            sns.distplot(data1[names[1]], hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=get_color(0))
        
        elif mode == 2:
            ax = g.ax_joint.plot(data1[names[0]], data1[names[1]], "o", mew=0.2, ms=3, c=get_color(0), alpha=alpha)
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, shade=True, shade_lowest=False, cmap=palette_default[0])
            ax = sns.kdeplot(data1[names[0]], data1[names[1]], ax=g.ax_joint, n_levels=n, cmap=palette_default_r[0])
            
            sns.distplot(data1[names[0]], hist=hist, kde=kde, ax=g.ax_marg_x, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=get_color(0))
            sns.distplot(data1[names[1]], hist=hist, kde=kde, vertical=True, ax=g.ax_marg_y, kde_kws={'legend':False, 'shade':True}, axlabel=False, color=get_color(0))


    #Final modifications
    if title:
        plt.suptitle(title, fontsize=title_fontsize, y=1.01)

    if x_limits != 'default':
        ax.xlim(x_limits)
    
    if y_limits != 'default':
        ax.ylim(y_limits)



#---------------------------------------------------------------------------------------------------------------        
        
        
def calculate_rmsf(trajfiles, topfile, plot=True, ref=0, mask='@CA', name='default'):
    '''
        This function returns a numpy.ndarray object that contains rmsf information of the atoms in the mask.
        
        Arguments:
        
        - trajfiles: trajectory file or list of trajecrory files(see formats in the bottom part).
        
        - topfile: topology file used to read the trajectories(see formats in the bottom part).
        
        - ref: frame used as a reference while the RMSF is calculated(integer).
        
        - mask: string that indicates with cpptraj format the mask for the calculation (default: "@Ca").
        
        - plot: if True, plots a RMSF plot(default: True).
        
        - name: Name of the dataset provided. Will apear in the title in the plot.
        
        Returns: numpy.ndarray
        
        
        Formats that accept: 
        
        Amber Trajectory 	.crd
        Amber NetCDF 	.nc
        Amber Restart 	.rst7
        Amber NetCDF 	.ncrst
        Charmm DCD 	.dcd
        PDB 	.pdb
        Mol2 	.mol2
        Scripps 	.binpos
        Gromacs 	.trr
        SQM Input 	.sqm
        '''
    
    if name == 'default':
        if type(trajfiles) == list:
            name = trajfiles[0].split('/')[-1]
        else:
            name = trajfiles.split('/')[-1]
    
    
    MD = pt.iterload(trajfiles, topfile).superpose(trajfiles, ref=ref)
    rmsf = pt.rmsf(MD, mask)
    
    if plot:
        
        mpl.rcParams['font.size'] = 15
        plt.figure(figsize=(12,5))
        plt.plot(rmsf.T[1])
        plt.xlabel('{} number'.format(mask), fontsize=20)
        plt.ylabel('RMSF ($\AA$)', fontsize=20)
        plt.xlim((0,len(rmsf)))
        plt.title('{} RMSF'.format(name), fontsize=25)
    
    return rmsf.T[1]


#---------------------------------------------------------------------------------------------------------------

def read_pdb(pdb_path, header=False):
    '''
        This function returns a pandas.DataFrame object that contains the information if the .pdb.
        
        Arguments:
        
        - pdb_path: string with the file location
        
        - header: If header is True, you are saying that the PDB file has header that has to be ommited.
        
        Returns: pandas.DataFrame
        '''
    
    colspecs = [(0, 6), (6, 11), (12, 16), (16, 17), (17, 20), (21, 22), (22, 26),
                (26, 27), (30, 38), (38, 46), (46, 54), (54, 60), (60, 66), (76, 78),
                (78, 80)]
        
    names = ['ATOM', 'serial', 'name', 'altloc', 'resname', 'chainid', 'resid',
                 'icode', 'x', 'y', 'z', 'occupancy', 'tempfactor', 'element', 'charge']
            
    pdb = pd.read_fwf(pdb_path, names=names, colspecs=colspecs)
                
    if header:
        pdb = pdb[pdb[pdb['ATOM'] == 'ATOM'].index[0]:]

    return pdb


#---------------------------------------------------------------------------------------------------------------
def change_his(input_file, output_file, HIE=(), HID=(), HIP=()):
    '''
        Function that creates a new file with the provided changes from the HIS to HIE or HID.
        
        Arguments:
        
        - input_file: string with the input file location and name.
        
        - output_file: string with the output file locationand name.
        
        - output_file: string with the output file locationand name.
        
        - HIE: list of residues numbers that will be changed to from HIS to HIE. The list can be empty.
        
        - HID: list of residues numbers that will be changed to from HIS to HID. The list can be empty.
        
        - HIP: list of residues numbers that will be changed to from HIS to HIP. The list can be empty.
        
        Returns: nothing
        
        '''
    
    HIE = [str(x) for x in HIE]
    HID = [str(x) for x in HID]
    HIP = [str(x) for x in HIP]
    
    inp = open(input_file, 'r')
    out = open(output_file, 'w')
    i = 0
    for line in inp.readlines():
        i += 1
        try:
            if line.split()[4] in HIE:
                line = string.replace(line, 'HIS','HIE')
            elif line.split()[4] in HID:
                line = string.replace(line, 'HIS','HID')
            elif line.split()[4] in HIP:
                line = string.replace(line, 'HIS','HIP')
            out.write(line)
        except:
            out.write(line)
    inp.close()
    out.close()


#---------------------------------------------------------------------------------------------------------------
def change_asp(input_file, output_file, ASH):
    '''
        Function that creates a new file with the provided changes from the ASP to ASH.
        
        Arguments:
        
        - input_file: string with the input file location and name.
        
        - output_file: string with the output file locationand name.
        
        - ASH: list of residues numbers that will be changed to from ASP to ASH.
        
        
        Returns: nothing
        
        '''
    
    ASH = [str(x) for x in ASH]
    
    inp = open(input_file, 'r')
    out = open(output_file, 'w')
    i = 0
    for line in inp.readlines():
        i += 1
        try:
            if line.split()[4] in ASH:
                line = string.replace(line, 'ASP','ASH')
            out.write(line)
        except:
            out.write(line)
    inp.close()
    out.close()


#---------------------------------------------------------------------------------------------------------------

def change_glu(input_file, output_file, GLH):
    '''
        Function that creates a new file with the provided changes from the GLU to GLH.
        
        Arguments:
        
        - input_file: string with the input file location and name.
        
        - output_file: string with the output file locationand name.
        
        - GLH: list of residues numbers that will be changed to from GLU to GLH.
        
        
        Returns: nothing
        
        '''
    
    GLH = [str(x) for x in GLH]
    
    inp = open(input_file, 'r')
    out = open(output_file, 'w')
    i = 0
    for line in inp.readlines():
        i += 1
        try:
            if line.split()[4] in GLH:
                line = string.replace(line, 'GLU','GLH')
            out.write(line)
        except:
            out.write(line)
    inp.close()
    out.close()



#---------------------------------------------------------------------------------------------------------------


def plot_Nice_PES(P_test,bins=90,sigma=0.99, title=False, size = 1):
    '''
        Plots the Free Energy Surface(FES) in a more nicer way than "pyemma.plots.plot_free_energy". Also, the colorbar displayed is in kcal/mol(*0.592)
        
        Arguments:
        
        - P_test: Input data(tica.get_output) list of ndarray(T_i, d). From pyemma: "the mapped data, where T is the number of time steps of the input data, or if stride > 1, floor(T_in / stride). d is the output dimension of this transformer. If the input consists of a list of trajectories, Y will also be a corresponding list of trajectories"
        
        - bins: int or array_like or [int, int] or [array, array]. Number of ranges used in the histogram calculation. Default = 90
        
            The bin specification:
            
            If int, the number of bins for the two dimensions (nx=ny=bins).
            If array_like, the bin edges for the two dimensions (x_edges=y_edges=bins).
            If [int, int], the number of bins in each dimension (nx, ny = bins).
            If [array, array], the bin edges in each dimension (x_edges, y_edges = bins).
            A combination [int, array] or [array, int], where int is the number of bins and array is the bin edges.
        
        - sigma: integer for the Standard deviation for Gaussian kernel(filter applied). The standard deviations of the Gaussian filter are given for each axis as a sequence, or as a single number, in which case it is equal for all axes. Default = 0.99
        
        - title: string that contains the output title. If False, the title is not displayed. Default = False
        
        - size: factor that multiplies the size values of the plot (figsize and fontsize). Default = 1
        
        
        Returns: nothing
        
        In jupyter-notebook (RECOMENDED!)
        put the command '%matplotlib inline' before this command to automatically display the output

        '''
    
    mpl.style.use("seaborn-paper")
    plt.figure(figsize=(6*size,5*size))
    alldata=np.vstack(P_test)
    min1=np.min(alldata[:,0])
    max1=np.max(alldata[:,0])
    min2=np.min(alldata[:,1])
    max2=np.max(alldata[:,1])
    
    tickspacing1=1.0
    tickspacing2=1.0
    z,x,y = np.histogram2d(alldata[:,0], alldata[:,1], bins=bins)
    z += 0.1
    
    # compute free energies
    F = -np.log(z)
    
    
    # contour plot
    extent = [x[0], x[-1], y[0], y[-1]]
    
    plt.xticks(np.arange(int(min1), int(max1)+1, tickspacing1),fontsize=10*size)
    plt.yticks(np.arange(int(min2), int(max2)+1, tickspacing2),fontsize=10*size)
    #    sigma = 0.99 # this depends on how noisy your data is, play with it!
    data = gaussian_filter((F.T)*0.592-np.min(F.T)*0.592, sigma)
    levels=np.linspace(0,np.max(data)-0.5,num=10)
    plt.contour(data,colors='black',linestyles='solid',alpha=0.7,cmap=None, cbar=True, levels=levels,extent=extent)
    plt.contourf(data,alpha=0.5,cmap='jet', cbar=True,levels=levels,extent=extent)
    if title:
        plt.title(title, fontsize = 20*size, y=1.02)
    plt.subplots_adjust(bottom=0.1, right=0.8, top=0.8)
    cax = plt.axes([0.81, 0.1, 0.02, 0.7])
    plt.colorbar(cax=cax, format='%.1f').set_label('Free energy (kcal/mol)', fontsize=10*size, labelpad=5, y= 0.5)
    cax.axes.tick_params(labelsize=10*size)
