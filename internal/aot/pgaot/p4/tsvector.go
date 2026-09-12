package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_reset_tsvector_parser(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	return
}
func F_tsvector_delete_by_indices(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	v4 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(2) <= l2 {
		F_pg_qsort(m, l1, l2, int32(4), int32(1537))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v29 = int32(1)
			v31 = v4
			for {
				v41 = int32(2)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+v29<<(uint(v41)%32))))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+v31<<(uint(v41)%32))))
				if v44 == v48 {
					v58 = v31
				} else {
					v51 = v31 + int32(1)
					if v29 == v51 {
						v58 = v29
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1+v51<<(uint(int32(2))%32)))) = v44
						v58 = v51
					}
				}
				v60 = v29 + int32(1)
				if v60 != l2 {
					v29 = v60
					v31 = v58
					continue
				} else {
					break
				}
				break
			}
			v66 = v58 + int32(1)
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v82 = F_palloc0(m, int32(base.Ui32(v79)>>(uint(int32(2))%32)))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v85 = v84 - v66
				*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v85
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v87 <= int32(0) {
					v215 = v85
					v217 = int32(32)
				} else {
					v91 = int32(8)
					v92 = l0 + v91
					v93 = int32(2)
					v97 = v82 + v91
					v100 = v97 + v85<<(uint(v93)%32)
					v101 = int32(0)
					v107 = v101
					v108 = v87
					v109 = v101
					v111 = v101
					v113 = v4
					for {
						if v66 <= v113 {
							v130 = v92 + v109<<(uint(int32(2))%32)
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v138 = int32(base.Ui32(v131)>>(uint(int32(1))%32)) & int32(2047)
							if v138 != 0 {
								v139 = F__emscripten_memcpy_bulkmem(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v131)>>(uint(int32(12))%32)), v138)
								mBase = m.M
							} else {
							}
							v143 = v97 + v111<<(uint(int32(2))%32)
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v145 = int32(1)
							v146 = v144 & v145
							v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							*(*int32)(unsafe.Add(mBase, uint32(v143))) = v146 | v147&int32(-2)
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							*(*int32)(unsafe.Add(mBase, uint32(v143))) = v152&int32(4094) | v107<<(uint(int32(12))%32) | v146
							v160 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v164 = int32(base.Ui32(v160)>>(uint(v145)%32)) & int32(2047)
							v165 = v164 + v107
							if v160&v145 != 0 {
								v168 = int32(1)
								v171 = (v165 + v168) & int32(-2)
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v174 = int32(2)
								v184 = v92 + v173<<(uint(v174)%32) + (int32(base.Ui32(v160)>>(uint(int32(12))%32))+v164+v168)&int32(4194302)
								v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
								v189 = v185<<(uint(v168)%32) + v174
								if v189 != 0 {
									v190 = F__emscripten_memcpy_bulkmem(m, v100+v171, v184, v189)
									mBase = m.M
								} else {
								}
								v193 = v171 + v189
							} else {
								v193 = v165
							}
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v198 = v193
							v199 = v197
							v201 = v111 + int32(1)
							v202 = v113
						} else {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(int32(2))%32))))
							if v109 != v123 {
								v130 = v92 + v109<<(uint(int32(2))%32)
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
								v138 = int32(base.Ui32(v131)>>(uint(int32(1))%32)) & int32(2047)
								if v138 != 0 {
									v139 = F__emscripten_memcpy_bulkmem(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v131)>>(uint(int32(12))%32)), v138)
									mBase = m.M
								} else {
								}
								v143 = v97 + v111<<(uint(int32(2))%32)
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
								v145 = int32(1)
								v146 = v144 & v145
								v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
								*(*int32)(unsafe.Add(mBase, uint32(v143))) = v146 | v147&int32(-2)
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
								*(*int32)(unsafe.Add(mBase, uint32(v143))) = v152&int32(4094) | v107<<(uint(int32(12))%32) | v146
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
								v164 = int32(base.Ui32(v160)>>(uint(v145)%32)) & int32(2047)
								v165 = v164 + v107
								if v160&v145 != 0 {
									v168 = int32(1)
									v171 = (v165 + v168) & int32(-2)
									v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v174 = int32(2)
									v184 = v92 + v173<<(uint(v174)%32) + (int32(base.Ui32(v160)>>(uint(int32(12))%32))+v164+v168)&int32(4194302)
									v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
									v189 = v185<<(uint(v168)%32) + v174
									if v189 != 0 {
										v190 = F__emscripten_memcpy_bulkmem(m, v100+v171, v184, v189)
										mBase = m.M
									} else {
									}
									v193 = v171 + v189
								} else {
									v193 = v165
								}
								v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v198 = v193
								v199 = v197
								v201 = v111 + int32(1)
								v202 = v113
							} else {
								v198 = v107
								v199 = v108
								v201 = v111
								v202 = v113 + int32(1)
							}
						}
						v205 = v109 + int32(1)
						if v205 < v199 {
							v107 = v198
							v108 = v199
							v109 = v205
							v111 = v201
							v113 = v202
							continue
						} else {
							break
						}
						break
					}
					v211 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
					v215 = v211
					v217 = v198<<(uint(int32(2))%32) + int32(32)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v82))) = v215<<(uint(int32(4))%32) + v217
				return v82
			}
		}
	} else {
		v66 = l2
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v82 = F_palloc0(m, int32(base.Ui32(v79)>>(uint(int32(2))%32)))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v85 = v84 - v66
			*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v85
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v87 <= int32(0) {
				v215 = v85
				v217 = int32(32)
			} else {
				v91 = int32(8)
				v92 = l0 + v91
				v93 = int32(2)
				v97 = v82 + v91
				v100 = v97 + v85<<(uint(v93)%32)
				v101 = int32(0)
				v107 = v101
				v108 = v87
				v109 = v101
				v111 = v101
				v113 = v4
				for {
					if v66 <= v113 {
						v130 = v92 + v109<<(uint(int32(2))%32)
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						v138 = int32(base.Ui32(v131)>>(uint(int32(1))%32)) & int32(2047)
						if v138 != 0 {
							v139 = F__emscripten_memcpy_bulkmem(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v131)>>(uint(int32(12))%32)), v138)
							mBase = m.M
						} else {
						}
						v143 = v97 + v111<<(uint(int32(2))%32)
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						v145 = int32(1)
						v146 = v144 & v145
						v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
						*(*int32)(unsafe.Add(mBase, uint32(v143))) = v146 | v147&int32(-2)
						v152 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						*(*int32)(unsafe.Add(mBase, uint32(v143))) = v152&int32(4094) | v107<<(uint(int32(12))%32) | v146
						v160 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						v164 = int32(base.Ui32(v160)>>(uint(v145)%32)) & int32(2047)
						v165 = v164 + v107
						if v160&v145 != 0 {
							v168 = int32(1)
							v171 = (v165 + v168) & int32(-2)
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v174 = int32(2)
							v184 = v92 + v173<<(uint(v174)%32) + (int32(base.Ui32(v160)>>(uint(int32(12))%32))+v164+v168)&int32(4194302)
							v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
							v189 = v185<<(uint(v168)%32) + v174
							if v189 != 0 {
								v190 = F__emscripten_memcpy_bulkmem(m, v100+v171, v184, v189)
								mBase = m.M
							} else {
							}
							v193 = v171 + v189
						} else {
							v193 = v165
						}
						v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v198 = v193
						v199 = v197
						v201 = v111 + int32(1)
						v202 = v113
					} else {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(int32(2))%32))))
						if v109 != v123 {
							v130 = v92 + v109<<(uint(int32(2))%32)
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v138 = int32(base.Ui32(v131)>>(uint(int32(1))%32)) & int32(2047)
							if v138 != 0 {
								v139 = F__emscripten_memcpy_bulkmem(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v131)>>(uint(int32(12))%32)), v138)
								mBase = m.M
							} else {
							}
							v143 = v97 + v111<<(uint(int32(2))%32)
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v145 = int32(1)
							v146 = v144 & v145
							v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							*(*int32)(unsafe.Add(mBase, uint32(v143))) = v146 | v147&int32(-2)
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							*(*int32)(unsafe.Add(mBase, uint32(v143))) = v152&int32(4094) | v107<<(uint(int32(12))%32) | v146
							v160 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							v164 = int32(base.Ui32(v160)>>(uint(v145)%32)) & int32(2047)
							v165 = v164 + v107
							if v160&v145 != 0 {
								v168 = int32(1)
								v171 = (v165 + v168) & int32(-2)
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v174 = int32(2)
								v184 = v92 + v173<<(uint(v174)%32) + (int32(base.Ui32(v160)>>(uint(int32(12))%32))+v164+v168)&int32(4194302)
								v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
								v189 = v185<<(uint(v168)%32) + v174
								if v189 != 0 {
									v190 = F__emscripten_memcpy_bulkmem(m, v100+v171, v184, v189)
									mBase = m.M
								} else {
								}
								v193 = v171 + v189
							} else {
								v193 = v165
							}
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v198 = v193
							v199 = v197
							v201 = v111 + int32(1)
							v202 = v113
						} else {
							v198 = v107
							v199 = v108
							v201 = v111
							v202 = v113 + int32(1)
						}
					}
					v205 = v109 + int32(1)
					if v205 < v199 {
						v107 = v198
						v108 = v199
						v109 = v205
						v111 = v201
						v113 = v202
						continue
					} else {
						break
					}
					break
				}
				v211 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
				v215 = v211
				v217 = v198<<(uint(int32(2))%32) + int32(32)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v82))) = v215<<(uint(int32(4))%32) + v217
			return v82
		}
	}
}
func F_tsvector_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = l0 + int32(28)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v33 = int32(base.Ui32(v31) >> (uint(v29) % 32))
	if base.Ui32(v30) < base.Ui32(v33) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != v7 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v223 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if base.Ui32(v33) < base.Ui32(v30) {
		v199 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v223 = v199
	goto L4
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v223 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v39 < v38 {
		v199 = v36
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v38 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v223 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v47 = v14 + v46
	v48 = int32(2)
	v50 = v47 + v39<<(uint(v48)%32)
	v52 = v7 + v46
	v55 = v52 + v38<<(uint(v48)%32)
	v63 = v47
	v64 = v52
	v68 = int32(0)
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = int32(1)
	v72 = v70 & v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v73 & v71
	if v72 != v75 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v199 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v75) < base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	v83 = int32(2047)
	v84 = int32(base.Ui32(v73)>>(uint(v81)%32)) & v83
	v85 = int32(12)
	v86 = int32(base.Ui32(v73) >> (uint(v85) % 32))
	v88 = int32(base.Ui32(v70) >> (uint(v85) % 32))
	v92 = int32(base.Ui32(v70)>>(uint(v81)%32)) & v83
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v80 = int32(-1)
	goto L24
L23:
	;
	v80 = int32(1)
	goto L24
L24:
	;
	v223 = v80
	goto L4
L25:
	;
	if v72 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v223 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v98 = base.B2i32(base.Ui32(v92) < base.Ui32(v84))
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v92
	goto L34
L33:
	;
	v99 = v84
	goto L34
L34:
	;
	v100 = F_memcmp(m, v88+v55, v86+v50, v99)
	mBase = m.M
	if v100 != 0 {
		v199 = v100
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v92 == v84 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(-1)
	goto L39
L38:
	;
	v104 = int32(1)
	goto L39
L39:
	;
	v223 = v104
	goto L4
L40:
	;
	v223 = int32(-1)
	goto L4
L41:
	;
	v187 = int32(4)
	v193 = v68 + int32(1)
	if v193 != v38 {
		v63 = v63 + v187
		v64 = v64 + v187
		v68 = v193
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v113 = int32(1)
	v115 = int32(4194302)
	v117 = v50 + (v84+v86+v113)&v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v124 = v55 + (v92+v88+v113)&v115
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v118 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v132 = v124
	v133 = v117
	v134 = int32(0)
	goto L51
L44:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v118) < base.Ui32(v125) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v131 = int32(-1)
	goto L50
L49:
	;
	v131 = int32(1)
	goto L50
L50:
	;
	v223 = v131
	goto L4
L51:
	;
	v146 = int32(2)
	v147 = v132 + v146
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(16383)
	v150 = v148 & v149
	v152 = v133 + v146
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & v149
	if v150 != v155 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v155) < base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v161 = int32(14)
	v162 = int32(base.Ui32(v148) >> (uint(v161) % 32))
	v164 = int32(base.Ui32(v153) >> (uint(v161) % 32))
	if v162 == v164 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v160 = int32(-1)
	goto L58
L57:
	;
	v160 = int32(1)
	goto L58
L58:
	;
	v223 = v160
	goto L4
L59:
	;
	v167 = v134 + int32(1)
	if v167 == v125 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v132 = v147
	v133 = v152
	v134 = v167
	goto L51
L63:
	;
	v172 = int32(-1)
	goto L65
L64:
	;
	v172 = int32(1)
	goto L65
L65:
	;
	v199 = v172
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v7)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v228 != v14 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.B2i32(v223 <= int32(0))
L74:
	;
	goto L73
}
func F_tsvector_to_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v16 = F_palloc(m, v13<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if int32(0) < v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v9 + int32(8)
	v25 = v18
	v26 = int32(0)
	goto L7
L5:
	;
	v54 = v18
	goto L6
L6:
	;
	v60 = F_construct_array_builtin(m, v16, v54, int32(25))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L11
	}
L7:
	;
	v30 = int32(2)
	v31 = v26 << (uint(v30) % 32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22+v31)))
	v45 = F_cstring_to_text_with_len(m, v22+v25<<(uint(v30)%32)+int32(base.Ui32(v37)>>(uint(int32(12))%32)), int32(base.Ui32(v37)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v54 = v50
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v31))) = v45
	v49 = v26 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v49 < v50 {
		v25 = v50
		v26 = v49
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_pfree(m, v16)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v64 != v9 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_pfree(m, v9)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	return v60
L16:
	;
	goto L15
}
func F_tsvector_unnest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L38
	}
L2:
	;
	v20 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	goto L15
L5:
	;
	return int32(0)
L6:
	;
	v24 = int32(4486928)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v27
	v30 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v30
	F_TupleDescInitEntry(m, v30, int32(1), int32(374016), int32(25), int32(-1), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v30, int32(2), int32(138238), int32(1005), int32(-1), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v30, int32(3), int32(123314), int32(1009), int32(-1), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v57 = F_get_call_result_type(m, l0, int32(0), v14+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v57 != int32(1) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v64 = F_pg_detoast_datum_copy(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v64
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v25
	goto L4
L14:
	;
	m.G0 = v14 + int32(32)
	return v244
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if base.Ui64(v74) < base.Ui64(base.I64_extend_i32_s(v76)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+30)) = uint8(v79)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+28)) = uint16(v79)
	v85 = v75 + int32(8)
	v86 = int32(2)
	v92 = v85 + base.I32_wrap_i64(v74)<<(uint(v86)%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v101 = F_cstring_to_text_with_len(m, v85+v76<<(uint(v86)%32)+int32(base.Ui32(v93)>>(uint(int32(12))%32)), int32(base.Ui32(v93)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L37
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v101
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v104&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	v213 = F_heap_form_tuple(m, v208, v14+int32(16), v14+int32(28))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L35
	}
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v108 = int32(2)
	v111 = int32(1)
	v122 = v85 + v107<<(uint(v108)%32) + (int32(base.Ui32(v104)>>(uint(v111)%32))&int32(2047)+int32(base.Ui32(v104)>>(uint(int32(12))%32))+v111)&int32(4194302)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	v126 = F_palloc(m, v123<<(uint(v108)%32))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v195 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+29)) = uint16(v195)
	goto L20
L24:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	v131 = F_palloc(m, v128<<(uint(int32(2))%32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	if v133 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v138 = int32(0)
	goto L29
L27:
	;
	v179 = v79
	goto L28
L28:
	;
	v187 = F_construct_array_builtin(m, v126, v179, int32(21))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L33
	}
L29:
	;
	v149 = v138 << (uint(int32(2)) % 32)
	v151 = int32(1)
	v153 = v122 + int32(2) + v138<<(uint(v151)%32)
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153))))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v149))) = v154 & int32(16383)
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153))))
	v162 = int32(68) - int32(base.Ui32(v159)>>(uint(int32(14))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)) = uint8(v162)
	v168 = F_cstring_to_text_with_len(m, v14+int32(15), v151)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v179 = v173
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149+v131))) = v168
	v172 = v138 + int32(1)
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	if base.Ui32(v172) < base.Ui32(v173) {
		v138 = v172
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v187
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	v192 = F_construct_array_builtin(m, v131, v190, int32(25))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v192
	goto L20
L35:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = v215 + int64(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+20)) = int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	v223 = F_HeapTupleHeaderGetDatum(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v244 = v223
	goto L14
L37:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+20)) = int32(2)
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v230)
	v244 = int32(0)
	goto L14
L38:
	;
	F_errmsg_internal(m, int32(365102), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(492630), int32(653), int32(77061))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
