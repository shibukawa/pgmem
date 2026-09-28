package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleAlias(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		v7 = v5 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v29 = v19
				v31 = v17
				v33 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v29) {
						v38 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v38
						v41 = int32(8)
					} else {
						v41 = v29
					}
					v43 = int32(1024) - v41
					if base.Ui32(v31) < base.Ui32(v43) {
						v45 = v31
					} else {
						v45 = v43
					}
					if v45 != 0 {
						base.MemoryCopy(m, v41+v18, v33, v45)
					} else {
					}
					v49 = v41 + v45
					v50 = v31 - v45
					if v50 != 0 {
						v29 = v49
						v31 = v50
						v33 = v45 + v33
						continue
					} else {
						break
					}
					break
				}
				v59 = v49
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v59 = v53 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
			v70 = v59
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v70) < base.Ui32(v7) {
			v80 = v4
			v81 = v7
			v82 = v70
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v82) {
					v91 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v91
					v94 = int32(8)
				} else {
					v94 = v82
				}
				v96 = int32(1024) - v94
				if base.Ui32(v81) < base.Ui32(v96) {
					v98 = v81
				} else {
					v98 = v96
				}
				if v98 != 0 {
					base.MemoryCopy(m, v94+v75, v80, v98)
				} else {
				}
				v102 = v94 + v98
				v103 = v81 - v98
				if v103 != 0 {
					v80 = v80 + v98
					v81 = v103
					v82 = v102
					continue
				} else {
					break
				}
				break
			}
			v111 = v102
		} else {
			if v7 != 0 {
				base.MemoryCopy(m, v70+v75, v4, v7)
			} else {
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v111 = v106 + v7
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v117 + int32(1)
	}
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		return
	} else {
		return
	}
}
func F__jumbleAlterTableSpaceOptionsStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		v7 = v5 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v29 = v19
				v31 = v17
				v33 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v29) {
						v38 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v38
						v41 = int32(8)
					} else {
						v41 = v29
					}
					v43 = int32(1024) - v41
					if base.Ui32(v31) < base.Ui32(v43) {
						v45 = v31
					} else {
						v45 = v43
					}
					if v45 != 0 {
						base.MemoryCopy(m, v41+v18, v33, v45)
					} else {
					}
					v49 = v41 + v45
					v50 = v31 - v45
					if v50 != 0 {
						v29 = v49
						v31 = v50
						v33 = v45 + v33
						continue
					} else {
						break
					}
					break
				}
				v59 = v49
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v59 = v53 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
			v70 = v59
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(int32(1024)-v70) < base.Ui32(v7) {
			v80 = v4
			v81 = v7
			v82 = v70
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v82) {
					v91 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v75))) = v91
					v94 = int32(8)
				} else {
					v94 = v82
				}
				v96 = int32(1024) - v94
				if base.Ui32(v81) < base.Ui32(v96) {
					v98 = v81
				} else {
					v98 = v96
				}
				if v98 != 0 {
					base.MemoryCopy(m, v94+v75, v80, v98)
				} else {
				}
				v102 = v94 + v98
				v103 = v81 - v98
				if v103 != 0 {
					v80 = v80 + v98
					v81 = v103
					v82 = v102
					continue
				} else {
					break
				}
				break
			}
			v111 = v102
		} else {
			if v7 != 0 {
				base.MemoryCopy(m, v70+v75, v4, v7)
			} else {
			}
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v111 = v106 + v7
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v111
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v117 + int32(1)
	}
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		return
	} else {
		v125 = l1 + int32(12)
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v131 == int32(0) {
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v185 = v134
		} else {
			v135 = int32(4)
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v137-int32(1021)) < base.Ui32(v135) {
				v146 = v137
				v147 = v135
				v150 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v146) {
						v155 = F_hash_bytes_extended(m, v136, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v136))) = v155
						v158 = int32(8)
					} else {
						v158 = v146
					}
					v160 = int32(1024) - v158
					if base.Ui32(v147) < base.Ui32(v160) {
						v162 = v147
					} else {
						v162 = v160
					}
					if v162 != 0 {
						base.MemoryCopy(m, v158+v136, v150, v162)
					} else {
					}
					v166 = v158 + v162
					v167 = v147 - v162
					if v167 != 0 {
						v146 = v166
						v147 = v167
						v150 = v162 + v150
						continue
					} else {
						break
					}
					break
				}
				v175 = v166
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v137+v136))) = v131
				v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v175 = v170 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v175
			v185 = v175
		}
		v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v185 != int32(1024) {
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
			*(*uint8)(unsafe.Add(mBase, uint32(v185+v190))) = uint8(v194)
			v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v196 + int32(1)
		} else {
			v202 = F_hash_bytes_extended(m, v190, int32(1024), int64(0))
			mBase = m.M
			*(*int64)(unsafe.Add(mBase, uint32(v190))) = v202
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
			*(*uint8)(unsafe.Add(mBase, uint32(v190)+8)) = uint8(v204)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
		}
		return
	}
}
func F__jumbleJsonReturning(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int64
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v7 = l1 + int32(8)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v67 = v16
		} else {
			v17 = int32(4)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v19-int32(1021)) < base.Ui32(v17) {
				v28 = v19
				v30 = v17
				v32 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v28) {
						v37 = F_hash_bytes_extended(m, v18, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v37
						v40 = int32(8)
					} else {
						v40 = v28
					}
					v42 = int32(1024) - v40
					if base.Ui32(v30) < base.Ui32(v42) {
						v44 = v30
					} else {
						v44 = v42
					}
					if v44 != 0 {
						base.MemoryCopy(m, v40+v18, v32, v44)
					} else {
					}
					v48 = v40 + v44
					v49 = v30 - v44
					if v49 != 0 {
						v28 = v48
						v30 = v49
						v32 = v44 + v32
						continue
					} else {
						break
					}
					break
				}
				v57 = v48
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19+v18))) = v13
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v57 = v52 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v57
			v67 = v57
		}
		v72 = int32(4)
		v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v67-int32(1021)) < base.Ui32(v72) {
			v79 = v7
			v80 = v67
			v82 = v72
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v80) {
					v89 = F_hash_bytes_extended(m, v73, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v73))) = v89
					v92 = int32(8)
				} else {
					v92 = v80
				}
				v94 = int32(1024) - v92
				if base.Ui32(v82) < base.Ui32(v94) {
					v96 = v82
				} else {
					v96 = v94
				}
				if v96 != 0 {
					base.MemoryCopy(m, v92+v73, v79, v96)
				} else {
				}
				v100 = v92 + v96
				v101 = v82 - v96
				if v101 != 0 {
					v79 = v79 + v96
					v80 = v100
					v82 = v101
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100
		} else {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*int32)(unsafe.Add(mBase, uint32(v67+v73))) = v104
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106 + int32(4)
		}
		v118 = l1 + int32(12)
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v124 == int32(0) {
			v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v178 = v127
		} else {
			v128 = int32(4)
			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v130-int32(1021)) < base.Ui32(v128) {
				v139 = v130
				v141 = v128
				v143 = l0 + int32(28)
				for {
					if base.Ui32(int32(1024)) <= base.Ui32(v139) {
						v148 = F_hash_bytes_extended(m, v129, int32(1024), int64(0))
						mBase = m.M
						*(*int64)(unsafe.Add(mBase, uint32(v129))) = v148
						v151 = int32(8)
					} else {
						v151 = v139
					}
					v153 = int32(1024) - v151
					if base.Ui32(v141) < base.Ui32(v153) {
						v155 = v141
					} else {
						v155 = v153
					}
					if v155 != 0 {
						base.MemoryCopy(m, v151+v129, v143, v155)
					} else {
					}
					v159 = v151 + v155
					v160 = v141 - v155
					if v160 != 0 {
						v139 = v159
						v141 = v160
						v143 = v155 + v143
						continue
					} else {
						break
					}
					break
				}
				v168 = v159
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v130+v129))) = v124
				v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v168 = v163 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168
			v178 = v168
		}
		v183 = int32(4)
		v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui32(v178-int32(1021)) < base.Ui32(v183) {
			v190 = v118
			v191 = v178
			v193 = v183
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v191) {
					v200 = F_hash_bytes_extended(m, v184, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v184))) = v200
					v203 = int32(8)
				} else {
					v203 = v191
				}
				v205 = int32(1024) - v203
				if base.Ui32(v193) < base.Ui32(v205) {
					v207 = v193
				} else {
					v207 = v205
				}
				if v207 != 0 {
					base.MemoryCopy(m, v203+v184, v190, v207)
				} else {
				}
				v211 = v203 + v207
				v212 = v193 - v207
				if v212 != 0 {
					v190 = v190 + v207
					v191 = v211
					v193 = v212
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v211
		} else {
			v215 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
			*(*int32)(unsafe.Add(mBase, uint32(v178+v184))) = v215
			v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v217 + int32(4)
		}
		return
	}
}
func F__jumblePartitionCmd(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v10 = l1 + int32(12)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v70 = v19
			} else {
				v20 = int32(4)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if base.Ui32(v22-int32(1021)) < base.Ui32(v20) {
					v31 = v22
					v32 = v20
					v35 = l0 + int32(28)
					for {
						if base.Ui32(int32(1024)) <= base.Ui32(v31) {
							v40 = F_hash_bytes_extended(m, v21, int32(1024), int64(0))
							mBase = m.M
							*(*int64)(unsafe.Add(mBase, uint32(v21))) = v40
							v43 = int32(8)
						} else {
							v43 = v31
						}
						v45 = int32(1024) - v43
						if base.Ui32(v32) < base.Ui32(v45) {
							v47 = v32
						} else {
							v47 = v45
						}
						if v47 != 0 {
							base.MemoryCopy(m, v43+v21, v35, v47)
						} else {
						}
						v51 = v43 + v47
						v52 = v32 - v47
						if v52 != 0 {
							v31 = v51
							v32 = v52
							v35 = v47 + v35
							continue
						} else {
							break
						}
						break
					}
					v60 = v51
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22+v21))) = v16
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v60 = v55 + int32(4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60
				v70 = v60
			}
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v70 != int32(1024) {
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				*(*uint8)(unsafe.Add(mBase, uint32(v70+v75))) = uint8(v79)
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 + int32(1)
			} else {
				v87 = F_hash_bytes_extended(m, v75, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v75))) = v87
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)) = uint8(v89)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
			}
			return
		}
	}
}
