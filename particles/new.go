package particles

import "project-particles/config"
import "math/rand"
import "time"

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

/*
func NewSystem() System {
	var NbParticule int=config.General.InitNumParticles
	var tab []Particle=[]Particle{}
	var P Particle
	var x float64 = float64(rand.Intn(config.General.WindowSizeX))
	var y float64 = float64(rand.Intn(config.General.WindowSizeY))

		for i:=0;i<=NbParticule;i++{
			P =Particle{
				PositionX: x,
				PositionY: y,
				ScaleX:    1, ScaleY: 1,
				ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
				Opacity: 1,
			}
			tab=append(tab,P)
		}
		
		return System{Content: tab}

	
}*/

func NewSystem() System {
	rand.Seed(time.Now().UnixNano())
	var a System

	if config.General.RandomSpawn==true{
		var Nbparticules = config.General.InitNumParticles
    	var tab []Particle= []Particle{}
    
    	for i:=1; i<=Nbparticules;i++{
        
			var xu float64=float64(rand.Intn((config.General.WindowSizeX)))
			var yu float64=float64(rand.Intn((config.General.WindowSizeY)))
			var p =Particle{
				PositionX: xu,
				PositionY: yu,
				ScaleX:    1, ScaleY: 1,
				ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
				Opacity: 1,
			}
			
			tab = append(tab,p)
    	}
		a=System{Content: tab}
		
	}
	if config.General.RandomSpawn==false{
		
		var Nbparticules = config.General.InitNumParticles
    	var tab []Particle= []Particle{}
    
    	for i:=1; i<=Nbparticules;i++{
        
			var xu float64=float64((config.General.SpawnX))
			var yu float64=float64((config.General.SpawnY))
			var p =Particle{
				PositionX: xu,
				PositionY: yu,
				ScaleX:    1, ScaleY: 1,
				ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
				Opacity: 1,
			}
			
			tab = append(tab,p)
    	}
		a=System{Content: tab}
	}
	return a
}