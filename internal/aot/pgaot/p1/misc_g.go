package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetConflictingVirtualXIDs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
	if v15 == v3 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v24 = F_emscripten_builtin_malloc(m, v19<<(uint(int32(3))%32)+int32(8))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1])) = v24
		if v24 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v133 = m.ExcPending
			if v133 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(_a_F_GetConflictingVirtualXIDs_0))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_GetConflictingVirtualXIDs_1), int32(0))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetConflictingVirtualXIDs_2), int32(3400), int32(_a_F_GetConflictingVirtualXIDs_3))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
			v34 = F_LWLockAcquire(m, v30+int32(512), int32(1))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if int32(0) < v38 {
					v44 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
					v52 = v38
					v53 = v3
					v54 = v3
					for {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v53<<(uint(int32(2))%32))))
						v64 = v46 + v61*int32(768)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
						if v65 == int32(0) {
							v100 = v52
							v101 = v54
						} else {
							if l1 != 0 {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
								if v68 != l1 {
									v100 = v52
									v101 = v54
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
									if l0 == int32(0) {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
										if v86 == int32(0) {
											v100 = v52
											v101 = v54
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
											v92 = v44 + v54<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
											*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											v100 = v97
											v101 = v54 + int32(1)
										}
									} else {
										if v70 == int32(0) {
											v100 = v52
											v101 = v54
										} else {
											v75 = int32(3)
											if base.B2i32(base.Ui32(l0) < base.Ui32(v75))|base.B2i32(base.Ui32(v70) < base.Ui32(v75)) == int32(0) {
												if v70-l0 <= int32(0) {
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
													if v86 == int32(0) {
														v100 = v52
														v101 = v54
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
														v92 = v44 + v54<<(uint(int32(3))%32)
														*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
														*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v100 = v97
														v101 = v54 + int32(1)
													}
												} else {
													v100 = v52
													v101 = v54
												}
											} else {
												if base.Ui32(l0) < base.Ui32(v70) {
													v100 = v52
													v101 = v54
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
													if v86 == int32(0) {
														v100 = v52
														v101 = v54
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
														v92 = v44 + v54<<(uint(int32(3))%32)
														*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
														*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v100 = v97
														v101 = v54 + int32(1)
													}
												}
											}
										}
									}
								}
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
								if l0 == int32(0) {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
									if v86 == int32(0) {
										v100 = v52
										v101 = v54
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
										v92 = v44 + v54<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
										*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v100 = v97
										v101 = v54 + int32(1)
									}
								} else {
									if v70 == int32(0) {
										v100 = v52
										v101 = v54
									} else {
										v75 = int32(3)
										if base.B2i32(base.Ui32(l0) < base.Ui32(v75))|base.B2i32(base.Ui32(v70) < base.Ui32(v75)) == int32(0) {
											if v70-l0 <= int32(0) {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
												if v86 == int32(0) {
													v100 = v52
													v101 = v54
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
													v92 = v44 + v54<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													v100 = v97
													v101 = v54 + int32(1)
												}
											} else {
												v100 = v52
												v101 = v54
											}
										} else {
											if base.Ui32(l0) < base.Ui32(v70) {
												v100 = v52
												v101 = v54
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
												if v86 == int32(0) {
													v100 = v52
													v101 = v54
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
													v92 = v44 + v54<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													v100 = v97
													v101 = v54 + int32(1)
												}
											}
										}
									}
								}
							}
						}
						v103 = v53 + int32(1)
						if v103 < v100 {
							v52 = v100
							v53 = v103
							v54 = v101
							continue
						} else {
							break
						}
						break
					}
					v112 = v101
				} else {
					v112 = v3
				}
				v117 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
				F_LWLockRelease(m, v117+int32(512))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					v123 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
					*(*int64)(unsafe.Add(mBase, uint32(v123+v112<<(uint(int32(3))%32)))) = int64(4294967295)
					return v123
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
		v34 = F_LWLockAcquire(m, v30+int32(512), int32(1))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if int32(0) < v38 {
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
				v52 = v38
				v53 = v3
				v54 = v3
				for {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(36)+v53<<(uint(int32(2))%32))))
					v64 = v46 + v61*int32(768)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
					if v65 == int32(0) {
						v100 = v52
						v101 = v54
					} else {
						if l1 != 0 {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
							if v68 != l1 {
								v100 = v52
								v101 = v54
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
								if l0 == int32(0) {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
									if v86 == int32(0) {
										v100 = v52
										v101 = v54
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
										v92 = v44 + v54<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
										*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v100 = v97
										v101 = v54 + int32(1)
									}
								} else {
									if v70 == int32(0) {
										v100 = v52
										v101 = v54
									} else {
										v75 = int32(3)
										if base.B2i32(base.Ui32(l0) < base.Ui32(v75))|base.B2i32(base.Ui32(v70) < base.Ui32(v75)) == int32(0) {
											if v70-l0 <= int32(0) {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
												if v86 == int32(0) {
													v100 = v52
													v101 = v54
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
													v92 = v44 + v54<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													v100 = v97
													v101 = v54 + int32(1)
												}
											} else {
												v100 = v52
												v101 = v54
											}
										} else {
											if base.Ui32(l0) < base.Ui32(v70) {
												v100 = v52
												v101 = v54
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
												if v86 == int32(0) {
													v100 = v52
													v101 = v54
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
													v92 = v44 + v54<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													v100 = v97
													v101 = v54 + int32(1)
												}
											}
										}
									}
								}
							}
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+52))
							if l0 == int32(0) {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
								if v86 == int32(0) {
									v100 = v52
									v101 = v54
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
									v92 = v44 + v54<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
									*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v100 = v97
									v101 = v54 + int32(1)
								}
							} else {
								if v70 == int32(0) {
									v100 = v52
									v101 = v54
								} else {
									v75 = int32(3)
									if base.B2i32(base.Ui32(l0) < base.Ui32(v75))|base.B2i32(base.Ui32(v70) < base.Ui32(v75)) == int32(0) {
										if v70-l0 <= int32(0) {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
											if v86 == int32(0) {
												v100 = v52
												v101 = v54
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
												v92 = v44 + v54<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
												*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												v100 = v97
												v101 = v54 + int32(1)
											}
										} else {
											v100 = v52
											v101 = v54
										}
									} else {
										if base.Ui32(l0) < base.Ui32(v70) {
											v100 = v52
											v101 = v54
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
											if v86 == int32(0) {
												v100 = v52
												v101 = v54
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
												v92 = v44 + v54<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v86
												*(*int32)(unsafe.Add(mBase, uint32(v92))) = v89
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												v100 = v97
												v101 = v54 + int32(1)
											}
										}
									}
								}
							}
						}
					}
					v103 = v53 + int32(1)
					if v103 < v100 {
						v52 = v100
						v53 = v103
						v54 = v101
						continue
					} else {
						break
					}
					break
				}
				v112 = v101
			} else {
				v112 = v3
			}
			v117 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
			F_LWLockRelease(m, v117+int32(512))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int32(0)
			} else {
				v123 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
				*(*int64)(unsafe.Add(mBase, uint32(v123+v112<<(uint(int32(3))%32)))) = int64(4294967295)
				return v123
			}
		}
	}
}
func F_GetHugePageSize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetHugePageSize[0]))
	if v4 != 0 {
		v8 = v4 << (uint(int32(10)) % 32)
	} else {
		v8 = int32(_a_F_GetHugePageSize_0)
	}
	if v8 != 0 {
	} else {
	}
	if l0 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v8
	} else {
	}
	return
}
func F_GetIncludedPublicationRelations(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_get_publication_relations(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_GetIntoRelEFlags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	return v2 << (uint(int32(6)) % 32) & int32(64)
}
func F_GetLockmodeName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v3 = int32(2)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v3)%32))+uint32(_c_F_GetLockmodeName[0])))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6+l1<<(uint(v3)%32))))
	return v10
}
func F_GetScanItems(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	F_tuplesort_reset(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v29 <= v28 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	F_tuplesort_performsort(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	v35 = v28
	v36 = v29
	v43 = int32(0)
	goto L5
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v47 <= v43 {
		goto L3
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v50 = v35 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v50
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(int32(2))%32)+v54)))
	if v56 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v61 = v56
	goto L11
L9:
	;
	v251 = v50
	v252 = v36
	goto L10
L10:
	;
	if v251 < v252 {
		v35 = v251
		v36 = v252
		v43 = v43 + int32(1)
		goto L5
	} else {
		goto L51
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v79 = F_ReadBufferExtended(m, v75, v76, v61, v76, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	v251 = v246
	v252 = v245
	goto L10
L13:
	;
	F_LockBufferInternal(m, v79, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v79 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+16)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v101+v238)))
	F_UnlockReleaseBuffer(m, v79)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L49
	}
L16:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+12)))
	if base.Ui32(v102) < base.Ui32(int32(25)) {
		goto L15
	} else {
		goto L20
	}
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_GetScanItems[0]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v79^int32(-1))<<(uint(int32(2))%32))))
	v101 = v93
	goto L16
L18:
	;
	goto L19
L19:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_GetScanItems[1]))
	v101 = v95 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L16
L20:
	;
	v110 = int32(base.Ui32(v102+int32(_a_F_GetScanItems_0))>>(uint(int32(2))%32)) & int32(_a_F_GetScanItems_1)
	if v110 == int32(0) {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v120 = int32(1)
	goto L22
L22:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v101+int32(20)+v120&int32(_a_F_GetScanItems_1)<<(uint(int32(2))%32))))
	v140 = v101 + v137&int32(_a_F_GetScanItems_2)
	v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(v140)+6)))
	if int32(0) <= v141 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L15
L24:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	m.T0[v190].(func(*base.Module, int32))(m, v22)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L44
	}
L25:
	;
	v184 = F_nocache_index_getattr(m, v140, int32(1), v24)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L43
	}
L26:
	;
	v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+28)))
	if v144 < int32(0) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+8)))
	if v177&int32(1) == int32(0) {
		v188 = int64(0)
		goto L24
	} else {
		goto L42
	}
L29:
	;
	v149 = v144 + v140 + int32(8)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+32)))
	if v150 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+30)))
	if base.I32_popcnt(v153) != int32(1) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v188 = base.I64_extend_i32_u(v149)
	goto L24
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L39
	}
L34:
	;
	switch base.I32_ctz(v153) {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	default:
		goto L33
	}
L35:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
	v188 = v161
	goto L24
L36:
	;
	v160 = int64(*(*int32)(unsafe.Add(mBase, uint32(v149))))
	v188 = v160
	goto L24
L37:
	;
	v159 = int64(*(*int16)(unsafe.Add(mBase, uint32(v149))))
	v188 = v159
	goto L24
L38:
	;
	v158 = int64(*(*int8)(unsafe.Add(mBase, uint32(v149))))
	v188 = v158
	goto L24
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v153
	F_errmsg_internal(m, int32(_a_F_GetScanItems_3), v19)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_GetScanItems_4), int32(123), int32(_a_F_GetScanItems_5))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	goto L25
L43:
	;
	v188 = v184
	goto L24
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v196 = m.T0[v195].(func(*base.Module, int32, int32, int64, int64) int64)(m, v193, v194, v188, l1)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v198))) = v196
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v201)
	*(*int64)(unsafe.Add(mBase, uint32(v198)+8)) = base.I64_extend_i32_u(v140)
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)) = uint8(v201)
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	v209 = v207 & int32(_a_F_GetScanItems_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)) = uint16(v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)) = uint16(v212)
	goto L46
L46:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	F_tuplesort_puttupleslot(m, v214, v22)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v218 = v120 + int32(1)
	if base.Ui32(v218&int32(_a_F_GetScanItems_1)) <= base.Ui32(v110) {
		v120 = v218
		goto L22
	} else {
		goto L48
	}
L48:
	;
	goto L23
L49:
	;
	if v240 != int32(-1) {
		v61 = v240
		goto L11
	} else {
		goto L50
	}
L50:
	;
	goto L12
L51:
	;
	goto L6
L52:
	;
	m.G0 = v19 + int32(16)
	return
}
func F_GetStrictOldestNonRemovableTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
		v20 = F_LWLockAcquire(m, v16+int32(384), int32(1))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[3]))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
			F_LWLockRelease(m, v28+int32(384))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				return v26
			}
		}
	} else {
		if l0 != 0 {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+117)))
			if v35 != int32(1) {
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v54 != 0 {
					v72 = F_GetOldestNonRemovableTransactionId(m, l0)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						return v72
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v55 != 0 {
						v72 = F_GetOldestNonRemovableTransactionId(m, l0)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							return v72
						}
					} else {
						v56 = F_GetRunningTransactionData(m)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
							F_LWLockRelease(m, v59+int32(512))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
								F_LWLockRelease(m, v65+int32(384))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
									return v70
								}
							}
						}
					}
				}
			} else {
				v38 = F_GetRunningTransactionData(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
					F_LWLockRelease(m, v41+int32(512))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
						F_LWLockRelease(m, v47+int32(384))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
							return v52
						}
					}
				}
			}
		} else {
			v38 = F_GetRunningTransactionData(m)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
				F_LWLockRelease(m, v41+int32(512))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_GetStrictOldestNonRemovableTransactionId[2]))
					F_LWLockRelease(m, v47+int32(384))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
						return v52
					}
				}
			}
		}
	}
}
func F_GetUserIdAndSecContext(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetUserIdAndSecContext[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_GetUserIdAndSecContext[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
	return
}
func F_GlobalVisTestIsRemovableXid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v16 = v12 + base.I64_extend_i32_s(l1-base.I32_wrap_i64(v12))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v16) < base.Ui64(v17) {
		v35 = int32(1)
		m.G0 = v9 + int32(48)
		return v35
	} else {
		v19 = int32(0)
		if base.Ui64(v12) <= base.Ui64(v16) {
			v35 = v19
			m.G0 = v9 + int32(48)
			return v35
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestIsRemovableXid[0]))
			if v22 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestIsRemovableXid[1]))
				if v24 == v22 {
					v35 = v19
					m.G0 = v9 + int32(48)
					return v35
				} else {
					F_ComputeXidHorizons(m, v9+int32(8))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						v35 = base.B2i32(base.Ui64(v16) < base.Ui64(v32))
						m.G0 = v9 + int32(48)
						return v35
					}
				}
			} else {
				F_ComputeXidHorizons(m, v9+int32(8))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					v35 = base.B2i32(base.Ui64(v16) < base.Ui64(v32))
					m.G0 = v9 + int32(48)
					return v35
				}
			}
		}
	}
}
func F_GrantLock(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v7 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v6 + v7
	v12 = l0 + l2<<(uint(int32(2))%32)
	v14 = v12 + int32(88)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v15 + v7
	v20 = v7 << (uint(l2) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v20 | v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v24 == v25 {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v27 & (v20 ^ int32(-1))
	} else {
	}
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v32 | v20
	return
}
func F_gai_strerror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = int32(_a_F_gai_strerror_0)
	v6 = l0 + int32(1)
	if v6 == int32(0) {
		v26 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	return v26 + base.B2i32(v28 == int32(0))
L2:
	;
	v10 = v4
	v11 = v6
	goto L3
L3:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v12 == int32(0) {
		v26 = v10
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v26 = v22
	goto L1
L5:
	;
	v16 = v10
	goto L6
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v20 != 0 {
		v16 = v16 + int32(1)
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v22 = v16 + int32(2)
	v24 = v11 + int32(1)
	if v24 != 0 {
		v10 = v22
		v11 = v24
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	goto L4
}
func F_gbtreekey_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_gbtreekey_out_0), int32(45), int32(_a_F_gbtreekey_out_1), int32(_a_F_gbtreekey_out_2), int32(_a_F_gbtreekey_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_gdb_date_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v6 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0))))
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(l1))))
	v8 = F_DirectFunctionCall2Coll(m, int32(2640), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return float64(0)
	} else {
		v12 = base.I32_wrap_i64(v8)
		v14 = v12 >> (uint(int32(31)) % 32)
		return base.F64_convert_i32_u(v12 ^ v14 - v14)
	}
}
func F_generate_grouped_paths(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 float64
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 float64
	_ = v242
	var v243 int32
	_ = v243
	var v245 float64
	_ = v245
	var v246 int32
	_ = v246
	var v247 float64
	_ = v247
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v250 int32
	_ = v250
	var v252 float64
	_ = v252
	var v253 int32
	_ = v253
	var v254 float64
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	v4 = int32(0)
	v16 = float64(0)
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v25 == v4 {
		v46 = v4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v20 + int32(48)
	return
L2:
	;
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	goto L2
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v29 = v28
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if base.Ui32(int32(2)) <= base.Ui32(v33-int32(303)) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v46 = int32(1)
	goto L3
L7:
	;
	if v33 != int32(293) {
		v46 = v4
		goto L3
	} else {
		goto L10
	}
L8:
	;
	v29 = v32 + int32(72)
	goto L5
L9:
	;
	goto L6
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v40 != 0 {
		v46 = v4
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_mark_dummy_rel(m, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v51 = int32(0)
	if base.B2i32(v49 == v51)|base.B2i32(v50 == v51) != 0 {
		v97 = base.B2i32(v49|v50 == v51)
		goto L18
	} else {
		goto L19
	}
L15:
	;
	return
L16:
	;
	goto L1
L17:
	;
	if v97 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L18:
	;
	goto L17
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v65 != v66 {
		v97 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v68 = int32(1)
	if v65 <= v68 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = v68
	goto L23
L22:
	;
	v71 = v65
	goto L23
L23:
	;
	v72 = int32(8)
	v77 = int32(0)
	goto L24
L24:
	;
	v85 = v77 << (uint(int32(2)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v49+v72+v85)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v50+v72+v85)))
	v90 = base.B2i32(v87 == v89)
	if v87 != v89 {
		v97 = v90
		goto L18
	} else {
		goto L26
	}
L25:
	;
	v97 = v90
	goto L18
L26:
	;
	v93 = v77 + int32(1)
	if v93 != v71 {
		v77 = v93
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+32)))
	if v104 != int32(1) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v107 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v107
	F_get_agg_clause_costs(m, l0, int32(6), v20+int32(8))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v122 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v167 != 0 {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v167 = int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v131 <= int32(0) {
		v159 = int32(1)
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v167 = v159
	goto L31
L36:
	;
	v134 = int32(0)
	if v134 < v131 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v137 = v131
	goto L39
L38:
	;
	v137 = v134
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v140 = int32(0)
	goto L40
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138+v140<<(uint(int32(2))%32))))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v150 = int32(0)
	v151 = base.B2i32(v149 != v150)
	if v149 == v150 {
		v159 = v151
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v159 = v151
	goto L35
L42:
	;
	v155 = v140 + int32(1)
	if v155 != v137 {
		v140 = v155
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(int32(5)) < base.Ui32(v168) {
		v179 = l1
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v187 = v4
	goto L46
L46:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v188 != 0 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+236))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v182 = F_make_tlist_from_pathtarget(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L15
	} else {
		goto L50
	}
L48:
	;
	if int32(1)<<(uint(v168)%32)&int32(44) == int32(0) {
		v179 = l1
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+240))
	v179 = v178
	goto L47
L50:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v185 = F_make_pathkeys_for_sortclauses(m, l0, v184, v182)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	v187 = v185
	goto L46
L52:
	;
	v189 = int32(0)
	if v188 == v189 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v227 = v4
	goto L54
L54:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v228 != 0 {
		goto L68
	} else {
		goto L69
	}
L55:
	;
	v227 = v226
	goto L54
L56:
	;
	v226 = int32(1)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v196 <= int32(0) {
		v220 = int32(1)
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v226 = v220
	goto L55
L60:
	;
	v199 = int32(0)
	if v199 < v196 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v202 = v196
	goto L63
L62:
	;
	v202 = v199
	goto L63
L63:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v207 = v189
	goto L64
L64:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v203+v207<<(uint(int32(2))%32))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+18)))
	if v212 != int32(1) {
		v220 = v212
		goto L59
	} else {
		goto L66
	}
L65:
	;
	v220 = v212
	goto L59
L66:
	;
	v216 = v207 + int32(1)
	if v216 != v202 {
		v207 = v216
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v230 = v229
	goto L70
L69:
	;
	v230 = v4
	goto L70
L70:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v231 != int32(1) {
		v240 = v4
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v230 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v234 == int32(0) {
		v240 = v4
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v240 = v238
	goto L71
L74:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v242 = *(*float64)(unsafe.Add(mBase, uint32(v230)+32))
	v243 = int32(0)
	v245 = F_estimate_num_groups(m, l0, v241, v242, v243, v243)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L15
	} else {
		goto L77
	}
L75:
	;
	v247 = v16
	goto L76
L76:
	;
	if v240 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v247 = v245
	goto L76
L78:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v240)+32))
	v250 = int32(0)
	v252 = F_estimate_num_groups(m, l0, v248, v249, v250, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L15
	} else {
		goto L81
	}
L79:
	;
	v254 = v16
	goto L80
L80:
	;
	v256 = v167 ^ int32(1)
	v258 = base.B2i32(v230 == int32(0))
	if v256|v258 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v254 = v252
	goto L80
L82:
	;
	v437 = base.B2i32(v240 == int32(0))
	if v256|v437 != 0 {
		goto L143
	} else {
		goto L144
	}
L83:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v260 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v263 <= int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v277 = v4
	goto L86
L86:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283+v277<<(uint(int32(2))%32))))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	if v287 != v230 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L82
L88:
	;
	v416 = v277 + int32(1)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v416 < v417 {
		v277 = v416
		goto L86
	} else {
		goto L142
	}
L89:
	;
	v291 = v288
	goto L91
L90:
	;
	v291 = int32(0)
	goto L91
L91:
	;
	if v291 != 0 {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v287)+64))
	v294 = v20 + int32(4)
	if v187 == v292 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	if v372|base.B2i32(v287 == v230) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L94:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v360
	v372 = int32(1)
	goto L93
L95:
	;
	if v187 != 0 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	if v187 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = int32(0)
	v372 = int32(1)
	goto L93
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = int32(0)
	v372 = int32(1)
	goto L93
L100:
	;
	goto L101
L101:
	;
	if v292 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v312 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v312
	v372 = v312
	goto L93
L103:
	;
	goto L104
L104:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	v316 = int32(0)
	if v316 < v315 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v319 = v315
	goto L107
L106:
	;
	v319 = v316
	goto L107
L107:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v325 = int32(0)
	goto L108
L108:
	;
	if v325 < v320 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v336 = v332 + v325<<(uint(int32(2))%32)
	goto L112
L111:
	;
	v336 = int32(0)
	goto L112
L112:
	;
	if v325 == v319 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v319
	v372 = base.B2i32(v336 == int32(0))
	goto L93
L114:
	;
	goto L115
L115:
	;
	v342 = base.B2i32(v336 == int32(0))
	if v336 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v325
	v372 = v342
	goto L93
L117:
	;
	goto L118
L118:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	if v346 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v325
	v372 = v342
	goto L93
L120:
	;
	goto L121
L121:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v346+v325<<(uint(int32(2))%32))))
	if v350 != v354 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v325
	v372 = int32(0)
	goto L93
L123:
	;
	v325 = v325 + int32(1)
	goto L108
L125:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v377 == int32(0) {
		goto L88
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v387 = F_create_projection_path(m, l0, l1, v287, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L15
	} else {
		goto L131
	}
L128:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_grouped_paths[0])))
	if v381&int32(1) == int32(0) {
		goto L88
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v409 = F_create_agg_path(m, l0, l1, v401, v402, int32(1), int32(6), v405, int32(0), v20+int32(8), v247)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L15
	} else {
		goto L140
	}
L131:
	;
	if v372 != 0 {
		v401 = v387
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v389 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v398 = F_create_incremental_sort_path(m, l0, l1, v387, v187, v389, float64(-1))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L15
	} else {
		goto L139
	}
L134:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_grouped_paths[0])))
	if v391&int32(1) != 0 {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v395 = F_create_sort_path(m, l1, v387, v187, float64(-1))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L15
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v401 = v395
	goto L130
L139:
	;
	v401 = v398
	goto L130
L140:
	;
	F_add_path(m, l1, v409)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L15
	} else {
		goto L141
	}
L141:
	;
	goto L88
L142:
	;
	goto L87
L143:
	;
	v613 = v227 ^ int32(1)
	if v613|v258 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L144:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	if v439 == int32(0) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v442 <= int32(0) {
		goto L143
	} else {
		goto L146
	}
L146:
	;
	v448 = int32(0)
	goto L147
L147:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463+v448<<(uint(int32(2))%32))))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+64))
	v470 = v20 + int32(4)
	if v187 == v468 {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	goto L143
L149:
	;
	v592 = v448 + int32(1)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v592 < v593 {
		v448 = v592
		goto L147
	} else {
		goto L199
	}
L150:
	;
	if v548|base.B2i32(v467 == v240) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L151:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v536
	v548 = int32(1)
	goto L150
L152:
	;
	if v187 != 0 {
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v187 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = int32(0)
	v548 = int32(1)
	goto L150
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = int32(0)
	v548 = int32(1)
	goto L150
L157:
	;
	goto L158
L158:
	;
	if v468 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v488 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v488
	v548 = v488
	goto L150
L160:
	;
	goto L161
L161:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	v492 = int32(0)
	if v492 < v491 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v495 = v491
	goto L164
L163:
	;
	v495 = v492
	goto L164
L164:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v501 = int32(0)
	goto L165
L165:
	;
	if v501 < v496 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v512 = v508 + v501<<(uint(int32(2))%32)
	goto L169
L168:
	;
	v512 = int32(0)
	goto L169
L169:
	;
	if v501 == v495 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v495
	v548 = base.B2i32(v512 == int32(0))
	goto L150
L171:
	;
	goto L172
L172:
	;
	v518 = base.B2i32(v512 == int32(0))
	if v512 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v501
	v548 = v518
	goto L150
L174:
	;
	goto L175
L175:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v468)+12))
	if v522 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v501
	v548 = v518
	goto L150
L177:
	;
	goto L178
L178:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v522+v501<<(uint(int32(2))%32))))
	if v526 != v530 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v501
	v548 = int32(0)
	goto L150
L180:
	;
	v501 = v501 + int32(1)
	goto L165
L182:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v553 == int32(0) {
		goto L149
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v563 = F_create_projection_path(m, l0, l1, v467, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L15
	} else {
		goto L188
	}
L185:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_grouped_paths[0])))
	if v557&int32(1) == int32(0) {
		goto L149
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v585 = F_create_agg_path(m, l0, l1, v577, v578, int32(1), int32(6), v581, int32(0), v20+int32(8), v254)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L15
	} else {
		goto L197
	}
L188:
	;
	if v548 != 0 {
		v577 = v563
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v565 != 0 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v574 = F_create_incremental_sort_path(m, l0, l1, v563, v187, v565, float64(-1))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L15
	} else {
		goto L196
	}
L191:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_generate_grouped_paths[0])))
	if v567&int32(1) != 0 {
		goto L190
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v571 = F_create_sort_path(m, l1, v563, v187, float64(-1))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L15
	} else {
		goto L195
	}
L194:
	;
	goto L193
L195:
	;
	v577 = v571
	goto L187
L196:
	;
	v577 = v574
	goto L187
L197:
	;
	F_add_partial_path(m, l1, v585)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L15
	} else {
		goto L198
	}
L198:
	;
	goto L149
L199:
	;
	goto L148
L200:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v618 = F_create_projection_path(m, l0, l1, v230, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L15
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if v613|v437 != 0 {
		goto L1
	} else {
		goto L206
	}
L203:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v627 = F_create_agg_path(m, l0, l1, v618, v620, int32(2), int32(6), v623, int32(0), v20+int32(8), v247)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L15
	} else {
		goto L204
	}
L204:
	;
	F_add_path(m, l1, v627)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L15
	} else {
		goto L205
	}
L205:
	;
	goto L202
L206:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v633 = F_create_projection_path(m, l0, l1, v240, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L15
	} else {
		goto L207
	}
L207:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v642 = F_create_agg_path(m, l0, l1, v633, v635, int32(2), int32(6), v638, int32(0), v20+int32(8), v254)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L15
	} else {
		goto L208
	}
L208:
	;
	F_add_partial_path(m, l1, v642)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L15
	} else {
		goto L209
	}
L209:
	;
	goto L1
}
func F_generate_matching_part_pairs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	v8 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v22 = F_palloc_mul(m, int32(4), l4)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = F_palloc_mul(m, int32(4), l4)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l4 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v19 < v20 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v30 = l4 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(l4) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v42 = v8
	v51 = v8
	goto L9
L7:
	;
	v99 = v8
	goto L8
L8:
	;
	v117 = v99
	v127 = v8
	goto L13
L9:
	;
	v54 = v42 << (uint(int32(2)) % 32)
	v56 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v54))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v54+v22))) = v56
	v61 = int32(4)
	v62 = v54 | v61
	*(*int32)(unsafe.Add(mBase, uint32(v25+v62))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v22+v62))) = v56
	v70 = v54 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v70))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v22+v70))) = v56
	v78 = v54 | int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v78))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v78+v22))) = v56
	v86 = v42 + v61
	v88 = v51 + v61
	if v88 != l4&int32(2147483644) {
		v42 = v86
		v51 = v88
		goto L9
	} else {
		goto L11
	}
L10:
	;
	if v30 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	v99 = v86
	goto L8
L13:
	;
	v129 = v117 << (uint(int32(2)) % 32)
	v131 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v129))) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v129+v22))) = v131
	v136 = int32(1)
	v139 = v127 + v136
	if v139 != v30 {
		v117 = v117 + v136
		v127 = v139
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L4
L15:
	;
	goto L14
L16:
	;
	v160 = v20
	goto L18
L17:
	;
	v160 = v19
	goto L18
L18:
	;
	if int32(0) < v160 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v171 = int32(0)
	goto L22
L20:
	;
	goto L21
L21:
	;
	if int32(0) < l4 {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	if v20 <= v171 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	if v19 <= v171 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v171<<(uint(int32(2))%32))))
	if v187 < int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+v187<<(uint(int32(2))%32)))) = v171
	goto L24
L27:
	;
	v209 = v171 + int32(1)
	if v209 != v160 {
		v171 = v209
		goto L22
	} else {
		goto L30
	}
L28:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196+v171<<(uint(int32(2))%32))))
	if v200 < int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v200<<(uint(int32(2))%32)))) = v171
	goto L27
L30:
	;
	goto L23
L31:
	;
	v239 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	F_pfree(m, v22)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L48
	}
L34:
	;
	v251 = v239 << (uint(int32(2)) % 32)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v25+v251)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v22)))
	if v253&v255 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if int32(0) <= v255 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v286 = v239 + int32(1)
	if v286 != l4 {
		v239 = v286
		goto L34
	} else {
		goto L47
	}
L39:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262+v255<<(uint(int32(2))%32))))
	v268 = v266
	goto L41
L40:
	;
	v268 = int32(0)
	goto L41
L41:
	;
	v269 = F_lappend(m, v259, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v269
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if int32(0) <= v253 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+276))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275+v253<<(uint(int32(2))%32))))
	v281 = v279
	goto L45
L44:
	;
	v281 = int32(0)
	goto L45
L45:
	;
	v282 = F_lappend(m, v272, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v282
	goto L38
L47:
	;
	goto L35
L48:
	;
	F_pfree(m, v25)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	return
}
func F_generate_mergejoin_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v319 int32
	_ = v319
	var v320 float64
	_ = v320
	var v321 float64
	_ = v321
	var v325 float64
	_ = v325
	var v326 float64
	_ = v326
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v483 int32
	_ = v483
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 float64
	_ = v520
	var v521 float64
	_ = v521
	var v525 float64
	_ = v525
	var v526 float64
	_ = v526
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v20 = F_find_mergeclauses_for_outer_pathkeys(m, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if base.B2i32(v20 == int32(0))&base.B2i32(l4 != int32(2)) != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l6 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v38 = F_make_inner_pathkeys_for_merge(m, l0, v20, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L15
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v28 = v27
	goto L10
L9:
	;
	v28 = int32(0)
	goto L10
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v32 = v30
	goto L13
L12:
	;
	v32 = int32(0)
	goto L13
L13:
	;
	if v32 != v28 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	F_try_mergejoin_path(m, l0, l1, l3, l7, l8, v20, int32(0), v38, l4, l5, l9)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l7)+64))
	if v38 == v42 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v38 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L18:
	;
	v95 = int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v51 = int32(0)
	goto L22
L21:
	;
	v95 = v87
	goto L17
L22:
	;
	v55 = int32(0)
	if v38 == v55 {
		v65 = v55
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v87 = int32(0)
	goto L21
L24:
	;
	if v42 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v59 <= v51 {
		v65 = int32(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v65 = v61 + v51<<(uint(int32(2))%32)
	goto L24
L27:
	;
	v71 = base.B2i32(v65 == int32(0))
	if v65 == int32(0) {
		v87 = v71
		goto L21
	} else {
		goto L32
	}
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v51 < v66 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v95 = base.B2i32(v65 == int32(0))
	goto L17
L31:
	;
	goto L30
L32:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v74 == int32(0) {
		v87 = v71
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v51<<(uint(int32(2))%32)+v74)))
	if v81 == v83 {
		v51 = v51 + int32(1)
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L23
L35:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if l6|base.B2i32(v98 < int32(2)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v95 != 0 {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v104 = F_list_copy(m, v38)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v98 <= int32(0) {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v108 = v104
	goto L36
L41:
	;
	v108 = v38
	goto L36
L42:
	;
	v110 = l7
	goto L44
L43:
	;
	v110 = int32(0)
	goto L44
L44:
	;
	v121 = v108
	v122 = v98
	v126 = v110
	v127 = v110
	goto L45
L45:
	;
	v128 = int32(0)
	if base.B2i32(v121 == v128)|base.B2i32(v122 <= v128) != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L1
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v140 = int32(0)
	if v139 == v140 {
		goto L56
	} else {
		goto L57
	}
L48:
	;
	v138 = int32(0)
	goto L50
L49:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v122 < v135 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L47
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v122
	goto L53
L52:
	;
	goto L53
L53:
	;
	v138 = v121
	goto L50
L54:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v355 = int32(0)
	if v354 == v355 {
		goto L136
	} else {
		goto L137
	}
L55:
	;
	if v287 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L56:
	;
	v287 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if int32(0) < v153 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v163 = v140
	v166 = v140
	goto L62
L60:
	;
	v268 = v140
	goto L61
L61:
	;
	v287 = v268
	goto L55
L62:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169+v166<<(uint(int32(2))%32))))
	if l9 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v268 = v251
	goto L61
L64:
	;
	v258 = v166 + int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v258 < v259 {
		v163 = v251
		v166 = v258
		goto L62
	} else {
		goto L97
	}
L65:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+21)))
	if v174 != int32(1) {
		v251 = v163
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if v163 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L67
L69:
	;
	v177 = F_compare_path_costs(m, v163, v173, int32(1))
	mBase = m.M
	if v177 <= int32(0) {
		v251 = v163
		goto L64
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v173)+64))
	if v138 == v180 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	if v238 != 0 {
		goto L91
	} else {
		goto L92
	}
L74:
	;
	v188 = int32(0)
	goto L75
L75:
	;
	v196 = int32(0)
	if v138 == v196 {
		v206 = v196
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v206 != 0 {
		v251 = v163
		goto L64
	} else {
		goto L90
	}
L77:
	;
	if v180 != 0 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v200 <= v188 {
		v206 = int32(0)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v206 = v202 + v188<<(uint(int32(2))%32)
	goto L77
L80:
	;
	if v206 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v188 < v207 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v206 == int32(0) {
		goto L73
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	v251 = v163
	goto L64
L86:
	;
	goto L76
L87:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	if v213 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213+v188<<(uint(int32(2))%32))))
	if v220 == v222 {
		v188 = v188 + int32(1)
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v251 = v163
	goto L64
L90:
	;
	goto L73
L91:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v241 = v239
	goto L93
L92:
	;
	v241 = int32(0)
	goto L93
L93:
	;
	v242 = F_bms_is_subset(m, v241, v140)
	mBase = m.M
	if v242 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v243 = v173
	goto L96
L95:
	;
	v243 = v163
	goto L96
L96:
	;
	v251 = v243
	goto L64
L97:
	;
	goto L63
L98:
	;
	v352 = int32(0)
	v353 = v126
	goto L54
L99:
	;
	goto L100
L100:
	;
	if v126 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v287)+40))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v126)+40))
	if v297 != v298 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	if v122 < v98 {
		goto L129
	} else {
		goto L130
	}
L104:
	;
	if int32(0) <= v340 {
		v352 = int32(0)
		v353 = v126
		goto L54
	} else {
		goto L128
	}
L105:
	;
	if v297 < v298 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	goto L114
L108:
	;
	v303 = int32(-1)
	goto L110
L109:
	;
	v303 = int32(1)
	goto L110
L110:
	;
	v340 = v303
	goto L104
L111:
	;
	v340 = v334
	goto L104
L112:
	;
	v334 = int32(0)
	goto L111
L114:
	;
	goto L115
L115:
	;
	v319 = int32(-1)
	v320 = *(*float64)(unsafe.Add(mBase, uint32(v287)+56))
	v321 = *(*float64)(unsafe.Add(mBase, uint32(v126)+56))
	if base.F64_lt(v320, v321) != 0 {
		v334 = v319
		goto L111
	} else {
		goto L122
	}
L122:
	;
	if base.F64_gt(v320, v321) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v340 = int32(1)
	goto L104
L124:
	;
	goto L125
L125:
	;
	v325 = *(*float64)(unsafe.Add(mBase, uint32(v287)+48))
	v326 = *(*float64)(unsafe.Add(mBase, uint32(v126)+48))
	if base.F64_lt(v325, v326) != 0 {
		v334 = v319
		goto L111
	} else {
		goto L126
	}
L126:
	;
	if base.F64_gt(v325, v326) != 0 {
		v334 = int32(1)
		goto L111
	} else {
		goto L127
	}
L127:
	;
	goto L112
L128:
	;
	goto L103
L129:
	;
	v345 = F_trim_mergeclauses_for_inner_pathkeys(m, v20, v138)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L132
	}
L130:
	;
	v347 = v20
	goto L131
L131:
	;
	v348 = int32(0)
	F_try_mergejoin_path(m, l0, l1, l3, v287, l8, v347, v348, v348, l4, l5, l9)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L133
	}
L132:
	;
	v347 = v345
	goto L131
L133:
	;
	v352 = v347
	v353 = v287
	goto L54
L134:
	;
	if v122 < int32(2) {
		goto L1
	} else {
		goto L215
	}
L135:
	;
	if v502 == int32(0) {
		v567 = v127
		goto L134
	} else {
		goto L178
	}
L136:
	;
	v502 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	if int32(0) < v368 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v378 = v355
	v381 = v355
	goto L142
L140:
	;
	v483 = v355
	goto L141
L141:
	;
	v502 = v483
	goto L135
L142:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384+v381<<(uint(int32(2))%32))))
	if l9 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v483 = v466
	goto L141
L144:
	;
	v473 = v381 + int32(1)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	if v473 < v474 {
		v378 = v466
		v381 = v473
		goto L142
	} else {
		goto L177
	}
L145:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+21)))
	if v389 != int32(1) {
		v466 = v378
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if v378 != 0 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	goto L147
L149:
	;
	v392 = F_compare_path_costs(m, v378, v388, v355)
	mBase = m.M
	if v392 <= int32(0) {
		v466 = v378
		goto L144
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v388)+64))
	if v138 == v395 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v388)+16))
	if v453 != 0 {
		goto L171
	} else {
		goto L172
	}
L154:
	;
	v403 = int32(0)
	goto L155
L155:
	;
	v411 = int32(0)
	if v138 == v411 {
		v421 = v411
		goto L157
	} else {
		goto L158
	}
L156:
	;
	if v421 != 0 {
		v466 = v378
		goto L144
	} else {
		goto L170
	}
L157:
	;
	if v395 != 0 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v415 <= v403 {
		v421 = int32(0)
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v421 = v417 + v403<<(uint(int32(2))%32)
	goto L157
L160:
	;
	if v421 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L161:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v403 < v422 {
		goto L160
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v421 == int32(0) {
		goto L153
	} else {
		goto L165
	}
L164:
	;
	goto L163
L165:
	;
	v466 = v378
	goto L144
L166:
	;
	goto L156
L167:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	if v428 == int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v428+v403<<(uint(int32(2))%32))))
	if v435 == v437 {
		v403 = v403 + int32(1)
		goto L155
	} else {
		goto L169
	}
L169:
	;
	v466 = v378
	goto L144
L170:
	;
	goto L153
L171:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v456 = v454
	goto L173
L172:
	;
	v456 = int32(0)
	goto L173
L173:
	;
	v457 = F_bms_is_subset(m, v456, v355)
	mBase = m.M
	if v457 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v458 = v388
	goto L176
L175:
	;
	v458 = v378
	goto L176
L176:
	;
	v466 = v458
	goto L144
L177:
	;
	goto L143
L178:
	;
	if v127 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v502)+40))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v127)+40))
	if v510 != v511 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	goto L181
L181:
	;
	if v502 != v353 {
		goto L207
	} else {
		goto L208
	}
L182:
	;
	if int32(0) <= v553 {
		v567 = v127
		goto L134
	} else {
		goto L206
	}
L183:
	;
	if v510 < v511 {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	goto L185
L185:
	;
	goto L191
L186:
	;
	v516 = int32(-1)
	goto L188
L187:
	;
	v516 = int32(1)
	goto L188
L188:
	;
	v553 = v516
	goto L182
L189:
	;
	v553 = v547
	goto L182
L190:
	;
	v547 = int32(0)
	goto L189
L191:
	;
	v519 = int32(-1)
	v520 = *(*float64)(unsafe.Add(mBase, uint32(v502)+48))
	v521 = *(*float64)(unsafe.Add(mBase, uint32(v127)+48))
	if base.F64_lt(v520, v521) != 0 {
		v547 = v519
		goto L189
	} else {
		goto L194
	}
L194:
	;
	if base.F64_gt(v520, v521) != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v553 = int32(1)
	goto L182
L196:
	;
	goto L197
L197:
	;
	v525 = *(*float64)(unsafe.Add(mBase, uint32(v502)+56))
	v526 = *(*float64)(unsafe.Add(mBase, uint32(v127)+56))
	if base.F64_lt(v525, v526) != 0 {
		v547 = v519
		goto L189
	} else {
		goto L198
	}
L198:
	;
	if base.F64_gt(v525, v526) == int32(0) {
		goto L190
	} else {
		goto L199
	}
L199:
	;
	v547 = int32(1)
	goto L189
L206:
	;
	goto L181
L207:
	;
	if v352 != 0 {
		v560 = v352
		goto L210
	} else {
		goto L211
	}
L208:
	;
	goto L209
L209:
	;
	v567 = v502
	goto L134
L210:
	;
	v561 = int32(0)
	F_try_mergejoin_path(m, l0, l1, l3, v502, l8, v560, v561, v561, l4, l5, l9)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L2
	} else {
		goto L214
	}
L211:
	;
	if v98 <= v122 {
		v560 = v20
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v558 = F_trim_mergeclauses_for_inner_pathkeys(m, v20, v138)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L2
	} else {
		goto L213
	}
L213:
	;
	v560 = v558
	goto L210
L214:
	;
	goto L209
L215:
	;
	if l6 == int32(0) {
		v121 = v138
		v122 = v122 - int32(1)
		v126 = v353
		v127 = v567
		goto L45
	} else {
		goto L216
	}
L216:
	;
	goto L46
}
func F_generate_qualified_relation_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_0), v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_generate_qualified_relation_name_1), int32(_a_F_generate_qualified_relation_name_2), int32(_a_F_generate_qualified_relation_name_3))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
			v35 = v33 + v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
			v39 = m.G0
			v41 = v39 - int32(16)
			m.G0 = v41
			v45 = *(*int32)(unsafe.Add(mBase, _c_F_generate_qualified_relation_name[0]))
			if base.B2i32(v45 != int32(0))&base.B2i32(v36 == v45) != 0 {
				v51 = F_pstrdup(m, int32(_a_F_generate_qualified_relation_name_4))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v68 = v51
					if v68 != 0 {
						v85 = F_quote_qualified_identifier(m, v68, v35+int32(4))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							m.G0 = v41 + int32(16)
							F_ReleaseCatCache(m, v14)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v85
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = v36
							F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_5), v41)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_generate_qualified_relation_name_6), int32(3725), int32(_a_F_generate_qualified_relation_name_7))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v55 = F_SearchSysCache1(m, int32(38), base.I64_extend_i32_u(v36))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v55 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = v36
							F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_5), v41)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_generate_qualified_relation_name_6), int32(3725), int32(_a_F_generate_qualified_relation_name_7))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
						v64 = F_pstrdup(m, v59+v60+int32(4))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v55)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = v64
								if v68 != 0 {
									v85 = F_quote_qualified_identifier(m, v68, v35+int32(4))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										m.G0 = v41 + int32(16)
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v85
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v41))) = v36
										F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_5), v41)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_generate_qualified_relation_name_6), int32(3725), int32(_a_F_generate_qualified_relation_name_7))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_gensign(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	if l2 <= int32(0) {
	} else {
		v12 = l3 << (uint(int32(3)) % 32)
		if l2 != int32(1) {
			v21 = l1
			v22 = int32(0)
			for {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v29 = base.I32_rem_u_s(v28, v12)
				v30 = int32(3)
				v32 = l0 + int32(base.Ui32(v29)>>(uint(v30)%32))
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
				v34 = int32(1)
				v35 = int32(7)
				v38 = v33 | v34<<(uint(v29&v35)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v38)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v41 = base.I32_rem_u_s(v40, v12)
				v44 = l0 + int32(base.Ui32(v41)>>(uint(v30)%32))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
				v50 = v45 | v34<<(uint(v41&v35)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v50)
				v53 = v21 + int32(8)
				v55 = v22 + int32(2)
				if v55 != l2&int32(2147483646) {
					v21 = v53
					v22 = v55
					continue
				} else {
					break
				}
				break
			}
			if l2&int32(1) == int32(0) {
			} else {
				v60 = v53
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				v68 = base.I32_rem_u_s(v67, v12)
				v71 = l0 + int32(base.Ui32(v68)>>(uint(int32(3))%32))
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
				v77 = v72 | int32(1)<<(uint(v68&int32(7))%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v77)
			}
		} else {
			v60 = l1
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v68 = base.I32_rem_u_s(v67, v12)
			v71 = l0 + int32(base.Ui32(v68)>>(uint(int32(3))%32))
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v77 = v72 | int32(1)<<(uint(v68&int32(7))%32)
			*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v77)
		}
	}
	return
}
func F_geo_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v18 float64
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v38 float64
	_ = v38
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v60 float64
	_ = v60
	var v67 float64
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 float64
	_ = v87
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v123 float64
	_ = v123
	var v127 float64
	_ = v127
	var v129 float64
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v149 float64
	_ = v149
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v164 float64
	_ = v164
	var v166 float64
	_ = v166
	var v168 float64
	_ = v168
	var v178 float64
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v198 float64
	_ = v198
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v217 float64
	_ = v217
	var v228 float64
	_ = v228
	var v231 float64
	_ = v231
	var v237 int64
	_ = v237
	var v242 int32
	_ = v242
	var v265 float64
	_ = v265
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v274 float64
	_ = v274
	var v279 float64
	_ = v279
	var v284 float64
	_ = v284
	var v288 float64
	_ = v288
	var v297 float64
	_ = v297
	var v306 float64
	_ = v306
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v319 float64
	_ = v319
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	v18 = base.F64_mul(base.F64_div(v14, float64(360)), float64(6.283185307179586))
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v31 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v18))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v31) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v31) < base.Ui32(int32(1044816030)) {
			v60 = float64(1)
		} else {
			v38 = F___cos(m, v18, float64(0))
			mBase = m.M
			v60 = v38
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v31) {
			v60 = base.F64_sub(v18, v18)
		} else {
			v42 = F___rem_pio2(m, v18, v24)
			mBase = m.M
			v43 = *(*float64)(unsafe.Add(mBase, uint32(v24)+8))
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v24)))
			switch v42&int32(3) - int32(1) {
			case 0:
				v51 = F___sin(m, v44, v43, int32(1))
				mBase = m.M
				v60 = base.F64_neg(v51)
			case 1:
				v53 = F___cos(m, v44, v43)
				mBase = m.M
				v60 = base.F64_neg(v53)
			case 2:
				v56 = F___sin(m, v44, v43, int32(1))
				mBase = m.M
				v60 = v56
			default:
				v49 = F___cos(m, v44, v43)
				mBase = m.M
				v60 = v49
			}
		}
	}
	m.G0 = v24 + int32(16)
	v67 = base.F64_mul(base.F64_div(v13, float64(360)), float64(6.283185307179586))
	v71 = m.G0
	v73 = v71 - int32(16)
	m.G0 = v73
	v80 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v67))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v80) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v80) < base.Ui32(int32(1044816030)) {
			v109 = float64(1)
		} else {
			v87 = F___cos(m, v67, float64(0))
			mBase = m.M
			v109 = v87
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v80) {
			v109 = base.F64_sub(v67, v67)
		} else {
			v91 = F___rem_pio2(m, v67, v73)
			mBase = m.M
			v92 = *(*float64)(unsafe.Add(mBase, uint32(v73)+8))
			v93 = *(*float64)(unsafe.Add(mBase, uint32(v73)))
			switch v91&int32(3) - int32(1) {
			case 0:
				v100 = F___sin(m, v93, v92, int32(1))
				mBase = m.M
				v109 = base.F64_neg(v100)
			case 1:
				v102 = F___cos(m, v93, v92)
				mBase = m.M
				v109 = base.F64_neg(v102)
			case 2:
				v105 = F___sin(m, v93, v92, int32(1))
				mBase = m.M
				v109 = v105
			default:
				v98 = F___cos(m, v93, v92)
				mBase = m.M
				v109 = v98
			}
		}
	}
	m.G0 = v73 + int32(16)
	v113 = float64(6.283185307179586)
	v114 = float64(360)
	v123 = base.F64_abs(base.F64_sub(base.F64_mul(base.F64_div(v12, v114), v113), base.F64_mul(base.F64_div(v10, v114), v113)))
	if base.F64_gt(v123, float64(3.141592653589793)) != 0 {
		v127 = base.F64_sub(v113, v123)
	} else {
		v127 = v123
	}
	v129 = base.F64_mul(v127, float64(0.5))
	v133 = m.G0
	v135 = v133 - int32(16)
	m.G0 = v135
	v142 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v129))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v142) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v142) < base.Ui32(int32(1045430272)) {
			v168 = v129
		} else {
			v149 = F___sin(m, v129, float64(0), int32(0))
			mBase = m.M
			v168 = v149
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v142) {
			v168 = base.F64_sub(v129, v129)
		} else {
			v153 = F___rem_pio2(m, v129, v135)
			mBase = m.M
			v154 = *(*float64)(unsafe.Add(mBase, uint32(v135)+8))
			v155 = *(*float64)(unsafe.Add(mBase, uint32(v135)))
			switch v153&int32(3) - int32(1) {
			case 0:
				v162 = F___cos(m, v155, v154)
				mBase = m.M
				v168 = v162
			case 1:
				v164 = F___sin(m, v155, v154, int32(1))
				mBase = m.M
				v168 = base.F64_neg(v164)
			case 2:
				v166 = F___cos(m, v155, v154)
				mBase = m.M
				v168 = base.F64_neg(v166)
			default:
				v161 = F___sin(m, v155, v154, int32(1))
				mBase = m.M
				v168 = v161
			}
		}
	}
	m.G0 = v135 + int32(16)
	v178 = base.F64_mul(base.F64_abs(base.F64_sub(v18, v67)), float64(0.5))
	v182 = m.G0
	v184 = v182 - int32(16)
	m.G0 = v184
	v191 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v178))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v191) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v191) < base.Ui32(int32(1045430272)) {
			v217 = v178
		} else {
			v198 = F___sin(m, v178, float64(0), int32(0))
			mBase = m.M
			v217 = v198
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v191) {
			v217 = base.F64_sub(v178, v178)
		} else {
			v202 = F___rem_pio2(m, v178, v184)
			mBase = m.M
			v203 = *(*float64)(unsafe.Add(mBase, uint32(v184)+8))
			v204 = *(*float64)(unsafe.Add(mBase, uint32(v184)))
			switch v202&int32(3) - int32(1) {
			case 0:
				v211 = F___cos(m, v204, v203)
				mBase = m.M
				v217 = v211
			case 1:
				v213 = F___sin(m, v204, v203, int32(1))
				mBase = m.M
				v217 = base.F64_neg(v213)
			case 2:
				v215 = F___cos(m, v204, v203)
				mBase = m.M
				v217 = base.F64_neg(v215)
			default:
				v210 = F___sin(m, v204, v203, int32(1))
				mBase = m.M
				v217 = v210
			}
		}
	}
	m.G0 = v184 + int32(16)
	v228 = base.F64_sqrt(base.F64_add(base.F64_mul(v217, v217), base.F64_mul(v168, base.F64_mul(v168, base.F64_mul(v60, v109)))))
	if base.F64_gt(v228, float64(1)) != 0 {
		v231 = float64(1)
	} else {
		v231 = v228
	}
	v237 = base.I64_reinterpret_f64(v231)
	v242 = base.I32_wrap_i64(int64(base.Ui64(v237)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v242) {
		if base.I32_wrap_i64(v237)|(v242-int32(1072693248)) == int32(0) {
			v319 = base.F64_add(base.F64_mul(v231, float64(1.5707963267948966)), float64(7.52316384526264e-37))
		} else {
			v319 = base.F64_div(float64(0), base.F64_sub(v231, v231))
		}
	} else {
		if base.Ui32(v242) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v242+int32(-1048576)) < base.Ui32(int32(1044381696)) {
				v311 = v231
				v319 = v311
			} else {
				v265 = F_R(m, base.F64_mul(v231, v231))
				mBase = m.M
				v319 = base.F64_add(base.F64_mul(v231, v265), v231)
			}
		} else {
			v272 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v231)), float64(0.5))
			v273 = base.F64_sqrt(v272)
			v274 = F_R(m, v272)
			mBase = m.M
			if base.Ui32(int32(1072640819)) <= base.Ui32(v242) {
				v279 = base.F64_add(base.F64_mul(v273, v274), v273)
				v306 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v279, v279), float64(-6.123233995736766e-17)))
			} else {
				v284 = float64(0.7853981633974483)
				v288 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v273) & int64(-4294967296))
				v297 = base.F64_div(base.F64_sub(v272, base.F64_mul(v288, v288)), base.F64_add(v273, v288))
				v306 = base.F64_add(base.F64_sub(base.F64_sub(v284, base.F64_add(v288, v288)), base.F64_sub(base.F64_mul(base.F64_add(v273, v273), v274), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v297, v297)))), v284)
			}
			if v237 < int64(0) {
				v310 = base.F64_neg(v306)
			} else {
				v310 = v306
			}
			v311 = v310
			v319 = v311
		}
	}
	return base.I64_reinterpret_f64(base.F64_mul(v319, float64(7917.495432)))
}
func F_geqo_eval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0]))
	v16 = F_AllocSetContextCreateInternal(m, v11, int32(_a_F_geqo_eval_0), v5, int32(_a_F_geqo_eval_1), int32(_a_F_geqo_eval_2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = int32(_a_F_geqo_eval_3)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0])) = v16
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		if v22 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			v24 = v23
		} else {
			v24 = v5
		}
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = int32(0)
		v28 = F_gimme_tree(m, l1, l2, l3)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			if v28 == int32(0) {
				v38 = float64(1.7976931348623157e+308)
				v39 = int32(2147483647)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
				v35 = *(*float64)(unsafe.Add(mBase, uint32(v34)+56))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
				v38 = v35
				v39 = v36
			}
			*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v39
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			v43 = int32(0)
			if base.B2i32(v42 == v43)|base.B2i32(v24 <= v43) != 0 {
				v53 = int32(0)
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
				if v24 < v50 {
					*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v24
				} else {
				}
				v53 = v42
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v53
			*(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0])) = v19
			F_MemoryContextDelete(m, v16)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_getKeyJsonValueFromContainer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v371 int32
	_ = v371
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	v5 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = v17 & int32(268435455)
	if v19 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v25 = l0 + int32(4)
	v28 = v25 + v19<<(uint(int32(3))%32)
	v34 = v5
	v38 = v19
	v39 = v5
	goto L4
L4:
	;
	v48 = int32(base.Ui32(v38-v39)>>(uint(int32(1))%32)) + v39
	v53 = v48
	v54 = v34
	goto L6
L5:
	;
	return v397
L6:
	;
	v66 = v53 - int32(1)
	if int32(0) <= v66 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v25+v48<<(uint(int32(2))%32))))
	if v84 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0+v53<<(uint(int32(2))%32))))
	v75 = v72&int32(268435455) + v54
	if int32(0) <= v72 {
		v53 = v66
		v54 = v75
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v79 = v54
	goto L10
L10:
	;
	goto L7
L11:
	;
	v79 = v75
	goto L10
L12:
	;
	goto L5
L13:
	;
	v384 = int32(0)
	v388 = base.B2i32(v383 < v384)
	if v383 < v384 {
		goto L86
	} else {
		goto L87
	}
L14:
	;
	if v141 != l2 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v92 = v48
	v95 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v141 = v84 & int32(268435455)
	goto L14
L18:
	;
	v105 = v92 - int32(1)
	if int32(0) <= v105 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v141 = v84&int32(268435455) - v118
	goto L14
L20:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0+v92<<(uint(int32(2))%32))))
	v114 = v111&int32(268435455) + v95
	if int32(0) <= v111 {
		v92 = v105
		v95 = v114
		goto L18
	} else {
		goto L23
	}
L21:
	;
	v118 = v95
	goto L22
L22:
	;
	goto L19
L23:
	;
	v118 = v114
	goto L22
L24:
	;
	if l2 < v141 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v147 = v79 + v28
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v146 = int32(1)
	goto L29
L28:
	;
	v146 = int32(-1)
	goto L29
L29:
	;
	v383 = v146
	goto L13
L30:
	;
	if v209 != 0 {
		v383 = v209
		goto L13
	} else {
		goto L48
	}
L31:
	;
	v209 = int32(0)
	goto L30
L32:
	;
	v183 = v178
	v184 = v179
	v185 = v180
	goto L42
L33:
	;
	if (v147|l1)&int32(3) != 0 {
		v178 = v147
		v179 = l1
		v180 = l2
		goto L32
	} else {
		goto L36
	}
L34:
	;
	v171 = v147
	v172 = l1
	v173 = l2
	goto L35
L35:
	;
	if v173 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L36:
	;
	v155 = v147
	v156 = l1
	v157 = l2
	goto L37
L37:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v160 != v161 {
		v178 = v155
		v179 = v156
		v180 = v157
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v171 = v166
	v172 = v164
	v173 = v168
	goto L35
L39:
	;
	v163 = int32(4)
	v164 = v156 + v163
	v166 = v155 + v163
	v168 = v157 - v163
	if base.Ui32(int32(3)) < base.Ui32(v168) {
		v155 = v166
		v156 = v164
		v157 = v168
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v178 = v171
	v179 = v172
	v180 = v173
	goto L32
L42:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v188 == v189 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v209 = v188 - v189
	goto L30
L44:
	;
	v191 = int32(1)
	v196 = v185 - v191
	if v196 != 0 {
		v183 = v183 + v191
		v184 = v184 + v191
		v185 = v196
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	goto L31
L48:
	;
	if l3 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v213 = F_palloc(m, int32(32))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v217 = l3
	goto L51
L51:
	;
	v219 = v48 + v19
	v224 = v219
	v225 = int32(0)
	goto L54
L52:
	;
	return int32(0)
L53:
	;
	v217 = v213
	goto L51
L54:
	;
	v237 = v224 - int32(1)
	if int32(0) <= v237 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v259 = l0 + v219<<(uint(int32(2))%32) + int32(4)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	switch int32(base.Ui32(v260)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L62
	case 3:
		goto L63
	case 4:
		goto L66
	default:
		goto L61
	}
L56:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0+v224<<(uint(int32(2))%32))))
	v246 = v243&int32(268435455) + v225
	if int32(0) <= v243 {
		v224 = v237
		v225 = v246
		goto L54
	} else {
		goto L59
	}
L57:
	;
	v250 = v225
	goto L58
L58:
	;
	goto L55
L59:
	;
	v250 = v246
	goto L58
L60:
	;
	v397 = v217
	goto L12
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(18)
	v327 = (v250 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = v28 + v327
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v330 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v318 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)) = uint8(v318)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(3)
	goto L60
L63:
	;
	v314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)) = uint8(v314)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(3)
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v28 + (v250+int32(3))&int32(-4)
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = v28 + v250
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v271 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(0)
	goto L60
L67:
	;
	v276 = v219
	v277 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v271 & int32(268435455)
	goto L60
L70:
	;
	v284 = v276 - int32(1)
	if int32(0) <= v284 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v271&int32(268435455) - v297
	goto L60
L72:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0+v276<<(uint(int32(2))%32))))
	v293 = v290&int32(268435455) + v277
	if int32(0) <= v290 {
		v276 = v284
		v277 = v293
		goto L70
	} else {
		goto L75
	}
L73:
	;
	v297 = v277
	goto L74
L74:
	;
	goto L71
L75:
	;
	v297 = v293
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v371 + (v250 - v327)
	goto L60
L77:
	;
	v335 = v219
	v336 = int32(0)
	goto L80
L78:
	;
	goto L79
L79:
	;
	v371 = v330 & int32(268435455)
	goto L76
L80:
	;
	v343 = v335 - int32(1)
	if int32(0) <= v343 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v371 = v330&int32(268435455) - v356
	goto L76
L82:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0+v335<<(uint(int32(2))%32))))
	v352 = v349&int32(268435455) + v336
	if int32(0) <= v349 {
		v335 = v343
		v336 = v352
		goto L80
	} else {
		goto L85
	}
L83:
	;
	v356 = v336
	goto L84
L84:
	;
	goto L81
L85:
	;
	v356 = v352
	goto L84
L86:
	;
	v389 = v48 + int32(1)
	goto L88
L87:
	;
	v389 = v39
	goto L88
L88:
	;
	if v383 < v384 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v390 = v38
	goto L91
L90:
	;
	v390 = v48
	goto L91
L91:
	;
	if base.Ui32(v389) < base.Ui32(v390) {
		v34 = v384
		v38 = v390
		v39 = v389
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v397 = v384
	goto L12
}
func F_get_attstatsslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v6
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v19
	v27 = v15 + v16
	v29 = v27 + int32(32)
	v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+20)))
	if v30 == l2 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L33
	} else {
		goto L55
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L33
	} else {
		goto L52
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return v154
L4:
	;
	v70 = v68 << (uint(int32(2)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v29+v70)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70+v27)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v75
	v77 = int32(1)
	if l4&v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L5:
	;
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+24)))
	if v44 == l2 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v41 != l3 {
		goto L5
	} else {
		goto L15
	}
L7:
	;
	if l3 == int32(0) {
		v68 = v6
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+22)))
	if l2 != v38 {
		goto L5
	} else {
		goto L13
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v34 == l3 {
		v68 = v6
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+22)))
	if l2 != v36 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	if l3 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v68 = int32(1)
	goto L4
L15:
	;
	v68 = int32(1)
	goto L4
L16:
	;
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+28)))
	if l2 != v61 {
		v154 = v6
		goto L3
	} else {
		goto L27
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v57 != l3 {
		goto L16
	} else {
		goto L26
	}
L18:
	;
	v46 = int32(2)
	if l3 == int32(0) {
		v68 = v46
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+26)))
	if l2 != v53 {
		goto L16
	} else {
		goto L24
	}
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v49 == l3 {
		v68 = v46
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+26)))
	if l2 != v51 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	if l3 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v68 = int32(3)
	goto L4
L26:
	;
	v68 = int32(3)
	goto L4
L27:
	;
	v63 = int32(4)
	if l3 == int32(0) {
		v68 = v63
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v66 != l3 {
		v154 = v6
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v68 = v63
	goto L4
L30:
	;
	v83 = F_SysCacheGetAttrNotNull(m, int32(65), l1, v68+int32(27))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if l4&int32(2) == int32(0) {
		v154 = v77
		goto L3
	} else {
		goto L45
	}
L33:
	;
	return int32(0)
L34:
	;
	v88 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v83))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v90
	v94 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v90))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	if v94 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+22)))
	v100 = v98 + v99
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+76)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+78)))
	v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v100)+128)))
	F_deconstruct_array(m, v88, v101, v102, v103, l0+int32(12), int32(0), l0+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+78)))
	if v111 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	F_ReleaseCatCache(m, v94)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L33
	} else {
		goto L44
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v88
	goto L39
L41:
	;
	goto L42
L42:
	;
	F_pfree(m, v88)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	goto L32
L45:
	;
	v129 = F_SysCacheGetAttrNotNull(m, int32(65), l1, v68+int32(22))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v132 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v129))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v134 != int32(1) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
	if v137 <= int32(0) {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	if v140 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	if v141 != int32(700) {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v132 + int32(24)
	v154 = v77
	goto L3
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v90
	F_errmsg_internal(m, int32(_a_F_get_attstatsslot_0), v13)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L33
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_get_attstatsslot_1), int32(3595), int32(_a_F_get_attstatsslot_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_get_attstatsslot_3), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L33
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_get_attstatsslot_1), int32(3640), int32(_a_F_get_attstatsslot_2))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L33
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_atttypetypmodcoll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_SearchSysCache2(m, int32(7), base.I64_extend_i32_u(l0), base.I64_extend_i32_s(l1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
				F_errmsg_internal(m, int32(_a_F_get_atttypetypmodcoll_0), v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_atttypetypmodcoll_1), int32(1178), int32(_a_F_get_atttypetypmodcoll_2))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
			v35 = v33 + v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v40
			F_ReleaseCatCache(m, v15)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
func F_get_coercion_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v66 = F_format_type_with_typemod(m, l2, l3)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L24
	}
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v27&int32(1) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 != int32(7) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 != l2 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 != int32(-1) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	F_get_const_expr(m, l0, l1, int32(-1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	goto L1
L9:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v59&int32(1) != 0 {
		goto L1
	} else {
		goto L22
	}
L10:
	;
	F_get_rule_expr(m, l0, l1, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L21
	}
L11:
	;
	F_appendStringInfoChar(m, v13, int32(40))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v40 = v27
	goto L13
L13:
	;
	v41 = F_isSimpleNode(m, l0, l4, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v35&int32(1) == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v40 = v35
	goto L13
L16:
	;
	if v41 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoChar(m, v43, int32(40))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	F_get_rule_expr(m, l0, l1, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoChar(m, v50, int32(41))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	goto L9
L22:
	;
	F_appendStringInfoChar(m, v13, int32(41))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L1
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v66
	F_appendStringInfo(m, v13, int32(_a_F_get_coercion_expr_0), v11)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_get_memoize_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int64
	_ = v19
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v554 int32
	_ = v554
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 float64
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1082 float64
	_ = v1082
	var v1091 float64
	_ = v1091
	var v1095 float64
	_ = v1095
	var v1096 int64
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1103 float64
	_ = v1103
	var v1105 float64
	_ = v1105
	var v1108 float64
	_ = v1108
	var v1111 float64
	_ = v1111
	var v1120 int32
	_ = v1120
	var v1132 int32
	_ = v1132
	var v1140 int32
	_ = v1140
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	v8 = int32(0)
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l6)+40))
	if v19&int64(1024) == int64(0) {
		v1120 = v8
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	F_list_free(m, v1140)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L46
	} else {
		goto L265
	}
L3:
	;
	return v1120
L4:
	;
	if v19&int64(256) != int64(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v28)+16))
	if base.F64_lt(v29, float64(2)) != 0 {
		v1120 = v8
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+333)))
	if v32 != int32(1) {
		v390 = v8
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v396 != 0 {
		goto L100
	} else {
		goto L101
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v36 = int32(0)
	if v35 == v36 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v81 == int32(2) {
		v390 = v8
		goto L9
	} else {
		goto L27
	}
L12:
	;
	v81 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v44 = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v45 <= v44 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = v44
	goto L17
L16:
	;
	v48 = v45
	goto L17
L17:
	;
	v52 = int32(0)
	v54 = v36
	goto L18
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(8)+v52<<(uint(int32(2))%32))))
	if v61 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v81 = v73
	goto L11
L20:
	;
	goto L19
L21:
	;
	v62 = int32(2)
	if v54 != 0 {
		v73 = v62
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v68 = v54
	goto L23
L23:
	;
	v70 = v52 + int32(1)
	if v70 != v48 {
		v52 = v70
		v54 = v68
		goto L18
	} else {
		goto L26
	}
L24:
	;
	v63 = int32(1)
	if base.Ui32(v63) < base.Ui32(base.I32_popcnt(v61)) {
		v73 = v62
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v68 = v63
	goto L23
L26:
	;
	v73 = v68
	goto L20
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v84 == int32(0) {
		v390 = v8
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v87 <= int32(0) {
		v390 = v8
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v98 = v8
	v102 = v8
	goto L30
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v98<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v113 == int32(0) {
		v368 = v102
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v390 = v368
	goto L9
L32:
	;
	v375 = v98 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v375 < v376 {
		v98 = v375
		v102 = v368
		goto L30
	} else {
		goto L98
	}
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v117 = int32(0)
	if base.B2i32(v116 == v117)|base.B2i32(v35 == v117) != 0 {
		v163 = base.B2i32(v116|v35 == v117)
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v163 == int32(0) {
		v368 = v102
		goto L32
	} else {
		goto L45
	}
L35:
	;
	goto L34
L36:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v131 != v132 {
		v163 = int32(0)
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v134 = int32(1)
	if v131 <= v134 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v137 = v134
	goto L40
L39:
	;
	v137 = v131
	goto L40
L40:
	;
	v138 = int32(8)
	v143 = int32(0)
	goto L41
L41:
	;
	v151 = v143 << (uint(int32(2)) % 32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v116+v138+v151)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v35+v138+v151)))
	v156 = base.B2i32(v153 == v155)
	if v153 != v155 {
		v163 = v156
		goto L35
	} else {
		goto L43
	}
L42:
	;
	v163 = v156
	goto L35
L43:
	;
	v159 = v143 + int32(1)
	if v159 != v137 {
		v143 = v159
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v172 = F_pull_varnos(m, l0, v171)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return int32(0)
L47:
	;
	v176 = int32(0)
	if base.B2i32(v172 == v176)|base.B2i32(v35 == v176) != 0 {
		v221 = v176
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v221 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L49:
	;
	goto L48
L50:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v186 < v187 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v189 = v186
	goto L53
L52:
	;
	v189 = v187
	goto L53
L53:
	;
	if v189 <= int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v192 = int32(1)
	goto L56
L55:
	;
	v192 = v189
	goto L56
L56:
	;
	v193 = int32(8)
	v198 = int32(0)
	goto L57
L57:
	;
	v205 = v198 << (uint(int32(2)) % 32)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v35+v193+v205)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v172+v193+v205)))
	v210 = v207 & v209
	v212 = base.B2i32(v210 != int32(0))
	if v210 != 0 {
		v221 = v212
		goto L49
	} else {
		goto L59
	}
L58:
	;
	v221 = v212
	goto L49
L59:
	;
	v214 = v198 + int32(1)
	if v214 != v192 {
		v198 = v214
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v226 = F_lappend(m, v102, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L46
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v229 = F_pull_vars_of_level(m, v223, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L46
	} else {
		goto L66
	}
L64:
	;
	v368 = v226
	goto L32
L65:
	;
	F_list_free(m, v229)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L46
	} else {
		goto L97
	}
L66:
	;
	if v229 == int32(0) {
		v348 = v102
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v233 = int32(0)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v234 <= v233 {
		v348 = v102
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v244 = v233
	v249 = v102
	goto L69
L69:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v244<<(uint(int32(2))%32))))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v260 != int32(6) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v348 = v331
	goto L65
L71:
	;
	v333 = v244 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v333 < v334 {
		v244 = v333
		v249 = v331
		goto L69
	} else {
		goto L96
	}
L72:
	;
	v329 = F_lappend(m, v249, v259)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L46
	} else {
		goto L95
	}
L73:
	;
	if v260 != int32(321) {
		v331 = v249
		goto L71
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	v325 = F_bms_is_member(m, v323, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L46
	} else {
		goto L93
	}
L76:
	;
	v265 = F_find_placeholder_info(m, l0, v259)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L46
	} else {
		goto L77
	}
L77:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	v269 = int32(0)
	if v267 == v269 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v322 != 0 {
		goto L72
	} else {
		goto L92
	}
L79:
	;
	v322 = int32(1)
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v268 == int32(0) {
		v315 = v269
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v322 = v315
	goto L78
L83:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v279 < v278 {
		v315 = v269
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v281 = int32(1)
	if v278 <= v281 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v284 = v281
	goto L87
L86:
	;
	v284 = v278
	goto L87
L87:
	;
	v285 = int32(8)
	v290 = int32(0)
	goto L88
L88:
	;
	v297 = v290 << (uint(int32(2)) % 32)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v267+v285+v297)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v268+v285+v297)))
	v304 = v299 & (v301 ^ int32(-1))
	v306 = base.B2i32(v304 == int32(0))
	if v304 != 0 {
		v315 = v306
		goto L82
	} else {
		goto L90
	}
L89:
	;
	v315 = v306
	goto L82
L90:
	;
	v308 = v290 + int32(1)
	if v308 != v284 {
		v290 = v308
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v331 = v249
	goto L71
L93:
	;
	if v325 == int32(0) {
		v331 = v249
		goto L71
	} else {
		goto L94
	}
L94:
	;
	goto L72
L95:
	;
	v331 = v329
	goto L71
L96:
	;
	goto L70
L97:
	;
	v368 = v348
	goto L32
L98:
	;
	goto L31
L99:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+8)))
	if l5&int32(-2) == int32(4) {
		goto L107
	} else {
		goto L108
	}
L100:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
	if v397 != 0 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v398|v390 != 0 {
		goto L99
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	goto L1
L105:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v477 = F_contain_volatile_functions(m, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L46
	} else {
		goto L122
	}
L106:
	;
	if v396 == int32(0) {
		goto L1
	} else {
		goto L112
	}
L107:
	;
	if v400&int32(1) == int32(0) {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v412 = l6 + int32(8)
	if v400&int32(1) == int32(0) {
		v469 = v412
		goto L105
	} else {
		goto L111
	}
L110:
	;
	v417 = l6 + int32(8)
	goto L106
L111:
	;
	v417 = v412
	goto L106
L112:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v420 == int32(0) {
		v469 = v417
		goto L105
	} else {
		goto L113
	}
L113:
	;
	v423 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v424 <= v423 {
		v469 = v417
		goto L105
	} else {
		goto L114
	}
L114:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v396)+20))
	v435 = v423
	goto L115
L115:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446+v435<<(uint(int32(2))%32))))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+56))
	v452 = F_bms_is_member(m, v451, v427)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L46
	} else {
		goto L117
	}
L116:
	;
	goto L1
L117:
	;
	if v452 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v455 = v435 + int32(1)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	if v455 < v456 {
		v435 = v455
		goto L115
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	goto L116
L121:
	;
	v469 = v417
	goto L105
L122:
	;
	if v477 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	if v479 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v535 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L125:
	;
	v482 = int32(0)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v483 <= v482 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v493 = v482
	goto L127
L127:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v504+v493<<(uint(int32(2))%32))))
	v509 = F_contain_volatile_functions(m, v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L46
	} else {
		goto L129
	}
L128:
	;
	goto L1
L129:
	;
	if v509 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v514 = v493 + int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v514 < v515 {
		v493 = v514
		goto L127
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	goto L128
L133:
	;
	goto L124
L134:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v968 = F_list_concat(m, v390, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L46
	} else {
		goto L239
	}
L135:
	;
	v946 = int32(0)
	v949 = v946
	v956 = v946
	v957 = v946
	goto L134
L136:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v535)+16))
	if v538 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v608)+16))
	if v620 == int32(0) {
		goto L135
	} else {
		goto L157
	}
L138:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	if v541 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	v543 = int32(0)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	if v544 <= v543 {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	v542 = v541
	goto L143
L142:
	;
	v542 = l2
	goto L143
L143:
	;
	v608 = v535
	v619 = v542
	goto L137
L144:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v596 == int32(0) {
		goto L135
	} else {
		goto L153
	}
L145:
	;
	v554 = v543
	goto L146
L146:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v565+v554<<(uint(int32(2))%32))))
	v570 = F_contain_volatile_functions(m, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L46
	} else {
		goto L148
	}
L147:
	;
	goto L1
L148:
	;
	if v570 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v575 = v554 + int32(1)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	if v575 < v576 {
		v554 = v575
		goto L146
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	goto L147
L152:
	;
	goto L144
L153:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	if v599 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v600 = v599
	goto L156
L155:
	;
	v600 = l2
	goto L156
L156:
	;
	v608 = v596
	v619 = v600
	goto L137
L157:
	;
	v623 = int32(0)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v620)+4))
	if v624 <= v623 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v627 = int32(0)
	v949 = v627
	v956 = v623
	v957 = v627
	goto L134
L159:
	;
	goto L160
L160:
	;
	v629 = int32(0)
	v632 = v629
	v638 = v629
	v639 = v623
	v640 = v629
	goto L161
L161:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v620)+12))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v650+v638<<(uint(int32(2))%32))))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v654)+4))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	if v656 != int32(17) {
		v1132 = v632
		v1140 = v640
		goto L2
	} else {
		goto L163
	}
L162:
	;
	v949 = v918
	v956 = v923
	v957 = v919
	goto L134
L163:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655)+28))
	if v659 == int32(0) {
		v1132 = v632
		v1140 = v640
		goto L2
	} else {
		goto L164
	}
L164:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v662 != int32(2) {
		v1132 = v632
		v1140 = v640
		goto L2
	} else {
		goto L165
	}
L165:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v654)+44))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v619)+8))
	v668 = int32(0)
	if v666 == v668 {
		goto L169
	} else {
		goto L170
	}
L166:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v904+v654)))
	if v906 == int32(0) {
		v1132 = v632
		v1140 = v640
		goto L2
	} else {
		goto L230
	}
L167:
	;
	v898 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v654)+120)) = uint8(v898)
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v655)+28))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+12))
	v903 = v901
	v904 = int32(160)
	goto L166
L168:
	;
	if v721 != 0 {
		goto L182
	} else {
		goto L183
	}
L169:
	;
	v721 = int32(1)
	goto L168
L170:
	;
	goto L171
L171:
	;
	if v667 == int32(0) {
		v714 = v668
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v721 = v714
	goto L168
L173:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v666)+4))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v678 < v677 {
		v714 = v668
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v680 = int32(1)
	if v677 <= v680 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v683 = v680
	goto L177
L176:
	;
	v683 = v677
	goto L177
L177:
	;
	v684 = int32(8)
	v689 = int32(0)
	goto L178
L178:
	;
	v696 = v689 << (uint(int32(2)) % 32)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v666+v684+v696)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v667+v684+v696)))
	v703 = v698 & (v700 ^ int32(-1))
	v705 = base.B2i32(v703 == int32(0))
	if v703 != 0 {
		v714 = v705
		goto L172
	} else {
		goto L180
	}
L179:
	;
	v714 = v705
	goto L172
L180:
	;
	v707 = v689 + int32(1)
	if v707 != v683 {
		v689 = v707
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v654)+48))
	v723 = int32(0)
	if v722 == v723 {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	goto L184
L184:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v654)+44))
	v778 = int32(0)
	if v777 == v778 {
		goto L201
	} else {
		goto L202
	}
L185:
	;
	if v776 != 0 {
		goto L167
	} else {
		goto L199
	}
L186:
	;
	v776 = int32(1)
	goto L185
L187:
	;
	goto L188
L188:
	;
	if v665 == int32(0) {
		v769 = v723
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v776 = v769
	goto L185
L190:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	if v733 < v732 {
		v769 = v723
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v735 = int32(1)
	if v732 <= v735 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v738 = v735
	goto L194
L193:
	;
	v738 = v732
	goto L194
L194:
	;
	v739 = int32(8)
	v744 = int32(0)
	goto L195
L195:
	;
	v751 = v744 << (uint(int32(2)) % 32)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v722+v739+v751)))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v665+v739+v751)))
	v758 = v753 & (v755 ^ int32(-1))
	v760 = base.B2i32(v758 == int32(0))
	if v758 != 0 {
		v769 = v760
		goto L189
	} else {
		goto L197
	}
L196:
	;
	v769 = v760
	goto L189
L197:
	;
	v762 = v744 + int32(1)
	if v762 != v738 {
		v744 = v762
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	goto L184
L200:
	;
	if v831 == int32(0) {
		v1132 = v632
		v1140 = v640
		goto L2
	} else {
		goto L214
	}
L201:
	;
	v831 = int32(1)
	goto L200
L202:
	;
	goto L203
L203:
	;
	if v665 == int32(0) {
		v824 = v778
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v831 = v824
	goto L200
L205:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v777)+4))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	if v788 < v787 {
		v824 = v778
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v790 = int32(1)
	if v787 <= v790 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v793 = v790
	goto L209
L208:
	;
	v793 = v787
	goto L209
L209:
	;
	v794 = int32(8)
	v799 = int32(0)
	goto L210
L210:
	;
	v806 = v799 << (uint(int32(2)) % 32)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v777+v794+v806)))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v665+v794+v806)))
	v813 = v808 & (v810 ^ int32(-1))
	v815 = base.B2i32(v813 == int32(0))
	if v813 != 0 {
		v824 = v815
		goto L204
	} else {
		goto L212
	}
L211:
	;
	v824 = v815
	goto L204
L212:
	;
	v817 = v799 + int32(1)
	if v817 != v793 {
		v799 = v817
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v654)+48))
	v835 = int32(0)
	if v834 == v835 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	if v888 == int32(0) {
		v1132 = v632
		v1140 = v640
		goto L2
	} else {
		goto L229
	}
L216:
	;
	v888 = int32(1)
	goto L215
L217:
	;
	goto L218
L218:
	;
	if v667 == int32(0) {
		v881 = v835
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v888 = v881
	goto L215
L220:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v845 < v844 {
		v881 = v835
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v847 = int32(1)
	if v844 <= v847 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v850 = v847
	goto L224
L223:
	;
	v850 = v844
	goto L224
L224:
	;
	v851 = int32(8)
	v856 = int32(0)
	goto L225
L225:
	;
	v863 = v856 << (uint(int32(2)) % 32)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v834+v851+v863)))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v667+v851+v863)))
	v870 = v865 & (v867 ^ int32(-1))
	v872 = base.B2i32(v870 == int32(0))
	if v870 != 0 {
		v881 = v872
		goto L219
	} else {
		goto L227
	}
L226:
	;
	v881 = v872
	goto L219
L227:
	;
	v874 = v856 + int32(1)
	if v874 != v850 {
		v856 = v874
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v891 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v654)+120)) = uint8(v891)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v655)+28))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v893)+12))
	v903 = v894 + int32(4)
	v904 = int32(164)
	goto L166
L230:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v903)))
	v910 = F_list_member(m, v632, v909)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L46
	} else {
		goto L231
	}
L231:
	;
	if v910 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v914 = F_lappend_oid(m, v640, v906)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L46
	} else {
		goto L235
	}
L233:
	;
	v918 = v632
	v919 = v640
	goto L234
L234:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v654)+124))
	v923 = base.B2i32(v920 == int32(0)) | v639
	v925 = v638 + int32(1)
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v620)+4))
	if v925 < v926 {
		v632 = v918
		v638 = v925
		v639 = v923
		v640 = v919
		goto L161
	} else {
		goto L237
	}
L235:
	;
	v916 = F_lappend(m, v632, v909)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L46
	} else {
		goto L236
	}
L236:
	;
	v918 = v916
	v919 = v914
	goto L234
L237:
	;
	goto L162
L238:
	;
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	v1048 = v1035 & int32(1)
	v1049 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v1051 = F_palloc0(m, int32(120))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L46
	} else {
		goto L257
	}
L239:
	;
	if v968 == int32(0) {
		v1028 = v949
		v1035 = v956
		v1036 = v957
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v968)+4))
	if v972 <= int32(0) {
		v1028 = v949
		v1035 = v956
		v1036 = v957
		goto L238
	} else {
		goto L241
	}
L241:
	;
	v976 = v949
	v982 = int32(0)
	v984 = v957
	goto L242
L242:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v968)+12))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v994+v982<<(uint(int32(2))%32))))
	v999 = F_contain_volatile_functions(m, v998)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L46
	} else {
		goto L244
	}
L243:
	;
	v1028 = v1021
	v1035 = v1023
	v1036 = v1022
	goto L238
L244:
	;
	if v999 != 0 {
		v1132 = v976
		v1140 = v984
		goto L2
	} else {
		goto L245
	}
L245:
	;
	v1001 = F_exprType(m, v998)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L46
	} else {
		goto L246
	}
L246:
	;
	v1004 = F_lookup_type_cache(m, v1001, int32(17))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L46
	} else {
		goto L247
	}
L247:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+68))
	if v1006 == int32(0) {
		v1132 = v976
		v1140 = v984
		goto L2
	} else {
		goto L248
	}
L248:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+52))
	if v1009 == int32(0) {
		v1132 = v976
		v1140 = v984
		goto L2
	} else {
		goto L249
	}
L249:
	;
	v1012 = F_list_member(m, v976, v998)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L46
	} else {
		goto L250
	}
L250:
	;
	if v1012 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+52))
	v1017 = F_lappend_oid(m, v984, v1016)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L46
	} else {
		goto L254
	}
L252:
	;
	v1021 = v976
	v1022 = v984
	goto L253
L253:
	;
	v1023 = int32(1)
	v1025 = v982 + v1023
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v968)+4))
	if v1025 < v1026 {
		v976 = v1021
		v982 = v1025
		v984 = v1022
		goto L242
	} else {
		goto L256
	}
L254:
	;
	v1019 = F_lappend(m, v976, v998)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L46
	} else {
		goto L255
	}
L255:
	;
	v1021 = v1019
	v1022 = v1017
	goto L253
L256:
	;
	goto L243
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1051))) = int64(1567663063337)
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+12)) = v1056
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v1059 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1051)+20)) = uint8(v1059)
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+16)) = v1058
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1062 == int32(1) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	v1067 = v1065
	goto L260
L259:
	;
	v1067 = int32(0)
	goto L260
L260:
	;
	v1069 = v1067 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1051)+21)) = uint8(v1069)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+24)) = v1071
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+88)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1051)+85)) = uint8(v1048)
	*(*uint8)(unsafe.Add(mBase, uint32(v1051)+84)) = uint8(v1046)
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+80)) = v1028
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+76)) = v1036
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+72)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+64)) = v1073
	v1082 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1049)&int64(9223372036854775807)))|base.F64_gt(v1049, v1082) != 0 {
		v1095 = v1082
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1096 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1051)+104)) = v1096
	*(*float64)(unsafe.Add(mBase, uint32(v1051)+96)) = v1095
	*(*int64)(unsafe.Add(mBase, uint32(v1051)+112)) = v1096
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1051)+40)) = v1101
	v1103 = *(*float64)(unsafe.Add(mBase, uint32(l3)+48))
	v1105 = *(*float64)(unsafe.Add(mBase, _c_F_get_memoize_path[0]))
	*(*float64)(unsafe.Add(mBase, uint32(v1051)+48)) = base.F64_add(v1103, v1105)
	v1108 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1051)+56)) = base.F64_add(v1105, v1108)
	v1111 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1051)+32)) = v1111
	v1120 = v1051
	goto L3
L262:
	;
	goto L261
L263:
	;
	v1091 = float64(1)
	if base.F64_le(v1049, v1091) != 0 {
		v1095 = v1091
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v1095 = base.F64_nearest(v1049)
	goto L262
L265:
	;
	F_list_free(m, v1132)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L46
	} else {
		goto L266
	}
L266:
	;
	return int32(0)
}
func F_get_mergejoin_opfamilies(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	v2 = int32(0)
	v10 = int64(0)
	v12 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(l0), v10, v10)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v16 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v28 = v2
	v29 = v2
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)+v28<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	v37 = v35 + v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v38 <= int32(2741) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L24
	}
L9:
	;
	v73 = v28 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v73 < v74 {
		v28 = v73
		v29 = v71
		goto L7
	} else {
		goto L23
	}
L10:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+16)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v63 = F_IndexAmTranslateStrategy(m, v61, v60, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L11:
	;
	v54 = F_GetIndexAmRoutineByAmId(m, v38, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	switch v38 - int32(403) {
	case 0:
		v60 = v38
		goto L10
	case 1:
		goto L11
	case 2:
		v71 = v29
		goto L9
	default:
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.B2i32(v38 == int32(2742))|base.B2i32(v38 == int32(3580))|base.B2i32(v38 == int32(4000)) != 0 {
		v71 = v29
		goto L9
	} else {
		goto L17
	}
L15:
	;
	if v38 != int32(783) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v71 = v29
	goto L9
L17:
	;
	goto L11
L18:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+10)))
	if v56 != int32(1) {
		v71 = v29
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v60 = v59
	goto L10
L20:
	;
	if v63 != int32(3) {
		v71 = v29
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v68 = F_lappend_oid(m, v29, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v71 = v68
	goto L9
L23:
	;
	goto L8
L24:
	;
	return v71
}
func F_get_negator(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14291(m, l0, int32(40))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_opname(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14286(m, l0, int32(40))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_ordering_op_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v18 = int64(0)
	v20 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(l0), v18, v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if int32(0) < v24 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v20)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(-64)+v30<<(uint(int32(2))%32))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v44 = v42 + v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	if v45 <= int32(2741) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	goto L5
L8:
	;
	v90 = v30 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if v90 < v91 {
		v30 = v90
		goto L6
	} else {
		goto L23
	}
L9:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+16)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v70 = F_IndexAmTranslateStrategy(m, v68, v67, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L10:
	;
	v61 = F_GetIndexAmRoutineByAmId(m, v45, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v45 == int32(783) {
		goto L8
	} else {
		goto L16
	}
L12:
	;
	switch v45 - int32(403) {
	case 0:
		v67 = v45
		goto L9
	case 1:
		goto L10
	case 2:
		goto L8
	default:
		goto L11
	}
L13:
	;
	goto L14
L14:
	;
	if base.B2i32(v45 == int32(2742))|base.B2i32(v45 == int32(3580))|base.B2i32(v45 == int32(4000)) != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	goto L10
L17:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+10)))
	if v63 != int32(1) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v67 = v66
	goto L9
L19:
	;
	if v70&int32(-5) != int32(1) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v76 != v77 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v81
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v70
	F_ReleaseCatCacheList(m, v20)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	return int32(1)
L23:
	;
	goto L7
L24:
	;
	return int32(0)
}
func F_get_proposed_default_constraint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = F_make_ands_explicit(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v7
		v18 = F_list_make1_impl(m, int32(1), v5+int32(8))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = F_makeBoolExpr(m, int32(2), v18, int32(-1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_eval_const_expressions(m, int32(0), v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v26 = F_canonicalize_qual(m, v23, int32(1))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = F_make_ands_implicit(m, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							m.G0 = v5 + int32(16)
							return v28
						}
					}
				}
			}
		}
	}
}
func F_get_relids_for_join(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v9 = F_find_jointree_node_for_rel(m, v8, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
				F_errmsg_internal(m, int32(_a_F_get_relids_for_join_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_relids_for_join_1), int32(_a_F_get_relids_for_join_2), int32(_a_F_get_relids_for_join_3))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v30 = F_get_relids_in_jointree(m, v9, int32(1), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v30
			}
		}
	}
}
func F_get_relids_in_jointree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v4 {
		v86 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v86
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v13 - int32(63) {
	case 0:
		goto L3
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L4
	}
L3:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v81 = F_bms_make_singleton(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L29
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L11
	} else {
		goto L26
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = F_get_relids_in_jointree(m, v44, l1, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L15
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == int32(0) {
		v86 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v20 <= v19 {
		v86 = v4
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v23 = v19
	v26 = v4
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v23<<(uint(int32(2))%32))))
	v34 = F_get_relids_in_jointree(m, v33, l1, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v86 = v38
	goto L1
L11:
	;
	return int32(0)
L12:
	;
	v38 = F_bms_join(m, v26, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v41 = v23 + int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v41 < v42 {
		v23 = v41
		v26 = v38
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v48 = F_get_relids_in_jointree(m, v47, l1, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v50 = F_bms_join(m, v45, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v52 == int32(0) {
		v86 = v50
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v55 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if l2 == int32(0) {
		v86 = v50
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l1 == int32(0) {
		v86 = v50
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v60 = F_bms_add_member(m, v50, v52)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v86 = v60
	goto L1
L24:
	;
	v64 = F_bms_add_member(m, v50, v52)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v86 = v64
	goto L1
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v70
	F_errmsg_internal(m, int32(_a_F_get_relids_in_jointree_0), v9)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_get_relids_in_jointree_1), int32(_a_F_get_relids_in_jointree_2), int32(_a_F_get_relids_in_jointree_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v86 = v81
	goto L1
}
func F_get_rels_with_domain(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = F_format_type_be(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v14, int32(4), int32(3), int32(184), int64(1247))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_ScanKeyInit(m, v14+int32(56), int32(5), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v44 = F_systable_beginscan(m, v24, int32(2674), int32(1), int32(0), int32(2), v14)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v46 = F_systable_getnext(m, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v46 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = v46
	v53 = v2
	goto L12
L10:
	;
	v252 = v2
	goto L11
L11:
	;
	F_systable_endscan(m, v44)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L62
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v61 = v59 + v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v64 = v62 - int32(1247)
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v252 = v239
	goto L11
L14:
	;
	v245 = F_systable_getnext(m, v44)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L60
	}
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v80 <= int32(0) {
		v239 = v53
		goto L14
	} else {
		goto L29
	}
L16:
	;
	if v64 == int32(12) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v68 = F_get_typtype(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	goto L15
L20:
	;
	v239 = v53
	goto L14
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v68 == int32(100) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v73 = F_get_rels_with_domain(m, v70)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_find_composite_type_dependencies(m, v70, int32(0), v16)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v75 = F_list_concat(m, v53, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v239 = v75
	goto L14
L28:
	;
	v239 = v53
	goto L14
L29:
	;
	if v53 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+48))
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v169)+120)))
	if v170 < v160 {
		v239 = v162
		goto L14
	} else {
		goto L51
	}
L31:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v127 = F_relation_open(m, v125, int32(5))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L42
	}
L32:
	;
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v85 < v86 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = v86
	goto L35
L34:
	;
	v90 = v85
	goto L35
L35:
	;
	v92 = v85
	goto L36
L36:
	;
	if v92 == v90 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v159 = v109
	v160 = v80
	v162 = v53
	goto L30
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v92<<(uint(int32(2))%32)+v107)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+56))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v111 != v112 {
		v92 = v92 + int32(1)
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v143 = F_palloc(m, int32(12))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L48
	}
L41:
	;
	F_relation_close(m, v127, int32(5))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L47
	}
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+48))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+72))
	if v130 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_find_composite_type_dependencies(m, v130, int32(0), v16)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v135 = v129
	goto L45
L45:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+119)))
	switch v136 - int32(109) {
	case 0, 5:
		goto L40
	default:
		goto L41
	}
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+48))
	v135 = v134
	goto L45
L47:
	;
	v239 = v53
	goto L14
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v127
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v127)+48))
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v149)+120)))
	v151 = F_palloc_mul(m, int32(4), v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = v151
	v154 = F_lappend(m, v53, v143)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v159 = v143
	v160 = v156
	v162 = v154
	goto L30
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+52))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v179 = v172 + v173<<(uint(int32(3))%32) + v160*int32(100)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+19)))
	if v180 != 0 {
		v239 = v162
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179-int32(72))+68))
	if v183 != l0 {
		v239 = v162
		goto L14
	} else {
		goto L53
	}
L53:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v185 + int32(1)
	if v185 <= int32(0) {
		v218 = v185
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v228+v218<<(uint(int32(2))%32)))) = v232
	v239 = v162
	goto L14
L55:
	;
	v192 = v185
	goto L56
L56:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v205 = v202 + v192<<(uint(int32(2))%32)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v205-int32(4))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v208 <= v209 {
		v218 = v192
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v218 = int32(0)
	goto L54
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v208
	v212 = int32(1)
	if v212 < v192 {
		v192 = v192 - v212
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if v245 != 0 {
		v49 = v245
		v53 = v239
		goto L12
	} else {
		goto L61
	}
L61:
	;
	goto L13
L62:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	m.G0 = v14 + int32(112)
	return v252
}
func F_get_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 != 0 {
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		if v42 == int32(1) {
			v45 = F_cstring_to_text(m, l1)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v47)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45
				return int32(0)
			}
		} else {
			return int32(0)
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v8 != 0 {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
			if v42 == int32(1) {
				v45 = F_cstring_to_text(m, l1)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v47)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45
					return int32(0)
				}
			} else {
				return int32(0)
			}
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
			v10 = int32(1)
			v12 = int32(0)
			if base.B2i32(v9&v10 == v12)|base.B2i32(l2 != v10) == v12 {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v19)
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
				if v42 == int32(1) {
					v45 = F_cstring_to_text(m, l1)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v47)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45
						return int32(0)
					}
				} else {
					return int32(0)
				}
			} else {
				v23 = int32(0)
				if base.B2i32(v9&int32(1) == v23)|base.B2i32(l2 != int32(11)) == v23 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
					if v42 == int32(1) {
						v45 = F_cstring_to_text(m, l1)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v47)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45
							return int32(0)
						}
					} else {
						return int32(0)
					}
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
					v35 = F_cstring_to_text_with_len(m, v32, v33-v32)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
						if v42 == int32(1) {
							v45 = F_cstring_to_text(m, l1)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v47)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45
								return int32(0)
							}
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_get_singleton_append_subpath(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 - int32(293) {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		if v7 == int32(0) {
			v35 = l0
			return v35
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			if v10 == int32(1) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				v23 = F_lappend(m, v20, v22)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					v29 = F_list_concat(m, v23, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v35 = v34
						return v35
					}
				}
			} else {
				v35 = l0
				return v35
			}
		}
	case 1:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		if v13 == int32(0) {
			v35 = l0
			return v35
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v16 != int32(1) {
				v35 = l0
				return v35
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				v23 = F_lappend(m, v20, v22)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					v29 = F_list_concat(m, v23, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v35 = v34
						return v35
					}
				}
			}
		}
	default:
		v35 = l0
		return v35
	}
}
func F_get_sortgroupclause_expr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(0)
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12+v15<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v11 != v23 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	return v28
L6:
	;
	v26 = v15 + int32(1)
	if v26 != v8 {
		v15 = v26
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	F_errmsg_internal(m, int32(_a_F_get_sortgroupclause_expr_0), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_get_sortgroupclause_expr_1), int32(366), int32(_a_F_get_sortgroupclause_expr_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_sortgroupref_clause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	if v11 < v8 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = v8
	goto L6
L5:
	;
	v14 = v11
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = v3
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15+v19<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if l0 != v25 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return v24
L9:
	;
	v28 = v19 + int32(1)
	if v14 != v28 {
		v19 = v28
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	F_errmsg_internal(m, int32(_a_F_get_sortgroupref_clause_0), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_get_sortgroupref_clause_1), int32(443), int32(_a_F_get_sortgroupref_clause_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_statistics_object_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14294(m, l0, l1, int32(_a_F_get_statistics_object_oid_0), int32(2689), int32(_a_F_get_statistics_object_oid_1), int32(63))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_get_th(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_strlen(m, l0)
	mBase = m.M
	v12 = l0 + v11
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12-int32(1)))))
	if base.Ui32((v15-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		if base.Ui32(int32(2)) <= base.Ui32(v11) {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12-int32(2)))))
			if v26 == int32(49) {
				if l1 == int32(1) {
					v50 = int32(_a_F_get_th_0)
				} else {
					v50 = int32(_a_F_get_th_1)
				}
				v51 = v50
			} else {
				switch v15 - int32(49) {
				case 0:
					if l1 == int32(1) {
						v35 = int32(_a_F_get_th_2)
					} else {
						v35 = int32(_a_F_get_th_3)
					}
					v51 = v35
				case 1:
					if l1 == int32(1) {
						v40 = int32(_a_F_get_th_4)
					} else {
						v40 = int32(_a_F_get_th_5)
					}
					v51 = v40
				case 2:
					if l1 == int32(1) {
						v45 = int32(_a_F_get_th_6)
					} else {
						v45 = int32(_a_F_get_th_7)
					}
					v51 = v45
				default:
					if l1 == int32(1) {
						v50 = int32(_a_F_get_th_0)
					} else {
						v50 = int32(_a_F_get_th_1)
					}
					v51 = v50
				}
			}
		} else {
			switch v15 - int32(49) {
			case 0:
				if l1 == int32(1) {
					v35 = int32(_a_F_get_th_2)
				} else {
					v35 = int32(_a_F_get_th_3)
				}
				v51 = v35
			case 1:
				if l1 == int32(1) {
					v40 = int32(_a_F_get_th_4)
				} else {
					v40 = int32(_a_F_get_th_5)
				}
				v51 = v40
			case 2:
				if l1 == int32(1) {
					v45 = int32(_a_F_get_th_6)
				} else {
					v45 = int32(_a_F_get_th_7)
				}
				v51 = v45
			default:
				if l1 == int32(1) {
					v50 = int32(_a_F_get_th_0)
				} else {
					v50 = int32(_a_F_get_th_1)
				}
				v51 = v50
			}
		}
		m.G0 = v9 + int32(16)
		return v51
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685634))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg(m, int32(_a_F_get_th_8), v9)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_th_9), int32(1568), int32(_a_F_get_th_10))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_get_typisdefined(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v10)+82)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = v12
				return v15 & int32(1)
			}
		} else {
			v15 = int32(0)
			return v15 & int32(1)
		}
	}
}
func F_get_typlenbyval(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_get_typlenbyval_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_typlenbyval_1), int32(2570), int32(_a_F_get_typlenbyval_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v32)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v34)
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_get_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v5 = l4
	v11 = F_palloc0(m, int32(40))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = F_palloc0(m, int32(36))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_pg_detoast_datum_packed(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v20 == int32(1) {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
					if v26 == int32(18) {
						v29 = int32(16)
					} else {
						v29 = int32(0)
					}
					if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v36 = int32(4)
					} else {
						v36 = v29
					}
					v49 = v36
				} else {
					v37 = int32(1)
					if v20&v37 != 0 {
						v49 = int32(base.Ui32(v20)>>(uint(v37)%32)) - v37
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v51 = int32(1)
				if v20&v51 != 0 {
					v55 = v51
				} else {
					v55 = int32(4)
				}
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_get_worker[0]))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
				v61 = F_makeJsonLexContextCstringLen(m, int32(0), v18+v55, v49, v59, int32(1))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
					*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)) = uint8(v5)
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v61
					v69 = F_palloc0_mul(m, int32(1), l3)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v69
						v73 = F_palloc_mul(m, int32(4), l3)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v73
							if int32(0) < l3 {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
								v79 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v79)
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
								v95 = int32(1475)
								v96 = int32(36)
								*(*int32)(unsafe.Add(mBase, uint32(v96+v11))) = v95
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(1475)
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
								if l3 != 0 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(1476)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(1477)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1478)
									v95 = int32(1479)
									v96 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(v96+v11))) = v95
								}
							}
							if l1 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(1480)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(1481)
							} else {
							}
							if l2 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(1482)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = int32(1483)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(1476)
							} else {
							}
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							v111 = F_pg_parse_json(m, v110, v11)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								if v111 != 0 {
									F_json_errsave_error(m, v111, v110, int32(0))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
										F_freeJsonLexContext(m, v116)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
											return v119
										}
									}
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
									F_freeJsonLexContext(m, v116)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
										return v119
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_getid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = l0
	goto L3
L1:
	;
	m.G0 = v11 + int32(16)
	return v139
L2:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v115))) = uint8(v119)
	v121 = v110
	goto L39
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(base.Ui32(v21-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v21 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v33 = v13
	v36 = v21
	v38 = v4
	v40 = v4
	goto L11
L5:
	;
	v13 = v13 + int32(1)
	goto L3
L6:
	;
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L4
L8:
	;
	v110 = v13
	v115 = v21
	goto L2
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	if base.B2i32(v36 == int32(34))|v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v110 = v108
	v115 = v105
	goto L2
L13:
	;
	v108 = v104 + int32(1)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 != 0 {
		v33 = v108
		v36 = v109
		v38 = v105
		v40 = v106
		goto L11
	} else {
		goto L38
	}
L14:
	;
	if int32(63) <= v38 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	if v36 != int32(34) {
		v73 = v33
		goto L14
	} else {
		goto L21
	}
L16:
	;
	if base.I32_extend8_s(v36) < int32(0) {
		v73 = v33
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L18
L18:
	;
	if v36 == int32(95) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	if base.B2i32(base.Ui32(v36-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v36|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
		v110 = v33
		v115 = v38
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	if v40 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v104 = v33
	v105 = v38
	v106 = int32(1)
	goto L13
L23:
	;
	goto L24
L24:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v67 != int32(34) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v104 = v33
	v105 = v38
	v106 = int32(0)
	goto L13
L26:
	;
	goto L27
L27:
	;
	v73 = v33 + int32(1)
	goto L14
L28:
	;
	v76 = int32(0)
	v77 = F_errsave_start(m, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v38))) = uint8(v36)
	v104 = v73
	v105 = v38 + int32(1)
	v106 = v40
	goto L13
L31:
	;
	return int32(0)
L32:
	;
	if v77 == int32(0) {
		v139 = v76
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(34103428))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_getid_0), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(64)
	v93 = F_errdetail(m, int32(_a_F_getid_1), v11)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, l2, int32(_a_F_getid_2), int32(208), int32(_a_F_getid_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v139 = v76
	goto L1
L38:
	;
	goto L12
L39:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v129-int32(9)))&base.B2i32(v129 != int32(32)) != 0 {
		v139 = v121
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v121 = v121 + int32(1)
	goto L39
}
func F_getinternalerrposition(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_getinternalerrposition[0]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_getinternalerrposition[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_getinternalerrposition_0), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_getinternalerrposition_1), int32(1813), int32(_a_F_getinternalerrposition_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_c_F_getinternalerrposition[1])))
		return v26
	}
}
func F_getlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v10 = F_LogicalTapeRead(m, l0, v5+int32(12), int32(4))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(4) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_getlen_0), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_getlen_1), int32(3175), int32(_a_F_getlen_2))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v5 + int32(16)
				return v16
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_getlen_3), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_getlen_1), int32(3173), int32(_a_F_getlen_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_getoffset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	v3 = int32(0)
	v9 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v10 - int32(43) {
	case 0:
		goto L2
	default:
		v18 = l0
		v19 = v9
		goto L1
	case 2:
		goto L3
	}
L1:
	;
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18))))
	if base.Ui32(int32(9)) < base.Ui32(v20-int32(48)) {
		v151 = v3
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v18 = l0 + int32(1)
	v19 = v9
	goto L1
L3:
	;
	v18 = l0 + int32(1)
	v19 = int32(0)
	goto L1
L4:
	;
	return v151
L5:
	;
	v25 = v18
	v27 = v3
	v28 = v20
	goto L6
L6:
	;
	v38 = base.I32_extend8_s(v28) + v27*int32(10) - int32(48)
	if int32(167) < v38 {
		v151 = v3
		goto L4
	} else {
		goto L8
	}
L7:
	;
	if v38 < int32(0) {
		v151 = v3
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v42 = v25 + int32(1)
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+1)))
	if base.Ui32(v43-int32(48)) < base.Ui32(int32(10)) {
		v25 = v42
		v27 = v38
		v28 = v43
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v51 = v38 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v53 != int32(58) {
		v131 = v42
		v136 = v51
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v19 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v56-int32(48)) {
		v151 = v3
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v64 = v25 + int32(2)
	v66 = int32(0)
	v67 = v56
	goto L14
L14:
	;
	v77 = base.I32_extend8_s(v67) + v66*int32(10) - int32(48)
	if int32(59) < v77 {
		v151 = v3
		goto L4
	} else {
		goto L16
	}
L15:
	;
	if v77 < int32(0) {
		v151 = v3
		goto L4
	} else {
		goto L18
	}
L16:
	;
	v81 = v64 + int32(1)
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+1)))
	if base.Ui32(v82-int32(48)) < base.Ui32(int32(10)) {
		v64 = v81
		v66 = v77
		v67 = v82
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v91 = v77*int32(60) + v51
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v93 != int32(58) {
		v131 = v81
		v136 = v91
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64)+2)))
	if base.Ui32(int32(9)) < base.Ui32(v96-int32(48)) {
		v151 = v3
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v104 = v64 + int32(2)
	v106 = v96
	v107 = int32(0)
	goto L21
L21:
	;
	v117 = base.I32_extend8_s(v106) + v107*int32(10) - int32(48)
	if int32(60) < v117 {
		v151 = v3
		goto L4
	} else {
		goto L23
	}
L22:
	;
	if v117 < int32(0) {
		v151 = v3
		goto L4
	} else {
		goto L25
	}
L23:
	;
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v104)+1)))
	v122 = v104 + int32(1)
	if base.Ui32(v120-int32(48)) < base.Ui32(int32(10)) {
		v104 = v122
		v106 = v120
		v107 = v117
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v129 = v117 + v91
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v129
	v131 = v122
	v136 = v129
	goto L11
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0) - v136
	goto L28
L27:
	;
	goto L28
L28:
	;
	v151 = v131
	goto L4
}
func F_getvacant(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int64
	_ = v289
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 < v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	if v188 != 0 {
		goto L41
	} else {
		goto L42
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v11 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v23 = v20 + v11<<(uint(int32(5))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)) = uint16(v19)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = int64(0)
	v29 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v17 + v11*v18<<(uint(v29)%32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v33 + v34*v11<<(uint(v29)%32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v40 + v41*v11<<(uint(int32(3))%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v47 <= v19 {
		v180 = v23
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v79 = base.I32_div_s(v12<<(uint(int32(1))%32), int32(3))
	v80 = int32(2)
	if v79 < (l2-l3)>>(uint(v80)%32) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v53 = v19
	goto L6
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60+v53<<(uint(int32(2))%32)))) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v66+v53<<(uint(int32(3))%32)))) = v64
	v73 = v53 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v73 < v74 {
		v53 = v73
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v180 = v23
	goto L1
L8:
	;
	goto L7
L9:
	;
	v87 = l2 - v79<<(uint(v80)%32)
	goto L11
L10:
	;
	v87 = l3
	goto L11
L11:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v92 = v89 + v12<<(uint(int32(5))%32)
	if base.Ui32(v88) < base.Ui32(v92) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v167 + int32(32)
	v180 = v167
	goto L1
L13:
	;
	v96 = v88
	goto L16
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(v89) < base.Ui32(v88) {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if base.Ui32(v87) <= base.Ui32(v104) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v107 = v104
	goto L20
L19:
	;
	v107 = int32(0)
	goto L20
L20:
	;
	if v107 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+8)))
	if v110&int32(4) == int32(0) {
		v167 = v96
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v116 = v96 + int32(32)
	if base.Ui32(v116) < base.Ui32(v92) {
		v96 = v116
		goto L16
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	goto L17
L26:
	;
	v132 = v89
	goto L29
L27:
	;
	goto L28
L28:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v159 != 0 {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	if base.Ui32(v87) <= base.Ui32(v139) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L28
L31:
	;
	v147 = v132 + int32(32)
	if base.Ui32(v147) < base.Ui32(v88) {
		v132 = v147
		goto L29
	} else {
		goto L37
	}
L32:
	;
	v142 = v139
	goto L34
L33:
	;
	v142 = int32(0)
	goto L34
L34:
	;
	if v142 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+8)))
	if v143&int32(4) != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v167 = v132
	goto L12
L37:
	;
	goto L30
L38:
	;
	v161 = v159
	goto L40
L39:
	;
	v161 = int32(15)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v161
	return int32(0)
L41:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+16)))
	v190 = v189
	v193 = v188
	goto L44
L42:
	;
	goto L43
L43:
	;
	v225 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v225 < v227 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v193)+24))
	v201 = base.I32_extend16_s(v190)
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v200+v201<<(uint(int32(2))%32)))) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v193)+28))
	v210 = v207 + v201<<(uint(int32(3))%32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v205
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210)+4)))
	if v211 != 0 {
		v190 = v214
		v193 = v211
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	goto L45
L47:
	;
	v231 = v227
	v236 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	if v335&int32(2) == int32(0) {
		v350 = v335
		goto L65
	} else {
		goto L66
	}
L50:
	;
	v242 = v236 << (uint(int32(2)) % 32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v180)+24))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242+v243)))
	if v245 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	if v246 != v180 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v312 = v231
	goto L54
L54:
	;
	v323 = v236 + int32(1)
	if v323 < v312 {
		v231 = v312
		v236 = v323
		goto L50
	} else {
		goto L64
	}
L55:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v180)+24))
	v303 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v301+v242))) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v180)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v305+v236<<(uint(int32(3))%32)))) = v303
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v312 = v311
	goto L54
L56:
	;
	v256 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+16)))
	v257 = v246
	v260 = v256
	goto L60
L57:
	;
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+16)))
	if v236 != v248 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v180)+28))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v250+v236<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v245)+12)) = v254
	goto L55
L59:
	;
	v282 = int32(3)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v180)+28))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v285+v236<<(uint(v282)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v281+v280<<(uint(v282)%32)))) = v289
	goto L55
L60:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v257)+28))
	v270 = v267 + v260<<(uint(int32(3))%32)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v271 == int32(0) {
		v280 = v260
		v281 = v267
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v257)+28))
	v280 = base.I32_extend16_s(v260)
	v281 = v278
	goto L59
L62:
	;
	v275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v270)+4)))
	if base.B2i32(v271 != v180)|base.B2i32(v236 != v275) != 0 {
		v257 = v271
		v260 = v275
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L51
L65:
	;
	if v350&int32(8) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L66:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if base.Ui32(v340) <= base.Ui32(v341) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v345 = v341
	goto L69
L68:
	;
	v345 = int32(0)
	goto L69
L69:
	;
	if base.B2i32(v340 == v341)|v345 != 0 {
		v350 = v335
		goto L65
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v340
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	v350 = v348
	goto L65
L71:
	;
	return v180
L72:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if base.Ui32(v356) <= base.Ui32(v357) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v361 = v357
	goto L75
L74:
	;
	v361 = int32(0)
	goto L75
L75:
	;
	if base.B2i32(v356 == v357)|v361 != 0 {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v356
	goto L71
}
func F_ginarrayconsistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
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
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v116 int64
	_ = v116
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = base.I32_wrap_i64(v16) & int32(_a_F_ginarrayconsistent_0)
	switch v19 - int32(1) {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L3
	case 3:
		goto L5
	default:
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v116
L2:
	;
	v116 = int64(1)
	goto L1
L3:
	;
	v100 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v100)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v66)
	v68 = int32(0)
	v69 = int64(1)
	if v14 <= v68 {
		v116 = v69
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v45)
	if v14 <= v45 {
		goto L2
	} else {
		goto L16
	}
L7:
	;
	v22 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v22)
	if v14 <= v22 {
		v116 = v7
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v27 = v22
	goto L9
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v15))))
	if v35 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v116 = v7
	goto L1
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v12))))
	if v39 != int32(1) {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v43 = v27 + int32(1)
	if v43 != v14 {
		v27 = v43
		goto L9
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L10
L16:
	;
	v50 = v45
	goto L17
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v15))))
	if v58 != int32(1) {
		v116 = v7
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L2
L19:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v12))))
	if v62 != 0 {
		v116 = v7
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v64 = v50 + int32(1)
	if v14 != v64 {
		v50 = v64
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v72 = v68
	goto L23
L23:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v15))))
	if v80 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v116 = int64(0)
	goto L1
L25:
	;
	v82 = v72 + int32(1)
	if v14 != v82 {
		v72 = v82
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v116 = v69
	goto L1
L29:
	;
	return int64(0)
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	F_errmsg_internal(m, int32(_a_F_ginarrayconsistent_1), v10)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_ginarrayconsistent_2), int32(219), int32(_a_F_ginarrayconsistent_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gingetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v315 int64
	_ = v315
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v382 int32
	_ = v382
	var v407 int64
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v519 int32
	_ = v519
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v769 int32
	_ = v769
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int64
	_ = v818
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int64
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1033 int64
	_ = v1033
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1151 int64
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1238 int32
	_ = v1238
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1275 int64
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1316 int32
	_ = v1316
	var v1326 int32
	_ = v1326
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1356 int32
	_ = v1356
	var v1368 int32
	_ = v1368
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1470 int64
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1485 int32
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1509 int32
	_ = v1509
	var v1510 int64
	_ = v1510
	var v1516 int64
	_ = v1516
	var v1517 int64
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1564 int32
	_ = v1564
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1963 int32
	_ = v1963
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v2027 int64
	_ = v2027
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2064 int64
	_ = v2064
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2100 int64
	_ = v2100
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2124 int32
	_ = v2124
	var v2145 int32
	_ = v2145
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2188 int32
	_ = v2188
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int64
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2218 int32
	_ = v2218
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2310 int32
	_ = v2310
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2328 int32
	_ = v2328
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2373 int32
	_ = v2373
	var v2375 int32
	_ = v2375
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2434 int32
	_ = v2434
	var v2440 int32
	_ = v2440
	var v2442 int32
	_ = v2442
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2454 int32
	_ = v2454
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2464 int64
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2481 int32
	_ = v2481
	var v2482 int64
	_ = v2482
	var v2483 int64
	_ = v2483
	var v2484 int64
	_ = v2484
	var v2485 int64
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int64
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2514 int64
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2538 int32
	_ = v2538
	var v2567 int32
	_ = v2567
	var v2573 int32
	_ = v2573
	var v2575 int32
	_ = v2575
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2584 int32
	_ = v2584
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2611 int32
	_ = v2611
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2623 int32
	_ = v2623
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2631 int32
	_ = v2631
	var v2635 int32
	_ = v2635
	var v2637 int32
	_ = v2637
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2670 int32
	_ = v2670
	var v2677 int32
	_ = v2677
	var v2708 int32
	_ = v2708
	var v2712 int32
	_ = v2712
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2735 int32
	_ = v2735
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2744 int32
	_ = v2744
	var v2746 int32
	_ = v2746
	var v2747 int32
	_ = v2747
	var v2752 int32
	_ = v2752
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2771 int32
	_ = v2771
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2779 int32
	_ = v2779
	var v2785 int32
	_ = v2785
	var v2787 int32
	_ = v2787
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2806 int32
	_ = v2806
	var v2809 int64
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2812 int32
	_ = v2812
	var v2814 int32
	_ = v2814
	var v2825 int32
	_ = v2825
	var v2826 int64
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2833 int32
	_ = v2833
	var v2835 int32
	_ = v2835
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2842 int32
	_ = v2842
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2859 int32
	_ = v2859
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2906 int32
	_ = v2906
	var v2909 int32
	_ = v2909
	var v2912 int32
	_ = v2912
	var v2914 int32
	_ = v2914
	var v2916 int32
	_ = v2916
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2956 int32
	_ = v2956
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2970 int32
	_ = v2970
	var v2974 int32
	_ = v2974
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2986 int32
	_ = v2986
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2991 int32
	_ = v2991
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3005 int32
	_ = v3005
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3019 int32
	_ = v3019
	var v3020 int32
	_ = v3020
	var v3030 int32
	_ = v3030
	var v3031 int32
	_ = v3031
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3037 int32
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3055 int32
	_ = v3055
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3076 int32
	_ = v3076
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3083 int32
	_ = v3083
	var v3089 int32
	_ = v3089
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3097 int32
	_ = v3097
	var v3098 int32
	_ = v3098
	var v3100 int32
	_ = v3100
	var v3134 int32
	_ = v3134
	var v3136 int32
	_ = v3136
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3179 int32
	_ = v3179
	var v3183 int32
	_ = v3183
	var v3187 int32
	_ = v3187
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3262 int32
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3269 int32
	_ = v3269
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3311 int32
	_ = v3311
	var v3313 int32
	_ = v3313
	var v3318 int32
	_ = v3318
	var v3348 int32
	_ = v3348
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3356 int32
	_ = v3356
	var v3360 int32
	_ = v3360
	var v3364 int32
	_ = v3364
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3371 int32
	_ = v3371
	var v3377 int32
	_ = v3377
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3412 int32
	_ = v3412
	var v3414 int32
	_ = v3414
	var v3417 int32
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3420 int32
	_ = v3420
	var v3425 int32
	_ = v3425
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3430 int32
	_ = v3430
	var v3432 int32
	_ = v3432
	var v3435 int32
	_ = v3435
	var v3471 int32
	_ = v3471
	var v3472 int32
	_ = v3472
	var v3478 int32
	_ = v3478
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3517 int32
	_ = v3517
	var v3548 int32
	_ = v3548
	var v3549 int32
	_ = v3549
	var v3552 int32
	_ = v3552
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3564 int32
	_ = v3564
	var v3593 int32
	_ = v3593
	var v3599 int32
	_ = v3599
	var v3630 int32
	_ = v3630
	var v3634 int32
	_ = v3634
	var v3636 int32
	_ = v3636
	var v3638 int32
	_ = v3638
	var v3639 int32
	_ = v3639
	var v3640 int32
	_ = v3640
	var v3644 int32
	_ = v3644
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3649 int32
	_ = v3649
	var v3653 int32
	_ = v3653
	var v3664 int32
	_ = v3664
	var v3689 int32
	_ = v3689
	var v3691 int32
	_ = v3691
	var v3694 int32
	_ = v3694
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3702 int32
	_ = v3702
	var v3705 int32
	_ = v3705
	var v3706 int32
	_ = v3706
	var v3708 int32
	_ = v3708
	var v3713 int32
	_ = v3713
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3746 int32
	_ = v3746
	var v3748 int32
	_ = v3748
	var v3750 int32
	_ = v3750
	var v3754 int32
	_ = v3754
	var v3757 int32
	_ = v3757
	var v3758 int32
	_ = v3758
	var v3762 int32
	_ = v3762
	var v3793 int32
	_ = v3793
	var v3799 int32
	_ = v3799
	var v3801 int32
	_ = v3801
	var v3830 int32
	_ = v3830
	var v3831 int32
	_ = v3831
	var v3834 int32
	_ = v3834
	var v3838 int32
	_ = v3838
	var v3842 int32
	_ = v3842
	var v3844 int32
	_ = v3844
	var v3847 int32
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3883 int32
	_ = v3883
	var v3885 int32
	_ = v3885
	var v3887 int32
	_ = v3887
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3936 int32
	_ = v3936
	var v3937 int32
	_ = v3937
	var v3972 int32
	_ = v3972
	var v3976 int32
	_ = v3976
	var v3977 int32
	_ = v3977
	var v3981 int32
	_ = v3981
	var v4004 int64
	_ = v4004
	var v4009 int32
	_ = v4009
	var v4011 int32
	_ = v4011
	var v4013 int32
	_ = v4013
	var v4015 int32
	_ = v4015
	var v4050 int32
	_ = v4050
	var v4052 int32
	_ = v4052
	var v4053 int32
	_ = v4053
	var v4057 int32
	_ = v4057
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4070 int32
	_ = v4070
	var v4074 int32
	_ = v4074
	var v4080 int32
	_ = v4080
	var v4097 int32
	_ = v4097
	var v4100 int32
	_ = v4100
	var v4101 int32
	_ = v4101
	var v4104 int32
	_ = v4104
	var v4105 int32
	_ = v4105
	var v4110 int32
	_ = v4110
	var v4114 int32
	_ = v4114
	var v4115 int32
	_ = v4115
	var v4117 int32
	_ = v4117
	var v4120 int32
	_ = v4120
	var v4122 int64
	_ = v4122
	var v4125 int64
	_ = v4125
	var v4128 int64
	_ = v4128
	var v4129 int64
	_ = v4129
	var v4130 int64
	_ = v4130
	var v4133 int64
	_ = v4133
	var v4139 int32
	_ = v4139
	var v4140 int32
	_ = v4140
	var v4150 int32
	_ = v4150
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4167 int32
	_ = v4167
	var v4180 int32
	_ = v4180
	var v4184 int32
	_ = v4184
	var v4185 int32
	_ = v4185
	var v4186 int32
	_ = v4186
	var v4188 int64
	_ = v4188
	var v4190 int32
	_ = v4190
	var v4196 int32
	_ = v4196
	var v4201 int64
	_ = v4201
	var v4203 int32
	_ = v4203
	var v4205 int32
	_ = v4205
	var v4210 int32
	_ = v4210
	var v4211 int32
	_ = v4211
	var v4212 int32
	_ = v4212
	var v4214 int64
	_ = v4214
	var v4216 int32
	_ = v4216
	var v4222 int32
	_ = v4222
	var v4228 int32
	_ = v4228
	var v4229 int32
	_ = v4229
	var v4230 int32
	_ = v4230
	var v4231 int64
	_ = v4231
	var v4232 int32
	_ = v4232
	var v4234 int64
	_ = v4234
	var v4248 int32
	_ = v4248
	var v4249 int32
	_ = v4249
	var v4250 int32
	_ = v4250
	var v4253 int32
	_ = v4253
	var v4256 int32
	_ = v4256
	var v4257 int32
	_ = v4257
	var v4294 int32
	_ = v4294
	var v4295 int32
	_ = v4295
	var v4297 int32
	_ = v4297
	var v4298 int32
	_ = v4298
	var v4306 int32
	_ = v4306
	var v4309 int32
	_ = v4309
	var v4314 int32
	_ = v4314
	var v4323 int32
	_ = v4323
	var v4361 int32
	_ = v4361
	var v4362 int32
	_ = v4362
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4371 int32
	_ = v4371
	var v4375 int32
	_ = v4375
	var v4378 int32
	_ = v4378
	var v4382 int32
	_ = v4382
	var v4396 int32
	_ = v4396
	var v4398 int64
	_ = v4398
	var v4414 int32
	_ = v4414
	var v4415 int32
	_ = v4415
	var v4418 int32
	_ = v4418
	var v4419 int32
	_ = v4419
	var v4444 int32
	_ = v4444
	var v4448 int32
	_ = v4448
	var v4449 int32
	_ = v4449
	var v4450 int32
	_ = v4450
	var v4452 int64
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4460 int32
	_ = v4460
	var v4465 int64
	_ = v4465
	var v4467 int32
	_ = v4467
	var v4469 int32
	_ = v4469
	var v4472 int32
	_ = v4472
	var v4473 int32
	_ = v4473
	var v4474 int32
	_ = v4474
	var v4476 int64
	_ = v4476
	var v4478 int32
	_ = v4478
	var v4484 int32
	_ = v4484
	var v4490 int32
	_ = v4490
	var v4491 int32
	_ = v4491
	var v4492 int32
	_ = v4492
	var v4493 int64
	_ = v4493
	var v4495 int64
	_ = v4495
	var v4509 int32
	_ = v4509
	var v4510 int32
	_ = v4510
	var v4511 int32
	_ = v4511
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4522 int32
	_ = v4522
	var v4523 int32
	_ = v4523
	var v4527 int32
	_ = v4527
	var v4555 int32
	_ = v4555
	var v4558 int32
	_ = v4558
	var v4559 int32
	_ = v4559
	var v4563 int64
	_ = v4563
	var v4573 int32
	_ = v4573
	var v4577 int32
	_ = v4577
	var v4588 int32
	_ = v4588
	var v4608 int32
	_ = v4608
	var v4612 int32
	_ = v4612
	var v4613 int32
	_ = v4613
	var v4614 int64
	_ = v4614
	var v4615 int64
	_ = v4615
	var v4618 int64
	_ = v4618
	var v4624 int32
	_ = v4624
	var v4625 int32
	_ = v4625
	var v4626 int32
	_ = v4626
	var v4628 int32
	_ = v4628
	var v4631 int32
	_ = v4631
	var v4634 int32
	_ = v4634
	var v4636 int32
	_ = v4636
	var v4639 int32
	_ = v4639
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4644 int32
	_ = v4644
	var v4645 int32
	_ = v4645
	var v4652 int32
	_ = v4652
	var v4653 int32
	_ = v4653
	var v4654 int32
	_ = v4654
	var v4655 int32
	_ = v4655
	var v4664 int32
	_ = v4664
	var v4680 int32
	_ = v4680
	var v4702 int32
	_ = v4702
	var v4704 int64
	_ = v4704
	var v4707 int64
	_ = v4707
	var v4710 int64
	_ = v4710
	var v4722 int32
	_ = v4722
	var v4753 int32
	_ = v4753
	var v4757 int32
	_ = v4757
	var v4758 int32
	_ = v4758
	var v4761 int32
	_ = v4761
	var v4763 int32
	_ = v4763
	var v4765 int64
	_ = v4765
	var v4766 int64
	_ = v4766
	var v4769 int64
	_ = v4769
	var v4773 int64
	_ = v4773
	var v4775 int32
	_ = v4775
	var v4777 int32
	_ = v4777
	var v4779 int32
	_ = v4779
	var v4780 int32
	_ = v4780
	var v4782 int32
	_ = v4782
	var v4784 int32
	_ = v4784
	var v4789 int32
	_ = v4789
	var v4790 int32
	_ = v4790
	var v4825 int32
	_ = v4825
	var v4826 int32
	_ = v4826
	var v4827 int32
	_ = v4827
	var v4830 int32
	_ = v4830
	var v4832 int32
	_ = v4832
	var v4834 int32
	_ = v4834
	var v4839 int32
	_ = v4839
	var v4873 int32
	_ = v4873
	var v4875 int32
	_ = v4875
	var v4876 int32
	_ = v4876
	var v4879 int32
	_ = v4879
	var v4880 int32
	_ = v4880
	var v4881 int32
	_ = v4881
	var v4886 int32
	_ = v4886
	var v4895 int32
	_ = v4895
	var v4901 int32
	_ = v4901
	var v4908 int32
	_ = v4908
	var v4913 int32
	_ = v4913
	var v4914 int32
	_ = v4914
	var v4915 int64
	_ = v4915
	var v4920 int32
	_ = v4920
	var v4927 int32
	_ = v4927
	var v4928 int32
	_ = v4928
	var v4930 int32
	_ = v4930
	var v4932 int32
	_ = v4932
	var v4935 int32
	_ = v4935
	var v4937 int32
	_ = v4937
	var v4942 int64
	_ = v4942
	var v4946 int64
	_ = v4946
	var v4958 int32
	_ = v4958
	var v4959 int32
	_ = v4959
	var v4960 int32
	_ = v4960
	var v4961 int32
	_ = v4961
	var v4968 int32
	_ = v4968
	var v4970 int32
	_ = v4970
	var v4973 int32
	_ = v4973
	var v5000 int32
	_ = v5000
	var v5001 int32
	_ = v5001
	var v5010 int32
	_ = v5010
	var v5014 int32
	_ = v5014
	var v5048 int32
	_ = v5048
	var v5050 int32
	_ = v5050
	var v5059 int32
	_ = v5059
	var v5085 int32
	_ = v5085
	var v5088 int32
	_ = v5088
	var v5089 int32
	_ = v5089
	var v5092 int32
	_ = v5092
	var v5096 int32
	_ = v5096
	var v5106 int32
	_ = v5106
	var v5136 int32
	_ = v5136
	var v5142 int32
	_ = v5142
	var v5145 int32
	_ = v5145
	var v5146 int32
	_ = v5146
	var v5147 int32
	_ = v5147
	var v5155 int32
	_ = v5155
	var v5160 int32
	_ = v5160
	var v5166 int32
	_ = v5166
	var v5189 int64
	_ = v5189
	v3 = int32(0)
	v27 = int64(0)
	v34 = m.G0
	v36 = v34 - int32(_a_F_gingetbitmap_0)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ginFreeScanKeys(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v43 = m.G0
	v45 = v43 - int32(96)
	m.G0 = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+88)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v45)+80)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v45)+72)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v45)+64)) = v49
	v57 = int32(_a_F_gingetbitmap_1)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v60
	v62 = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v63 <= v62 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v66 = v62
	goto L5
L4:
	;
	v66 = v63
	goto L5
L5:
	;
	v69 = F_palloc(m, v66*int32(104))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[2]))) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[4]))) = v69
	v77 = F_palloc(m, int32(128))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[5]))) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[6]))) = v77
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v82 <= v79 {
		v519 = v3
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3])))
	if v545 != 0 {
		goto L72
	} else {
		goto L73
	}
L9:
	;
	v111 = v3
	v119 = v3
	goto L10
L10:
	;
	v137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v45)+60)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v45)+56)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v45)+52)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v45)+48)) = v137
	v149 = v47 + v119*int32(56)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v150&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v510 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[5]))) = uint8(v510)
	v519 = v111
	goto L8
L12:
	;
	goto L11
L13:
	;
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v149)+4)))
	v155 = v153 - int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v48+int32(_a_F_gingetbitmap_2)+v155<<(uint(int32(2))%32))))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v149)+48))
	v164 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v149)+6)))
	v165 = F_FunctionCall7Coll(m, v48+int32(1936)+v155*int32(28), v162, v163, base.I64_extend_i32_u(v45+int32(16)), v164, base.I64_extend_i32_u(v45+int32(60)), base.I64_extend_i32_u(v45+int32(56)), base.I64_extend_i32_u(v45+int32(52)), base.I64_extend_i32_u(v45+int32(48)))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v167 = base.I32_wrap_i64(v165)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	if base.Ui32(int32(3)) <= base.Ui32(v168) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v171 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+48)) = v171
	v174 = v171
	goto L17
L16:
	;
	v174 = v168
	goto L17
L17:
	;
	if v167 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v187 = base.B2i32(v174 != int32(0)) | v111
	v188 = F_palloc0(m, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L24
	}
L19:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if int32(0) < v175 {
		v184 = v175
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v174 == int32(0) {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v181 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v181
	v184 = v181
	goto L18
L24:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	if v191 == int32(0) {
		v252 = v187
		v253 = v190
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+4)))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+6)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v149)+48))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v45)+60))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3])))
	v285 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3]))) = v284 + v285
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[4])))
	v291 = v288 + v284*int32(104)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v253
	v296 = v253 + v285
	v297 = F_palloc_mul(m, int32(4), v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L34
	}
L26:
	;
	v194 = int32(0)
	if v190 <= v194 {
		v252 = v187
		v253 = v190
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v203 = v194
	v204 = v187
	v205 = v190
	goto L28
L28:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v203))))
	if v232 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v252 = v240
	v253 = v241
	goto L25
L30:
	;
	v236 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v203+v188))) = uint8(v236)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v240 = v236
	v241 = v238
	goto L32
L31:
	;
	v240 = v204
	v241 = v205
	goto L32
L32:
	;
	v243 = v203 + int32(1)
	if v243 < v241 {
		v203 = v243
		v204 = v240
		v205 = v241
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v297
	v302 = F_palloc0_mul(m, int32(1), v296)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v291)+86)) = uint8(base.B2i32(v280 == int32(2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v291)+84)) = uint16(v278)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+80)) = v280
	*(*uint16)(unsafe.Add(mBase, uint32(v291)+76)) = uint16(v279)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+72)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v291)+68)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v291)+64)) = v167
	*(*int64)(unsafe.Add(mBase, uint32(v291)+56)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v291)+28)) = v302
	v315 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v291)+12)) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v291)+20)) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v291)+88)) = v315
	v321 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v291)+96)) = uint8(v321)
	v325 = v48 + int32(4)
	if v280 == int32(3) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v253 != 0 {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+32)) = int32(58)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+36)) = int32(59)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291)+84)))
	v338 = v325 + v335*int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+44)) = v338 + int32(3696)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+40)) = v338 + int32(2800)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v325+v335<<(uint(int32(2))%32))+uint32(_c_F_gingetbitmap[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+48)) = v350
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v338+int32(2804))))
	if v356 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v357 = int32(60)
	goto L42
L41:
	;
	v357 = int32(61)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+32)) = v357
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v338+int32(3700))))
	if v363 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v364 = int32(62)
	goto L45
L44:
	;
	v364 = int32(63)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+36)) = v364
	goto L36
L46:
	;
	v382 = v321
	goto L49
L47:
	;
	goto L48
L48:
	;
	switch v280 - int32(1) {
	case 0:
		v474 = int32(2)
		goto L60
	default:
		goto L59
	case 2:
		goto L61
	}
L49:
	;
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v167+v382<<(uint(int32(3))%32))))
	v409 = int32(*(*int8)(unsafe.Add(mBase, uint32(v382+v188))))
	v410 = int32(0)
	if v282 == v410 {
		v421 = v410
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	if v283 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v413 = int32(0)
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v278)+uint32(_c_F_gingetbitmap[8]))))
	if v414&int32(1) == v413 {
		v421 = v413
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382+v282))))
	v421 = v420
	goto L51
L54:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v283+v382<<(uint(int32(2))%32))))
	v427 = v425
	goto L56
L55:
	;
	v427 = int32(0)
	goto L56
L56:
	;
	v428 = F_ginFillScanEntry(m, v48, v278, v279, v280, v407, v409, v421, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v430+v382<<(uint(int32(2))%32)))) = v428
	v436 = v382 + int32(1)
	if v436 != v253 {
		v382 = v436
		goto L49
	} else {
		goto L58
	}
L58:
	;
	goto L50
L59:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	if v496 != int32(2) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = v475 + int32(1)
	v479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291)+84)))
	v480 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v291)+80))
	v486 = F_ginFillScanEntry(m, v48, v479, v480, v481, int64(0), base.I32_extend8_s(v474), v480, v480)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v474 = int32(255)
	goto L60
L62:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v488+v475<<(uint(int32(2))%32)))) = v486
	goto L59
L63:
	;
	v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v149)+4)))
	v501 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v499)+63)) = uint8(v501)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v504 = v119 + int32(1)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v504 < v505 {
		v111 = v252
		v119 = v504
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v519 = v252
	goto L8
L67:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1064)+uint32(_c_F_gingetbitmap[5]))))
	if v1065 != 0 {
		v5166 = v36
		v5189 = v27
		goto L135
	} else {
		goto L136
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L130
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v58
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+272))
	if v1016 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L70:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[5]))))
	if v971 != 0 {
		goto L69
	} else {
		goto L118
	}
L71:
	;
	if v519&int32(1) == int32(0) {
		goto L69
	} else {
		goto L117
	}
L72:
	;
	v549 = v3
	v555 = v545
	v558 = int32(0)
	goto L75
L73:
	;
	goto L74
L74:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[5]))))
	if v803 != 0 {
		goto L71
	} else {
		goto L103
	}
L75:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[4])))
	v583 = v580 + v549*int32(104)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+80))
	if v584 != int32(2) {
		v618 = v555
		v620 = v558
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if int32(0) < v620 {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	v622 = v549 + int32(1)
	if base.Ui32(v622) < base.Ui32(v618) {
		v549 = v622
		v555 = v618
		v558 = v620
		goto L75
	} else {
		goto L83
	}
L78:
	;
	v587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+84)))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v587)+63)))
	if v589 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v592 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v583)+86)) = uint8(v592)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	*(*int32)(unsafe.Add(mBase, uint32(v583))) = v594 + int32(1)
	v604 = F_ginFillScanEntry(m, v48, v587, v592, int32(2), int64(0), int32(-1), v592, v592)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v618 = v555
	v620 = v558 + int32(1)
	goto L77
L82:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v583)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v606+v594<<(uint(int32(2))%32)))) = v604
	v611 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+84)))
	v613 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v611)+63)) = uint8(v613)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3])))
	v618 = v615
	v620 = v558
	goto L77
L83:
	;
	goto L76
L84:
	;
	v628 = F_palloc(m, v618*int32(104))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v769 = v618
	goto L86
L86:
	;
	if v769 != 0 {
		goto L71
	} else {
		goto L102
	}
L87:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3])))
	if v630 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v632 = int32(0)
	v636 = v632
	v640 = v632
	v646 = v630 - v620
	goto L91
L89:
	;
	v730 = int32(0)
	goto L90
L90:
	;
	if v730 != 0 {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[4])))
	v670 = v667 + v640*int32(104)
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+86)))
	if v671 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v730 = v692 * int32(104)
	goto L90
L93:
	;
	v691 = v640 + int32(1)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v691) < base.Ui32(v692) {
		v636 = v688
		v640 = v691
		v646 = v689
		goto L91
	} else {
		goto L97
	}
L94:
	;
	v674 = int32(104)
	base.MemoryCopy(m, v628+v646*v674, v670, v674)
	v688 = v636
	v689 = v646 + int32(1)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v681 = int32(104)
	base.MemoryCopy(m, v628+v636*v681, v670, v681)
	v688 = v636 + int32(1)
	v689 = v646
	goto L93
L97:
	;
	goto L92
L98:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[4])))
	base.MemoryCopy(m, v731, v628, v730)
	goto L100
L99:
	;
	goto L100
L100:
	;
	F_pfree(m, v628)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3])))
	v769 = v735
	goto L86
L102:
	;
	goto L74
L103:
	;
	v804 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[3]))) = v804
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_gingetbitmap[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v806))) = int64(0)
	v811 = F_palloc_mul(m, int32(4), v804)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806)+8)) = v811
	v814 = int32(1)
	v816 = F_palloc0_mul(m, v814, v814)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v818 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v806)+56)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v806)+28)) = v816
	*(*int64)(unsafe.Add(mBase, uint32(v806)+64)) = v818
	*(*int64)(unsafe.Add(mBase, uint32(v806)+70)) = v818
	v825 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v806)+86)) = uint8(v825)
	v827 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v806)+84)) = uint16(v827)
	*(*int32)(unsafe.Add(mBase, uint32(v806)+80)) = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v806)+12)) = v818
	*(*int64)(unsafe.Add(mBase, uint32(v806)+20)) = v818
	*(*int64)(unsafe.Add(mBase, uint32(v806)+88)) = v818
	*(*uint8)(unsafe.Add(mBase, uint32(v806)+96)) = uint8(v825)
	goto L107
L106:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v806)))
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = v883 + int32(1)
	v887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v806)+84)))
	v888 = int32(0)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v806)+80))
	v894 = F_ginFillScanEntry(m, v48, v887, v888, v889, int64(0), int32(-1), v888, v888)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L116
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806)+32)) = int32(58)
	*(*int32)(unsafe.Add(mBase, uint32(v806)+36)) = int32(59)
	goto L106
L116:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v806)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v896+v883<<(uint(int32(2))%32)))) = v894
	goto L70
L117:
	;
	goto L70
L118:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ginGetStats(m, v972, v45+int32(16))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
	if v977 <= int32(0) {
		goto L68
	} else {
		goto L120
	}
L120:
	;
	goto L69
L121:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1032 != 0 {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015)+268)))
	if v1019 != int32(1) {
		goto L121
	} else {
		goto L125
	}
L123:
	;
	v1026 = v1016
	goto L124
L124:
	;
	v1027 = *(*int64)(unsafe.Add(mBase, uint32(v1026)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1026)+16)) = v1027 + int64(1)
	goto L121
L125:
	;
	F_pgstat_assoc_relation(m, v1015)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+272))
	v1026 = v1025
	goto L124
L127:
	;
	v1033 = *(*int64)(unsafe.Add(mBase, uint32(v1032)))
	*(*int64)(unsafe.Add(mBase, uint32(v1032))) = v1033 + int64(1)
	goto L129
L128:
	;
	goto L129
L129:
	;
	m.G0 = v45 + int32(96)
	goto L67
L130:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_gingetbitmap_3), int32(0))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v1052 + int32(4)
	F_errhint(m, int32(_a_F_gingetbitmap_4), v45)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_5), int32(480), int32(_a_F_gingetbitmap_6))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	m.G0 = v5166 + int32(_a_F_gingetbitmap_0)
	return v5189
L136:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1068 = F_ReadBuffer(m, v1066, int32(0))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v1070, int32(0), v1072)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_LockBufferInternal(m, v1068, int32(1))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	if v1068 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+36))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[2])))
	if v2106 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L141:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+24))
	if v1096 == int32(-1) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1081+(v1068^int32(-1))<<(uint(int32(2))%32))))
	v1095 = v1087
	goto L141
L143:
	;
	goto L144
L144:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v1095 = v1089 + v1068<<(uint(int32(13))%32) + int32(-8192)
	goto L141
L145:
	;
	F_UnlockReleaseBuffer(m, v1068)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1102 = F_ReadBuffer(m, v1101, v1096)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v2072 = l0
	v2073 = l1
	v2077 = v36
	v2100 = v27
	goto L140
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v1102
	F_LockBufferInternal(m, v1102, int32(1))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v1108 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+40)) = uint16(v1108)
	F_UnlockReleaseBuffer(m, v1068)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+uint32(_c_F_gingetbitmap[3])))
	v1114 = F_palloc_mul(m, int32(1), v1113)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v1114
	v1119 = F_scanGetCandidate(m, l0, v36+int32(36))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if v1119 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1123 = l0
	v1124 = l1
	v1128 = v36
	v1134 = v1064
	v1137 = v36 + int32(44)
	v1151 = v27
	goto L157
L155:
	;
	v2036 = l0
	v2037 = l1
	v2041 = v36
	v2064 = v27
	goto L156
L156:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2041)+52))
	F_pfree(m, v2069)
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L1
	} else {
		goto L285
	}
L157:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+36))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+uint32(_c_F_gingetbitmap[3])))
	if v1158 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v2036 = v1247
	v2037 = v1248
	v2041 = v1252
	v2064 = v2027
	goto L156
L159:
	;
	v1163 = int32(0)
	goto L162
L160:
	;
	v1207 = int32(0)
	goto L161
L161:
	;
	if v1207 != 0 {
		goto L168
	} else {
		goto L169
	}
L162:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+uint32(_c_F_gingetbitmap[4])))
	v1196 = v1193 + v1163*int32(104)
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1196)))
	if v1197 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1207 = v1203
	goto L161
L164:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+28))
	base.MemoryFill(m, v1198, int32(0), v1197)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v1202 = v1163 + int32(1)
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v1202) < base.Ui32(v1203) {
		v1163 = v1202
		goto L162
	} else {
		goto L167
	}
L167:
	;
	goto L163
L168:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+52))
	base.MemoryFill(m, v1238, int32(0), v1207)
	goto L170
L169:
	;
	goto L170
L170:
	;
	v1247 = v1123
	v1248 = v1124
	v1252 = v1128
	v1253 = v1157
	v1258 = v1134
	v1261 = v1137
	v1263 = v1157 + int32(4)
	v1265 = v1157 + int32(144)
	v1266 = v1157 + int32(_a_F_gingetbitmap_2)
	v1275 = v1151
	goto L172
L171:
	;
	if v1732 != 0 {
		goto L260
	} else {
		goto L261
	}
L172:
	;
	v1280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+42)))
	v1281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+40)))
	v1282 = v1280 - v1281
	if v1282 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L256
	}
L174:
	;
	base.MemoryFill(m, v1281+v1252+int32(63), int32(0), v1282)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+36))
	if v1288 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+uint32(_c_F_gingetbitmap[3])))
	if v1307 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1292+(v1288^int32(-1))<<(uint(int32(2))%32))))
	v1306 = v1298
	goto L177
L179:
	;
	goto L180
L180:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v1306 = v1300 + v1288<<(uint(int32(13))%32) + int32(-8192)
	goto L177
L181:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1252)+40)) = uint16(v1733)
	v1764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1306)+16)))
	v1766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306+v1764)+6)))
	if v1766&int32(32) != 0 {
		goto L171
	} else {
		goto L244
	}
L182:
	;
	v1732 = int32(0)
	v1733 = v1280
	goto L181
L183:
	;
	goto L184
L184:
	;
	v1316 = v1307
	v1326 = int32(0)
	goto L185
L185:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+uint32(_c_F_gingetbitmap[4])))
	v1350 = v1347 + v1326*int32(104)
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)))
	if v1351 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v1729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+42)))
	v1732 = v1695
	v1733 = v1729
	goto L181
L187:
	;
	v1356 = v1351
	v1368 = int32(0)
	goto L190
L188:
	;
	v1695 = v1316
	goto L189
L189:
	;
	v1727 = v1326 + int32(1)
	if base.Ui32(v1727) < base.Ui32(v1695) {
		v1316 = v1695
		v1326 = v1727
		goto L185
	} else {
		goto L243
	}
L190:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+28))
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1386+v1368))))
	if v1388 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+uint32(_c_F_gingetbitmap[3])))
	v1695 = v1692
	goto L189
L192:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+8))
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1391+v1368<<(uint(int32(2))%32))))
	v1396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+40)))
	v1397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+42)))
	if base.Ui32(v1397) <= base.Ui32(v1396) {
		v1564 = v1397
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v1659 = v1356
	goto L194
L194:
	;
	v1690 = v1368 + int32(1)
	if base.Ui32(v1690) < base.Ui32(v1659) {
		v1356 = v1659
		v1368 = v1690
		goto L190
	} else {
		goto L242
	}
L195:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+52))
	v1646 = v1645 + v1326
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1646))))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+28))
	v1650 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1648+v1368))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1646))) = uint8(base.B2i32(v1647|v1650 != int32(0)))
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1350)))
	v1659 = v1655
	goto L194
L196:
	;
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1395)+9)))
	if v1595 != int32(1) {
		goto L195
	} else {
		goto L240
	}
L197:
	;
	v1401 = v1397
	v1402 = v1396
	goto L198
L198:
	;
	v1437 = int32(base.Ui32((v1401-v1402)&int32(_a_F_gingetbitmap_7))>>(uint(int32(1))%32)) + v1402
	v1439 = v1437 & int32(_a_F_gingetbitmap_8)
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1306+int32(20)+v1439<<(uint(int32(2))%32))))
	v1446 = v1306 + v1443&int32(_a_F_gingetbitmap_9)
	v1447 = F_gintuple_get_attrnum(m, v1263, v1446)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L200
	}
L199:
	;
	v1564 = v1552
	goto L196
L200:
	;
	v1449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1350)+84)))
	if base.Ui32(v1449) < base.Ui32(v1447) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v1557 = int32(_a_F_gingetbitmap_8)
	if base.Ui32(v1553&v1557) < base.Ui32(v1552&v1557) {
		v1401 = v1552
		v1402 = v1553
		goto L198
	} else {
		goto L239
	}
L202:
	;
	v1552 = v1437
	v1553 = v1402
	goto L201
L203:
	;
	goto L204
L204:
	;
	if base.Ui32(v1447) < base.Ui32(v1449) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1552 = v1401
	v1553 = v1437 + int32(1)
	goto L201
L206:
	;
	goto L207
L207:
	;
	v1455 = v1439 - int32(1)
	v1458 = v1455 + (v1252 - int32(-64))
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1458))))
	if v1459 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1470 = F_gintuple_get_key(m, v1263, v1446, v1252+int32(1088)+v1455)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v1475 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1395)+8)))
	if v1475 == int32(-1) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1252+int32(2112)+v1455<<(uint(int32(3))%32)))) = v1470
	v1473 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1458))) = uint8(v1473)
	goto L210
L212:
	;
	v1549 = base.B2i32(v1544 < int32(0))
	if v1544 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L213:
	;
	v1522 = int32(1)
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1395)+9)))
	if v1523 == v1522 {
		goto L228
	} else {
		goto L229
	}
L214:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+20))
	if v1478 != int32(2) {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252+int32(1088)+v1455))))
	if v1491 != v1475&int32(255) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1252+int32(1088)+v1455))))
	if v1485 == int32(3) {
		v1544 = int32(-1)
		goto L212
	} else {
		goto L218
	}
L218:
	;
	goto L213
L219:
	;
	if v1475 < base.I32_extend8_s(v1491) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	goto L221
L221:
	;
	if v1475 != 0 {
		goto L213
	} else {
		goto L225
	}
L222:
	;
	v1499 = int32(-1)
	goto L224
L223:
	;
	v1499 = int32(1)
	goto L224
L224:
	;
	v1544 = v1499
	goto L212
L225:
	;
	v1500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1395)+24)))
	v1502 = v1500 - int32(1)
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1266+v1502<<(uint(int32(2))%32))))
	v1510 = *(*int64)(unsafe.Add(mBase, uint32(v1395)))
	v1516 = *(*int64)(unsafe.Add(mBase, uint32(v1252+int32(2112)+v1455<<(uint(int32(3))%32))))
	v1517 = F_FunctionCall2Coll(m, v1265+v1502*int32(28), v1509, v1510, v1516)
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v1519 = base.I32_wrap_i64(v1517)
	if v1519 != 0 {
		v1544 = v1519
		goto L212
	} else {
		goto L227
	}
L227:
	;
	goto L213
L228:
	;
	v1526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+42)))
	v1533 = F_matchPartialInPendingList(m, v1263, v1306, v1439, v1526, v1395, v1252+int32(2112), v1252+int32(1088), v1252-int32(-64))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	v1535 = v1522
	goto L230
L230:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1536+v1368))) = uint8(v1535)
	v1539 = int32(_a_F_gingetbitmap_8)
	if base.Ui32(v1402&v1539) < base.Ui32(v1401&v1539) {
		goto L195
	} else {
		goto L232
	}
L231:
	;
	v1535 = v1533
	goto L230
L232:
	;
	v1564 = v1401
	goto L196
L233:
	;
	v1550 = v1402
	goto L235
L234:
	;
	v1550 = v1437 + int32(1)
	goto L235
L235:
	;
	if v1544 < int32(0) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1551 = v1437
	goto L238
L237:
	;
	v1551 = v1401
	goto L238
L238:
	;
	v1552 = v1551
	v1553 = v1550
	goto L201
L239:
	;
	goto L199
L240:
	;
	v1600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1252)+42)))
	v1607 = F_matchPartialInPendingList(m, v1263, v1306, v1564&int32(_a_F_gingetbitmap_8), v1600, v1395, v1252+int32(2112), v1252+int32(1088), v1252-int32(-64))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1609+v1368))) = uint8(v1607)
	goto L195
L242:
	;
	goto L191
L243:
	;
	goto L186
L244:
	;
	v1769 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1261)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1252)+60)) = uint16(v1769)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1261)))
	*(*int32)(unsafe.Add(mBase, uint32(v1252)+56)) = v1771
	v1775 = F_scanGetCandidate(m, v1247, v1252+int32(36))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	if v1775 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1778 = v1252 + int32(56)
	v1779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1261)+2)))
	v1780 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1261))))
	v1781 = int32(16)
	v1784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1778)+2)))
	v1785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1778))))
	if v1779|v1780<<(uint(v1781)%32) == v1784|v1785<<(uint(v1781)%32) {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	goto L248
L248:
	;
	goto L173
L249:
	;
	if v1795 != 0 {
		goto L172
	} else {
		goto L255
	}
L250:
	;
	goto L249
L251:
	;
	v1791 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1261)+4)))
	v1792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1778)+4)))
	if v1791 == v1792 {
		v1795 = int32(1)
		goto L250
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1795 = int32(0)
	goto L250
L254:
	;
	goto L253
L255:
	;
	goto L248
L256:
	;
	F_errmsg_internal(m, int32(_a_F_gingetbitmap_10), int32(0))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_11), int32(1814), int32(_a_F_gingetbitmap_12))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	v2034 = F_scanGetCandidate(m, v1247, v1252+int32(36))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L283
	}
L260:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+52))
	v1814 = int32(0)
	goto L263
L261:
	;
	goto L262
L262:
	;
	v1891 = int32(0)
	v1892 = int32(_a_F_gingetbitmap_1)
	v1893 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1258)))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1895
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+uint32(_c_F_gingetbitmap[3])))
	if v1898 != 0 {
		goto L270
	} else {
		goto L271
	}
L263:
	;
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814+v1810))))
	if v1845 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	goto L262
L265:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+uint32(_c_F_gingetbitmap[4])))
	v1852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1848+v1814*int32(104))+86)))
	if v1852 != int32(1) {
		v2027 = v1275
		goto L259
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1856 = v1814 + int32(1)
	if v1856 != v1732 {
		v1814 = v1856
		goto L263
	} else {
		goto L269
	}
L268:
	;
	goto L267
L269:
	;
	goto L264
L270:
	;
	v1902 = v1891
	v1906 = v1891
	goto L273
L271:
	;
	v1963 = v1891
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1893
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1258)))
	F_MemoryContextReset(m, v1991)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L1
	} else {
		goto L281
	}
L273:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+uint32(_c_F_gingetbitmap[4])))
	v1935 = v1932 + v1902*int32(104)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1935)+32))
	v1937 = m.T0[v1936].(func(*base.Module, int32) int32)(m, v1935)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L275
	}
L274:
	;
	v1963 = v1951
	goto L272
L275:
	;
	if v1937 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1893
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1258)))
	F_MemoryContextReset(m, v1943)
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v1946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1935)+95)))
	v1947 = int32(1)
	v1951 = base.B2i32(v1946|v1906&v1947 != int32(0))
	v1953 = v1902 + v1947
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v1953) < base.Ui32(v1954) {
		v1902 = v1953
		v1906 = v1951
		goto L273
	} else {
		goto L280
	}
L279:
	;
	v2027 = v1275
	goto L259
L280:
	;
	goto L274
L281:
	;
	F_tbm_add_tuples(m, v1248, v1261, int32(1), v1963)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v2027 = v1275 + int64(1)
	goto L259
L283:
	;
	if v2034 != 0 {
		v1123 = v1247
		v1124 = v1248
		v1128 = v1252
		v1134 = v1258
		v1137 = v1261
		v1151 = v2027
		goto L157
	} else {
		goto L284
	}
L284:
	;
	goto L158
L285:
	;
	v2072 = v2036
	v2073 = v2037
	v2077 = v2041
	v2100 = v2064
	goto L140
L286:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5142 = m.ExcPending
	if v5142 != 0 {
		goto L1
	} else {
		goto L696
	}
L287:
	;
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[3])))
	if v3311 != 0 {
		goto L496
	} else {
		goto L497
	}
L288:
	;
	v2110 = v2105 + int32(4)
	v2124 = int32(0)
	goto L289
L289:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[6])))
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2145+v2124<<(uint(int32(2))%32))))
	v2151 = v2149 + int32(652)
	v2153 = v2149 + int32(28)
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+8))
	goto L295
L290:
	;
	if v3174 == int32(0) {
		goto L287
	} else {
		goto L487
	}
L291:
	;
	F_freeGinBtreeStack(m, v2250)
	mBase = m.M
	v3171 = m.ExcPending
	if v3171 != 0 {
		goto L1
	} else {
		goto L485
	}
L292:
	;
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	F_UnlockBuffer(m, v3134)
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L1
	} else {
		goto L484
	}
L293:
	;
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1100))
	v2968 = m.T0[v2967].(func(*base.Module, int32, int32) int32)(m, v2077+int32(1088), v2250)
	mBase = m.M
	v2969 = m.ExcPending
	if v2969 != 0 {
		goto L1
	} else {
		goto L454
	}
L294:
	;
	if v2923 == int32(0) {
		goto L292
	} else {
		goto L450
	}
L295:
	;
	v2188 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2149)+656)) = uint16(v2188)
	*(*uint16)(unsafe.Add(mBase, uint32(v2153)+8)) = uint16(v2188)
	*(*int64)(unsafe.Add(mBase, uint32(v2153))) = int64(0)
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+648))
	if v2194 != 0 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+40))
	v2923 = v2920
	goto L294
L297:
	;
	F_pfree(m, v2194)
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L1
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v2197 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+652)) = v2197
	*(*int64)(unsafe.Add(mBase, uint32(v2149)+644)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+40)) = v2197
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+660)) = v2197
	*(*uint8)(unsafe.Add(mBase, uint32(v2149)+659)) = uint8(v2197)
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+48)) = int32(-1)
	v2210 = v2077 + int32(1088)
	v2211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2149)+24)))
	v2212 = *(*int64)(unsafe.Add(mBase, uint32(v2149)))
	v2213 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2149)+8)))
	base.MemoryFill(m, v2210, v2197, int32(72))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2110)))
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+48)) = v2110
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+40)) = v2218
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+32)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+24)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+20)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+16)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+12)) = int32(50)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+8)) = int32(51)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+4)) = int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v2210))) = int32(53)
	*(*uint8)(unsafe.Add(mBase, uint32(v2210)+64)) = uint8(v2213)
	*(*int64)(unsafe.Add(mBase, uint32(v2210)+56)) = v2212
	*(*uint16)(unsafe.Add(mBase, uint32(v2210)+54)) = uint16(v2211)
	*(*uint16)(unsafe.Add(mBase, uint32(v2210)+52)) = uint16(v2197)
	*(*uint8)(unsafe.Add(mBase, uint32(v2210)+36)) = uint8(v2197)
	*(*int32)(unsafe.Add(mBase, uint32(v2210)+28)) = int32(54)
	goto L301
L300:
	;
	goto L299
L301:
	;
	v2250 = F_ginFindLeafPage(m, v2210, int32(1), int32(0))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L1
	} else {
		goto L303
	}
L302:
	;
	v2271 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2149)+658)) = uint8(v2271)
	v2273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2149)+9)))
	if v2273 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L303:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	if v2252 < int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2256+(v2252^int32(-1))<<(uint(int32(2))%32))))
	v2270 = v2262
	goto L302
L305:
	;
	goto L306
L306:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2270 = v2264 + v2252<<(uint(int32(13))%32) + int32(-8192)
	goto L302
L307:
	;
	v2276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2149)+8)))
	if v2276 != int32(255) {
		goto L293
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1100))
	v2282 = m.T0[v2281].(func(*base.Module, int32, int32) int32)(m, v2077+int32(1088), v2250)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L1
	} else {
		goto L311
	}
L310:
	;
	goto L309
L311:
	;
	v2285 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[11]))
	v2289 = F_tbm_create(m, v2285<<(uint(int32(10))%32), int32(0))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+40)) = v2289
	v2292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2149)+9)))
	if v2292 == int32(1) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v2295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2149)+8)))
	if v2295 != 0 {
		v2923 = v2289
		goto L294
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1136))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+8))
	v2298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2149)+24)))
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	if v2300 < int32(0) {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	goto L315
L317:
	;
	F_PredicateLockPage(m, v2299, v2319, v2154)
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L1
	} else {
		goto L321
	}
L318:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2304+(v2300^int32(-1))*int32(56))+16))
	v2319 = v2310
	goto L317
L319:
	;
	goto L320
L320:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2313 = int32(56)
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v2312+v2300*v2313-v2313)+16))
	v2319 = v2318
	goto L317
L321:
	;
	v2323 = v2298 - int32(1)
	v2328 = v2297 + v2323<<(uint(int32(3))%32) + int32(28)
	goto L323
L322:
	;
	goto L296
L323:
	;
	v2362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)))
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	if v2363 < int32(0) {
		goto L327
	} else {
		goto L328
	}
L324:
	;
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+40))
	if v2897 != 0 {
		goto L440
	} else {
		goto L441
	}
L325:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1136))
	v2450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2448+v2450<<(uint(int32(2))%32))+20))
	v2457 = v2448 + v2454&int32(_a_F_gingetbitmap_9)
	v2458 = F_gintuple_get_attrnum(m, v2449, v2457)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L1
	} else {
		goto L346
	}
L326:
	;
	v2382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2381)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2382) {
		goto L330
	} else {
		goto L331
	}
L327:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2367+(v2363^int32(-1))<<(uint(int32(2))%32))))
	v2381 = v2373
	goto L326
L328:
	;
	goto L329
L329:
	;
	v2375 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2381 = v2375 + v2363<<(uint(int32(13))%32) + int32(-8192)
	goto L326
L330:
	;
	v2390 = int32(base.Ui32(v2382+int32(_a_F_gingetbitmap_13)) >> (uint(int32(2)) % 32))
	goto L332
L331:
	;
	v2390 = int32(0)
	goto L332
L332:
	;
	if base.Ui32(v2390&int32(_a_F_gingetbitmap_8)) < base.Ui32(v2362) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v2394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2381)+16)))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2381+v2394)))
	if v2396 == int32(-1) {
		goto L322
	} else {
		goto L336
	}
L334:
	;
	v2430 = v2363
	goto L335
L335:
	;
	if v2430 < int32(0) {
		goto L343
	} else {
		goto L344
	}
L336:
	;
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	v2401 = F_ginStepRight(m, v2363, v2399, int32(1))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2250)+4)) = v2401
	if v2401 < int32(0) {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	v2423 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)) = uint16(v2423)
	*(*int32)(unsafe.Add(mBase, uint32(v2250))) = v2422
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	F_PredicateLockPage(m, v2426, v2422, v2154)
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L342
	}
L339:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2407+(v2401^int32(-1))*int32(56))+16))
	v2422 = v2413
	goto L338
L340:
	;
	goto L341
L341:
	;
	v2415 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2416 = int32(56)
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v2415+v2401*v2416-v2416)+16))
	v2422 = v2421
	goto L338
L342:
	;
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	v2430 = v2429
	goto L335
L343:
	;
	v2434 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v2434+(v2430^int32(-1))<<(uint(int32(2))%32))))
	v2448 = v2440
	goto L325
L344:
	;
	goto L345
L345:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2448 = v2442 + v2430<<(uint(int32(13))%32) + int32(-8192)
	goto L325
L346:
	;
	if v2458 != v2298 {
		goto L322
	} else {
		goto L347
	}
L347:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1136))
	v2464 = F_gintuple_get_key(m, v2461, v2457, v2077-int32(-64))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	v2466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2149)+9)))
	if v2466 == int32(1) {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	goto L324
L350:
	;
	v2893 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)))
	v2895 = v2893 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)) = uint16(v2895)
	goto L323
L351:
	;
	v2499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2457)+4)))
	if v2499 == int32(_a_F_gingetbitmap_8) {
		goto L361
	} else {
		goto L362
	}
L352:
	;
	v2469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077)+64)))
	if v2469 != 0 {
		goto L322
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+20))
	if v2492 != int32(2) {
		goto L351
	} else {
		goto L359
	}
L355:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1136))
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v2470+v2323<<(uint(int32(2))%32))+uint32(_c_F_gingetbitmap[14])))
	v2482 = *(*int64)(unsafe.Add(mBase, uint32(v2149)))
	v2483 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v2149)+16)))
	v2484 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2149)+12)))
	v2485 = F_FunctionCall4Coll(m, v2470+v2323*int32(28)+int32(_a_F_gingetbitmap_14), v2481, v2482, v2464, v2483, v2484)
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	v2487 = base.I32_wrap_i64(v2485)
	if int32(0) < v2487 {
		goto L322
	} else {
		goto L357
	}
L357:
	;
	if int32(0) <= v2487 {
		goto L351
	} else {
		goto L358
	}
L358:
	;
	goto L350
L359:
	;
	v2495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077)+64)))
	if v2495 == int32(3) {
		goto L322
	} else {
		goto L360
	}
L360:
	;
	goto L351
L361:
	;
	v2502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2457)+2)))
	v2503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2457))))
	v2506 = v2502 | v2503<<(uint(int32(16))%32)
	v2507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077)+64)))
	if v2507 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	goto L363
L363:
	;
	v2847 = F_ginReadTuple(m, v2457, v2077+int32(2112))
	mBase = m.M
	v2848 = m.ExcPending
	if v2848 != 0 {
		goto L1
	} else {
		goto L437
	}
L364:
	;
	v2510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328)+4)))
	v2511 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2328)+2)))
	v2512 = F_datumCopy(m, v2464, v2510, v2511)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L1
	} else {
		goto L367
	}
L365:
	;
	v2514 = v2464
	goto L366
L366:
	;
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	F_UnlockBuffer(m, v2515)
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L1
	} else {
		goto L368
	}
L367:
	;
	v2514 = v2512
	goto L366
L368:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	F_PredicateLockPage(m, v2518, v2506, v2154)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	v2524 = F_ginScanBeginPostingTree(m, v2077+int32(2112), v2523, v2506)
	mBase = m.M
	v2525 = m.ExcPending
	if v2525 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v2524)+4))
	F_IncrBufferRefCount(m, v2526)
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	F_freeGinBtreeStack(m, v2524)
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L1
	} else {
		goto L372
	}
L372:
	;
	v2538 = v2526
	goto L373
L373:
	;
	if v2538 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L374:
	;
	F_UnlockReleaseBuffer(m, v2538)
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L1
	} else {
		goto L395
	}
L375:
	;
	v2582 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2581)+16)))
	v2584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2582+v2581)+6)))
	if v2584&int32(4) != 0 {
		goto L379
	} else {
		goto L380
	}
L376:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v2567+(v2538^int32(-1))<<(uint(int32(2))%32))))
	v2581 = v2573
	goto L375
L377:
	;
	goto L378
L378:
	;
	v2575 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2581 = v2575 + v2538<<(uint(int32(13))%32) + int32(-8192)
	goto L375
L379:
	;
	v2635 = v2582
	goto L381
L380:
	;
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+40))
	v2588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2581)+16)))
	v2589 = v2581 + v2588
	v2590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2589)+6)))
	if v2590&int32(128) != 0 {
		goto L383
	} else {
		goto L384
	}
L381:
	;
	v2637 = *(*int32)(unsafe.Add(mBase, uint32(v2635+v2581)))
	if v2637 != int32(-1) {
		goto L391
	} else {
		goto L392
	}
L382:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+660))
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+660)) = v2627 + v2628
	v2631 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2581)+16)))
	v2635 = v2631
	goto L381
L383:
	;
	v2593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2581)+12)))
	v2594 = int32(32)
	v2596 = m.G0
	v2598 = v2596 - int32(16)
	m.G0 = v2598
	v2604 = F_ginPostingListDecodeAllSegments(m, v2581+v2594, v2593-v2594, v2598+int32(12))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L1
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v2615 = int32(0)
	v2616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2589)+4)))
	if v2616 == v2615 {
		v2627 = v2615
		goto L382
	} else {
		goto L389
	}
L386:
	;
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2598)+12))
	F_tbm_add_tuples(m, v2587, v2604, v2606, int32(0))
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	F_pfree(m, v2604)
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	m.G0 = v2598 + int32(16)
	v2627 = v2606
	goto L382
L389:
	;
	F_tbm_add_tuples(m, v2587, v2581+int32(32), v2616, int32(0))
	mBase = m.M
	v2623 = m.ExcPending
	if v2623 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	v2627 = v2616
	goto L382
L391:
	;
	v2641 = F_ginStepRight(m, v2538, v2523, int32(1))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L1
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	goto L374
L394:
	;
	v2538 = v2641
	goto L373
L395:
	;
	v2645 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	F_LockBufferInternal(m, v2645, int32(1))
	mBase = m.M
	v2648 = m.ExcPending
	if v2648 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	if v2649 < int32(0) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	v2668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2667)+16)))
	v2670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2668+v2667)+6)))
	if v2670&int32(2) == int32(0) {
		goto L349
	} else {
		goto L401
	}
L398:
	;
	v2653 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v2653+(v2649^int32(-1))<<(uint(int32(2))%32))))
	v2667 = v2659
	goto L397
L399:
	;
	goto L400
L400:
	;
	v2661 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2667 = v2661 + v2649<<(uint(int32(13))%32) + int32(-8192)
	goto L397
L401:
	;
	v2677 = v2649
	goto L402
L402:
	;
	v2708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)))
	if v2677 < int32(0) {
		goto L406
	} else {
		goto L407
	}
L403:
	;
	v2838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077)+64)))
	if v2838 != 0 {
		goto L350
	} else {
		goto L434
	}
L404:
	;
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1136))
	v2795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)))
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v2793+v2795<<(uint(int32(2))%32))+20))
	v2802 = v2793 + v2799&int32(_a_F_gingetbitmap_9)
	v2803 = F_gintuple_get_attrnum(m, v2794, v2802)
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L1
	} else {
		goto L427
	}
L405:
	;
	v2727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2726)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2727) {
		goto L409
	} else {
		goto L410
	}
L406:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(v2712+(v2677^int32(-1))<<(uint(int32(2))%32))))
	v2726 = v2718
	goto L405
L407:
	;
	goto L408
L408:
	;
	v2720 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2726 = v2720 + v2677<<(uint(int32(13))%32) + int32(-8192)
	goto L405
L409:
	;
	v2735 = int32(base.Ui32(v2727+int32(_a_F_gingetbitmap_13)) >> (uint(int32(2)) % 32))
	goto L411
L410:
	;
	v2735 = int32(0)
	goto L411
L411:
	;
	if base.Ui32(v2735&int32(_a_F_gingetbitmap_8)) < base.Ui32(v2708) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v2739 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2726)+16)))
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v2726+v2739)))
	if v2741 == int32(-1) {
		goto L286
	} else {
		goto L415
	}
L413:
	;
	v2775 = v2677
	goto L414
L414:
	;
	if v2775 < int32(0) {
		goto L422
	} else {
		goto L423
	}
L415:
	;
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	v2746 = F_ginStepRight(m, v2677, v2744, int32(1))
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2250)+4)) = v2746
	if v2746 < int32(0) {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	v2768 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)) = uint16(v2768)
	*(*int32)(unsafe.Add(mBase, uint32(v2250))) = v2767
	v2771 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	F_PredicateLockPage(m, v2771, v2767, v2154)
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
		goto L1
	} else {
		goto L421
	}
L418:
	;
	v2752 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v2752+(v2746^int32(-1))*int32(56))+16))
	v2767 = v2758
	goto L417
L419:
	;
	goto L420
L420:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2761 = int32(56)
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2760+v2746*v2761-v2761)+16))
	v2767 = v2766
	goto L417
L421:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	v2775 = v2774
	goto L414
L422:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v2779+(v2775^int32(-1))<<(uint(int32(2))%32))))
	v2793 = v2785
	goto L404
L423:
	;
	goto L424
L424:
	;
	v2787 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2793 = v2787 + v2775<<(uint(int32(13))%32) + int32(-8192)
	goto L404
L425:
	;
	goto L403
L426:
	;
	v2833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)))
	v2835 = v2833 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)) = uint16(v2835)
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	v2677 = v2837
	goto L402
L427:
	;
	if v2803 != v2298 {
		goto L426
	} else {
		goto L428
	}
L428:
	;
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1136))
	v2809 = F_gintuple_get_key(m, v2806, v2802, v2077+int32(2112))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	v2811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077)+2112)))
	v2812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077)+64)))
	if v2811 != v2812 {
		goto L426
	} else {
		goto L430
	}
L430:
	;
	if v2811 != 0 {
		goto L350
	} else {
		goto L431
	}
L431:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1136))
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2814+v2323<<(uint(int32(2))%32))+uint32(_c_F_gingetbitmap[14])))
	v2826 = F_FunctionCall2Coll(m, v2814+v2323*int32(28)+int32(140), v2825, v2809, v2514)
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	if base.I32_wrap_i64(v2826) == int32(0) {
		goto L425
	} else {
		goto L433
	}
L433:
	;
	goto L426
L434:
	;
	v2839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328)+4)))
	if v2839 != 0 {
		goto L350
	} else {
		goto L435
	}
L435:
	;
	F_pfree(m, base.I32_wrap_i64(v2514))
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	goto L350
L437:
	;
	v2849 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+40))
	v2850 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+2112))
	F_tbm_add_tuples(m, v2849, v2847, v2850, int32(0))
	mBase = m.M
	v2853 = m.ExcPending
	if v2853 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+660))
	v2855 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2457)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+660)) = v2854 + v2855
	F_pfree(m, v2847)
	mBase = m.M
	v2859 = m.ExcPending
	if v2859 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	goto L350
L440:
	;
	v2898 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+44))
	if v2898 != 0 {
		goto L443
	} else {
		goto L444
	}
L441:
	;
	v2912 = v2649
	goto L442
L442:
	;
	F_UnlockBuffer(m, v2912)
	mBase = m.M
	v2914 = m.ExcPending
	if v2914 != 0 {
		goto L1
	} else {
		goto L448
	}
L443:
	;
	F_pfree(m, v2898)
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L1
	} else {
		goto L446
	}
L444:
	;
	v2902 = v2897
	goto L445
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+44)) = int32(0)
	F_tbm_free(m, v2902)
	mBase = m.M
	v2906 = m.ExcPending
	if v2906 != 0 {
		goto L1
	} else {
		goto L447
	}
L446:
	;
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+40))
	v2902 = v2901
	goto L445
L447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+40)) = int32(0)
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	v2912 = v2909
	goto L442
L448:
	;
	F_freeGinBtreeStack(m, v2250)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	goto L295
L450:
	;
	v2956 = *(*int32)(unsafe.Add(mBase, uint32(v2923)+16))
	goto L451
L451:
	;
	if v2956 == int32(0) {
		goto L292
	} else {
		goto L452
	}
L452:
	;
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+40))
	v2960 = F_tbm_begin_private_iterate(m, v2959)
	mBase = m.M
	v2961 = m.ExcPending
	if v2961 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v2962 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2149)+658)) = uint8(v2962)
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+44)) = v2960
	goto L292
L454:
	;
	if v2968 != 0 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2970 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2250)+8)))
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(v2270+v2970<<(uint(int32(2))%32))+20))
	v2977 = v2270 + v2974&int32(_a_F_gingetbitmap_9)
	v2978 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2977)+4)))
	if v2978 == int32(_a_F_gingetbitmap_8) {
		goto L458
	} else {
		goto L459
	}
L456:
	;
	goto L457
L457:
	;
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(v2110)))
	v3079 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	if v3079 < int32(0) {
		goto L480
	} else {
		goto L481
	}
L458:
	;
	v2981 = *(*int32)(unsafe.Add(mBase, uint32(v2110)))
	v2982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2977)+2)))
	v2983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2977))))
	v2986 = v2982 | v2983<<(uint(int32(16))%32)
	F_PredicateLockPage(m, v2981, v2986, v2154)
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L1
	} else {
		goto L461
	}
L459:
	;
	goto L460
L460:
	;
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(v2110)))
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	if v3045 < int32(0) {
		goto L473
	} else {
		goto L474
	}
L461:
	;
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+4))
	F_UnlockBuffer(m, v2989)
	mBase = m.M
	v2991 = m.ExcPending
	if v2991 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v2994 = *(*int32)(unsafe.Add(mBase, uint32(v2110)))
	v2995 = F_ginScanBeginPostingTree(m, v2149+int32(664), v2994, v2986)
	mBase = m.M
	v2996 = m.ExcPending
	if v2996 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	v2997 = *(*int32)(unsafe.Add(mBase, uint32(v2995)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+28)) = v2997
	F_IncrBufferRefCount(m, v2997)
	mBase = m.M
	v3000 = m.ExcPending
	if v3000 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	v3001 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+28))
	if v3001 < int32(0) {
		goto L466
	} else {
		goto L467
	}
L465:
	;
	v3020 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2077)+2116)) = uint16(v3020)
	*(*int32)(unsafe.Add(mBase, uint32(v2077)+2112)) = v3020
	*(*int32)(unsafe.Add(mBase, uint32(v2077)+28)) = v3020
	*(*uint16)(unsafe.Add(mBase, uint32(v2077)+32)) = uint16(v3020)
	v3030 = F_GinDataLeafPageGetItems(m, v3019, v2151, v2077+int32(28))
	mBase = m.M
	v3031 = m.ExcPending
	if v3031 != 0 {
		goto L1
	} else {
		goto L469
	}
L466:
	;
	v3005 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v3005+(v3001^int32(-1))<<(uint(int32(2))%32))))
	v3019 = v3011
	goto L465
L467:
	;
	goto L468
L468:
	;
	v3013 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v3019 = v3013 + v3001<<(uint(int32(13))%32) + int32(-8192)
	goto L465
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+648)) = v3030
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+652))
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v2995)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+660)) = v3033 * v3034
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+28))
	F_UnlockBuffer(m, v3037)
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	F_freeGinBtreeStack(m, v2995)
	mBase = m.M
	v3041 = m.ExcPending
	if v3041 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	v3042 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2149)+658)) = uint8(v3042)
	goto L291
L472:
	;
	F_PredicateLockPage(m, v3044, v3064, v2154)
	mBase = m.M
	v3066 = m.ExcPending
	if v3066 != 0 {
		goto L1
	} else {
		goto L476
	}
L473:
	;
	v3049 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(v3049+(v3045^int32(-1))*int32(56))+16))
	v3064 = v3055
	goto L472
L474:
	;
	goto L475
L475:
	;
	v3057 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v3058 = int32(56)
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v3057+v3045*v3058-v3058)+16))
	v3064 = v3063
	goto L472
L476:
	;
	v3067 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2977)+4)))
	if v3067 == int32(0) {
		goto L292
	} else {
		goto L477
	}
L477:
	;
	v3071 = F_ginReadTuple(m, v2977, v2151)
	mBase = m.M
	v3072 = m.ExcPending
	if v3072 != 0 {
		goto L1
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+648)) = v3071
	v3074 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2149)+658)) = uint8(v3074)
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+652))
	*(*int32)(unsafe.Add(mBase, uint32(v2149)+660)) = v3076
	goto L292
L479:
	;
	F_PredicateLockPage(m, v3078, v3098, v2154)
	mBase = m.M
	v3100 = m.ExcPending
	if v3100 != 0 {
		goto L1
	} else {
		goto L483
	}
L480:
	;
	v3083 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v3083+(v3079^int32(-1))*int32(56))+16))
	v3098 = v3089
	goto L479
L481:
	;
	goto L482
L482:
	;
	v3091 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v3092 = int32(56)
	v3097 = *(*int32)(unsafe.Add(mBase, uint32(v3091+v3079*v3092-v3092)+16))
	v3098 = v3097
	goto L479
L483:
	;
	goto L292
L484:
	;
	goto L291
L485:
	;
	v3173 = v2124 + int32(1)
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[2])))
	if base.Ui32(v3173) < base.Ui32(v3174) {
		v2124 = v3173
		goto L289
	} else {
		goto L486
	}
L486:
	;
	goto L290
L487:
	;
	v3179 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[15]))
	if v3179 <= int32(0) {
		goto L287
	} else {
		goto L488
	}
L488:
	;
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[6])))
	v3187 = int32(0)
	goto L489
L489:
	;
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(v3183+v3187<<(uint(int32(2))%32))))
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3221)+660))
	if base.Ui32(v3222) <= base.Ui32(v3179*v3174) {
		goto L287
	} else {
		goto L491
	}
L490:
	;
	v3230 = int32(0)
	v3231 = v3174
	goto L493
L491:
	;
	v3225 = v3187 + int32(1)
	if v3225 != v3174 {
		v3187 = v3225
		goto L489
	} else {
		goto L492
	}
L492:
	;
	goto L490
L493:
	;
	v3262 = v3230 << (uint(int32(2)) % 32)
	v3263 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[6])))
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(v3262+v3263)))
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v3265)+660))
	v3267 = base.I32_div_u_s(v3266, v3231)
	*(*int32)(unsafe.Add(mBase, uint32(v3265)+660)) = v3267
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[6])))
	v3271 = *(*int32)(unsafe.Add(mBase, uint32(v3269+v3262)))
	v3272 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3271)+659)) = uint8(v3272)
	v3275 = v3230 + v3272
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[2])))
	if base.Ui32(v3275) < base.Ui32(v3276) {
		v3230 = v3275
		v3231 = v3276
		goto L493
	} else {
		goto L495
	}
L494:
	;
	goto L287
L495:
	;
	goto L494
L496:
	;
	v3313 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v3318 = int32(0)
	goto L499
L497:
	;
	goto L498
L498:
	;
	v3972 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2077)+68)) = uint16(v3972)
	*(*int32)(unsafe.Add(mBase, uint32(v2077)+64)) = v3972
	v3976 = v2072
	v3977 = v2073
	v3981 = v2077
	v4004 = v2100
	goto L557
L499:
	;
	v3348 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[4])))
	v3351 = v3348 + v3318*int32(104)
	v3352 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3351)+96)) = uint8(v3352)
	*(*int64)(unsafe.Add(mBase, uint32(v3351)+88)) = int64(0)
	v3356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3351)+86)))
	if v3356 == int32(1) {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	goto L498
L501:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3313
	v3936 = v3318 + int32(1)
	v3937 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v3936) < base.Ui32(v3937) {
		v3318 = v3936
		goto L499
	} else {
		goto L556
	}
L502:
	;
	v3360 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3360
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+16)) = int32(0)
	v3364 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+24)) = v3364
	v3368 = F_palloc(m, v3364<<(uint(int32(2))%32))
	mBase = m.M
	v3369 = m.ExcPending
	if v3369 != 0 {
		goto L1
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v3420 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	if base.Ui32(int32(2)) <= base.Ui32(v3420) {
		goto L510
	} else {
		goto L511
	}
L505:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+20)) = v3368
	v3371 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+24))
	if v3371 <= int32(0) {
		goto L501
	} else {
		goto L506
	}
L506:
	;
	v3377 = int32(0)
	goto L507
L507:
	;
	v3409 = v3377 << (uint(int32(2)) % 32)
	v3410 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+20))
	v3412 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+8))
	v3414 = *(*int32)(unsafe.Add(mBase, uint32(v3412+v3409)))
	*(*int32)(unsafe.Add(mBase, uint32(v3409+v3410))) = v3414
	v3417 = v3377 + int32(1)
	v3418 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+24))
	if v3417 < v3418 {
		v3377 = v3417
		goto L507
	} else {
		goto L509
	}
L508:
	;
	goto L501
L509:
	;
	goto L508
L510:
	;
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v2105)))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3425
	v3428 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	v3429 = F_palloc_mul(m, int32(4), v3428)
	mBase = m.M
	v3430 = m.ExcPending
	if v3430 != 0 {
		goto L1
	} else {
		goto L513
	}
L511:
	;
	goto L512
L512:
	;
	v3887 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3887
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+16)) = int32(1)
	v3894 = F_palloc(m, int32(4))
	mBase = m.M
	v3895 = m.ExcPending
	if v3895 != 0 {
		goto L1
	} else {
		goto L555
	}
L513:
	;
	v3432 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	if v3432 != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v3435 = int32(0)
	goto L517
L515:
	;
	v3478 = int32(0)
	goto L516
L516:
	;
	F_qsort_arg(m, v3429, v3478, int32(4), int32(55), v3351)
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L1
	} else {
		goto L520
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3429+v3435<<(uint(int32(2))%32)))) = v3435
	v3471 = v3435 + int32(1)
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	if base.Ui32(v3471) < base.Ui32(v3472) {
		v3435 = v3471
		goto L517
	} else {
		goto L519
	}
L518:
	;
	v3478 = v3472
	goto L516
L519:
	;
	goto L518
L520:
	;
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	if base.Ui32(int32(2)) <= base.Ui32(v3511) {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v3517 = int32(1)
	goto L524
L522:
	;
	v3564 = v3511
	goto L523
L523:
	;
	v3593 = int32(1)
	if v3564 != v3593 {
		goto L527
	} else {
		goto L528
	}
L524:
	;
	v3548 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+28))
	v3549 = int32(2)
	v3552 = *(*int32)(unsafe.Add(mBase, uint32(v3429+v3517<<(uint(v3549)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3548+v3552))) = uint8(v3549)
	v3557 = v3517 + int32(1)
	v3558 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	if base.Ui32(v3557) < base.Ui32(v3558) {
		v3517 = v3557
		goto L524
	} else {
		goto L526
	}
L525:
	;
	v3564 = v3558
	goto L523
L526:
	;
	goto L525
L527:
	;
	v3599 = int32(0)
	goto L530
L528:
	;
	v3664 = v3593
	goto L529
L529:
	;
	v3689 = int32(0)
	v3691 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3691
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+16)) = v3664
	v3694 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+24)) = v3694 - v3664
	v3699 = F_palloc(m, v3664<<(uint(int32(2))%32))
	mBase = m.M
	v3700 = m.ExcPending
	if v3700 != 0 {
		goto L1
	} else {
		goto L540
	}
L530:
	;
	v3630 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+28))
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v3429+v3599<<(uint(int32(2))%32))))
	v3636 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3630+v3634))) = uint8(v3636)
	v3638 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+36))
	v3639 = m.T0[v3638].(func(*base.Module, int32) int32)(m, v3351)
	mBase = m.M
	v3640 = m.ExcPending
	if v3640 != 0 {
		goto L1
	} else {
		goto L533
	}
L531:
	;
	v3664 = v3653 + int32(1)
	goto L529
L532:
	;
	goto L531
L533:
	;
	if v3639 == int32(0) {
		v3653 = v3599
		goto L532
	} else {
		goto L534
	}
L534:
	;
	v3644 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[16]))
	if v3644 != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3646 = m.ExcPending
	if v3646 != 0 {
		goto L1
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	v3647 = int32(1)
	v3648 = v3599 + v3647
	v3649 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	if base.Ui32(v3648) < base.Ui32(v3649-v3647) {
		v3599 = v3648
		goto L530
	} else {
		goto L539
	}
L538:
	;
	goto L537
L539:
	;
	v3653 = v3648
	goto L532
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+12)) = v3699
	v3702 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+24))
	v3705 = F_palloc(m, v3702<<(uint(int32(2))%32))
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L1
	} else {
		goto L541
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+20)) = v3705
	v3708 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+16))
	if int32(0) < v3708 {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v3713 = v3689
	goto L545
L543:
	;
	v3762 = v3689
	goto L544
L544:
	;
	v3793 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+24))
	if int32(0) < v3793 {
		goto L548
	} else {
		goto L549
	}
L545:
	;
	v3744 = int32(2)
	v3745 = v3713 << (uint(v3744) % 32)
	v3746 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+12))
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+8))
	v3750 = *(*int32)(unsafe.Add(mBase, uint32(v3745+v3429)))
	v3754 = *(*int32)(unsafe.Add(mBase, uint32(v3748+v3750<<(uint(v3744)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3745+v3746))) = v3754
	v3757 = v3713 + int32(1)
	v3758 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+16))
	if v3757 < v3758 {
		v3713 = v3757
		goto L545
	} else {
		goto L547
	}
L546:
	;
	v3762 = v3757
	goto L544
L547:
	;
	goto L546
L548:
	;
	v3799 = v3762
	v3801 = int32(0)
	goto L551
L549:
	;
	goto L550
L550:
	;
	v3883 = *(*int32)(unsafe.Add(mBase, uint32(v2105)))
	F_MemoryContextReset(m, v3883)
	mBase = m.M
	v3885 = m.ExcPending
	if v3885 != 0 {
		goto L1
	} else {
		goto L554
	}
L551:
	;
	v3830 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+20))
	v3831 = int32(2)
	v3834 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+8))
	v3838 = *(*int32)(unsafe.Add(mBase, uint32(v3429+v3799<<(uint(v3831)%32))))
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v3834+v3838<<(uint(v3831)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3830+v3801<<(uint(v3831)%32)))) = v3842
	v3844 = int32(1)
	v3847 = v3801 + v3844
	v3848 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+24))
	if v3847 < v3848 {
		v3799 = v3799 + v3844
		v3801 = v3847
		goto L551
	} else {
		goto L553
	}
L552:
	;
	goto L550
L553:
	;
	goto L552
L554:
	;
	goto L501
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3351)+12)) = v3894
	v3897 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+8))
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3897)))
	*(*int32)(unsafe.Add(mBase, uint32(v3894))) = v3898
	goto L501
L556:
	;
	goto L500
L557:
	;
	v4009 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+68)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1092)) = uint16(v4009)
	v4011 = *(*int32)(unsafe.Add(mBase, uint32(v3981)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+1088)) = v4011
	v4013 = *(*int32)(unsafe.Add(mBase, uint32(v3976)+36))
	v4015 = v4013 + int32(4)
	goto L560
L559:
	;
	F_tbm_add_tuples(m, v3977, v3981-int32(-64), int32(1), v5106)
	mBase = m.M
	v5136 = m.ExcPending
	if v5136 != 0 {
		goto L1
	} else {
		goto L695
	}
L560:
	;
	v4050 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[16]))
	if v4050 != 0 {
		goto L562
	} else {
		goto L563
	}
L561:
	;
	if v5001 == int32(0) {
		goto L685
	} else {
		goto L686
	}
L562:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4052 = m.ExcPending
	if v4052 != 0 {
		goto L1
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	v4053 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+68)) = uint16(v4053)
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+64)) = v4053
	v4057 = *(*int32)(unsafe.Add(mBase, uint32(v4013)+uint32(_c_F_gingetbitmap[3])))
	if v4057 == v4053 {
		goto L566
	} else {
		goto L567
	}
L565:
	;
	goto L564
L566:
	;
	v5106 = int32(0)
	goto L559
L567:
	;
	goto L568
L568:
	;
	v4062 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1090)))
	v4063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1088)))
	v4070 = v4062
	v4074 = v4063
	v4080 = int32(0)
	goto L569
L569:
	;
	v4097 = *(*int32)(unsafe.Add(mBase, uint32(v4013)+uint32(_c_F_gingetbitmap[4])))
	v4100 = v4097 + v4080*int32(104)
	v4101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+68)))
	if v4101 != int32(_a_F_gingetbitmap_8) {
		goto L572
	} else {
		goto L573
	}
L570:
	;
	if v4970 == int32(0) {
		goto L560
	} else {
		goto L683
	}
L571:
	;
	v5000 = v4080 + int32(1)
	v5001 = *(*int32)(unsafe.Add(mBase, uint32(v4013)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v5000) < base.Ui32(v5001) {
		goto L679
	} else {
		goto L680
	}
L572:
	;
	v4114 = *(*int32)(unsafe.Add(mBase, uint32(v4013)))
	v4115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1092)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2116)) = uint16(v4115)
	v4117 = *(*int32)(unsafe.Add(mBase, uint32(v3981)+1088))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+2112)) = v4117
	v4120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2114)))
	v4122 = int64(32)
	v4125 = int64(48)
	v4128 = base.I64_extend_i32_u(v4115) | (base.I64_extend_i32_u(v4120)<<(uint(v4122)%64) | base.I64_extend_i32_u(v4117)<<(uint(v4125)%64))
	v4129 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4100)+92)))
	v4130 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4100)+90)))
	v4133 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4100)+88)))
	if base.Ui64(v4128) < base.Ui64(v4129|(v4130<<(uint(v4122)%64)|v4133<<(uint(v4125)%64))) {
		goto L576
	} else {
		goto L577
	}
L573:
	;
	v4104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+64)))
	v4105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+66)))
	if v4104&v4105 == int32(_a_F_gingetbitmap_8) {
		goto L572
	} else {
		goto L574
	}
L574:
	;
	v4110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4100)+86)))
	if v4110 == int32(0) {
		goto L572
	} else {
		goto L575
	}
L575:
	;
	v4968 = v4070
	v4970 = int32(1)
	v4973 = v4074
	goto L571
L576:
	;
	v4873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4100)+96)))
	if v4873 != 0 {
		v5166 = v3981
		v5189 = v4004
		goto L135
	} else {
		goto L657
	}
L577:
	;
	v4139 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+16))
	if v4139 != 0 {
		goto L582
	} else {
		goto L583
	}
L578:
	;
	v4396 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+24))
	if v4396 != 0 {
		goto L603
	} else {
		goto L604
	}
L579:
	;
	v4366 = v4248
	v4367 = v4249
	v4371 = v4361
	v4375 = v4248
	v4378 = v4249
	v4382 = v4362
	goto L578
L580:
	;
	v4366 = v4117
	v4367 = v4120
	v4371 = v4115 + int32(1)
	v4375 = v4117
	v4378 = v4120
	v4382 = v4115
	goto L578
L581:
	;
	v4297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4100)+86)))
	if v4297 != 0 {
		goto L580
	} else {
		goto L598
	}
L582:
	;
	v4140 = int32(_a_F_gingetbitmap_8)
	v4150 = v4140
	v4151 = v4140
	v4154 = int32(0)
	v4155 = v4140
	v4167 = int32(1)
	goto L585
L583:
	;
	goto L584
L584:
	;
	v4294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4100)+86)))
	if v4294 != 0 {
		goto L580
	} else {
		goto L597
	}
L585:
	;
	v4180 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+12))
	v4184 = *(*int32)(unsafe.Add(mBase, uint32(v4180+v4154<<(uint(int32(2))%32))))
	v4185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4184)+658)))
	if v4185 != 0 {
		v4248 = v4150
		v4249 = v4151
		v4250 = v4155
		v4253 = v4167
		goto L587
	} else {
		goto L588
	}
L586:
	;
	if v4253 == int32(0) {
		goto L581
	} else {
		goto L596
	}
L587:
	;
	v4256 = v4154 + int32(1)
	v4257 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+16))
	if base.Ui32(v4256) < base.Ui32(v4257) {
		v4150 = v4248
		v4151 = v4249
		v4154 = v4256
		v4155 = v4250
		v4167 = v4253
		goto L585
	} else {
		goto L595
	}
L588:
	;
	v4186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4184)+36)))
	v4188 = int64(65535)
	v4190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4184)+34)))
	v4196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4184)+32)))
	v4201 = base.I64_extend_i32_u(v4186)&v4188 | (base.I64_extend_i32_u(v4190)&v4188<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4196)<<(uint(int64(48))%64))
	if base.Ui64(v4201) <= base.Ui64(v4128) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v4203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1092)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+12)) = uint16(v4203)
	v4205 = *(*int32)(unsafe.Add(mBase, uint32(v3981)+1088))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+8)) = v4205
	F_entryGetItem(m, v4015, v4184, v3981+int32(8))
	mBase = m.M
	v4210 = m.ExcPending
	if v4210 != 0 {
		goto L1
	} else {
		goto L592
	}
L590:
	;
	v4228 = v4186
	v4229 = v4196
	v4230 = v4190
	v4231 = v4201
	goto L591
L591:
	;
	v4232 = int32(0)
	v4234 = int64(65535)
	if base.Ui64(base.I64_extend_i32_u(v4155)&v4234|(base.I64_extend_i32_u(v4150)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v4151)&v4234<<(uint(int64(32))%64))) <= base.Ui64(v4231) {
		v4248 = v4150
		v4249 = v4151
		v4250 = v4155
		v4253 = v4232
		goto L587
	} else {
		goto L594
	}
L592:
	;
	v4211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4184)+658)))
	if v4211 != 0 {
		v4248 = v4150
		v4249 = v4151
		v4250 = v4155
		v4253 = v4167
		goto L587
	} else {
		goto L593
	}
L593:
	;
	v4212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4184)+36)))
	v4214 = int64(65535)
	v4216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4184)+34)))
	v4222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4184)+32)))
	v4228 = v4212
	v4229 = v4222
	v4230 = v4216
	v4231 = base.I64_extend_i32_u(v4212)&v4214 | (base.I64_extend_i32_u(v4216)&v4214<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4222)<<(uint(int64(48))%64))
	goto L591
L594:
	;
	v4248 = v4229
	v4249 = v4230
	v4250 = v4228
	v4253 = v4232
	goto L587
L595:
	;
	goto L586
L596:
	;
	goto L584
L597:
	;
	v4295 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4100)+96)) = uint8(v4295)
	goto L576
L598:
	;
	v4298 = int32(_a_F_gingetbitmap_8)
	if v4250&v4298 != v4298 {
		goto L599
	} else {
		goto L600
	}
L599:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2114)) = uint16(v4249)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2112)) = uint16(v4248)
	v4323 = v4250 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2116)) = uint16(v4323)
	v4361 = v4250
	v4362 = v4323
	goto L579
L600:
	;
	v4306 = v4249&int32(_a_F_gingetbitmap_8) | v4248<<(uint(int32(16))%32)
	if v4306 == int32(-1) {
		goto L599
	} else {
		goto L601
	}
L601:
	;
	v4309 = int32(_a_F_gingetbitmap_8)
	if base.Ui32(v4306) <= base.Ui32(v4117&v4140<<(uint(int32(16))%32)|v4120) {
		v4366 = v4248
		v4367 = v4249
		v4371 = v4309
		v4375 = v4117
		v4378 = v4120
		v4382 = v4115
		goto L578
	} else {
		goto L602
	}
L602:
	;
	v4314 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2116)) = uint16(v4314)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2114)) = uint16(v4249)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2112)) = uint16(v4248)
	v4361 = v4309
	v4362 = v4314
	goto L579
L603:
	;
	v4398 = int64(65535)
	v4414 = v4366
	v4415 = v4367
	v4418 = int32(0)
	v4419 = v4371
	goto L606
L604:
	;
	v4522 = v4366
	v4523 = v4367
	v4527 = v4371
	goto L605
L605:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v4100)+92)) = uint16(v4527)
	*(*uint16)(unsafe.Add(mBase, uint32(v4100)+90)) = uint16(v4523)
	*(*uint16)(unsafe.Add(mBase, uint32(v4100)+88)) = uint16(v4522)
	v4555 = *(*int32)(unsafe.Add(mBase, uint32(v4100)))
	if v4555 == int32(0) {
		goto L618
	} else {
		goto L619
	}
L606:
	;
	v4444 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+20))
	v4448 = *(*int32)(unsafe.Add(mBase, uint32(v4444+v4418<<(uint(int32(2))%32))))
	v4449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4448)+658)))
	if v4449 != 0 {
		v4509 = v4414
		v4510 = v4415
		v4511 = v4419
		goto L608
	} else {
		goto L609
	}
L607:
	;
	v4522 = v4509
	v4523 = v4510
	v4527 = v4511
	goto L605
L608:
	;
	v4516 = v4418 + int32(1)
	v4517 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+24))
	if base.Ui32(v4516) < base.Ui32(v4517) {
		v4414 = v4509
		v4415 = v4510
		v4418 = v4516
		v4419 = v4511
		goto L606
	} else {
		goto L616
	}
L609:
	;
	v4450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4448)+36)))
	v4452 = int64(65535)
	v4454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4448)+34)))
	v4460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4448)+32)))
	v4465 = base.I64_extend_i32_u(v4450)&v4452 | (base.I64_extend_i32_u(v4454)&v4452<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4460)<<(uint(int64(48))%64))
	if base.Ui64(v4465) <= base.Ui64(base.I64_extend_i32_u(v4382)&v4398|(base.I64_extend_i32_u(v4375)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v4378)&v4398<<(uint(int64(32))%64))) {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	v4467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+2116)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+4)) = uint16(v4467)
	v4469 = *(*int32)(unsafe.Add(mBase, uint32(v3981)+2112))
	*(*int32)(unsafe.Add(mBase, uint32(v3981))) = v4469
	F_entryGetItem(m, v4015, v4448, v3981)
	mBase = m.M
	v4472 = m.ExcPending
	if v4472 != 0 {
		goto L1
	} else {
		goto L613
	}
L611:
	;
	v4490 = v4450
	v4491 = v4460
	v4492 = v4454
	v4493 = v4465
	goto L612
L612:
	;
	v4495 = int64(65535)
	if base.Ui64(base.I64_extend_i32_u(v4419)&v4495|(base.I64_extend_i32_u(v4414)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v4415)&v4495<<(uint(int64(32))%64))) <= base.Ui64(v4493) {
		v4509 = v4414
		v4510 = v4415
		v4511 = v4419
		goto L608
	} else {
		goto L615
	}
L613:
	;
	v4473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4448)+658)))
	if v4473 != 0 {
		v4509 = v4414
		v4510 = v4415
		v4511 = v4419
		goto L608
	} else {
		goto L614
	}
L614:
	;
	v4474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4448)+36)))
	v4476 = int64(65535)
	v4478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4448)+34)))
	v4484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4448)+32)))
	v4490 = v4474
	v4491 = v4484
	v4492 = v4478
	v4493 = base.I64_extend_i32_u(v4474)&v4476 | (base.I64_extend_i32_u(v4478)&v4476<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4484)<<(uint(int64(48))%64))
	goto L612
L615:
	;
	v4509 = v4491
	v4510 = v4492
	v4511 = v4490
	goto L608
L616:
	;
	goto L607
L617:
	;
	v4702 = *(*int32)(unsafe.Add(mBase, uint32(v4100)))
	if v4702 != 0 {
		goto L635
	} else {
		goto L636
	}
L618:
	;
	v4558 = int32(_a_F_gingetbitmap_1)
	v4559 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4114
	v4680 = v4559
	goto L617
L619:
	;
	goto L620
L620:
	;
	v4563 = int64(65535)
	v4573 = int32(0)
	v4577 = v4573
	v4588 = v4573
	goto L621
L621:
	;
	v4608 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+8))
	v4612 = *(*int32)(unsafe.Add(mBase, uint32(v4608+v4577<<(uint(int32(2))%32))))
	v4613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4612)+658)))
	if v4613 != 0 {
		goto L624
	} else {
		goto L625
	}
L622:
	;
	v4644 = int32(_a_F_gingetbitmap_1)
	v4645 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4114
	if v4639&int32(1) == int32(0) {
		v4680 = v4645
		goto L617
	} else {
		goto L631
	}
L623:
	;
	v4641 = v4577 + int32(1)
	v4642 = *(*int32)(unsafe.Add(mBase, uint32(v4100)))
	if base.Ui32(v4641) < base.Ui32(v4642) {
		v4577 = v4641
		v4588 = v4639
		goto L621
	} else {
		goto L630
	}
L624:
	;
	v4634 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+28))
	v4636 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4634+v4577))) = uint8(v4636)
	v4639 = v4588
	goto L623
L625:
	;
	v4614 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4612)+36)))
	v4615 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4612)+34)))
	v4618 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4612)+32)))
	if v4614|(v4615<<(uint(int64(32))%64)|v4618<<(uint(int64(48))%64)) != base.I64_extend_i32_u(v4523)&v4563<<(uint(int64(32))%64)|base.I64_extend_i32_u(v4522)<<(uint(int64(48))%64)|v4563 {
		goto L624
	} else {
		goto L626
	}
L626:
	;
	v4624 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+28))
	v4625 = v4624 + v4577
	v4626 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+4))
	if base.Ui32(v4577) < base.Ui32(v4626) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v4628 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4625))) = uint8(v4628)
	v4639 = int32(1)
	goto L623
L628:
	;
	goto L629
L629:
	;
	v4631 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4625))) = uint8(v4631)
	v4639 = v4631
	goto L623
L630:
	;
	goto L622
L631:
	;
	v4652 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+36))
	v4653 = m.T0[v4652].(func(*base.Module, int32) int32)(m, v4100)
	mBase = m.M
	v4654 = m.ExcPending
	if v4654 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	v4655 = int32(1)
	if base.Ui32(v4655) < base.Ui32((v4653-v4655)&int32(255)) {
		v4680 = v4645
		goto L617
	} else {
		goto L633
	}
L633:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4645
	F_MemoryContextReset(m, v4114)
	mBase = m.M
	v4664 = m.ExcPending
	if v4664 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4100)+92)) = int32(16908287)
	*(*uint16)(unsafe.Add(mBase, uint32(v4100)+90)) = uint16(v4523)
	*(*uint16)(unsafe.Add(mBase, uint32(v4100)+88)) = uint16(v4522)
	goto L576
L635:
	;
	v4704 = int64(65535)
	v4707 = base.I64_extend_i32_u(v4523) & v4704 << (uint(int64(32)) % 64)
	v4710 = base.I64_extend_i32_u(v4522) << (uint(int64(48)) % 64)
	v4722 = int32(0)
	goto L638
L636:
	;
	goto L637
L637:
	;
	v4825 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+36))
	v4826 = m.T0[v4825].(func(*base.Module, int32) int32)(m, v4100)
	mBase = m.M
	v4827 = m.ExcPending
	if v4827 != 0 {
		goto L1
	} else {
		goto L655
	}
L638:
	;
	v4753 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+8))
	v4757 = *(*int32)(unsafe.Add(mBase, uint32(v4753+v4722<<(uint(int32(2))%32))))
	v4758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4757)+658)))
	if v4758 == int32(1) {
		goto L641
	} else {
		goto L642
	}
L639:
	;
	goto L637
L640:
	;
	v4789 = v4722 + int32(1)
	v4790 = *(*int32)(unsafe.Add(mBase, uint32(v4100)))
	if base.Ui32(v4789) < base.Ui32(v4790) {
		v4722 = v4789
		goto L638
	} else {
		goto L650
	}
L641:
	;
	v4761 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+28))
	v4763 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4761+v4722))) = uint8(v4763)
	goto L640
L642:
	;
	goto L643
L643:
	;
	v4765 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4757)+36)))
	v4766 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4757)+34)))
	v4769 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4757)+32)))
	v4773 = v4765 | (v4766<<(uint(int64(32))%64) | v4769<<(uint(int64(48))%64))
	if v4707|v4710|v4704 == v4773 {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v4775 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+28))
	v4777 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4775+v4722))) = uint8(v4777)
	goto L640
L645:
	;
	goto L646
L646:
	;
	v4779 = *(*int32)(unsafe.Add(mBase, uint32(v4100)+28))
	v4780 = v4779 + v4722
	if v4773 == v4707|(v4710|base.I64_extend_i32_u(v4527)&v4704) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v4782 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4780))) = uint8(v4782)
	goto L640
L648:
	;
	goto L649
L649:
	;
	v4784 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4780))) = uint8(v4784)
	goto L640
L650:
	;
	goto L639
L651:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4680
	F_MemoryContextReset(m, v4114)
	mBase = m.M
	v4839 = m.ExcPending
	if v4839 != 0 {
		goto L1
	} else {
		goto L656
	}
L652:
	;
	v4834 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v4100)+94)) = uint16(v4834)
	goto L651
L653:
	;
	v4832 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4100)+94)) = uint8(v4832)
	goto L651
L654:
	;
	v4830 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4100)+94)) = uint8(v4830)
	goto L651
L655:
	;
	switch v4826 & int32(255) {
	case 0:
		goto L653
	case 1:
		goto L654
	default:
		goto L652
	}
L656:
	;
	goto L576
L657:
	;
	v4875 = v4100 + int32(88)
	v4876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4100)+94)))
	if v4876 == int32(1) {
		goto L660
	} else {
		goto L661
	}
L658:
	;
	v4960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+66)))
	v4961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+64)))
	v4968 = v4958
	v4970 = base.B2i32(v4886 == v4960|v4961<<(uint(int32(16))%32))
	v4973 = v4959
	goto L571
L659:
	;
	v4942 = int64(48)
	v4946 = int64(32)
	v4968 = v4879
	v4970 = base.B2i32(base.I64_extend_i32_u(v4881)|base.I64_extend_i32_u(v4880)<<(uint(v4942)%64)|base.I64_extend_i32_u(v4879)<<(uint(v4946)%64) == base.I64_extend_i32_u(v4914)<<(uint(v4942)%64)|v4915|base.I64_extend_i32_u(v4913)<<(uint(v4946)%64))
	v4973 = v4880
	goto L571
L660:
	;
	v4879 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4100)+90)))
	v4880 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4100)+88)))
	v4881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4100)+92)))
	if v4881 == int32(_a_F_gingetbitmap_8) {
		goto L666
	} else {
		goto L667
	}
L661:
	;
	goto L662
L662:
	;
	v4935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4875)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1092)) = uint16(v4935)
	v4937 = *(*int32)(unsafe.Add(mBase, uint32(v4875)))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+1088)) = v4937
	goto L560
L663:
	;
	v4930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4875)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+68)) = uint16(v4930)
	v4932 = *(*int32)(unsafe.Add(mBase, uint32(v4875)))
	*(*int32)(unsafe.Add(mBase, uint32(v3981)+64)) = v4932
	v4968 = v4927
	v4970 = int32(1)
	v4973 = v4928
	goto L571
L664:
	;
	if v4080 != 0 {
		v4958 = v4070
		v4959 = v4074
		goto L658
	} else {
		goto L678
	}
L665:
	;
	v4913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+66)))
	v4914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+64)))
	v4915 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+68)))
	if v4915 != int64(65535) {
		goto L659
	} else {
		goto L676
	}
L666:
	;
	v4886 = v4880<<(uint(int32(16))%32) | v4879
	if v4886 != int32(-1) {
		goto L669
	} else {
		goto L670
	}
L667:
	;
	goto L668
L668:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1090)) = uint16(v4879)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1088)) = uint16(v4880)
	v4908 = v4881 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1092)) = uint16(v4908)
	if v4080 == int32(0) {
		v4927 = v4879
		v4928 = v4880
		goto L663
	} else {
		goto L675
	}
L669:
	;
	if base.Ui32(v4886) <= base.Ui32(v4070&int32(_a_F_gingetbitmap_8)|v4074<<(uint(int32(16))%32)) {
		goto L664
	} else {
		goto L672
	}
L670:
	;
	goto L671
L671:
	;
	v4901 = int32(_a_F_gingetbitmap_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1092)) = uint16(v4901)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1090)) = uint16(v4879)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1088)) = uint16(v4880)
	if v4080 != 0 {
		goto L665
	} else {
		goto L674
	}
L672:
	;
	v4895 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1092)) = uint16(v4895)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1090)) = uint16(v4879)
	*(*uint16)(unsafe.Add(mBase, uint32(v3981)+1088)) = uint16(v4880)
	if v4080 == v4895 {
		v4927 = v4879
		v4928 = v4880
		goto L663
	} else {
		goto L673
	}
L673:
	;
	v4958 = v4879
	v4959 = v4880
	goto L658
L674:
	;
	v4927 = v4879
	v4928 = v4880
	goto L663
L675:
	;
	goto L665
L676:
	;
	v4920 = v4914<<(uint(int32(16))%32) | v4913
	if v4920 == int32(-1) {
		goto L659
	} else {
		goto L677
	}
L677:
	;
	v4968 = v4879
	v4970 = base.B2i32(v4880<<(uint(int32(16))%32)|v4879 == v4920)
	v4973 = v4880
	goto L571
L678:
	;
	v4927 = v4070
	v4928 = v4074
	goto L663
L679:
	;
	if v4970 != 0 {
		v4070 = v4968
		v4074 = v4973
		v4080 = v5000
		goto L569
	} else {
		goto L682
	}
L680:
	;
	goto L681
L681:
	;
	goto L570
L682:
	;
	goto L681
L683:
	;
	goto L561
L684:
	;
	v5085 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+68)))
	if v5085 != int32(_a_F_gingetbitmap_8) {
		v5106 = v5059
		goto L559
	} else {
		goto L692
	}
L685:
	;
	v5059 = int32(0)
	goto L684
L686:
	;
	goto L687
L687:
	;
	v5010 = *(*int32)(unsafe.Add(mBase, uint32(v4013)+uint32(_c_F_gingetbitmap[4])))
	v5014 = int32(0)
	goto L688
L688:
	;
	v5048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5010+v5014*int32(104))+95)))
	if v5048 != 0 {
		v5059 = v5048
		goto L684
	} else {
		goto L690
	}
L689:
	;
	v5059 = v5048
	goto L684
L690:
	;
	v5050 = v5014 + int32(1)
	if v5050 != v5001 {
		v5014 = v5050
		goto L688
	} else {
		goto L691
	}
L691:
	;
	goto L689
L692:
	;
	v5088 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+66)))
	v5089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3981)+64)))
	v5092 = v5088 | v5089<<(uint(int32(16))%32)
	if v5092 == int32(-1) {
		v5106 = v5059
		goto L559
	} else {
		goto L693
	}
L693:
	;
	F_tbm_add_page(m, v3977, v5092)
	mBase = m.M
	v5096 = m.ExcPending
	if v5096 != 0 {
		goto L1
	} else {
		goto L694
	}
L694:
	;
	v4004 = v4004 + int64(1)
	goto L557
L695:
	;
	v4004 = v4004 + int64(1)
	goto L557
L696:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v5145 = m.ExcPending
	if v5145 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	v5146 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+1128))
	v5147 = *(*int32)(unsafe.Add(mBase, uint32(v5146)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v2077)+16)) = v5147 + int32(4)
	F_errmsg(m, int32(_a_F_gingetbitmap_15), v2077+int32(16))
	mBase = m.M
	v5155 = m.ExcPending
	if v5155 != 0 {
		goto L1
	} else {
		goto L698
	}
L698:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_11), int32(272), int32(_a_F_gingetbitmap_16))
	mBase = m.M
	v5160 = m.ExcPending
	if v5160 != 0 {
		goto L1
	} else {
		goto L699
	}
L699:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ginint4_queryextract(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v148 int32
	_ = v148
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v215 int64
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v288 int64
	_ = v288
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v2
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int64(0)
	} else {
		v29 = v24 + int32(8)
		v30 = base.I32_wrap_i64(v19)
		if v30&int32(_a_F_ginint4_queryextract_0) == int32(20) {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			if v36 <= int32(0) {
				v288 = int64(0)
				m.G0 = v16 + int32(16)
				return v288
			} else {
				v41 = F_query_has_required_values(m, v24)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					if v41 != 0 {
						v43 = int32(0)
					} else {
						v43 = int32(2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v43
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					v47 = F_palloc_mul(m, int32(8), v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v49 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v49
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						if v51 <= v49 {
							v264 = v47
						} else {
							v55 = int32(0)
							v57 = v2
							v63 = v51
							for {
								v70 = v29 + v55<<(uint(int32(3))%32)
								v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70))))
								if v71 == int32(2) {
									v77 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+4)))
									*(*int64)(unsafe.Add(mBase, uint32(v47+v57<<(uint(int32(3))%32)))) = v77
									v80 = v57 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v20))) = v80
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
									v83 = v80
									v84 = v82
								} else {
									v83 = v57
									v84 = v63
								}
								v86 = v55 + int32(1)
								if v86 < v84 {
									v55 = v86
									v57 = v83
									v63 = v84
									continue
								} else {
									break
								}
								break
							}
							v264 = v47
						}
						v288 = base.I64_extend_i32_u(v264)
						m.G0 = v16 + int32(16)
						return v288
					}
				}
			}
		} else {
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			if v88 != 0 {
				v89 = F_array_contains_nulls(m, v24)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int64(0)
				} else {
					if v89 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v296 = m.ExcPending
						if v296 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v299 = m.ExcPending
							if v299 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_ginint4_queryextract_1), int32(0))
								mBase = m.M
								v303 = m.ExcPending
								if v303 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_ginint4_queryextract_2), int32(61), int32(_a_F_ginint4_queryextract_3))
									mBase = m.M
									v308 = m.ExcPending
									if v308 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						v94 = F_ArrayGetNItemsSafe(m, v91, v24+int32(16))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20))) = v94
							if v94 <= int32(0) {
								v225 = v2
								v226 = v2
								v237 = v30 & int32(_a_F_ginint4_queryextract_0)
								switch v237 - int32(3) {
								case 0:
									v259 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
									v264 = v226
									v288 = base.I64_extend_i32_u(v264)
									m.G0 = v16 + int32(16)
									return v288
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v248 = m.ExcPending
									if v248 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v237
										F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_4), v16)
										mBase = m.M
										v252 = m.ExcPending
										if v252 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_ginint4_queryextract_2), int32(100), int32(_a_F_ginint4_queryextract_3))
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								case 3:
									v259 = v225 ^ int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
									v264 = v226
									v288 = base.I64_extend_i32_u(v264)
									m.G0 = v16 + int32(16)
									return v288
								case 4, 10:
									if v225 != 0 {
										v244 = int32(0)
									} else {
										v244 = int32(2)
									}
									v259 = v244
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
									v264 = v226
									v288 = base.I64_extend_i32_u(v264)
									m.G0 = v16 + int32(16)
									return v288
								case 5, 11:
									v259 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
									v264 = v226
									v288 = base.I64_extend_i32_u(v264)
									m.G0 = v16 + int32(16)
									return v288
								}
							} else {
								v100 = F_palloc_mul(m, int32(8), v94)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int64(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
									if v102 == int32(0) {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
										v112 = (v105<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									} else {
										v112 = v102
									}
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
									if v113 <= int32(0) {
										v225 = int32(0)
										v226 = v100
									} else {
										v117 = v112 + v24
										v119 = v113 & int32(3)
										v120 = int32(0)
										if base.Ui32(v113) < base.Ui32(int32(4)) {
											v183 = v120
											v196 = v183
											v204 = v2
											for {
												v215 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v196<<(uint(int32(2))%32)))))
												*(*int64)(unsafe.Add(mBase, uint32(v100+v196<<(uint(int32(3))%32)))) = v215
												v217 = int32(1)
												v221 = v204 + v217
												if v221 != v119 {
													v196 = v196 + v217
													v204 = v221
													continue
												} else {
													break
												}
												break
											}
											v225 = v217
											v226 = v100
										} else {
											v126 = v120
											v128 = int32(0)
											for {
												v139 = int32(3)
												v142 = int32(2)
												v145 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v126<<(uint(v142)%32)))))
												*(*int64)(unsafe.Add(mBase, uint32(v100+v126<<(uint(v139)%32)))) = v145
												v148 = v126 | int32(1)
												v155 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v148<<(uint(v142)%32)))))
												*(*int64)(unsafe.Add(mBase, uint32(v100+v148<<(uint(v139)%32)))) = v155
												v158 = v126 | v142
												v165 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v158<<(uint(v142)%32)))))
												*(*int64)(unsafe.Add(mBase, uint32(v100+v158<<(uint(v139)%32)))) = v165
												v168 = v126 | v139
												v175 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v168<<(uint(v142)%32)))))
												*(*int64)(unsafe.Add(mBase, uint32(v100+v168<<(uint(v139)%32)))) = v175
												v177 = int32(4)
												v178 = v126 + v177
												v180 = v128 + v177
												if v180 != v113&int32(2147483644) {
													v126 = v178
													v128 = v180
													continue
												} else {
													break
												}
												break
											}
											if v119 != 0 {
												v183 = v178
												v196 = v183
												v204 = v2
												for {
													v215 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v196<<(uint(int32(2))%32)))))
													*(*int64)(unsafe.Add(mBase, uint32(v100+v196<<(uint(int32(3))%32)))) = v215
													v217 = int32(1)
													v221 = v204 + v217
													if v221 != v119 {
														v196 = v196 + v217
														v204 = v221
														continue
													} else {
														break
													}
													break
												}
												v225 = v217
												v226 = v100
											} else {
												v225 = int32(1)
												v226 = v100
											}
										}
									}
									v237 = v30 & int32(_a_F_ginint4_queryextract_0)
									switch v237 - int32(3) {
									case 0:
										v259 = v2
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
										v264 = v226
										v288 = base.I64_extend_i32_u(v264)
										m.G0 = v16 + int32(16)
										return v288
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v248 = m.ExcPending
										if v248 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v237
											F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_4), v16)
											mBase = m.M
											v252 = m.ExcPending
											if v252 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_ginint4_queryextract_2), int32(100), int32(_a_F_ginint4_queryextract_3))
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									case 3:
										v259 = v225 ^ int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
										v264 = v226
										v288 = base.I64_extend_i32_u(v264)
										m.G0 = v16 + int32(16)
										return v288
									case 4, 10:
										if v225 != 0 {
											v244 = int32(0)
										} else {
											v244 = int32(2)
										}
										v259 = v244
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
										v264 = v226
										v288 = base.I64_extend_i32_u(v264)
										m.G0 = v16 + int32(16)
										return v288
									case 5, 11:
										v259 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
										v264 = v226
										v288 = base.I64_extend_i32_u(v264)
										m.G0 = v16 + int32(16)
										return v288
									}
								}
							}
						}
					}
				}
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v94 = F_ArrayGetNItemsSafe(m, v91, v24+int32(16))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v94
					if v94 <= int32(0) {
						v225 = v2
						v226 = v2
						v237 = v30 & int32(_a_F_ginint4_queryextract_0)
						switch v237 - int32(3) {
						case 0:
							v259 = v2
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
							v264 = v226
							v288 = base.I64_extend_i32_u(v264)
							m.G0 = v16 + int32(16)
							return v288
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v248 = m.ExcPending
							if v248 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = v237
								F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_4), v16)
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_ginint4_queryextract_2), int32(100), int32(_a_F_ginint4_queryextract_3))
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v259 = v225 ^ int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
							v264 = v226
							v288 = base.I64_extend_i32_u(v264)
							m.G0 = v16 + int32(16)
							return v288
						case 4, 10:
							if v225 != 0 {
								v244 = int32(0)
							} else {
								v244 = int32(2)
							}
							v259 = v244
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
							v264 = v226
							v288 = base.I64_extend_i32_u(v264)
							m.G0 = v16 + int32(16)
							return v288
						case 5, 11:
							v259 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
							v264 = v226
							v288 = base.I64_extend_i32_u(v264)
							m.G0 = v16 + int32(16)
							return v288
						}
					} else {
						v100 = F_palloc_mul(m, int32(8), v94)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int64(0)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
							if v102 == int32(0) {
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
								v112 = (v105<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v112 = v102
							}
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
							if v113 <= int32(0) {
								v225 = int32(0)
								v226 = v100
							} else {
								v117 = v112 + v24
								v119 = v113 & int32(3)
								v120 = int32(0)
								if base.Ui32(v113) < base.Ui32(int32(4)) {
									v183 = v120
									v196 = v183
									v204 = v2
									for {
										v215 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v196<<(uint(int32(2))%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v100+v196<<(uint(int32(3))%32)))) = v215
										v217 = int32(1)
										v221 = v204 + v217
										if v221 != v119 {
											v196 = v196 + v217
											v204 = v221
											continue
										} else {
											break
										}
										break
									}
									v225 = v217
									v226 = v100
								} else {
									v126 = v120
									v128 = int32(0)
									for {
										v139 = int32(3)
										v142 = int32(2)
										v145 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v126<<(uint(v142)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v100+v126<<(uint(v139)%32)))) = v145
										v148 = v126 | int32(1)
										v155 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v148<<(uint(v142)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v100+v148<<(uint(v139)%32)))) = v155
										v158 = v126 | v142
										v165 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v158<<(uint(v142)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v100+v158<<(uint(v139)%32)))) = v165
										v168 = v126 | v139
										v175 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v168<<(uint(v142)%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v100+v168<<(uint(v139)%32)))) = v175
										v177 = int32(4)
										v178 = v126 + v177
										v180 = v128 + v177
										if v180 != v113&int32(2147483644) {
											v126 = v178
											v128 = v180
											continue
										} else {
											break
										}
										break
									}
									if v119 != 0 {
										v183 = v178
										v196 = v183
										v204 = v2
										for {
											v215 = int64(*(*int32)(unsafe.Add(mBase, uint32(v117+v196<<(uint(int32(2))%32)))))
											*(*int64)(unsafe.Add(mBase, uint32(v100+v196<<(uint(int32(3))%32)))) = v215
											v217 = int32(1)
											v221 = v204 + v217
											if v221 != v119 {
												v196 = v196 + v217
												v204 = v221
												continue
											} else {
												break
											}
											break
										}
										v225 = v217
										v226 = v100
									} else {
										v225 = int32(1)
										v226 = v100
									}
								}
							}
							v237 = v30 & int32(_a_F_ginint4_queryextract_0)
							switch v237 - int32(3) {
							case 0:
								v259 = v2
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
								v264 = v226
								v288 = base.I64_extend_i32_u(v264)
								m.G0 = v16 + int32(16)
								return v288
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v248 = m.ExcPending
								if v248 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v237
									F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_4), v16)
									mBase = m.M
									v252 = m.ExcPending
									if v252 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_ginint4_queryextract_2), int32(100), int32(_a_F_ginint4_queryextract_3))
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 3:
								v259 = v225 ^ int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
								v264 = v226
								v288 = base.I64_extend_i32_u(v264)
								m.G0 = v16 + int32(16)
								return v288
							case 4, 10:
								if v225 != 0 {
									v244 = int32(0)
								} else {
									v244 = int32(2)
								}
								v259 = v244
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
								v264 = v226
								v288 = base.I64_extend_i32_u(v264)
								m.G0 = v16 + int32(16)
								return v288
							case 5, 11:
								v259 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v259
								v264 = v226
								v288 = base.I64_extend_i32_u(v264)
								m.G0 = v16 + int32(16)
								return v288
							}
						}
					}
				}
			}
		}
	}
}
func F_gistbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int64
	_ = v56
	v6 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		v11 = F_initGISTstate(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = int32(_a_F_gistbeginscan_0)
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			*(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0])) = v16
			v19 = F_palloc0(m, int32(_a_F_gistbeginscan_1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v11
				v22 = F_createTempGistContext(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v22
					*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v27
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
					v32 = F_palloc(m, v29<<(uint(int32(4))%32))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)) = uint8(v34)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v32
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
						if v37 <= int32(0) {
							v56 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v19
							*(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0])) = v14
							return v6
						} else {
							v41 = F_palloc0_mul(m, int32(8), v37)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v41
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
								v46 = F_palloc_mul(m, int32(1), v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v46
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
									if v49 == int32(0) {
									} else {
										base.MemoryFill(m, v46, int32(1), v49)
									}
									v56 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = int32(-1)
									*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v19
									*(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0])) = v14
									return v6
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_gistbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 float64
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 float64
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v703 float64
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int64
	_ = v717
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	v21 = F_RelationGetNumberOfBlocksInFork(m, l1, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v19
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	F_MemoryContextDelete(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L3
	} else {
		goto L159
	}
L2:
	;
	v447 = F_gistNewBuffer(m, l1, l0)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L3
	} else {
		goto L101
	}
L3:
	;
	return int32(0)
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	v31 = F_initGISTstate(m, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L3
	} else {
		goto L97
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v31
	v34 = F_createTempGistContext(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v34
	if v17 != 0 {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[1]))
	v139 = m.G0
	v141 = v139 - int32(16)
	m.G0 = v141
	v143 = int32(0)
	v145 = F_tuplesort_begin_common(m, v138, v143, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L34
	}
L11:
	;
	v118 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = int32(819)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v118
	if v58 != v54 {
		goto L2
	} else {
		goto L33
	}
L12:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v107 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v107
	v116 = base.I32_div_s(int32(_a_F_gistbuild_0)-v106<<(uint(int32(13))%32), int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v116
	if v102 != 0 {
		goto L10
	} else {
		goto L32
	}
L13:
	;
	v48 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+10)))
	if v48 < v50 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(1)
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(3)
	v102 = v4
	goto L12
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	switch v37 - int32(1) {
	case 0:
		goto L15
	case 1:
		goto L14
	default:
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(2)
	goto L13
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(2)
	goto L13
L20:
	;
	v54 = v50
	goto L22
L21:
	;
	v54 = v48
	goto L22
L22:
	;
	v58 = v48
	goto L24
L23:
	;
	if v17 == int32(0) {
		goto L11
	} else {
		goto L31
	}
L24:
	;
	v67 = base.B2i32(v58 == v54)
	if v67 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(0)
	goto L23
L26:
	;
	v70 = int32(1)
	v71 = v58 + v70
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+6)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v74+v76*(base.I32_extend16_s(v71)-v70)<<(uint(int32(2))%32)+int32(44)-int32(4))))
	goto L29
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	if v88 != 0 {
		v58 = v71
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L23
L31:
	;
	v102 = v67
	goto L12
L32:
	;
	goto L2
L33:
	;
	goto L10
L34:
	;
	v147 = int32(_a_F_gistbuild_1)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v150
	v153 = F_palloc(m, int32(12))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gistbuild[2])))
	if v156 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v176)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = int32(2054)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+40)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v145)+60)) = v153
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+36)) = uint8(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+16)) = int32(2055)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = int32(2056)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = int32(2057)
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = int32(2058)
	v192 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v153)+8)) = uint16(v192)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = l0
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	v199 = F_palloc0(m, v196*int32(36))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L42
	}
L37:
	;
	v161 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if v161 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = int32(102)
	F_errmsg_internal(m, int32(_a_F_gistbuild_2), v141)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_gistbuild_3), int32(512), int32(_a_F_gistbuild_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+44)) = v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	if v202 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v148
	m.G0 = v141 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v145
	v277 = int32(1)
	v278 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+140))
	v288 = m.T0[v287].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v277, v278, v277, v278, int32(-1), int32(95), v15+int32(32), v278)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L3
	} else {
		goto L51
	}
L44:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v210 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v199)+10)) = uint16(v210)
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+9)) = uint8(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v209
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+20)) = uint8(v216)
	F_PrepareSortSupportFromGistIndexRel(m, l1, v199)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	if v220 < int32(2) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v229 = v210
	goto L47
L47:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v145)+44))
	v238 = v235 + v229*int32(36)
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242+v229<<(uint(int32(2))%32))))
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+20)) = uint8(v247)
	v250 = v229 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v238)+10)) = uint16(v250)
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+9)) = uint8(v247)
	*(*int32)(unsafe.Add(mBase, uint32(v238)+4)) = v246
	F_PrepareSortSupportFromGistIndexRel(m, l1, v238)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L49
	}
L48:
	;
	goto L43
L49:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v145)+40))
	if v250 < v257 {
		v229 = v250
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	F_tuplesort_performsort(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v297 = F_smgr_bulk_start_rel(m, v295, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v297
	v301 = F_palloc0(m, int32(28))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v304 = F_palloc(m, int32(_a_F_gistbuild_5))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v301)+12)) = v304
	v309 = int32(1)
	F_PageInit(m, v304, int32(_a_F_gistbuild_5), int32(16))
	mBase = m.M
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v304)+16)))
	v314 = v304 + v313
	v315 = int32(_a_F_gistbuild_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v314)+14)) = uint16(v315)
	*(*uint16)(unsafe.Add(mBase, uint32(v314)+12)) = uint16(v309)
	*(*int32)(unsafe.Add(mBase, uint32(v314)+8)) = int32(-1)
	goto L56
L56:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v321 = F_tuplesort_getheaptuple(m, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	if v321 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v325 = v321
	goto L61
L59:
	;
	goto L60
L60:
	;
	v361 = v301
	goto L67
L61:
	;
	F_gist_indexsortbuild_levelstate_add(m, v15+int32(32), v301, v325)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L3
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	F_MemoryContextReset(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v344 = F_tuplesort_getheaptuple(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	if v344 != 0 {
		v325 = v344
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	if v370 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v395))) = int64(4294967296)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	v399 = F_smgr_bulk_get_buf(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L3
	} else {
		goto L92
	}
L69:
	;
	goto L68
L70:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	if v373 == int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_gist_indexsortbuild_levelstate_flush(m, v15+int32(32), v361)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	if v381 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_pfree(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L3
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	if v384 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	F_pfree(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	if v387 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L3
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	if v390 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L3
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	F_pfree(m, v361)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L3
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v361 = v380
	goto L67
L92:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	base.MemoryCopy(m, v399, v401, int32(_a_F_gistbuild_5))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	F_smgr_bulk_write(m, v404, int32(0), v399, int32(1))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L3
	} else {
		goto L93
	}
L93:
	;
	F_pfree(m, v361)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	F_smgr_bulk_finish(m, v411)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	F_tuplesort_end(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	v703 = v288
	goto L1
L97:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v421 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistbuild_7), v15+int32(16))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_gistbuild_8), int32(195), int32(_a_F_gistbuild_9))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v467 = int32(_a_F_gistbuild_10)
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3]))
	v470 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3])) = v469 + v470
	if v447 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	if v447 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[4]))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v452+(v447^int32(-1))<<(uint(int32(2))%32))))
	v466 = v458
	goto L100
L103:
	;
	goto L104
L104:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[5]))
	v466 = v460 + v447<<(uint(int32(13))%32) + int32(-8192)
	goto L100
L105:
	;
	F_MarkBufferDirty(m, v447)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L3
	} else {
		goto L110
	}
L106:
	;
	F_PageInit(m, v491, int32(_a_F_gistbuild_5), int32(16))
	mBase = m.M
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491)+16)))
	v496 = v491 + v495
	v497 = int32(_a_F_gistbuild_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v496)+14)) = uint16(v497)
	*(*uint16)(unsafe.Add(mBase, uint32(v496)+12)) = uint16(v470)
	*(*int32)(unsafe.Add(mBase, uint32(v496)+8)) = int32(-1)
	goto L105
L107:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[4]))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v477+(v447^int32(-1))<<(uint(int32(2))%32))))
	v491 = v483
	goto L106
L108:
	;
	goto L109
L109:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[5]))
	v491 = v485 + v447<<(uint(int32(13))%32) + int32(-8192)
	goto L106
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v466))) = int64(4294967296)
	F_UnlockReleaseBuffer(m, v447)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	v508 = int32(_a_F_gistbuild_10)
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3]))
	v511 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3])) = v510 - v511
	v515 = int32(0)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+140))
	v525 = m.T0[v524].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v511, v515, v511, v515, int32(-1), int32(96), v15+int32(32), v515)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v527 == int32(4) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v532 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L3
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+118)))
	if v676 != int32(112) {
		v703 = v525
		goto L1
	} else {
		goto L151
	}
L116:
	;
	if v532 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_errmsg_internal(m, int32(_a_F_gistbuild_11), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L3
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v543 = int32(_a_F_gistbuild_1)
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v547
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+44))
	v552 = v550 - int32(1)
	if int32(0) <= v552 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	F_errfinish(m, int32(_a_F_gistbuild_8), int32(323), int32(_a_F_gistbuild_9))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v555 = v552
	goto L125
L123:
	;
	v648 = v549
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v544
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v648)+4))
	F_BufFileClose(m, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L3
	} else {
		goto L150
	}
L125:
	;
	v568 = v555 << (uint(int32(2)) % 32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v549)+40))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v568+v569)))
	if v571 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v648 = v645
	goto L124
L127:
	;
	v575 = v571
	goto L130
L128:
	;
	goto L129
L129:
	;
	v630 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L3
	} else {
		goto L143
	}
L130:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	if v586 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	goto L129
L132:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v549)+40))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v613+v568)))
	if v615 != 0 {
		v575 = v615
		goto L130
	} else {
		goto L142
	}
L133:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+16)))
	if v587 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v607 = F_list_delete_first(m, v575)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L3
	} else {
		goto L141
	}
L136:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v591
	v593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v585)+16)) = uint8(v593)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v549)+28))
	v596 = F_lcons(m, v585, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L3
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	F_gistProcessEmptyingQueue(m, v15+int32(32))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L3
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v549)+28)) = v596
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v601
	goto L138
L140:
	;
	goto L132
L141:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v549)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v609+v568))) = v607
	goto L132
L142:
	;
	goto L131
L143:
	;
	if v630 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v555
	F_errmsg_internal(m, int32(_a_F_gistbuild_12), v15)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L3
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if int32(0) < v555 {
		v555 = v555 - int32(1)
		goto L125
	} else {
		goto L149
	}
L147:
	;
	F_errfinish(m, int32(_a_F_gistbuild_8), int32(1418), int32(_a_F_gistbuild_13))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	goto L126
L150:
	;
	goto L115
L151:
	;
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[6]))
	if v680 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v683 != 0 {
		v703 = v525
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v685 = int32(0)
	v687 = F_RelationGetNumberOfBlocksInFork(m, l1, v685)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L3
	} else {
		goto L157
	}
L155:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v684 != 0 {
		v703 = v525
		goto L1
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	F_log_newpage_range(m, l1, v685, v687, int32(1))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L3
	} else {
		goto L158
	}
L158:
	;
	v703 = v525
	goto L1
L159:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_brin_free_desc(m, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	v714 = F_palloc(m, int32(16))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L3
	} else {
		goto L161
	}
L161:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v714))) = v703
	v717 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v714)+8)) = base.F64_convert_i64_s(v717)
	m.G0 = v15 + int32(96)
	return v714
}
func F_gistdentryinit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	v7 = l6
	v9 = int32(0)
	if l7 == v9 {
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+18)) = uint8(v12)
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = l4
		*(*int64)(unsafe.Add(mBase, uint32(l2))) = l3
		v20 = l0 + l1*int32(28)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(2712))))
		if v23 == v12 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+uint32(_c_F_gistdentryinit[0])))
			v35 = F_FunctionCall1Coll(m, v20+int32(2708), v33, base.I64_extend_i32_u(l2))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v37 = base.I32_wrap_i64(v35)
				if l2 == v37 {
				} else {
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v43
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+16)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)) = uint16(v45)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+18)))
					v55 = v47
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+18)) = uint8(v55)
				}
				return
			}
		}
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = l4
		*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
		v55 = v9
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+18)) = uint8(v55)
		return
	}
}
func F_gistendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+28)) = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	F_brin_free_desc(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_gistfitpage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v80 int32
	_ = v80
	v3 = int32(0)
	if v3 < l1 {
		if l1 != int32(1) {
			v18 = int32(0)
			v19 = v3
			v20 = v3
			for {
				v24 = int32(2)
				v26 = l0 + v19<<(uint(v24)%32)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+6)))
				v29 = int32(_a_F_gistfitpage_0)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+6)))
				v38 = v20 + v28&v29 + v33&v29 + int32(8)
				v40 = v19 + v24
				v42 = v18 + v24
				if v42 != l1&int32(2147483646) {
					v18 = v42
					v19 = v40
					v20 = v38
					continue
				} else {
					break
				}
				break
			}
			if l1&int32(1) == int32(0) {
				v66 = v38
			} else {
				v48 = v40
				v49 = v38
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v48<<(uint(int32(2))%32))))
				v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+6)))
				v66 = v57&int32(_a_F_gistfitpage_0) + v49 + int32(4)
			}
		} else {
			v48 = v3
			v49 = v3
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0+v48<<(uint(int32(2))%32))))
			v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+6)))
			v66 = v57&int32(_a_F_gistfitpage_0) + v49 + int32(4)
		}
		v80 = base.B2i32(base.Ui32(v66) < base.Ui32(int32(_a_F_gistfitpage_1)))
	} else {
		v80 = int32(1)
	}
	return v80
}
func F_gistgetadjusted(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(2176)
	m.G0 = v27
	F_gistDeCompressAtt(m, l3, l0, l1, v27+int32(1120), v27+int32(320))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_gistDeCompressAtt(m, l3, l0, l2, v27+int32(352), v27+int32(288))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+10)))
	if v44 <= int32(0) {
		v305 = v5
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v27 + int32(2176)
	return v305
L5:
	;
	v50 = l3 + int32(_a_F_gistgetadjusted_0)
	v54 = v27 + int32(1920)
	v56 = v27 + int32(1896)
	v59 = base.I64_extend_i32_u(v27 + int32(1888))
	v62 = base.I64_extend_i32_u(v27 + int32(2152))
	v68 = v5
	v74 = v5
	goto L6
L6:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(288)+v68))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+int32(320)+v68))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+1888)) = int32(2)
	v97 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+2152)) = v97
	v103 = v27 + int32(32) + v68<<(uint(int32(3))%32)
	v104 = v27 + v68
	if v90&v94&int32(1) == v97 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v182 = int32(0)
	if v170 == v182 {
		v305 = v182
		goto L4
	} else {
		goto L29
	}
L8:
	;
	v178 = v68 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179)+10)))
	if v178 < v180 {
		v68 = v178
		v74 = v170
		goto L6
	} else {
		goto L28
	}
L9:
	;
	v111 = v68 * int32(24)
	v114 = v111 + (v27 + int32(352))
	v117 = v27 + int32(1120) + v111
	v119 = v94 & int32(1)
	if v119 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v163)
	*(*int64)(unsafe.Add(mBase, uint32(v103))) = int64(0)
	v170 = v74
	goto L8
L12:
	;
	v120 = v114
	goto L14
L13:
	;
	v120 = v117
	goto L14
L14:
	;
	v121 = v90 | v94
	if v121 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v122 = v120
	goto L17
L16:
	;
	v122 = v117
	goto L17
L17:
	;
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v122)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v123
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v122)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v122)))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v127
	if v121 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v129 = v120
	goto L20
L19:
	;
	v129 = v114
	goto L20
L20:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v129)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v130
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v129)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v132
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v134
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v136)
	v139 = v68 * int32(28)
	v143 = v50 + v68<<(uint(int32(2))%32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v145 = F_FunctionCall2Coll(m, l3+int32(916)+v139, v144, v59, v62)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v103))) = v145
	if v90|v74 != 0 {
		v170 = v74
		goto L8
	} else {
		goto L22
	}
L22:
	;
	if v119 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v117)))
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+1888)) = uint8(v152)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v157 = F_FunctionCall3Coll(m, v139+(l3+int32(_a_F_gistgetadjusted_1)), v156, v151, v145, v59)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v170 = int32(1)
	goto L8
L26:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1888)))
	if v159 != 0 {
		v170 = v152
		goto L8
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	goto L7
L29:
	;
	if int32(0) < v180 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v191 = v179
	v194 = v182
	goto L33
L31:
	;
	goto L32
L32:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v292 = F_index_form_tuple(m, v289, v27+int32(1888), v27)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L44
	}
L33:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v194))))
	if v214 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L32
L35:
	;
	v262 = v194 + int32(1)
	v263 = int32(*(*int16)(unsafe.Add(mBase, uint32(v257)+10)))
	if v262 < v263 {
		v191 = v257
		v194 = v262
		goto L33
	} else {
		goto L43
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(1888)+v194<<(uint(int32(3))%32)))) = int64(0)
	v257 = v191
	goto L35
L37:
	;
	goto L38
L38:
	;
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2170)) = uint8(v224)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+2168)) = uint16(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+2164)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v27)+2160)) = l0
	v232 = v194 << (uint(int32(3)) % 32)
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v232+(v27+int32(32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+2152)) = v236
	v240 = l3 + int32(1812) + v194*int32(28)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v241 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v50+v194<<(uint(int32(2))%32))))
	v246 = F_FunctionCall1Coll(m, v240, v245, v62)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v251 = v191
	v252 = v236
	goto L41
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27+int32(1888)+v232))) = v252
	v257 = v251
	goto L35
L42:
	;
	v249 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v246))))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v251 = v250
	v252 = v249
	goto L41
L43:
	;
	goto L34
L44:
	;
	v294 = int32(_a_F_gistgetadjusted_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v292)+4)) = uint16(v294)
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v292)+4)) = uint16(v296)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = v298
	v305 = v292
	goto L4
}
func F_gistgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l1 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v461
L2:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v395 != 0 {
		goto L91
	} else {
		goto L92
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v17 != int32(1) {
		v461 = int32(0)
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L15
	} else {
		goto L88
	}
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+17)))
	if v20 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+272))
	if v24 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v67 {
		goto L2
	} else {
		goto L25
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+268)))
	if v27 != int32(1) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v36 = v24
	goto L13
L13:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37 + int64(1)
	goto L10
L14:
	;
	F_pgstat_assoc_relation(m, v23)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+272))
	v36 = v35
	goto L13
L17:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v43 + int64(1)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[0]))) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+17)) = uint8(v47)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v47
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[1])))
	if v53 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_MemoryContextReset(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
	v58 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v58
	F_gistScanPage(m, l0, v12, v58, v58, v58)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	goto L9
L25:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[0]))))
	if base.Ui32(v71) <= base.Ui32(v70) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v75 = v70
	v76 = v71
	goto L29
L27:
	;
	v298 = v70
	goto L28
L28:
	;
	if v298 == int32(0) {
		v347 = v298
		goto L77
	} else {
		goto L78
	}
L29:
	;
	v82 = int32(_a_F_gistgettuple_0)
	v83 = v75 & v82
	if base.B2i32(v83 != v76&v82)|base.B2i32(v83 == int32(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v298 = v294
	goto L28
L31:
	;
	goto L39
L32:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v90&int32(1) == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v95 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v98 = int32(_a_F_gistgettuple_1)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v102
	v105 = F_palloc(m, int32(816))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L37
	}
L35:
	;
	v110 = v95
	goto L36
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if int32(407) < v112 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v105
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v99
	v110 = v105
	goto L36
L38:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+v115<<(uint(int32(4))%32))+44)))
	v120 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v112 + v120
	*(*uint16)(unsafe.Add(mBase, uint32(v110+v112<<(uint(v120)%32)))) = uint16(v119)
	goto L31
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v139 == int32(-1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	if base.Ui32(v291) <= base.Ui32(v294) {
		v75 = v294
		v76 = v291
		goto L29
	} else {
		goto L76
	}
L41:
	;
	v268 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+8))
	if v270 == v268 {
		v461 = v268
		goto L1
	} else {
		goto L66
	}
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if v142 <= int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
	v148 = F_ReadBuffer(m, v145, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	if v148 == int32(0) {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	F_LockBufferInternal(m, v148, int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_gistcheckpage(m, v155, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	if v148 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v176 = F_BufferGetLSNAtomic(m, v148)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L15
	} else {
		goto L53
	}
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[4]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161+(v148^int32(-1))<<(uint(int32(2))%32))))
	v175 = v167
	goto L48
L50:
	;
	goto L51
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[5]))
	v175 = v169 + v148<<(uint(int32(13))%32) + int32(-8192)
	goto L48
L52:
	;
	F_UnlockReleaseBuffer(m, v148)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L15
	} else {
		goto L65
	}
L53:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v146)+40))
	if v176 != v178 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v146)+28))
	if v180 <= int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v183 = F_BufferBeginSetHintBits(m, v148)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	if v183 == int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v188 = v175 + int32(20)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189))))
	v191 = int32(2)
	v193 = v188 + v190<<(uint(v191)%32)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v194 | int32(_a_F_gistgettuple_2)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v146)+28))
	if v191 <= v199 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v206 = int32(1)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+16)))
	v237 = v175 + v236
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)))
	v240 = v238 | int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)) = uint16(v240)
	v242 = int32(1)
	F_BufferFinishSetHintBits(m, v148, v242, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L15
	} else {
		goto L64
	}
L61:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	v212 = int32(1)
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211+v206<<(uint(v212)%32)))))
	v218 = v188 + v215<<(uint(int32(2))%32)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v219 | int32(_a_F_gistgettuple_2)
	v224 = v206 + v212
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v146)+28))
	if v224 < v225 {
		v206 = v224
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	goto L62
L64:
	;
	goto L52
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+28)) = int32(0)
	goto L41
L66:
	;
	v273 = F_pairingheap_remove_first(m, v269)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	if v273 == int32(0) {
		v461 = v268
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[6]))
	if v278 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L15
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v281
	v285 = int32(0)
	F_gistScanPage(m, l0, v273, v273+int32(32), v285, v285)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L15
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_pfree(m, v273)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L15
	} else {
		goto L74
	}
L74:
	;
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[0]))))
	if v291 == int32(0) {
		goto L39
	} else {
		goto L75
	}
L75:
	;
	goto L40
L76:
	;
	goto L30
L77:
	;
	v351 = v16 + int32(48)
	v354 = int32(4)
	v356 = v351 + v347&int32(_a_F_gistgettuple_0)<<(uint(v354)%32)
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v356)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v357)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v359
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351+v361<<(uint(v354)%32))+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v365)
	v367 = int32(1)
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v368 == v367 {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v307&int32(1) == int32(0) {
		v347 = v298
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v312 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v315 = int32(_a_F_gistgettuple_1)
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3]))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v319
	v322 = F_palloc(m, int32(816))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L15
	} else {
		goto L83
	}
L81:
	;
	v328 = v298
	v329 = v312
	goto L82
L82:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if int32(407) < v330 {
		v347 = v328
		goto L77
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v322
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v316
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v328 = v327
	v329 = v322
	goto L82
L84:
	;
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+v328&int32(_a_F_gistgettuple_0)<<(uint(int32(4))%32))+44)))
	v339 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v330 + v339
	*(*uint16)(unsafe.Add(mBase, uint32(v329+v330<<(uint(v339)%32)))) = uint16(v338)
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v347 = v346
	goto L77
L85:
	;
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v351+v371<<(uint(int32(4))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v375
	goto L87
L86:
	;
	goto L87
L87:
	;
	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v379 = v377 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))) = uint16(v379)
	v461 = v367
	goto L1
L88:
	;
	F_errmsg_internal(m, int32(_a_F_gistgettuple_3), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L15
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_gistgettuple_4), int32(629), int32(_a_F_gistgettuple_5))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L15
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_pfree(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L15
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	if v401 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	goto L93
L95:
	;
	v461 = int32(0)
	goto L1
L96:
	;
	goto L97
L97:
	;
	v406 = l0 + int32(60)
	v408 = v400
	goto L98
L98:
	;
	v416 = F_pairingheap_remove_first(m, v408)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L15
	} else {
		goto L100
	}
L99:
	;
	v461 = v448
	goto L1
L100:
	;
	if v416 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v461 = int32(0)
	goto L1
L102:
	;
	goto L103
L103:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	if v421 == int32(-1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v416)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v406)+4)) = uint16(v424)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v416)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v406))) = v426
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v428)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+23)))
	F_index_store_float8_orderby_distances(m, l0, v430, v416+int32(32), v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L15
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[6]))
	if v445 != 0 {
		goto L112
	} else {
		goto L113
	}
L107:
	;
	v436 = int32(1)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v437 == v436 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v416)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v440
	goto L110
L109:
	;
	goto L110
L110:
	;
	F_pfree(m, v416)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L15
	} else {
		goto L111
	}
L111:
	;
	v461 = v436
	goto L1
L112:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L15
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v448 = int32(0)
	F_gistScanPage(m, l0, v416, v416+int32(32), v448, v448)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L15
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	F_pfree(m, v416)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L15
	} else {
		goto L117
	}
L117:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+8))
	if v458 != 0 {
		v408 = v457
		goto L98
	} else {
		goto L118
	}
L118:
	;
	goto L99
}
func F_gistinserttuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v11 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 < v11 {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_gistinserttuples[0]))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+(v18^int32(-1))*int32(56))+16))
		v37 = v28
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_gistinserttuples[1]))
		v31 = int32(56)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+v18*v31-v31)+16))
		v37 = v36
	}
	F_CheckForSerializableConflictIn(m, v16, v11, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		v51 = F_gistplacetopage(m, v42, v43, l2, v44, l3, l4, l5, int32(0), l6, v14+int32(12), int32(1), v49, v50)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			if l7 != 0 {
				F_UnlockReleaseBuffer(m, l7)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = int32(0)
					if base.B2i32(l6 == v55)|base.B2i32(l9 == v55) == v55 {
						F_UnlockBuffer(m, l6)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
							if v64 != 0 {
								F_gistfinishsplit(m, l0, l1, l2, v64, l8)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(16)
									return v51
								}
							} else {
								if l8 == int32(0) {
									m.G0 = v14 + int32(16)
									return v51
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									F_UnlockBuffer(m, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										m.G0 = v14 + int32(16)
										return v51
									}
								}
							}
						}
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						if v64 != 0 {
							F_gistfinishsplit(m, l0, l1, l2, v64, l8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								m.G0 = v14 + int32(16)
								return v51
							}
						} else {
							if l8 == int32(0) {
								m.G0 = v14 + int32(16)
								return v51
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								F_UnlockBuffer(m, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(16)
									return v51
								}
							}
						}
					}
				}
			} else {
				v55 = int32(0)
				if base.B2i32(l6 == v55)|base.B2i32(l9 == v55) == v55 {
					F_UnlockBuffer(m, l6)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						if v64 != 0 {
							F_gistfinishsplit(m, l0, l1, l2, v64, l8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								m.G0 = v14 + int32(16)
								return v51
							}
						} else {
							if l8 == int32(0) {
								m.G0 = v14 + int32(16)
								return v51
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								F_UnlockBuffer(m, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(16)
									return v51
								}
							}
						}
					}
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					if v64 != 0 {
						F_gistfinishsplit(m, l0, l1, l2, v64, l8)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							m.G0 = v14 + int32(16)
							return v51
						}
					} else {
						if l8 == int32(0) {
							m.G0 = v14 + int32(16)
							return v51
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							F_UnlockBuffer(m, v69)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								m.G0 = v14 + int32(16)
								return v51
							}
						}
					}
				}
			}
		}
	}
}
func F_gistoptions(m *base.Module, l0 int64, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(32), int32(12), int32(_a_F_gistoptions_0), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gisttranslatecmptype(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = int32(2276)
	v6 = F_get_opfamily_proc(m, l1, v3, v3, int32(12))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v12 = F_OidFunctionCall1Coll(m, v6, int32(0), base.I64_extend_i32_s(l0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v16 = base.I32_wrap_i64(v12)
				return v16 & int32(_a_F_gisttranslatecmptype_0)
			}
		} else {
			v16 = int32(0)
			return v16 & int32(_a_F_gisttranslatecmptype_0)
		}
	}
}
func F_gistvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v638 int32
	_ = v638
	var v649 int64
	_ = v649
	var v650 int64
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int64
	_ = v696
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	v18 = m.G0
	v20 = v18 - int32(304)
	m.G0 = v20
	v24 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if int32(0) < v332 {
		goto L73
	} else {
		goto L74
	}
L2:
	;
	return int32(0)
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
	v30 = v28 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+92))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	v34 = F_get_opfamily_name(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L70
	}
L7:
	;
	v38 = base.I64_extend_i32_u(v33)
	v39 = int64(0)
	v41 = F_SearchSysCacheList(m, int32(4), int32(1), v38, v39, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v43 = int32(1)
	v46 = int64(0)
	v48 = F_SearchSysCacheList(m, int32(5), v43, v38, v46, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v50 <= int32(0) {
		v320 = v43
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v32
	goto L13
L12:
	;
	v53 = v31
	goto L13
L13:
	;
	v60 = int32(0)
	v61 = v43
	goto L14
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v48-int32(-64)+v60<<(uint(int32(2))%32))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+72))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
	v79 = v77 + v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v80 == v81 {
		v110 = v61
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v320 = v296
	goto L1
L16:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v111 != v31 {
		v296 = v110
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v83 = int32(0)
	v86 = F_errstart(m, int32(17), v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v86 == int32(0) {
		v110 = v83
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v94 = F_format_procedure(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+296)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v20)+292)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+288)) = v34
	F_errmsg(m, int32(_a_F_gistvalidate_1), v20+int32(288))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(86), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v110 = v83
	goto L16
L24:
	;
	v299 = v60 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v299 < v300 {
		v60 = v299
		v61 = v296
		goto L14
	} else {
		goto L69
	}
L25:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+16)))
	switch v113 - int32(1) {
	case 0:
		goto L38
	case 1:
		goto L28
	case 2, 3, 8:
		goto L37
	case 4:
		goto L36
	case 5:
		goto L35
	case 6:
		goto L34
	case 7:
		goto L33
	case 9:
		goto L32
	case 10:
		goto L31
	case 11:
		goto L30
	default:
		goto L29
	}
L26:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L65
	}
L27:
	;
	v264 = int32(0)
	v267 = F_errstart(m, int32(17), v264)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L2
	} else {
		goto L63
	}
L28:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+144)) = int64(9796820404457)
	v258 = int32(2)
	v262 = F_check_amproc_signature(m, v254, v53, int32(0), v258, v258, v20+int32(144))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L2
	} else {
		goto L61
	}
L29:
	;
	v245 = int32(0)
	v248 = F_errstart(m, int32(17), v245)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L2
	} else {
		goto L59
	}
L30:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = int32(23)
	v230 = int32(1)
	v235 = F_check_amproc_signature(m, v226, int32(21), v230, v230, v230, v20+int32(272))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L55
	}
L31:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = int32(2281)
	v217 = int32(1)
	v222 = F_check_amproc_signature(m, v213, int32(2278), v217, v217, v217, v20+int32(256))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L53
	}
L32:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v209 = F_check_amoptsproc_signature(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L51
	}
L33:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v191 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v20)+232)) = int64(111669149717)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v191
	v200 = int32(5)
	v204 = F_check_amproc_signature(m, v190, int32(701), int32(0), v200, v200, v20+int32(224))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L49
	}
L34:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v176 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+216)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v53
	v182 = int32(3)
	v186 = F_check_amproc_signature(m, v175, v176, int32(0), v182, v182, v20+int32(208))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L47
	}
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+192)) = int64(9796820404457)
	v167 = int32(2)
	v171 = F_check_amproc_signature(m, v162, int32(2281), int32(1), v167, v167, v20+int32(192))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L45
	}
L36:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v148 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(9796820404457)
	v154 = int32(3)
	v158 = F_check_amproc_signature(m, v147, v148, int32(1), v154, v154, v20+int32(176))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L43
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v135 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v135
	v138 = int32(1)
	v143 = F_check_amproc_signature(m, v134, v135, v138, v138, v138, v20+int32(160))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L41
	}
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v117 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = int64(111669149717)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v117
	v126 = int32(5)
	v130 = F_check_amproc_signature(m, v116, int32(16), int32(0), v126, v126, v20+int32(112))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	if v130 == int32(0) {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v296 = v110
	goto L24
L41:
	;
	if v143 == int32(0) {
		goto L27
	} else {
		goto L42
	}
L42:
	;
	v296 = v110
	goto L24
L43:
	;
	if v158 == int32(0) {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	v296 = v110
	goto L24
L45:
	;
	if v171 == int32(0) {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v296 = v110
	goto L24
L47:
	;
	if v186 == int32(0) {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	v296 = v110
	goto L24
L49:
	;
	if v204 == int32(0) {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v296 = v110
	goto L24
L51:
	;
	if v209 == int32(0) {
		goto L27
	} else {
		goto L52
	}
L52:
	;
	v296 = v110
	goto L24
L53:
	;
	if v222 == int32(0) {
		goto L27
	} else {
		goto L54
	}
L54:
	;
	v296 = v110
	goto L24
L55:
	;
	if v235 == int32(0) {
		goto L27
	} else {
		goto L56
	}
L56:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v239 != int32(2276) {
		goto L27
	} else {
		goto L57
	}
L57:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v242 != int32(2276) {
		goto L27
	} else {
		goto L58
	}
L58:
	;
	v296 = v110
	goto L24
L59:
	;
	if v248 == int32(0) {
		v296 = v245
		goto L24
	} else {
		goto L60
	}
L60:
	;
	v273 = int32(153)
	v274 = int32(_a_F_gistvalidate_4)
	goto L26
L61:
	;
	if v262 != 0 {
		v296 = v110
		goto L24
	} else {
		goto L62
	}
L62:
	;
	goto L27
L63:
	;
	if v267 == int32(0) {
		v296 = v264
		goto L24
	} else {
		goto L64
	}
L64:
	;
	v273 = int32(165)
	v274 = int32(_a_F_gistvalidate_5)
	goto L26
L65:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v279 = F_format_procedure(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v34
	F_errmsg(m, v274, v20+int32(96))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), v273, int32(_a_F_gistvalidate_3))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v296 = int32(0)
	goto L24
L69:
	;
	goto L15
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_gistvalidate_6), v20)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(52), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v342 = int32(0)
	v343 = v320
	goto L76
L74:
	;
	v509 = v320
	goto L75
L75:
	;
	v521 = F_identify_opfamily_groups(m, v41, v48)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L2
	} else {
		goto L117
	}
L76:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(-64)+v342<<(uint(int32(2))%32))))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+72))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+22)))
	v361 = v359 + v360
	v362 = int32(*(*int16)(unsafe.Add(mBase, uint32(v361)+16)))
	if int32(0) < v362 {
		v395 = v343
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v509 = v499
	goto L75
L78:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+18)))
	if v397 == int32(115) {
		v465 = int32(16)
		v466 = v395
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v365 = int32(0)
	v368 = F_errstart(m, int32(17), v365)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	if v368 == int32(0) {
		v395 = v365
		goto L78
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v376 = F_format_operator(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v378 = int32(*(*int16)(unsafe.Add(mBase, uint32(v361)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v34
	F_errmsg(m, int32(_a_F_gistvalidate_7), v20+int32(80))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(185), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v395 = v365
	goto L78
L86:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	v470 = F_check_amop_signature(m, v467, v465, v468, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L2
	} else {
		goto L107
	}
L87:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v402 = F_get_opfamily_proc(m, v33, v400, v400, int32(8))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L2
	} else {
		goto L89
	}
L88:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v433 = F_get_op_rettype(m, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L2
	} else {
		goto L97
	}
L89:
	;
	if v402 != 0 {
		v431 = v395
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v404 = int32(0)
	v407 = F_errstart(m, int32(17), v404)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	if v407 == int32(0) {
		v431 = v404
		goto L88
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v415 = F_format_operator(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v415
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v34
	F_errmsg(m, int32(_a_F_gistvalidate_8), v20-int32(-64))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(202), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	v431 = v404
	goto L88
L97:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v361)+28))
	v436 = F_opfamily_can_sort_type(m, v435, v433)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	if v436 != 0 {
		v465 = v433
		v466 = v431
		goto L86
	} else {
		goto L99
	}
L99:
	;
	v438 = int32(0)
	v441 = F_errstart(m, int32(17), v438)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v441 == int32(0) {
		v465 = v433
		v466 = v438
		goto L86
	} else {
		goto L101
	}
L101:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v449 = F_format_operator(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v34
	F_errmsg(m, int32(_a_F_gistvalidate_9), v20+int32(48))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(213), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v465 = v433
	v466 = v438
	goto L86
L106:
	;
	v501 = v342 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if v501 < v502 {
		v342 = v501
		v343 = v499
		goto L76
	} else {
		goto L115
	}
L107:
	;
	if v470 != 0 {
		v499 = v466
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v472 = int32(0)
	v475 = F_errstart(m, int32(17), v472)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	if v475 == int32(0) {
		v499 = v472
		goto L106
	} else {
		goto L110
	}
L110:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v483 = F_format_operator(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v34
	F_errmsg(m, int32(_a_F_gistvalidate_10), v20+int32(32))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(232), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v499 = v472
	goto L106
L115:
	;
	goto L77
L116:
	;
	v638 = v509
	v649 = int64(1)
	goto L150
L117:
	;
	if v521 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v613 = int32(0)
	goto L116
L119:
	;
	goto L120
L120:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v526 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v613 = int32(0)
	goto L116
L122:
	;
	goto L123
L123:
	;
	v530 = int32(0)
	if v526 != int32(1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v534 = int32(0)
	if v534 < v526 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v586 = v530
	v590 = v530
	goto L126
L126:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v521)+12))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603+v590<<(uint(int32(2))%32))))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	if v608 != v31 {
		v613 = v586
		goto L116
	} else {
		goto L146
	}
L127:
	;
	v537 = v526
	goto L129
L128:
	;
	v537 = v534
	goto L129
L129:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v521)+12))
	v543 = int32(0)
	v545 = v543
	v548 = v543
	v549 = v530
	goto L130
L130:
	;
	v564 = v542 + v549<<(uint(int32(2))%32)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	if v31 == v566 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v537&int32(1) == int32(0) {
		v613 = v578
		goto L116
	} else {
		goto L145
	}
L132:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	if v568 == v31 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v571 = v545
	goto L134
L134:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	if v31 == v573 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v570 = v565
	goto L137
L136:
	;
	v570 = v545
	goto L137
L137:
	;
	v571 = v570
	goto L134
L138:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v575 == v31 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v578 = v571
	goto L140
L140:
	;
	v579 = int32(2)
	v580 = v549 + v579
	v582 = v548 + v579
	if v582 != v537&int32(2147483646) {
		v545 = v578
		v548 = v582
		v549 = v580
		goto L130
	} else {
		goto L144
	}
L141:
	;
	v577 = v572
	goto L143
L142:
	;
	v577 = v571
	goto L143
L143:
	;
	v578 = v577
	goto L140
L144:
	;
	goto L131
L145:
	;
	v586 = v578
	v590 = v580
	goto L126
L146:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	if v610 == v31 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v612 = v607
	goto L149
L148:
	;
	v612 = v586
	goto L149
L149:
	;
	v613 = v612
	goto L116
L150:
	;
	if v613 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	F_ReleaseCatCacheList(m, v48)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L2
	} else {
		goto L167
	}
L152:
	;
	v696 = v649 + int64(1)
	if v696 != int64(13) {
		v638 = v694
		v649 = v696
		goto L150
	} else {
		goto L166
	}
L153:
	;
	v650 = *(*int64)(unsafe.Add(mBase, uint32(v613)+16))
	if base.I32_wrap_i64(int64(base.Ui64(v650)>>(uint(v649)%64)))&int32(1) != 0 {
		v694 = v638
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v655 = base.I32_wrap_i64(v649)
	v656 = int32(12)
	if int32(1)<<(uint(v655)%32)&int32(_a_F_gistvalidate_11) != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L155
L157:
	;
	v667 = base.B2i32(base.Ui32(v655) <= base.Ui32(v656))
	goto L159
L158:
	;
	v667 = int32(0)
	goto L159
L159:
	;
	if base.B2i32(v655&v656 == int32(8))|v667 != 0 {
		v694 = v638
		goto L152
	} else {
		goto L160
	}
L160:
	;
	v669 = int32(0)
	v672 = F_errstart(m, int32(17), v669)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L2
	} else {
		goto L161
	}
L161:
	;
	if v672 == int32(0) {
		v694 = v669
		goto L152
	} else {
		goto L162
	}
L162:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L2
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v655
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v30 + int32(8)
	F_errmsg(m, int32(_a_F_gistvalidate_12), v20+int32(16))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(273), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	v694 = v669
	goto L152
L166:
	;
	goto L151
L167:
	;
	F_ReleaseCatCacheList(m, v41)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	F_ReleaseCatCache(m, v24)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	m.G0 = v20 + int32(304)
	return v694 & int32(1)
}
func F_gseg_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	v2 = int32(0)
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v2)
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+16)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v16)+12)))
	if v18&int32(1) != 0 {
		v21 = int32(1)
		v22 = v8 - v21
		v24 = v22 & int32(_a_F_gseg_consistent_0)
		if base.B2i32(base.Ui32(int32(13)) < base.Ui32(v24))|base.B2i32(int32(base.Ui32(int32(_a_F_gseg_consistent_1))>>(uint(v24)%32))&v21 == int32(0)) != 0 {
			v95 = int64(0)
			return v95
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v22&int32(_a_F_gseg_consistent_0)<<(uint(int32(2))%32))+uint32(_c_F_gseg_consistent[0])))
			v42 = F_DirectFunctionCall2Coll(m, v40, int32(0), v14, v9)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int64(0)
			} else {
				return v42
			}
		}
	} else {
		switch v8 - int32(1) {
		case 0:
			v51 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_2), int32(0), v14, v9)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				v91 = base.B2i32(v51 == int64(0))
				v95 = base.I64_extend_i32_u(v91)
				return v95
			}
		case 1:
			v57 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_3), int32(0), v14, v9)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				v91 = base.B2i32(v57 == int64(0))
				v95 = base.I64_extend_i32_u(v91)
				return v95
			}
		case 2:
			v63 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_4), int32(0), v14, v9)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int64(0)
			} else {
				v91 = base.B2i32(v63 != int64(0))
				v95 = base.I64_extend_i32_u(v91)
				return v95
			}
		case 3:
			v69 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_5), int32(0), v14, v9)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int64(0)
			} else {
				v91 = base.B2i32(v69 == int64(0))
				v95 = base.I64_extend_i32_u(v91)
				return v95
			}
		case 4:
			v75 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_6), int32(0), v14, v9)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int64(0)
			} else {
				v91 = base.B2i32(v75 == int64(0))
				v95 = base.I64_extend_i32_u(v91)
				return v95
			}
		case 5, 6, 12:
			v81 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_7), int32(0), v14, v9)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int64(0)
			} else {
				v91 = base.B2i32(v81 != int64(0))
				v95 = base.I64_extend_i32_u(v91)
				return v95
			}
		case 7, 13:
			v87 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_4), int32(0), v14, v9)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int64(0)
			} else {
				v91 = base.B2i32(v87 != int64(0))
				v95 = base.I64_extend_i32_u(v91)
				return v95
			}
		default:
			v91 = v2
			v95 = base.I64_extend_i32_u(v91)
			return v95
		}
	}
}
func F_gtsquery_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)))) = uint8(base.B2i32(v5 == v6))
	return v3 & int64(4294967295)
}
func F_gtsvectorout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v150 int64
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int64
	_ = v179
	var v189 int64
	_ = v189
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		if v20&int32(1) != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v24 = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(int32(base.Ui32(v23)>>(uint(v24)%32))-int32(8)) >> (uint(v24) % 32))
			v34 = F_psprintf(m, int32(_a_F_gtsvectorout_0), v13+int32(16))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				v219 = v34
				v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v220 != v16 {
					F_pfree(m, v16)
					mBase = m.M
					v223 = m.ExcPending
					if v223 != 0 {
						return int64(0)
					} else {
						m.G0 = v13 + int32(32)
						return base.I64_extend_i32_u(v219)
					}
				} else {
					m.G0 = v13 + int32(32)
					return base.I64_extend_i32_u(v219)
				}
			}
		} else {
			if v20&int32(4) != 0 {
				v39 = F_pstrdup(m, int32(_a_F_gtsvectorout_1))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					v219 = v39
					v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v220 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int64(0)
						} else {
							m.G0 = v13 + int32(32)
							return base.I64_extend_i32_u(v219)
						}
					} else {
						m.G0 = v13 + int32(32)
						return base.I64_extend_i32_u(v219)
					}
				}
			} else {
				v41 = int32(8)
				v42 = v16 + v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
				v47 = v45 - v41
				if base.Ui32(v43) <= base.Ui32(int32(63)) {
					if v47 == int32(0) {
						v199 = v10
					} else {
						v52 = int32(3)
						v53 = v45 & v52
						if base.Ui32(v52) <= base.Ui32(v45-int32(9)) {
							v61 = v42
							v67 = int32(0)
							v69 = v10
							for {
								v70 = int32(4)
								v71 = v61 + v70
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
								v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_gtsvectorout[0]))))
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
								v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_gtsvectorout[0]))))
								v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
								v77 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_gtsvectorout[0]))))
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								v79 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v78)+uint32(_c_F_gtsvectorout[0]))))
								v83 = v73 + (v75 + (v77 + (v69 + v79)))
								v85 = v67 + v70
								if v85 != v47&int32(-4) {
									v61 = v71
									v67 = v85
									v69 = v83
									continue
								} else {
									break
								}
								break
							}
							if v53 == int32(0) {
								v199 = v83
							} else {
								v90 = v71
								v98 = v83
								v101 = v90
								v102 = int32(0)
								v109 = v98
								for {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
									v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_gtsvectorout[0]))))
									v112 = v109 + v111
									v113 = int32(1)
									v116 = v102 + v113
									if v116 != v53 {
										v101 = v101 + v113
										v102 = v116
										v109 = v112
										continue
									} else {
										break
									}
									break
								}
								v199 = v112
							}
						} else {
							v90 = v42
							v98 = v10
							v101 = v90
							v102 = int32(0)
							v109 = v98
							for {
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
								v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_gtsvectorout[0]))))
								v112 = v109 + v111
								v113 = int32(1)
								v116 = v102 + v113
								if v116 != v53 {
									v101 = v101 + v113
									v102 = v116
									v109 = v112
									continue
								} else {
									break
								}
								break
							}
							v199 = v112
						}
					}
				} else {
					v118 = int64(0)
					v119 = int32(0)
					if v47 == v119 {
						v189 = int64(0)
					} else {
						v126 = v47 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v47) {
							v131 = v42
							v133 = v118
							v136 = v119
							for {
								v137 = int32(4)
								v138 = v131 + v137
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+3)))
								v140 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_gtsvectorout[0]))))
								v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+2)))
								v142 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_gtsvectorout[0]))))
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
								v144 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_gtsvectorout[0]))))
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
								v146 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v145)+uint32(_c_F_gtsvectorout[0]))))
								v150 = v140 + (v142 + (v144 + (v133 + v146)))
								v152 = v136 + v137
								if v152 != v47&int32(-4) {
									v131 = v138
									v133 = v150
									v136 = v152
									continue
								} else {
									break
								}
								break
							}
							if v126 == int32(0) {
								v179 = v150
							} else {
								v156 = v138
								v158 = v150
								v163 = v156
								v164 = int32(0)
								v165 = v158
								for {
									v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
									v170 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v169)+uint32(_c_F_gtsvectorout[0]))))
									v171 = v165 + v170
									v172 = int32(1)
									v175 = v164 + v172
									if v175 != v126 {
										v163 = v163 + v172
										v164 = v175
										v165 = v171
										continue
									} else {
										break
									}
									break
								}
								v179 = v171
							}
						} else {
							v156 = v42
							v158 = v118
							v163 = v156
							v164 = int32(0)
							v165 = v158
							for {
								v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
								v170 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v169)+uint32(_c_F_gtsvectorout[0]))))
								v171 = v165 + v170
								v172 = int32(1)
								v175 = v164 + v172
								if v175 != v126 {
									v163 = v163 + v172
									v164 = v175
									v165 = v171
									continue
								} else {
									break
								}
								break
							}
							v179 = v171
						}
						v189 = v179
					}
					v199 = v189
				}
				v200 = base.I32_wrap_i64(v199)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v200
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v47<<(uint(int32(3))%32) - v200
				v207 = F_psprintf(m, int32(_a_F_gtsvectorout_2), v13)
				mBase = m.M
				v208 = m.ExcPending
				if v208 != 0 {
					return int64(0)
				} else {
					v219 = v207
					v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v220 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int64(0)
						} else {
							m.G0 = v13 + int32(32)
							return base.I64_extend_i32_u(v219)
						}
					} else {
						m.G0 = v13 + int32(32)
						return base.I64_extend_i32_u(v219)
					}
				}
			}
		}
	}
}
