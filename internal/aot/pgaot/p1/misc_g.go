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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
	if v14 == v3 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v23 = F_emscripten_builtin_malloc(m, v18<<(uint(int32(3))%32)+int32(8))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1])) = v23
		if v23 == int32(0) {
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
						F_errfinish(m, int32(_a_F_GetConflictingVirtualXIDs_2), int32(3436), int32(_a_F_GetConflictingVirtualXIDs_3))
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
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
			v33 = F_LWLockAcquire(m, v29+int32(512), int32(1))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if int32(0) < v37 {
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
					v49 = v3
					v50 = v3
					v51 = v43
					for {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v49<<(uint(int32(2))%32))))
						v60 = v51 + v57*int32(640)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+44))
						if v61 == int32(0) {
							v99 = v50
							v100 = v51
						} else {
							if l1 != 0 {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+60))
								if v64 != l1 {
									v99 = v50
									v100 = v51
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
									if l0 != 0 {
										if v66 == int32(0) {
											v99 = v50
											v100 = v51
										} else {
											if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v66)) == int32(0) {
												v80 = base.B2i32(base.Ui32(l0) < base.Ui32(v66))
											} else {
												v80 = base.B2i32(int32(0) < v66-l0)
											}
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
											if v80 != 0 {
												v99 = v50
												v100 = v82
											} else {
												v83 = v82
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
												if v84 == int32(0) {
													v99 = v50
													v100 = v83
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
													v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
													v92 = v89 + v50<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
													*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
													v99 = v50 + int32(1)
													v100 = v83
												}
											}
										}
									} else {
										v83 = v51
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
										if v84 == int32(0) {
											v99 = v50
											v100 = v83
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
											v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
											v92 = v89 + v50<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
											*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
											v99 = v50 + int32(1)
											v100 = v83
										}
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
								if l0 != 0 {
									if v66 == int32(0) {
										v99 = v50
										v100 = v51
									} else {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v66)) == int32(0) {
											v80 = base.B2i32(base.Ui32(l0) < base.Ui32(v66))
										} else {
											v80 = base.B2i32(int32(0) < v66-l0)
										}
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
										if v80 != 0 {
											v99 = v50
											v100 = v82
										} else {
											v83 = v82
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
											if v84 == int32(0) {
												v99 = v50
												v100 = v83
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
												v92 = v89 + v50<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
												*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
												v99 = v50 + int32(1)
												v100 = v83
											}
										}
									}
								} else {
									v83 = v51
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
									if v84 == int32(0) {
										v99 = v50
										v100 = v83
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
										v92 = v89 + v50<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
										*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
										v99 = v50 + int32(1)
										v100 = v83
									}
								}
							}
						}
						v103 = v49 + int32(1)
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v103 < v104 {
							v49 = v103
							v50 = v99
							v51 = v100
							continue
						} else {
							break
						}
						break
					}
					v112 = v99
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
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
		v33 = F_LWLockAcquire(m, v29+int32(512), int32(1))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if int32(0) < v37 {
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
				v49 = v3
				v50 = v3
				v51 = v43
				for {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v49<<(uint(int32(2))%32))))
					v60 = v51 + v57*int32(640)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+44))
					if v61 == int32(0) {
						v99 = v50
						v100 = v51
					} else {
						if l1 != 0 {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+60))
							if v64 != l1 {
								v99 = v50
								v100 = v51
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
								if l0 != 0 {
									if v66 == int32(0) {
										v99 = v50
										v100 = v51
									} else {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v66)) == int32(0) {
											v80 = base.B2i32(base.Ui32(l0) < base.Ui32(v66))
										} else {
											v80 = base.B2i32(int32(0) < v66-l0)
										}
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
										if v80 != 0 {
											v99 = v50
											v100 = v82
										} else {
											v83 = v82
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
											if v84 == int32(0) {
												v99 = v50
												v100 = v83
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
												v92 = v89 + v50<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
												*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
												v99 = v50 + int32(1)
												v100 = v83
											}
										}
									}
								} else {
									v83 = v51
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
									if v84 == int32(0) {
										v99 = v50
										v100 = v83
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
										v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
										v92 = v89 + v50<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
										*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
										v99 = v50 + int32(1)
										v100 = v83
									}
								}
							}
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
							if l0 != 0 {
								if v66 == int32(0) {
									v99 = v50
									v100 = v51
								} else {
									if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v66)) == int32(0) {
										v80 = base.B2i32(base.Ui32(l0) < base.Ui32(v66))
									} else {
										v80 = base.B2i32(int32(0) < v66-l0)
									}
									v82 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
									if v80 != 0 {
										v99 = v50
										v100 = v82
									} else {
										v83 = v82
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
										if v84 == int32(0) {
											v99 = v50
											v100 = v83
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
											v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
											v92 = v89 + v50<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
											*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
											v99 = v50 + int32(1)
											v100 = v83
										}
									}
								}
							} else {
								v83 = v51
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
								if v84 == int32(0) {
									v99 = v50
									v100 = v83
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
									v89 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
									v92 = v89 + v50<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v84
									*(*int32)(unsafe.Add(mBase, uint32(v92))) = v87
									v99 = v50 + int32(1)
									v100 = v83
								}
							}
						}
					}
					v103 = v49 + int32(1)
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v103 < v104 {
						v49 = v103
						v50 = v99
						v51 = v100
						continue
					} else {
						break
					}
					break
				}
				v112 = v99
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
func F_GetOldestMultiXactId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[0]))
	v13 = F_LWLockAcquire(m, v9+int32(1664), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(1)
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[1]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		if base.Ui32(v20) <= base.Ui32(v17) {
			v23 = v17
		} else {
			v23 = v20
		}
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[2]))
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[3]))
		v28 = v25 + v27
		if int32(0) < v28 {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[4]))
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[5]))
			v35 = v23
			v37 = int32(0)
			for {
				v43 = v37 << (uint(int32(2)) % 32)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v32+v43)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v34)))
				if v47-v35 < int32(0) {
					v51 = v47
				} else {
					v51 = v35
				}
				if v47 != 0 {
					v52 = v51
				} else {
					v52 = v35
				}
				if v45-v52 < int32(0) {
					v56 = v45
				} else {
					v56 = v52
				}
				if v45 != 0 {
					v57 = v56
				} else {
					v57 = v52
				}
				v59 = v37 + int32(1)
				if v59 != v28 {
					v35 = v57
					v37 = v59
					continue
				} else {
					break
				}
				break
			}
			v61 = v57
		} else {
			v61 = v23
		}
		v69 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[0]))
		F_LWLockRelease(m, v69+int32(1664))
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			return v61
		}
	}
}
func F_GetOldestTransactionIdConsideredRunning(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	F_ComputeXidHorizons(m, v5+int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
		m.G0 = v5 + int32(48)
		return v13
	}
}
func F_GetScanItems(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
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
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_tuplesort_reset(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v28 <= v27 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_tuplesort_performsort(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L48
	}
L4:
	;
	v33 = v27
	v36 = v28
	v42 = int32(0)
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v45 <= v42 {
		goto L3
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v48 = v33 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v48
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v33<<(uint(int32(2))%32)+v52)))
	if v54 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v62 = v54
	goto L11
L9:
	;
	v241 = v48
	v244 = v36
	goto L10
L10:
	;
	if v241 < v244 {
		v33 = v241
		v36 = v244
		v42 = v42 + int32(1)
		goto L5
	} else {
		goto L47
	}
L11:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v76 = F_ReadBufferExtended(m, v72, v73, v62, v73, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	v241 = v237
	v244 = v236
	goto L10
L13:
	;
	F_LockBuffer(m, v76, int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v76 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v98+v229)))
	F_UnlockReleaseBuffer(m, v76)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L45
	}
L16:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+12)))
	if base.Ui32(v99) < base.Ui32(int32(25)) {
		goto L15
	} else {
		goto L20
	}
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_GetScanItems[0]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+(v76^int32(-1))<<(uint(int32(2))%32))))
	v98 = v90
	goto L16
L18:
	;
	goto L19
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_GetScanItems[1]))
	v98 = v92 + v76<<(uint(int32(13))%32) + int32(-8192)
	goto L16
L20:
	;
	v107 = int32(base.Ui32(v99+int32(_a_F_GetScanItems_0))>>(uint(int32(2))%32)) & int32(_a_F_GetScanItems_1)
	if v107 == int32(0) {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v116 = int32(1)
	goto L22
L22:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v98+int32(20)+v116&int32(_a_F_GetScanItems_1)<<(uint(int32(2))%32))))
	v136 = v98 + v133&int32(_a_F_GetScanItems_2)
	v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v136)+6)))
	if int32(0) <= v137 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L15
L24:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	m.T0[v183].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L40
	}
L25:
	;
	v177 = F_nocache_index_getattr(m, v136, int32(1), v23)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L39
	}
L26:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v140 < int32(0) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v169 = int32(0)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+8)))
	if v170&int32(1) == v169 {
		v181 = v169
		goto L24
	} else {
		goto L38
	}
L29:
	;
	v145 = v140 + v136 + int32(8)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+26)))
	if v146 != int32(1) {
		v181 = v145
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+24)))
	switch v149 - int32(1) {
	case 0:
		goto L34
	case 1:
		goto L33
	default:
		goto L31
	case 3:
		goto L32
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v181 = v154
	goto L24
L33:
	;
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145))))
	v181 = v153
	goto L24
L34:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v145))))
	v181 = v152
	goto L24
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = base.I32_extend16_s(v149)
	F_errmsg_internal(m, int32(_a_F_GetScanItems_3), v18)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_GetScanItems_4), int32(70), int32(_a_F_GetScanItems_5))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	goto L25
L39:
	;
	v181 = v177
	goto L24
L40:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v189 = m.T0[v188].(func(*base.Module, int32, int32, int32, int32) int32)(m, v186, v187, v181, l1)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v189
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = v136
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)) = uint8(v194)
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v201 = v199 & int32(_a_F_GetScanItems_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v201)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v204)
	goto L42
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	F_tuplesort_puttupleslot(m, v206, v21)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v210 = v116 + int32(1)
	if base.Ui32(v210&int32(_a_F_GetScanItems_1)) <= base.Ui32(v107) {
		v116 = v210
		goto L22
	} else {
		goto L44
	}
L44:
	;
	goto L23
L45:
	;
	if v231 != int32(-1) {
		v62 = v231
		goto L11
	} else {
		goto L46
	}
L46:
	;
	goto L12
L47:
	;
	goto L6
L48:
	;
	m.G0 = v18 + int32(16)
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
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
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
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
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
	var v34 int32
	_ = v34
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v15 = v11 + base.I64_extend_i32_s(l1-base.I32_wrap_i64(v11))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui64(v15) < base.Ui64(v16) {
		v34 = int32(1)
		m.G0 = v9 + int32(48)
		return v34
	} else {
		v19 = int32(0)
		if base.Ui64(v11) <= base.Ui64(v15) {
			v34 = v19
			m.G0 = v9 + int32(48)
			return v34
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestIsRemovableXid[0]))
			if v22 != 0 {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_GlobalVisTestIsRemovableXid[1]))
				if v24 == v22 {
					v34 = v19
					m.G0 = v9 + int32(48)
					return v34
				} else {
					F_ComputeXidHorizons(m, v9+int32(8))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = base.B2i32(base.Ui64(v15) < base.Ui64(v32))
						m.G0 = v9 + int32(48)
						return v34
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
					v34 = base.B2i32(base.Ui64(v15) < base.Ui64(v32))
					m.G0 = v9 + int32(48)
					return v34
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
func F_gbtreekey_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_gbtreekey_out_0), int32(45), int32(_a_F_gbtreekey_out_1), int32(_a_F_gbtreekey_out_2), int32(_a_F_gbtreekey_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_gdb_date_dist(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_DirectFunctionCall2Coll(m, int32(2427), int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return float64(0)
	} else {
		v13 = v8 >> (uint(int32(31)) % 32)
		return base.F64_convert_i32_u(v8 ^ v13 - v13)
	}
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	v8 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v22 = l4 << (uint(int32(2)) % 32)
	v23 = F_palloc(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = F_palloc(m, v22)
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
	v31 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l4) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = v31
	v52 = v8
	goto L9
L7:
	;
	v100 = v31
	goto L8
L8:
	;
	v118 = v100
	v128 = v8
	goto L13
L9:
	;
	v55 = v43 << (uint(int32(2)) % 32)
	v57 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v55))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v55+v23))) = v57
	v62 = int32(4)
	v63 = v55 | v62
	*(*int32)(unsafe.Add(mBase, uint32(v25+v63))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v23+v63))) = v57
	v71 = v55 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v71))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v23+v71))) = v57
	v79 = v55 | int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v79))) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v79+v23))) = v57
	v87 = v43 + v62
	v89 = v52 + v62
	if v89 != l4&int32(2147483644) {
		v43 = v87
		v52 = v89
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
	v100 = v87
	goto L8
L13:
	;
	v130 = v118 << (uint(int32(2)) % 32)
	v132 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25+v130))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v130+v23))) = v132
	v137 = int32(1)
	v140 = v128 + v137
	if v140 != v30 {
		v118 = v118 + v137
		v128 = v140
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
	v161 = v20
	goto L18
L17:
	;
	v161 = v19
	goto L18
L18:
	;
	if int32(0) < v161 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v172 = int32(0)
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
	if v20 <= v172 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	if v19 <= v172 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v172<<(uint(int32(2))%32))))
	if v188 < int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+v188<<(uint(int32(2))%32)))) = v172
	goto L24
L27:
	;
	v210 = v172 + int32(1)
	if v210 != v161 {
		v172 = v210
		goto L22
	} else {
		goto L30
	}
L28:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v172<<(uint(int32(2))%32))))
	if v201 < int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v201<<(uint(int32(2))%32)))) = v172
	goto L27
L30:
	;
	goto L23
L31:
	;
	v240 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	F_pfree(m, v23)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L48
	}
L34:
	;
	v252 = v240 << (uint(int32(2)) % 32)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v25+v252)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252+v23)))
	if v254&v256 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if int32(0) <= v256 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v287 = v240 + int32(1)
	if v287 != l4 {
		v240 = v287
		goto L34
	} else {
		goto L47
	}
L39:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263+v256<<(uint(int32(2))%32))))
	v269 = v267
	goto L41
L40:
	;
	v269 = int32(0)
	goto L41
L41:
	;
	v270 = F_lappend(m, v260, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v270
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if int32(0) <= v254 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v254<<(uint(int32(2))%32))))
	v282 = v280
	goto L45
L44:
	;
	v282 = int32(0)
	goto L45
L45:
	;
	v283 = F_lappend(m, v273, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v283
	goto L38
L47:
	;
	goto L35
L48:
	;
	F_pfree(m, v25)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v329 float64
	_ = v329
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 float64
	_ = v528
	var v529 float64
	_ = v529
	var v533 float64
	_ = v533
	var v534 float64
	_ = v534
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
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
	v22 = int32(0)
	if l4&int32(-2) != int32(8) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = l4
	goto L6
L5:
	;
	v29 = v22
	goto L6
L6:
	;
	if base.B2i32(v20 == v22)&base.B2i32(v29 != int32(2)) != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if l6 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v20 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v44 = F_make_inner_pathkeys_for_merge(m, l0, v20, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L18
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v34 = v33
	goto L13
L12:
	;
	v34 = int32(0)
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v38 = v36
	goto L16
L15:
	;
	v38 = int32(0)
	goto L16
L16:
	;
	if v38 != v34 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	F_try_mergejoin_path(m, l0, l1, l3, l7, l8, v20, int32(0), v44, v29, l5, l9)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if l4 == int32(9) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l7)+64))
	if v44 == v50 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v44 == int32(0) {
		goto L1
	} else {
		goto L39
	}
L22:
	;
	v103 = int32(1)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v59 = int32(0)
	goto L26
L25:
	;
	v103 = v95
	goto L21
L26:
	;
	v63 = int32(0)
	if v44 == v63 {
		v73 = v63
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v95 = int32(0)
	goto L25
L28:
	;
	if v50 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v67 <= v59 {
		v73 = int32(0)
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v73 = v69 + v59<<(uint(int32(2))%32)
	goto L28
L31:
	;
	v79 = base.B2i32(v73 == int32(0))
	if v73 == int32(0) {
		v95 = v79
		goto L25
	} else {
		goto L36
	}
L32:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v59 < v74 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v103 = base.B2i32(v73 == int32(0))
	goto L21
L35:
	;
	goto L34
L36:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v82 == int32(0) {
		v95 = v79
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v59<<(uint(int32(2))%32)+v82)))
	if v89 == v91 {
		v59 = v59 + int32(1)
		goto L26
	} else {
		goto L38
	}
L38:
	;
	goto L27
L39:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if l6|base.B2i32(v106 < int32(2)) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v103 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v112 = F_list_copy(m, v44)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v106 <= int32(0) {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v116 = v112
	goto L40
L45:
	;
	v116 = v44
	goto L40
L46:
	;
	v118 = l7
	goto L48
L47:
	;
	v118 = int32(0)
	goto L48
L48:
	;
	v123 = v106
	v126 = v118
	v130 = v116
	v135 = v118
	goto L49
L49:
	;
	v136 = int32(0)
	if base.B2i32(v130 == v136)|base.B2i32(v123 <= v136) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L1
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v148 = int32(0)
	if v147 == v148 {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v146 = int32(0)
	goto L54
L53:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v123 < v143 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L51
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v123
	goto L57
L56:
	;
	goto L57
L57:
	;
	v146 = v130
	goto L54
L58:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v363 = int32(0)
	if v362 == v363 {
		goto L140
	} else {
		goto L141
	}
L59:
	;
	if v295 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L60:
	;
	v295 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if int32(0) < v161 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v171 = v148
	v174 = v148
	goto L66
L64:
	;
	v276 = v148
	goto L65
L65:
	;
	v295 = v276
	goto L59
L66:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v174<<(uint(int32(2))%32))))
	if l9 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v276 = v259
	goto L65
L68:
	;
	v266 = v174 + int32(1)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	if v266 < v267 {
		v171 = v259
		v174 = v266
		goto L66
	} else {
		goto L101
	}
L69:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+21)))
	if v182 != int32(1) {
		v259 = v171
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v171 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v185 = F_compare_path_costs(m, v171, v181, int32(1))
	mBase = m.M
	if v185 <= int32(0) {
		v259 = v171
		goto L68
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+64))
	if v146 == v188 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	if v246 != 0 {
		goto L95
	} else {
		goto L96
	}
L78:
	;
	v196 = int32(0)
	goto L79
L79:
	;
	v204 = int32(0)
	if v146 == v204 {
		v214 = v204
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v214 != 0 {
		v259 = v171
		goto L68
	} else {
		goto L94
	}
L81:
	;
	if v188 != 0 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v208 <= v196 {
		v214 = int32(0)
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v214 = v210 + v196<<(uint(int32(2))%32)
	goto L81
L84:
	;
	if v214 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v196 < v215 {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v214 == int32(0) {
		goto L77
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v259 = v171
	goto L68
L90:
	;
	goto L80
L91:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	if v221 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v221+v196<<(uint(int32(2))%32))))
	if v228 == v230 {
		v196 = v196 + int32(1)
		goto L79
	} else {
		goto L93
	}
L93:
	;
	v259 = v171
	goto L68
L94:
	;
	goto L77
L95:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v249 = v247
	goto L97
L96:
	;
	v249 = int32(0)
	goto L97
L97:
	;
	v250 = F_bms_is_subset(m, v249, v148)
	mBase = m.M
	if v250 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v251 = v181
	goto L100
L99:
	;
	v251 = v171
	goto L100
L100:
	;
	v259 = v251
	goto L68
L101:
	;
	goto L67
L102:
	;
	v360 = v126
	v361 = int32(0)
	goto L58
L103:
	;
	goto L104
L104:
	;
	if v126 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v295)+40))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v126)+40))
	if v305 != v306 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L107
L107:
	;
	if v123 < v106 {
		goto L133
	} else {
		goto L134
	}
L108:
	;
	if int32(0) <= v348 {
		v360 = v126
		v361 = int32(0)
		goto L58
	} else {
		goto L132
	}
L109:
	;
	if v305 < v306 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	goto L118
L112:
	;
	v311 = int32(-1)
	goto L114
L113:
	;
	v311 = int32(1)
	goto L114
L114:
	;
	v348 = v311
	goto L108
L115:
	;
	v348 = v342
	goto L108
L116:
	;
	v342 = int32(0)
	goto L115
L118:
	;
	goto L119
L119:
	;
	v327 = int32(-1)
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v295)+56))
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v126)+56))
	if base.F64_lt(v328, v329) != 0 {
		v342 = v327
		goto L115
	} else {
		goto L126
	}
L126:
	;
	if base.F64_gt(v328, v329) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v348 = int32(1)
	goto L108
L128:
	;
	goto L129
L129:
	;
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v295)+48))
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v126)+48))
	if base.F64_lt(v333, v334) != 0 {
		v342 = v327
		goto L115
	} else {
		goto L130
	}
L130:
	;
	if base.F64_gt(v333, v334) != 0 {
		v342 = int32(1)
		goto L115
	} else {
		goto L131
	}
L131:
	;
	goto L116
L132:
	;
	goto L107
L133:
	;
	v353 = F_trim_mergeclauses_for_inner_pathkeys(m, v20, v146)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L2
	} else {
		goto L136
	}
L134:
	;
	v355 = v20
	goto L135
L135:
	;
	v356 = int32(0)
	F_try_mergejoin_path(m, l0, l1, l3, v295, l8, v355, v356, v356, v29, l5, l9)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L137
	}
L136:
	;
	v355 = v353
	goto L135
L137:
	;
	v360 = v295
	v361 = v355
	goto L58
L138:
	;
	if v123 < int32(2) {
		goto L1
	} else {
		goto L219
	}
L139:
	;
	if v510 == int32(0) {
		v575 = v135
		goto L138
	} else {
		goto L182
	}
L140:
	;
	v510 = int32(0)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if int32(0) < v376 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v386 = v363
	v389 = v363
	goto L146
L144:
	;
	v491 = v363
	goto L145
L145:
	;
	v510 = v491
	goto L139
L146:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392+v389<<(uint(int32(2))%32))))
	if l9 != 0 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v491 = v474
	goto L145
L148:
	;
	v481 = v389 + int32(1)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	if v481 < v482 {
		v386 = v474
		v389 = v481
		goto L146
	} else {
		goto L181
	}
L149:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+21)))
	if v397 != int32(1) {
		v474 = v386
		goto L148
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	if v386 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v400 = F_compare_path_costs(m, v386, v396, v363)
	mBase = m.M
	if v400 <= int32(0) {
		v474 = v386
		goto L148
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v396)+64))
	if v146 == v403 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L155
L157:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v396)+16))
	if v461 != 0 {
		goto L175
	} else {
		goto L176
	}
L158:
	;
	v411 = int32(0)
	goto L159
L159:
	;
	v419 = int32(0)
	if v146 == v419 {
		v429 = v419
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v429 != 0 {
		v474 = v386
		goto L148
	} else {
		goto L174
	}
L161:
	;
	if v403 != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v423 <= v411 {
		v429 = int32(0)
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v429 = v425 + v411<<(uint(int32(2))%32)
	goto L161
L164:
	;
	if v429 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L165:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	if v411 < v430 {
		goto L164
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if v429 == int32(0) {
		goto L157
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	v474 = v386
	goto L148
L170:
	;
	goto L160
L171:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	if v436 == int32(0) {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v436+v411<<(uint(int32(2))%32))))
	if v443 == v445 {
		v411 = v411 + int32(1)
		goto L159
	} else {
		goto L173
	}
L173:
	;
	v474 = v386
	goto L148
L174:
	;
	goto L157
L175:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	v464 = v462
	goto L177
L176:
	;
	v464 = int32(0)
	goto L177
L177:
	;
	v465 = F_bms_is_subset(m, v464, v363)
	mBase = m.M
	if v465 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v466 = v396
	goto L180
L179:
	;
	v466 = v386
	goto L180
L180:
	;
	v474 = v466
	goto L148
L181:
	;
	goto L147
L182:
	;
	if v135 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v510)+40))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v135)+40))
	if v518 != v519 {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	goto L185
L185:
	;
	if v360 != v510 {
		goto L211
	} else {
		goto L212
	}
L186:
	;
	if int32(0) <= v561 {
		v575 = v135
		goto L138
	} else {
		goto L210
	}
L187:
	;
	if v518 < v519 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	goto L195
L190:
	;
	v524 = int32(-1)
	goto L192
L191:
	;
	v524 = int32(1)
	goto L192
L192:
	;
	v561 = v524
	goto L186
L193:
	;
	v561 = v555
	goto L186
L194:
	;
	v555 = int32(0)
	goto L193
L195:
	;
	v527 = int32(-1)
	v528 = *(*float64)(unsafe.Add(mBase, uint32(v510)+48))
	v529 = *(*float64)(unsafe.Add(mBase, uint32(v135)+48))
	if base.F64_lt(v528, v529) != 0 {
		v555 = v527
		goto L193
	} else {
		goto L198
	}
L198:
	;
	if base.F64_gt(v528, v529) != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v561 = int32(1)
	goto L186
L200:
	;
	goto L201
L201:
	;
	v533 = *(*float64)(unsafe.Add(mBase, uint32(v510)+56))
	v534 = *(*float64)(unsafe.Add(mBase, uint32(v135)+56))
	if base.F64_lt(v533, v534) != 0 {
		v555 = v527
		goto L193
	} else {
		goto L202
	}
L202:
	;
	if base.F64_gt(v533, v534) == int32(0) {
		goto L194
	} else {
		goto L203
	}
L203:
	;
	v555 = int32(1)
	goto L193
L210:
	;
	goto L185
L211:
	;
	if v361 != 0 {
		v568 = v361
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v575 = v510
	goto L138
L214:
	;
	v569 = int32(0)
	F_try_mergejoin_path(m, l0, l1, l3, v510, l8, v568, v569, v569, v29, l5, l9)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L2
	} else {
		goto L218
	}
L215:
	;
	if v106 <= v123 {
		v568 = v20
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v566 = F_trim_mergeclauses_for_inner_pathkeys(m, v20, v146)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L2
	} else {
		goto L217
	}
L217:
	;
	v568 = v566
	goto L214
L218:
	;
	goto L213
L219:
	;
	if l6 == int32(0) {
		v123 = v123 - int32(1)
		v126 = v360
		v130 = v146
		v135 = v575
		goto L49
	} else {
		goto L220
	}
L220:
	;
	goto L50
}
func F_generate_qualified_relation_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
			v19 = F_get_namespace_name_or_temp(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v65
						F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_0), v6+int32(-48))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_generate_qualified_relation_name_1), int32(_a_F_generate_qualified_relation_name_2), int32(_a_F_generate_qualified_relation_name_3))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v24 = v6 + int32(-16)
					F_initStringInfo(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = F_quote_identifier(m, v19)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v27
							F_appendStringInfo(m, v24, int32(_a_F_generate_qualified_relation_name_4), v6+int32(-32))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v37 = F_quote_identifier(m, v17+int32(4))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									F_appendStringInfoString(m, v24, v37)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
										F_ReleaseCatCache(m, v11)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int32(0)
										} else {
											m.G0 = v8 - int32(-64)
											return v41
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_5), v8)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_generate_qualified_relation_name_1), int32(_a_F_generate_qualified_relation_name_6), int32(_a_F_generate_qualified_relation_name_3))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
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
func F_geo_distance(m *base.Module, l0 int32) int32 {
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
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v322 = F_Float8GetDatum(m, base.F64_mul(v319, float64(7917.495432)))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		return int32(0)
	} else {
		return v322
	}
}
func F_geqo_eval(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0]))
	v15 = F_AllocSetContextCreateInternal(m, v10, int32(_a_F_geqo_eval_0), v4, int32(_a_F_geqo_eval_1), int32(_a_F_geqo_eval_2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return float64(0)
	} else {
		v19 = int32(_a_F_geqo_eval_3)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0])) = v15
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if v23 != 0 {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v25 = v24
		} else {
			v25 = v4
		}
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
		v29 = F_gimme_tree(m, l0, l1, l2)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return float64(0)
		} else {
			if v29 != 0 {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
				v32 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
				v34 = v32
			} else {
				v34 = float64(1.7976931348623157e+308)
			}
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v36 = int32(0)
			if base.B2i32(v35 == v36)|base.B2i32(v25 <= v36) != 0 {
				v46 = int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				if v25 < v43 {
					*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v25
				} else {
				}
				v46 = v35
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
			*(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0])) = v20
			F_MemoryContextDelete(m, v15)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return float64(0)
			} else {
				return v34
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
	v213 = F_palloc(m, int32(20))
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
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v28 + v327
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v330 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v318 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+4)) = uint8(v318)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(3)
	goto L60
L63:
	;
	v314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+4)) = uint8(v314)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(3)
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v28 + (v250+int32(3))&int32(-4)
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v28 + v250
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
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v271 & int32(268435455)
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
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v271&int32(268435455) - v297
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
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v371 + (v250 - v327)
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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
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
	v173 = m.ExcPending
	if v173 != 0 {
		goto L33
	} else {
		goto L55
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L33
	} else {
		goto L52
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return v151
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
		v151 = v6
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
		v151 = v6
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
		v151 = v77
		goto L3
	} else {
		goto L45
	}
L33:
	;
	return int32(0)
L34:
	;
	v87 = F_pg_detoast_datum_copy(m, v83)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89
	v92 = F_SearchSysCache1(m, int32(82), v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	if v92 == int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	v98 = v96 + v97
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+76)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+78)))
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98)+128)))
	F_deconstruct_array(m, v87, v99, v100, v101, l0+int32(12), int32(0), l0+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+78)))
	if v109 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	F_ReleaseCatCache(m, v92)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L33
	} else {
		goto L44
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v87
	goto L39
L41:
	;
	goto L42
L42:
	;
	F_pfree(m, v87)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
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
	v127 = F_SysCacheGetAttrNotNull(m, int32(65), l1, v68+int32(22))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v129 = F_pg_detoast_datum_copy(m, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L33
	} else {
		goto L47
	}
L47:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v131 != int32(1) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	if v134 <= int32(0) {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	if v137 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	if v138 != int32(700) {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v129 + int32(24)
	v151 = v77
	goto L3
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v89
	F_errmsg_internal(m, int32(_a_F_get_attstatsslot_0), v13)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L33
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_get_attstatsslot_1), int32(3421), int32(_a_F_get_attstatsslot_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
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
	v177 = m.ExcPending
	if v177 != 0 {
		goto L33
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_get_attstatsslot_1), int32(3466), int32(_a_F_get_attstatsslot_2))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_SearchSysCache2(m, int32(7), l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
				F_errmsg_internal(m, int32(_a_F_get_atttypetypmodcoll_0), v10)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_atttypetypmodcoll_1), int32(1046), int32(_a_F_get_atttypetypmodcoll_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
			v33 = v31 + v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v38
			F_ReleaseCatCache(m, v13)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v489 int32
	_ = v489
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v552 int32
	_ = v552
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v593 int32
	_ = v593
	var v602 int32
	_ = v602
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v969 int32
	_ = v969
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1053 int32
	_ = v1053
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 float64
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1119 float64
	_ = v1119
	var v1128 float64
	_ = v1128
	var v1132 float64
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1138 float64
	_ = v1138
	var v1140 float64
	_ = v1140
	var v1143 float64
	_ = v1143
	var v1146 float64
	_ = v1146
	var v1155 int32
	_ = v1155
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_memoize_path[0])))
	if v24 != int32(1) {
		v1155 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(16)
	return v1155
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	if base.F64_lt(v28, float64(2)) != 0 {
		v1155 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v31 != int32(1) {
		v393 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v395 != 0 {
		goto L95
	} else {
		goto L96
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v35 = int32(0)
	if v34 == v35 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v80 == int32(2) {
		v393 = v8
		goto L4
	} else {
		goto L22
	}
L7:
	;
	v80 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v43 = int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v44 <= v43 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v43
	goto L12
L11:
	;
	v47 = v44
	goto L12
L12:
	;
	v51 = int32(0)
	v53 = v35
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(8)+v51<<(uint(int32(2))%32))))
	if v60 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v80 = v72
	goto L6
L15:
	;
	goto L14
L16:
	;
	v61 = int32(2)
	if v53 != 0 {
		v72 = v61
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v67 = v53
	goto L18
L18:
	;
	v69 = v51 + int32(1)
	if v69 != v47 {
		v51 = v69
		v53 = v67
		goto L13
	} else {
		goto L21
	}
L19:
	;
	v62 = int32(1)
	if base.Ui32(v62) < base.Ui32(base.I32_popcnt(v60)) {
		v72 = v61
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v67 = v62
	goto L18
L21:
	;
	v72 = v67
	goto L15
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v83 == int32(0) {
		v393 = v8
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v86 <= int32(0) {
		v393 = v8
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v98 = v8
	v105 = v8
	goto L25
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v98<<(uint(int32(2))%32))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	if v112 == int32(0) {
		v371 = v105
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v393 = v371
	goto L4
L27:
	;
	v374 = v98 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v374 < v375 {
		v98 = v374
		v105 = v371
		goto L25
	} else {
		goto L93
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v116 = int32(0)
	if base.B2i32(v115 == v116)|base.B2i32(v34 == v116) != 0 {
		v162 = base.B2i32(v115|v34 == v116)
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v162 == int32(0) {
		v371 = v105
		goto L27
	} else {
		goto L40
	}
L30:
	;
	goto L29
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v130 != v131 {
		v162 = int32(0)
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v133 = int32(1)
	if v130 <= v133 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v136 = v133
	goto L35
L34:
	;
	v136 = v130
	goto L35
L35:
	;
	v137 = int32(8)
	v142 = int32(0)
	goto L36
L36:
	;
	v150 = v142 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v115+v137+v150)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v34+v137+v150)))
	v155 = base.B2i32(v152 == v154)
	if v152 != v154 {
		v162 = v155
		goto L30
	} else {
		goto L38
	}
L37:
	;
	v162 = v155
	goto L30
L38:
	;
	v158 = v142 + int32(1)
	if v158 != v136 {
		v142 = v158
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v171 = F_pull_varnos(m, l0, v170)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return int32(0)
L42:
	;
	v175 = int32(0)
	if base.B2i32(v171 == v175)|base.B2i32(v34 == v175) != 0 {
		v220 = v175
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v220 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L44:
	;
	goto L43
L45:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v185 < v186 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v188 = v185
	goto L48
L47:
	;
	v188 = v186
	goto L48
L48:
	;
	if v188 <= int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v191 = int32(1)
	goto L51
L50:
	;
	v191 = v188
	goto L51
L51:
	;
	v192 = int32(8)
	v197 = int32(0)
	goto L52
L52:
	;
	v204 = v197 << (uint(int32(2)) % 32)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v34+v192+v204)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v171+v192+v204)))
	v209 = v206 & v208
	v211 = base.B2i32(v209 != int32(0))
	if v209 != 0 {
		v220 = v211
		goto L44
	} else {
		goto L54
	}
L53:
	;
	v220 = v211
	goto L44
L54:
	;
	v213 = v197 + int32(1)
	if v213 != v191 {
		v197 = v213
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v225 = F_lappend(m, v105, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L41
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v228 = F_pull_vars_of_level(m, v222, int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L41
	} else {
		goto L61
	}
L59:
	;
	v371 = v225
	goto L27
L60:
	;
	F_list_free(m, v228)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L41
	} else {
		goto L92
	}
L61:
	;
	if v228 == int32(0) {
		v351 = v105
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v232 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v233 <= v232 {
		v351 = v105
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v243 = v232
	v252 = v105
	goto L64
L64:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254+v243<<(uint(int32(2))%32))))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v259 != int32(6) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v351 = v330
	goto L60
L66:
	;
	v332 = v243 + int32(1)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v332 < v333 {
		v243 = v332
		v252 = v330
		goto L64
	} else {
		goto L91
	}
L67:
	;
	v328 = F_lappend(m, v252, v258)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L41
	} else {
		goto L90
	}
L68:
	;
	if v259 != int32(319) {
		v330 = v252
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v324 = F_bms_is_member(m, v322, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L41
	} else {
		goto L88
	}
L71:
	;
	v264 = F_find_placeholder_info(m, l0, v258)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L41
	} else {
		goto L72
	}
L72:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v268 = int32(0)
	if v266 == v268 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v321 != 0 {
		goto L67
	} else {
		goto L87
	}
L74:
	;
	v321 = int32(1)
	goto L73
L75:
	;
	goto L76
L76:
	;
	if v267 == int32(0) {
		v314 = v268
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v321 = v314
	goto L73
L78:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v278 < v277 {
		v314 = v268
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v280 = int32(1)
	if v277 <= v280 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v283 = v280
	goto L82
L81:
	;
	v283 = v277
	goto L82
L82:
	;
	v284 = int32(8)
	v289 = int32(0)
	goto L83
L83:
	;
	v296 = v289 << (uint(int32(2)) % 32)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v266+v284+v296)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v267+v284+v296)))
	v303 = v298 & (v300 ^ int32(-1))
	v305 = base.B2i32(v303 == int32(0))
	if v303 != 0 {
		v314 = v305
		goto L77
	} else {
		goto L85
	}
L84:
	;
	v314 = v305
	goto L77
L85:
	;
	v307 = v289 + int32(1)
	if v307 != v283 {
		v289 = v307
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v330 = v252
	goto L66
L88:
	;
	if v324 == int32(0) {
		v330 = v252
		goto L66
	} else {
		goto L89
	}
L89:
	;
	goto L67
L90:
	;
	v330 = v328
	goto L66
L91:
	;
	goto L65
L92:
	;
	v371 = v351
	goto L27
L93:
	;
	goto L26
L94:
	;
	v400 = int32(0)
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+8)))
	if base.B2i32(v401&int32(1) == v400)&base.B2i32(l5&int32(-2) == int32(4)) != 0 {
		v1155 = v400
		goto L1
	} else {
		goto L100
	}
L95:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	if v396 != 0 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v397|v393 != 0 {
		goto L94
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v1155 = int32(0)
	goto L1
L100:
	;
	if v401&int32(1) == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v474 = F_contain_volatile_functions(m, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L41
	} else {
		goto L113
	}
L102:
	;
	if v395 == int32(0) {
		v1155 = v400
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v417 == int32(0) {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v420 = int32(0)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	if v421 <= v420 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v395)+20))
	v433 = v420
	goto L106
L106:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v443+v433<<(uint(int32(2))%32))))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+56))
	v449 = F_bms_is_member(m, v448, v424)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L41
	} else {
		goto L108
	}
L107:
	;
	v1155 = v400
	goto L1
L108:
	;
	if v449 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v452 = v433 + int32(1)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	if v452 < v453 {
		v433 = v452
		goto L106
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	goto L107
L112:
	;
	goto L101
L113:
	;
	if v474 != 0 {
		v1155 = v400
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v476 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v532 = int32(0)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v534 == v532 {
		v602 = v532
		goto L125
	} else {
		goto L126
	}
L116:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v479 <= int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v489 = v400
	goto L118
L118:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v489<<(uint(int32(2))%32))))
	v505 = F_contain_volatile_functions(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L41
	} else {
		goto L120
	}
L119:
	;
	v1155 = int32(0)
	goto L1
L120:
	;
	if v505 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v510 = v489 + int32(1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v510 < v511 {
		v489 = v510
		goto L118
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	goto L119
L124:
	;
	goto L115
L125:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	if v612 != 0 {
		goto L138
	} else {
		goto L139
	}
L126:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v534)+16))
	if v537 == int32(0) {
		v602 = v534
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v540 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	if v541 <= v540 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v602 = v593
	goto L125
L129:
	;
	v552 = v540
	goto L130
L130:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v562+v552<<(uint(int32(2))%32))))
	v567 = F_contain_volatile_functions(m, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L41
	} else {
		goto L132
	}
L131:
	;
	v1155 = v532
	goto L1
L132:
	;
	if v567 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v572 = v552 + int32(1)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	if v572 < v573 {
		v552 = v572
		goto L130
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	goto L131
L136:
	;
	goto L128
L137:
	;
	if v1081 == int32(0) {
		v1155 = v532
		goto L1
	} else {
		goto L247
	}
L138:
	;
	v613 = v612
	goto L140
L139:
	;
	v613 = l2
	goto L140
L140:
	;
	v614 = int32(0)
	v616 = v21 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v616))) = v614
	v620 = v21 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v614
	v624 = v21 + int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v614)
	if v602 == v614 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v1081 = v1053
	goto L137
L142:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	F_list_free(m, v1038)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L41
	} else {
		goto L245
	}
L143:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v955 = F_list_concat(m, v393, v954)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L41
	} else {
		goto L227
	}
L144:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v602)+16))
	if v629 == int32(0) {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v629)+4))
	if v632 <= int32(0) {
		goto L143
	} else {
		goto L146
	}
L146:
	;
	v640 = v614
	goto L147
L147:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v629)+12))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v653+v640<<(uint(int32(2))%32))))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	if v659 != int32(17) {
		goto L142
	} else {
		goto L149
	}
L148:
	;
	goto L143
L149:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v658)+28))
	if v662 == int32(0) {
		goto L142
	} else {
		goto L150
	}
L150:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v662)+4))
	if v665 != int32(2) {
		goto L142
	} else {
		goto L151
	}
L151:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v657)+44))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v613)+8))
	v671 = int32(0)
	if v669 == v671 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v907+v657)))
	if v909 == int32(0) {
		goto L142
	} else {
		goto L216
	}
L153:
	;
	v901 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v657)+120)) = uint8(v901)
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v658)+28))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)+12))
	v906 = v904
	v907 = int32(160)
	goto L152
L154:
	;
	if v724 != 0 {
		goto L168
	} else {
		goto L169
	}
L155:
	;
	v724 = int32(1)
	goto L154
L156:
	;
	goto L157
L157:
	;
	if v670 == int32(0) {
		v717 = v671
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v724 = v717
	goto L154
L159:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v670)+4))
	if v681 < v680 {
		v717 = v671
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v683 = int32(1)
	if v680 <= v683 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v686 = v683
	goto L163
L162:
	;
	v686 = v680
	goto L163
L163:
	;
	v687 = int32(8)
	v692 = int32(0)
	goto L164
L164:
	;
	v699 = v692 << (uint(int32(2)) % 32)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v669+v687+v699)))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v670+v687+v699)))
	v706 = v701 & (v703 ^ int32(-1))
	v708 = base.B2i32(v706 == int32(0))
	if v706 != 0 {
		v717 = v708
		goto L158
	} else {
		goto L166
	}
L165:
	;
	v717 = v708
	goto L158
L166:
	;
	v710 = v692 + int32(1)
	if v710 != v686 {
		v692 = v710
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v657)+48))
	v726 = int32(0)
	if v725 == v726 {
		goto L172
	} else {
		goto L173
	}
L169:
	;
	goto L170
L170:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v657)+44))
	v781 = int32(0)
	if v780 == v781 {
		goto L187
	} else {
		goto L188
	}
L171:
	;
	if v779 != 0 {
		goto L153
	} else {
		goto L185
	}
L172:
	;
	v779 = int32(1)
	goto L171
L173:
	;
	goto L174
L174:
	;
	if v668 == int32(0) {
		v772 = v726
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v779 = v772
	goto L171
L176:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v725)+4))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	if v736 < v735 {
		v772 = v726
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v738 = int32(1)
	if v735 <= v738 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v741 = v738
	goto L180
L179:
	;
	v741 = v735
	goto L180
L180:
	;
	v742 = int32(8)
	v747 = int32(0)
	goto L181
L181:
	;
	v754 = v747 << (uint(int32(2)) % 32)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v725+v742+v754)))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v668+v742+v754)))
	v761 = v756 & (v758 ^ int32(-1))
	v763 = base.B2i32(v761 == int32(0))
	if v761 != 0 {
		v772 = v763
		goto L175
	} else {
		goto L183
	}
L182:
	;
	v772 = v763
	goto L175
L183:
	;
	v765 = v747 + int32(1)
	if v765 != v741 {
		v747 = v765
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	goto L170
L186:
	;
	if v834 == int32(0) {
		goto L142
	} else {
		goto L200
	}
L187:
	;
	v834 = int32(1)
	goto L186
L188:
	;
	goto L189
L189:
	;
	if v668 == int32(0) {
		v827 = v781
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v834 = v827
	goto L186
L191:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	if v791 < v790 {
		v827 = v781
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v793 = int32(1)
	if v790 <= v793 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v796 = v793
	goto L195
L194:
	;
	v796 = v790
	goto L195
L195:
	;
	v797 = int32(8)
	v802 = int32(0)
	goto L196
L196:
	;
	v809 = v802 << (uint(int32(2)) % 32)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v780+v797+v809)))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v668+v797+v809)))
	v816 = v811 & (v813 ^ int32(-1))
	v818 = base.B2i32(v816 == int32(0))
	if v816 != 0 {
		v827 = v818
		goto L190
	} else {
		goto L198
	}
L197:
	;
	v827 = v818
	goto L190
L198:
	;
	v820 = v802 + int32(1)
	if v820 != v796 {
		v802 = v820
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v657)+48))
	v838 = int32(0)
	if v837 == v838 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	if v891 == int32(0) {
		goto L142
	} else {
		goto L215
	}
L202:
	;
	v891 = int32(1)
	goto L201
L203:
	;
	goto L204
L204:
	;
	if v670 == int32(0) {
		v884 = v838
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v891 = v884
	goto L201
L206:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v837)+4))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v670)+4))
	if v848 < v847 {
		v884 = v838
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v850 = int32(1)
	if v847 <= v850 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v853 = v850
	goto L210
L209:
	;
	v853 = v847
	goto L210
L210:
	;
	v854 = int32(8)
	v859 = int32(0)
	goto L211
L211:
	;
	v866 = v859 << (uint(int32(2)) % 32)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v837+v854+v866)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v670+v854+v866)))
	v873 = v868 & (v870 ^ int32(-1))
	v875 = base.B2i32(v873 == int32(0))
	if v873 != 0 {
		v884 = v875
		goto L205
	} else {
		goto L213
	}
L212:
	;
	v884 = v875
	goto L205
L213:
	;
	v877 = v859 + int32(1)
	if v877 != v853 {
		v859 = v877
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v894 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v657)+120)) = uint8(v894)
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v658)+28))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+12))
	v906 = v897 + int32(4)
	v907 = int32(164)
	goto L152
L216:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v906)))
	v914 = F_list_member(m, v912, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L41
	} else {
		goto L217
	}
L217:
	;
	if v914 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v919 = F_lappend_oid(m, v918, v909)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L41
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v657)+124))
	if v926 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v919
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v923 = F_lappend(m, v922, v913)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L41
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616))) = v923
	goto L220
L223:
	;
	v929 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v929)
	goto L225
L224:
	;
	goto L225
L225:
	;
	v932 = v640 + int32(1)
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v629)+4))
	if v932 < v933 {
		v640 = v932
		goto L147
	} else {
		goto L226
	}
L226:
	;
	goto L148
L227:
	;
	if v955 == int32(0) {
		v1081 = int32(1)
		goto L137
	} else {
		goto L228
	}
L228:
	;
	v959 = int32(1)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	if v960 <= int32(0) {
		v1053 = v959
		goto L141
	} else {
		goto L229
	}
L229:
	;
	v969 = int32(0)
	goto L230
L230:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v955)+12))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v982+v969<<(uint(int32(2))%32))))
	v987 = F_contain_volatile_functions(m, v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L41
	} else {
		goto L232
	}
L231:
	;
	v1053 = v959
	goto L141
L232:
	;
	if v987 != 0 {
		goto L142
	} else {
		goto L233
	}
L233:
	;
	v989 = F_exprType(m, v986)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L41
	} else {
		goto L234
	}
L234:
	;
	v992 = F_lookup_type_cache(m, v989, int32(17))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L41
	} else {
		goto L235
	}
L235:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v992)+68))
	if v994 == int32(0) {
		goto L142
	} else {
		goto L236
	}
L236:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v992)+52))
	if v997 == int32(0) {
		goto L142
	} else {
		goto L237
	}
L237:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v1001 = F_list_member(m, v1000, v986)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L41
	} else {
		goto L238
	}
L238:
	;
	if v1001 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v992)+52))
	v1007 = F_lappend_oid(m, v1005, v1006)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L41
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1014 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v1014)
	v1017 = v969 + v1014
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	if v1017 < v1018 {
		v969 = v1017
		goto L230
	} else {
		goto L244
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v1007
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v1011 = F_lappend(m, v1010, v986)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L41
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616))) = v1011
	goto L241
L244:
	;
	goto L231
L245:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	F_list_free(m, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L41
	} else {
		goto L246
	}
L246:
	;
	v1053 = int32(0)
	goto L141
L247:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+8)))
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)))
	v1088 = *(*float64)(unsafe.Add(mBase, uint32(l4)+32))
	v1090 = F_palloc0(m, int32(104))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L41
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v1090))) = int64(1550483194150)
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+12)) = v1095
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v1098 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1090)+20)) = uint8(v1098)
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+16)) = v1097
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v1101 == int32(1) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	v1106 = v1104
	goto L251
L250:
	;
	v1106 = int32(0)
	goto L251
L251:
	;
	v1108 = v1106 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1090)+21)) = uint8(v1108)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+24)) = v1110
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v1090)+85)) = uint8(v1087)
	*(*uint8)(unsafe.Add(mBase, uint32(v1090)+84)) = uint8(v1086)
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+80)) = v1084
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+76)) = v1085
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+72)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+64)) = v1112
	v1119 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1088)&int64(9223372036854775807)))|base.F64_gt(v1088, v1119) != 0 {
		v1132 = v1119
		goto L253
	} else {
		goto L254
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+96)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(v1090)+88)) = v1132
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+40)) = v1136
	v1138 = *(*float64)(unsafe.Add(mBase, uint32(l3)+48))
	v1140 = *(*float64)(unsafe.Add(mBase, _c_F_get_memoize_path[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v1090)+48)) = base.F64_add(v1138, v1140)
	v1143 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v1090)+56)) = base.F64_add(v1140, v1143)
	v1146 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v1090)+32)) = v1146
	v1155 = v1090
	goto L1
L253:
	;
	goto L252
L254:
	;
	v1128 = float64(1)
	if base.F64_le(v1088, v1128) != 0 {
		v1132 = v1128
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v1132 = base.F64_nearest(v1088)
	goto L253
}
func F_get_mergejoin_opfamilies(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	v2 = int32(0)
	v12 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v2, v2)
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if int32(0) < v16 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = v2
	v25 = v2
	goto L6
L4:
	;
	v79 = v2
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(48)+v24<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+56))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v35 <= int32(2741) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v79 = v70
	goto L5
L8:
	;
	v72 = v24 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v72 < v73 {
		v24 = v72
		v25 = v70
		goto L6
	} else {
		goto L23
	}
L9:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+16)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v62 = F_IndexAmTranslateStrategy(m, v60, v59, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v51 = F_GetIndexAmRoutineByAmId(m, v35, int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	switch v35 - int32(403) {
	case 0:
		v59 = v35
		goto L9
	case 1:
		goto L10
	case 2:
		v70 = v25
		goto L8
	default:
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.B2i32(v35 == int32(2742))|base.B2i32(v35 == int32(3580))|base.B2i32(v35 == int32(4000)) != 0 {
		v70 = v25
		goto L8
	} else {
		goto L16
	}
L14:
	;
	if v35 != int32(783) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v70 = v25
	goto L8
L16:
	;
	goto L10
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+10)))
	F_pfree(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v53 != int32(1) {
		v70 = v25
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v59 = v58
	goto L9
L20:
	;
	if v62 != int32(3) {
		v70 = v25
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v67 = F_lappend_oid(m, v25, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v70 = v67
	goto L8
L23:
	;
	goto L7
L24:
	;
	return v79
}
func F_get_name_for_var_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	v14 = m.G0
	v16 = v14 - int32(256)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != int32(8) {
		goto L15
	} else {
		goto L16
	}
L1:
	;
	m.G0 = v16 + int32(256)
	return v860
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L22
	} else {
		goto L249
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L22
	} else {
		goto L246
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L22
	} else {
		goto L243
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L22
	} else {
		goto L240
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L22
	} else {
		goto L237
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L22
	} else {
		goto L234
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L22
	} else {
		goto L231
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L22
	} else {
		goto L228
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L22
	} else {
		goto L225
	}
L11:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v16)+248))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	v695 = v16 + int32(168)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v16)+252))
	base.MemoryCopy(m, v695, v696, int32(80))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v696)+44))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+12))
	v706 = F_list_copy_tail(m, v699, (v692-v700)>>(uint(int32(2))%32)+int32(1))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L22
	} else {
		goto L221
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v68 = v67 + l2
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v69 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v56 = F_get_expr_result_tupdesc(m, l0, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L22
	} else {
		goto L27
	}
L14:
	;
	if v47 != int32(6) {
		goto L13
	} else {
		goto L25
	}
L15:
	;
	if v18 != int32(36) {
		v47 = v18
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v42 = F_find_param_referent(m, l0, l3, v16+int32(252), v16+int32(248))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if l1 <= int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v25 == int32(0) {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 < l1 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+l1<<(uint(int32(2))%32)-int32(4))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v860 = v37
	goto L1
L22:
	;
	return int32(0)
L23:
	;
	if v42 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = v46
	goto L14
L25:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v50 == int32(2249) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v860 = v56 + v58<<(uint(int32(4))%32) + l1*int32(100) - int32(76)
	goto L1
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v72 = v70
	goto L30
L29:
	;
	v72 = int32(0)
	goto L30
L30:
	;
	if v72 <= v68 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v68<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+252)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v80 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88+l0))))
	if v87 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = v85
	v88 = int32(8)
	goto L32
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
	if v83 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v87 = v80
	v88 = int32(40)
	goto L32
L36:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	switch v294 - int32(1) {
	case 0:
		goto L106
	case 1:
		goto L105
	default:
		v667 = l0
		goto L103
	case 5:
		goto L104
	}
L37:
	;
	switch v87 + int32(3) {
	case 0:
		goto L44
	case 1:
		goto L46
	case 2:
		goto L45
	default:
		goto L43
	}
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v93 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v96 < v87 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v87<<(uint(int32(2))%32)-int32(4))))
	if v90 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v105 = F_get_rte_attribute_name(m, v104, l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v860 = v105
	goto L1
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L22
	} else {
		goto L100
	}
L44:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v78)+64))
	if v233 == int32(0) {
		goto L43
	} else {
		goto L84
	}
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	if v168 == int32(0) {
		goto L43
	} else {
		goto L65
	}
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+56))
	if v110 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v110 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v150 == int32(0) {
		goto L9
	} else {
		goto L61
	}
L49:
	;
	goto L48
L50:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v116 <= int32(0) {
		v150 = int32(0)
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v150 = int32(0)
	goto L49
L53:
	;
	v119 = int32(0)
	if v119 < v116 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v122 = v116
	goto L56
L55:
	;
	v122 = v119
	goto L56
L56:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v127 = int32(0)
	goto L57
L57:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v123+v127<<(uint(int32(2))%32))))
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+8)))
	if v136 == v90&int32(_a_F_get_name_for_var_field_0) {
		v150 = v135
		goto L49
	} else {
		goto L59
	}
L58:
	;
	goto L52
L59:
	;
	v139 = v127 + int32(1)
	if v139 != v122 {
		v127 = v139
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
	v156 = v16 + int32(168)
	F_push_child_plan(m, v78, v154, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v160 = F_get_name_for_var_field(m, v159, l1, l2, l3)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L22
	} else {
		goto L63
	}
L63:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v163 = F_list_delete_first(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L22
	} else {
		goto L64
	}
L64:
	;
	base.MemoryCopy(m, v78, v156, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v163
	v860 = v160
	goto L1
L65:
	;
	if v168 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	if v208 == int32(0) {
		goto L8
	} else {
		goto L79
	}
L67:
	;
	goto L66
L68:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v174 <= int32(0) {
		v208 = int32(0)
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v208 = int32(0)
	goto L67
L71:
	;
	v177 = int32(0)
	if v177 < v174 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v180 = v174
	goto L74
L73:
	;
	v180 = v177
	goto L74
L74:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v185 = int32(0)
	goto L75
L75:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v181+v185<<(uint(int32(2))%32))))
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+8)))
	if v194 == v90&int32(_a_F_get_name_for_var_field_0) {
		v208 = v193
		goto L67
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v197 = v185 + int32(1)
	if v197 != v180 {
		v185 = v197
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	v214 = v16 + int32(168)
	base.MemoryCopy(m, v214, v78, int32(80))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v219 = F_lcons(m, v217, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L22
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v219
	F_set_deparse_plan(m, v78, v212)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L22
	} else {
		goto L81
	}
L81:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v225 = F_get_name_for_var_field(m, v224, l1, l2, l3)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L22
	} else {
		goto L82
	}
L82:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v228 = F_list_delete_first(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L22
	} else {
		goto L83
	}
L83:
	;
	base.MemoryCopy(m, v78, v214, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v228
	v860 = v225
	goto L1
L84:
	;
	if v233 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	if v273 == int32(0) {
		goto L7
	} else {
		goto L98
	}
L86:
	;
	goto L85
L87:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v239 <= int32(0) {
		v273 = int32(0)
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v273 = int32(0)
	goto L86
L90:
	;
	v242 = int32(0)
	if v242 < v239 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v245 = v239
	goto L93
L92:
	;
	v245 = v242
	goto L93
L93:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v250 = int32(0)
	goto L94
L94:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246+v250<<(uint(int32(2))%32))))
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+8)))
	if v259 == v90&int32(_a_F_get_name_for_var_field_0) {
		v273 = v258
		goto L86
	} else {
		goto L96
	}
L95:
	;
	goto L89
L96:
	;
	v262 = v250 + int32(1)
	if v262 != v245 {
		v250 = v262
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v278 = F_get_name_for_var_field(m, v277, l1, l2, l3)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	v860 = v278
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v87
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_1), v16)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L22
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_3), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L22
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v681 = F_get_expr_result_tupdesc(m, v667, int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L22
	} else {
		goto L220
	}
L104:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v104)+88))
	v447 = v446 + v68
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.Ui32(v448) <= base.Ui32(v447) {
		goto L155
	} else {
		goto L156
	}
L105:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v104)+52))
	if v431 == int32(0) {
		goto L4
	} else {
		goto L152
	}
L106:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v104)+36))
	if v297 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+76))
	if v298 != 0 {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	goto L109
L109:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	if v362 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L110:
	;
	if v336 == int32(0) {
		goto L6
	} else {
		goto L123
	}
L111:
	;
	goto L110
L112:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v302 <= int32(0) {
		v336 = int32(0)
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v336 = int32(0)
	goto L111
L115:
	;
	v305 = int32(0)
	if v305 < v302 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v308 = v302
	goto L118
L117:
	;
	v308 = v305
	goto L118
L118:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v313 = int32(0)
	goto L119
L119:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v309+v313<<(uint(int32(2))%32))))
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v321)+8)))
	if v322 == v90&int32(_a_F_get_name_for_var_field_0) {
		v336 = v321
		goto L111
	} else {
		goto L121
	}
L120:
	;
	goto L114
L121:
	;
	v325 = v313 + int32(1)
	if v325 != v308 {
		v313 = v325
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+26)))
	if v340 == int32(1) {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v344 != int32(6) {
		v667 = v343
		goto L103
	} else {
		goto L125
	}
L125:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v348 = F_list_copy_tail(m, v347, v68)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L22
	} else {
		goto L126
	}
L126:
	;
	v351 = v16 + int32(168)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v104)+36))
	F_set_deparse_for_query(m, v351, v352, v348)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L22
	} else {
		goto L127
	}
L127:
	;
	v355 = F_lcons(m, v351, v348)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v355
	v359 = F_get_name_for_var_field(m, v343, l1, int32(0), l3)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v347
	v860 = v359
	goto L1
L130:
	;
	v366 = F_palloc(m, int32(32))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L22
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	if v375 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l1
	v373 = F_pg_snprintf(m, v366, int32(32), int32(_a_F_get_name_for_var_field_5), v16+int32(112))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L22
	} else {
		goto L134
	}
L134:
	;
	v860 = v366
	goto L1
L135:
	;
	if v413 == int32(0) {
		goto L5
	} else {
		goto L148
	}
L136:
	;
	goto L135
L137:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v379 <= int32(0) {
		v413 = int32(0)
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v413 = int32(0)
	goto L136
L140:
	;
	v382 = int32(0)
	if v382 < v379 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v385 = v379
	goto L143
L142:
	;
	v385 = v382
	goto L143
L143:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	v390 = int32(0)
	goto L144
L144:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v386+v390<<(uint(int32(2))%32))))
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+8)))
	if v399 == v90&int32(_a_F_get_name_for_var_field_0) {
		v413 = v398
		goto L136
	} else {
		goto L146
	}
L145:
	;
	goto L139
L146:
	;
	v402 = v390 + int32(1)
	if v402 != v385 {
		v390 = v402
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	v419 = v16 + int32(168)
	F_push_child_plan(m, v78, v417, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L22
	} else {
		goto L149
	}
L149:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	v423 = F_get_name_for_var_field(m, v422, l1, l2, l3)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L22
	} else {
		goto L150
	}
L150:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v426 = F_list_delete_first(m, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L22
	} else {
		goto L151
	}
L151:
	;
	base.MemoryCopy(m, v78, v419, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v426
	v860 = v423
	goto L1
L152:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v434+v90<<(uint(int32(2))%32)-int32(4))))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if v441 != int32(6) {
		v667 = v440
		goto L103
	} else {
		goto L153
	}
L153:
	;
	v444 = F_get_name_for_var_field(m, v440, l1, v68, l3)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L22
	} else {
		goto L154
	}
L154:
	;
	v860 = v444
	goto L1
L155:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	if v598 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L156:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v74+v447<<(uint(int32(2))%32))))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+16))
	if v454 == int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	if v457 <= int32(0) {
		goto L155
	} else {
		goto L158
	}
L158:
	;
	v460 = int32(0)
	if v460 < v457 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v464 = v457
	goto L161
L160:
	;
	v464 = v460
	goto L161
L161:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v104)+84))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	v467 = v460
	goto L162
L162:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v466+v467<<(uint(int32(2))%32))))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	if base.B2i32(v487 == int32(0))|base.B2i32(v487 != v490) != 0 {
		v508 = v487
		v509 = v490
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v517 == int32(1) {
		goto L175
	} else {
		goto L176
	}
L164:
	;
	if v508-v509 != 0 {
		goto L171
	} else {
		goto L172
	}
L165:
	;
	goto L164
L166:
	;
	v493 = v484
	v494 = v465
	goto L167
L167:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
	if v498 == int32(0) {
		v508 = v498
		v509 = v497
		goto L165
	} else {
		goto L169
	}
L168:
	;
	v508 = v498
	v509 = v497
	goto L165
L169:
	;
	v501 = int32(1)
	if v498 == v497 {
		v493 = v493 + v501
		v494 = v494 + v501
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v512 = v467 + int32(1)
	if v464 != v512 {
		v467 = v512
		goto L162
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	goto L163
L174:
	;
	goto L155
L175:
	;
	v520 = int32(76)
	goto L177
L176:
	;
	v520 = int32(96)
	goto L177
L177:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v514+v520)))
	if v522 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	if v560 == int32(0) {
		goto L3
	} else {
		goto L191
	}
L179:
	;
	goto L178
L180:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v526 <= int32(0) {
		v560 = int32(0)
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v560 = int32(0)
	goto L179
L183:
	;
	v529 = int32(0)
	if v529 < v526 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v532 = v526
	goto L186
L185:
	;
	v532 = v529
	goto L186
L186:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v537 = int32(0)
	goto L187
L187:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v533+v537<<(uint(int32(2))%32))))
	v546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v545)+8)))
	if v546 == v90&int32(_a_F_get_name_for_var_field_0) {
		v560 = v545
		goto L179
	} else {
		goto L189
	}
L188:
	;
	goto L182
L189:
	;
	v549 = v537 + int32(1)
	if v549 != v532 {
		v537 = v549
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+26)))
	if v564 == int32(1) {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if v568 != int32(6) {
		v667 = v567
		goto L103
	} else {
		goto L193
	}
L193:
	;
	v572 = v16 + int32(168)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v574 = F_list_copy_tail(m, v573, v447)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L22
	} else {
		goto L194
	}
L194:
	;
	F_set_deparse_for_query(m, v572, v514, v574)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L22
	} else {
		goto L195
	}
L195:
	;
	v578 = F_lcons(m, v572, v574)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L22
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v578
	v582 = F_get_name_for_var_field(m, v567, l1, int32(0), l3)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L22
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v573
	v860 = v582
	goto L1
L198:
	;
	v602 = F_palloc(m, int32(32))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L22
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	if v611 != 0 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l1
	v609 = F_pg_snprintf(m, v602, int32(32), int32(_a_F_get_name_for_var_field_5), v16-int32(-64))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L22
	} else {
		goto L202
	}
L202:
	;
	v860 = v602
	goto L1
L203:
	;
	if v649 == int32(0) {
		goto L2
	} else {
		goto L216
	}
L204:
	;
	goto L203
L205:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if v615 <= int32(0) {
		v649 = int32(0)
		goto L204
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v649 = int32(0)
	goto L204
L208:
	;
	v618 = int32(0)
	if v618 < v615 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v621 = v615
	goto L211
L210:
	;
	v621 = v618
	goto L211
L211:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v611)+12))
	v626 = int32(0)
	goto L212
L212:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v622+v626<<(uint(int32(2))%32))))
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v634)+8)))
	if v635 == v90&int32(_a_F_get_name_for_var_field_0) {
		v649 = v634
		goto L204
	} else {
		goto L214
	}
L213:
	;
	goto L207
L214:
	;
	v638 = v626 + int32(1)
	if v638 != v621 {
		v626 = v638
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	v655 = v16 + int32(168)
	F_push_child_plan(m, v78, v653, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L22
	} else {
		goto L217
	}
L217:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	v659 = F_get_name_for_var_field(m, v658, l1, l2, l3)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L22
	} else {
		goto L218
	}
L218:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v662 = F_list_delete_first(m, v661)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L22
	} else {
		goto L219
	}
L219:
	;
	base.MemoryCopy(m, v78, v655, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v662
	v860 = v659
	goto L1
L220:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v860 = v681 + v683<<(uint(int32(4))%32) + l1*int32(100) - int32(76)
	goto L1
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+44)) = v706
	F_set_deparse_plan(m, v696, v693)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L22
	} else {
		goto L222
	}
L222:
	;
	v712 = F_get_name_for_var_field(m, v42, l1, int32(0), l3)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L22
	} else {
		goto L223
	}
L223:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v696)+44))
	F_list_free(m, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L22
	} else {
		goto L224
	}
L224:
	;
	base.MemoryCopy(m, v696, v695, int32(80))
	v860 = v712
	goto L1
L225:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+164)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v723
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_6), v16+int32(160))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L22
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_7), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L22
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_8), v16+int32(16))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L22
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_9), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L22
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_10), v16+int32(32))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L22
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_11), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L22
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_12), v16+int32(48))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L22
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_13), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L22
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v785)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+148)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v786
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_14), v16+int32(144))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L22
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_15), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L22
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_16), v16+int32(128))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L22
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_17), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L22
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_18), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L22
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_19), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L22
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v832
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_20), v16+int32(96))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L22
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_21), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L22
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_16), v16+int32(80))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L22
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_22), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L22
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_negator(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13922(m, l0, int32(40))
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
	v3 = Fn13917(m, l0, int32(40))
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v20 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v5, v5)
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L25
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(48)+v30<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+56))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	if v46 <= int32(2741) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	goto L5
L8:
	;
	v93 = v30 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v93 < v94 {
		v30 = v93
		goto L6
	} else {
		goto L24
	}
L9:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+16)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v73 = F_IndexAmTranslateStrategy(m, v71, v70, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v62 = F_GetIndexAmRoutineByAmId(m, v46, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v46 == int32(783) {
		goto L8
	} else {
		goto L16
	}
L12:
	;
	switch v46 - int32(403) {
	case 0:
		v70 = v46
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
	if base.B2i32(v46 == int32(2742))|base.B2i32(v46 == int32(3580))|base.B2i32(v46 == int32(4000)) != 0 {
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
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+10)))
	F_pfree(m, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v64 != int32(1) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	v70 = v69
	goto L9
L20:
	;
	if v73&int32(-5) != int32(1) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v79 != v80 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v84
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v73
	F_ReleaseCatCacheList(m, v20)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	return int32(1)
L24:
	;
	goto L7
L25:
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	v14 = v12 - int32(96)
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
	F_ScanKeyInit(m, v14, int32(4), int32(3), int32(184), int32(1247))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_ScanKeyInit(m, v14+int32(48), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v43 = F_systable_beginscan(m, v24, int32(2674), int32(1), int32(0), int32(2), v14)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v45 = F_systable_getnext(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = v45
	v52 = v2
	goto L12
L10:
	;
	v252 = v2
	goto L11
L11:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L62
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v60 = v58 + v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v63 = v61 - int32(1247)
	if v63 != 0 {
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
	v245 = F_systable_getnext(m, v43)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L60
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v79 <= int32(0) {
		v239 = v52
		goto L14
	} else {
		goto L29
	}
L16:
	;
	if v63 == int32(12) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v67 = F_get_typtype(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	goto L15
L20:
	;
	v239 = v52
	goto L14
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v67 == int32(100) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v72 = F_get_rels_with_domain(m, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_find_composite_type_dependencies(m, v69, int32(0), v16)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v74 = F_list_concat(m, v52, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v239 = v74
	goto L14
L28:
	;
	v239 = v52
	goto L14
L29:
	;
	if v52 == int32(0) {
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
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v126 = F_relation_open(m, v124, int32(5))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L42
	}
L32:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v84 < v85 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v89 = v85
	goto L35
L34:
	;
	v89 = v84
	goto L35
L35:
	;
	v91 = v84
	goto L36
L36:
	;
	if v91 == v89 {
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v159 = v108
	v160 = v79
	v162 = v52
	goto L30
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v91<<(uint(int32(2))%32)+v106)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+56))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v110 != v111 {
		v91 = v91 + int32(1)
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v142 = F_palloc(m, int32(12))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L48
	}
L41:
	;
	F_relation_close(m, v126, int32(5))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L47
	}
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+72))
	if v129 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_find_composite_type_dependencies(m, v129, int32(0), v16)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v134 = v128
	goto L45
L45:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+119)))
	switch v135 - int32(109) {
	case 0, 5:
		goto L40
	default:
		goto L41
	}
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
	v134 = v133
	goto L45
L47:
	;
	v239 = v52
	goto L14
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v126
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v147)+120)))
	v151 = F_palloc(m, v148<<(uint(int32(2))%32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = v151
	v154 = F_lappend(m, v52, v142)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v159 = v142
	v160 = v156
	v162 = v154
	goto L30
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+52))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v179 = v172 + v173<<(uint(int32(4))%32) + v160*int32(100)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+11)))
	if v180 != 0 {
		v239 = v162
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179-int32(80))+68))
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
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
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
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
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
		v48 = v245
		v52 = v239
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
	m.G0 = v14 + int32(96)
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
func F_get_setop_query(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_get_setop_query[0]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 != int32(142) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L63
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L60
	}
L9:
	;
	m.G0 = v10 + int32(32)
	return
L10:
	;
	if v19 != int32(63) {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 != int32(142) {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+v26<<(uint(int32(2))%32)-int32(4))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	F_get_query_def(m, v33, v12, v46, v47, v48, v49, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L23
	}
L15:
	;
	F_appendStringInfoChar(m, v12, int32(40))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L22
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	if v35 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+140))
	if v36 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
	if v37 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
	if v38 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+144))
	if v39 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v45 = int32(0)
	goto L14
L22:
	;
	v45 = int32(1)
	goto L14
L23:
	;
	if v45 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	F_appendStringInfoChar(m, v12, int32(41))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L9
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = v101 - int32(1)
	if base.Ui32(int32(3)) <= base.Ui32(v103) {
		goto L8
	} else {
		goto L43
	}
L27:
	;
	F_get_setop_query(m, v59, l1, l2)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L37
	}
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v63 == v64 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)))
	if v66 == v67 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_appendStringInfoChar(m, v12, int32(40))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v74 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_4), int32(8), v74, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_get_setop_query(m, v78, l1, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v83 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_9), int32(-8), v83, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L26
L37:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v89&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v93 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_4), v93, v93, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_appendStringInfoChar(m, v12, int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L26
L42:
	;
	goto L26
L43:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103<<(uint(int32(2))%32))+uint32(_c_F_get_setop_query[1])))
	F_appendStringInfoString(m, v12, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v111 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_get_setop_query_7))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v119 == int32(142) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_appendStringInfoChar(m, v12, int32(40))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	v126 = int32(0)
	goto L51
L51:
	;
	v128 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_4), v126, v128, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	v126 = int32(8)
	goto L51
L53:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)))
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)) = uint8(v133)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_setop_query(m, v135, l1, l2)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)) = uint8(v132)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v139&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v142 - v126
	goto L57
L56:
	;
	goto L57
L57:
	;
	if v119 != int32(142) {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v148 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_8), v148, v148, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	goto L9
L60:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v164
	F_errmsg_internal(m, int32(_a_F_get_setop_query_5), v10+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_get_setop_query_1), int32(_a_F_get_setop_query_6), int32(_a_F_get_setop_query_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v180
	F_errmsg_internal(m, int32(_a_F_get_setop_query_0), v10)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_get_setop_query_1), int32(_a_F_get_setop_query_2), int32(_a_F_get_setop_query_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
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
	v7 = Fn13925(m, l0, l1, int32(_a_F_get_statistics_object_oid_0), int32(2620), int32(_a_F_get_statistics_object_oid_1), int32(63))
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
		if int32(2) <= v11 {
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
					F_errfinish(m, int32(_a_F_get_th_9), int32(1572), int32(_a_F_get_th_10))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9)+82)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14 & int32(1)
			}
		} else {
			v14 = int32(0)
			return v14 & int32(1)
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_get_typlenbyval_0), v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_typlenbyval_1), int32(2398), int32(_a_F_get_typlenbyval_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v31)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v33)
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	v12 = F_palloc0(m, int32(40))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = F_palloc0(m, int32(36))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = F_pg_detoast_datum_packed(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = int32(1)
				v23 = v20 + v22
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				v28 = v26 & v22
				if v28 != 0 {
					v29 = v23
				} else {
					v29 = v20 + int32(4)
				}
				if v26 == int32(1) {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
					if v35 == int32(18) {
						v38 = int32(16)
					} else {
						v38 = int32(0)
					}
					if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v45 = int32(4)
					} else {
						v45 = v38
					}
					v56 = v45
				} else {
					v46 = int32(1)
					if v28 != 0 {
						v56 = int32(base.Ui32(v26)>>(uint(v46)%32)) - v46
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
						v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_get_worker[0]))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
				v61 = F_makeJsonLexContextCstringLen(m, int32(0), v29, v56, v59, int32(1))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l3
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v5)
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v61
					v68 = F_palloc0(m, l3)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v68
						v73 = F_palloc(m, l3<<(uint(int32(2))%32))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v73
							if int32(0) < l3 {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
								v79 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v79)
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
								v95 = int32(1335)
								v96 = int32(36)
								*(*int32)(unsafe.Add(mBase, uint32(v96+v12))) = v95
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(1335)
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
								if l3 != 0 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(1336)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(1337)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(1338)
									v95 = int32(1339)
									v96 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(v96+v12))) = v95
								}
							}
							if l1 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1340)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1341)
							} else {
							}
							if l2 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(1342)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(1343)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(1336)
							} else {
							}
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
							v111 = F_pg_parse_json(m, v110, v12)
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
										v116 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
										F_freeJsonLexContext(m, v116)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
											return v119
										}
									}
								} else {
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									F_freeJsonLexContext(m, v116)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
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
	F_errdetail(m, int32(_a_F_getid_1), v11)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, l2, int32(_a_F_getid_2), int32(206), int32(_a_F_getid_3))
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
				F_errfinish(m, int32(_a_F_getinternalerrposition_1), int32(1622), int32(_a_F_getinternalerrposition_2))
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
						F_errfinish(m, int32(_a_F_getlen_1), int32(2864), int32(_a_F_getlen_2))
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
					F_errfinish(m, int32(_a_F_getlen_1), int32(2862), int32(_a_F_getlen_2))
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
func F_ginarrayconsistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	switch v16 - int32(1) {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L2
	case 3:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v107
L2:
	;
	v103 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v103)
	v107 = v103
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L4:
	;
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v69)
	v72 = int32(0)
	if v14 <= v72 {
		v107 = v69
		goto L1
	} else {
		goto L23
	}
L5:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v45)
	if v14 <= v45 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v19)
	if v14 <= v19 {
		v107 = int32(0)
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v15))))
	if v32 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = int32(0)
	goto L1
L10:
	;
	v35 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v12))))
	if v37 != v35 {
		v107 = v35
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v42 = v24 + int32(1)
	if v42 != v14 {
		v24 = v42
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	goto L9
L15:
	;
	v107 = int32(1)
	goto L1
L16:
	;
	goto L17
L17:
	;
	v51 = v45
	goto L18
L18:
	;
	v58 = int32(0)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v15))))
	if v60 != int32(1) {
		v107 = v58
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v107 = v65
	goto L1
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v12))))
	if v64 != 0 {
		v107 = v58
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v65 = int32(1)
	v67 = v51 + v65
	if v14 != v67 {
		v51 = v67
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v75 = v72
	goto L24
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v15))))
	if v83 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v107 = int32(0)
	goto L1
L26:
	;
	v85 = v75 + int32(1)
	if v14 != v85 {
		v75 = v85
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	v107 = v69
	goto L1
L30:
	;
	return int32(0)
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v16
	F_errmsg_internal(m, int32(_a_F_ginarrayconsistent_0), v10)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ginarrayconsistent_1), int32(215), int32(_a_F_ginarrayconsistent_2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
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
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
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
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v264 int32
	_ = v264
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
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v301 int64
	_ = v301
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v368 int32
	_ = v368
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v737 int32
	_ = v737
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int64
	_ = v782
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int64
	_ = v985
	var v990 int32
	_ = v990
	var v991 int64
	_ = v991
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1109 int64
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1223 int64
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1261 int32
	_ = v1261
	var v1272 int32
	_ = v1272
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1483 int32
	_ = v1483
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1929 int64
	_ = v1929
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1964 int64
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1998 int64
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2005 int32
	_ = v2005
	var v2023 int32
	_ = v2023
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2079 int32
	_ = v2079
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2109 int32
	_ = v2109
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2147 int32
	_ = v2147
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2219 int32
	_ = v2219
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2281 int32
	_ = v2281
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2298 int32
	_ = v2298
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2325 int32
	_ = v2325
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2345 int32
	_ = v2345
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2352 int32
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2419 int32
	_ = v2419
	var v2426 int32
	_ = v2426
	var v2454 int32
	_ = v2454
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2485 int32
	_ = v2485
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2498 int32
	_ = v2498
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2510 int32
	_ = v2510
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2518 int32
	_ = v2518
	var v2522 int32
	_ = v2522
	var v2524 int32
	_ = v2524
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2540 int32
	_ = v2540
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2557 int32
	_ = v2557
	var v2564 int32
	_ = v2564
	var v2593 int32
	_ = v2593
	var v2597 int32
	_ = v2597
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2620 int32
	_ = v2620
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2629 int32
	_ = v2629
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2637 int32
	_ = v2637
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2656 int32
	_ = v2656
	var v2658 int32
	_ = v2658
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2664 int32
	_ = v2664
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2704 int32
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2729 int32
	_ = v2729
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2774 int32
	_ = v2774
	var v2777 int32
	_ = v2777
	var v2780 int32
	_ = v2780
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2789 int32
	_ = v2789
	var v2792 int32
	_ = v2792
	var v2823 int32
	_ = v2823
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2841 int32
	_ = v2841
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2853 int32
	_ = v2853
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2859 int32
	_ = v2859
	var v2862 int32
	_ = v2862
	var v2863 int32
	_ = v2863
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2873 int32
	_ = v2873
	var v2879 int32
	_ = v2879
	var v2881 int32
	_ = v2881
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2908 int32
	_ = v2908
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2913 int32
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2918 int32
	_ = v2918
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2943 int32
	_ = v2943
	var v2945 int32
	_ = v2945
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2952 int32
	_ = v2952
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2969 int32
	_ = v2969
	var v3001 int32
	_ = v3001
	var v3004 int32
	_ = v3004
	var v3037 int32
	_ = v3037
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3045 int32
	_ = v3045
	var v3049 int32
	_ = v3049
	var v3053 int32
	_ = v3053
	var v3085 int32
	_ = v3085
	var v3086 int32
	_ = v3086
	var v3089 int32
	_ = v3089
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3131 int32
	_ = v3131
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3178 int32
	_ = v3178
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3214 int32
	_ = v3214
	var v3218 int32
	_ = v3218
	var v3222 int32
	_ = v3222
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3229 int32
	_ = v3229
	var v3235 int32
	_ = v3235
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3268 int32
	_ = v3268
	var v3270 int32
	_ = v3270
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3289 int32
	_ = v3289
	var v3292 int32
	_ = v3292
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3333 int32
	_ = v3333
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3370 int32
	_ = v3370
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3403 int32
	_ = v3403
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3415 int32
	_ = v3415
	var v3442 int32
	_ = v3442
	var v3448 int32
	_ = v3448
	var v3477 int32
	_ = v3477
	var v3481 int32
	_ = v3481
	var v3483 int32
	_ = v3483
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3491 int32
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3500 int32
	_ = v3500
	var v3511 int32
	_ = v3511
	var v3534 int32
	_ = v3534
	var v3536 int32
	_ = v3536
	var v3539 int32
	_ = v3539
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3547 int32
	_ = v3547
	var v3550 int32
	_ = v3550
	var v3551 int32
	_ = v3551
	var v3553 int32
	_ = v3553
	var v3558 int32
	_ = v3558
	var v3587 int32
	_ = v3587
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3591 int32
	_ = v3591
	var v3593 int32
	_ = v3593
	var v3597 int32
	_ = v3597
	var v3600 int32
	_ = v3600
	var v3601 int32
	_ = v3601
	var v3605 int32
	_ = v3605
	var v3634 int32
	_ = v3634
	var v3640 int32
	_ = v3640
	var v3642 int32
	_ = v3642
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3673 int32
	_ = v3673
	var v3677 int32
	_ = v3677
	var v3681 int32
	_ = v3681
	var v3683 int32
	_ = v3683
	var v3686 int32
	_ = v3686
	var v3687 int32
	_ = v3687
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3724 int32
	_ = v3724
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3805 int32
	_ = v3805
	var v3809 int32
	_ = v3809
	var v3810 int32
	_ = v3810
	var v3814 int32
	_ = v3814
	var v3838 int64
	_ = v3838
	var v3840 int32
	_ = v3840
	var v3842 int32
	_ = v3842
	var v3844 int32
	_ = v3844
	var v3846 int32
	_ = v3846
	var v3879 int32
	_ = v3879
	var v3881 int32
	_ = v3881
	var v3882 int32
	_ = v3882
	var v3886 int32
	_ = v3886
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3900 int32
	_ = v3900
	var v3903 int32
	_ = v3903
	var v3910 int32
	_ = v3910
	var v3924 int32
	_ = v3924
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3931 int32
	_ = v3931
	var v3932 int32
	_ = v3932
	var v3937 int32
	_ = v3937
	var v3941 int32
	_ = v3941
	var v3942 int32
	_ = v3942
	var v3944 int32
	_ = v3944
	var v3947 int32
	_ = v3947
	var v3949 int64
	_ = v3949
	var v3952 int64
	_ = v3952
	var v3955 int64
	_ = v3955
	var v3956 int64
	_ = v3956
	var v3957 int64
	_ = v3957
	var v3960 int64
	_ = v3960
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3977 int32
	_ = v3977
	var v3978 int32
	_ = v3978
	var v3980 int32
	_ = v3980
	var v3982 int32
	_ = v3982
	var v3995 int32
	_ = v3995
	var v4005 int32
	_ = v4005
	var v4009 int32
	_ = v4009
	var v4010 int32
	_ = v4010
	var v4011 int32
	_ = v4011
	var v4013 int64
	_ = v4013
	var v4015 int32
	_ = v4015
	var v4021 int32
	_ = v4021
	var v4026 int64
	_ = v4026
	var v4028 int32
	_ = v4028
	var v4030 int32
	_ = v4030
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4037 int32
	_ = v4037
	var v4039 int64
	_ = v4039
	var v4041 int32
	_ = v4041
	var v4047 int32
	_ = v4047
	var v4053 int32
	_ = v4053
	var v4054 int32
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4056 int64
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4059 int64
	_ = v4059
	var v4072 int32
	_ = v4072
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4078 int32
	_ = v4078
	var v4081 int32
	_ = v4081
	var v4082 int32
	_ = v4082
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4129 int32
	_ = v4129
	var v4132 int32
	_ = v4132
	var v4137 int32
	_ = v4137
	var v4146 int32
	_ = v4146
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4186 int32
	_ = v4186
	var v4187 int32
	_ = v4187
	var v4188 int32
	_ = v4188
	var v4192 int32
	_ = v4192
	var v4197 int32
	_ = v4197
	var v4200 int32
	_ = v4200
	var v4215 int32
	_ = v4215
	var v4217 int64
	_ = v4217
	var v4233 int32
	_ = v4233
	var v4234 int32
	_ = v4234
	var v4236 int32
	_ = v4236
	var v4238 int32
	_ = v4238
	var v4261 int32
	_ = v4261
	var v4265 int32
	_ = v4265
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4269 int64
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4277 int32
	_ = v4277
	var v4282 int64
	_ = v4282
	var v4284 int32
	_ = v4284
	var v4286 int32
	_ = v4286
	var v4289 int32
	_ = v4289
	var v4290 int32
	_ = v4290
	var v4291 int32
	_ = v4291
	var v4293 int64
	_ = v4293
	var v4295 int32
	_ = v4295
	var v4301 int32
	_ = v4301
	var v4307 int32
	_ = v4307
	var v4308 int32
	_ = v4308
	var v4309 int32
	_ = v4309
	var v4310 int64
	_ = v4310
	var v4312 int64
	_ = v4312
	var v4325 int32
	_ = v4325
	var v4326 int32
	_ = v4326
	var v4327 int32
	_ = v4327
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4339 int32
	_ = v4339
	var v4340 int32
	_ = v4340
	var v4344 int32
	_ = v4344
	var v4370 int32
	_ = v4370
	var v4373 int32
	_ = v4373
	var v4374 int32
	_ = v4374
	var v4378 int64
	_ = v4378
	var v4388 int32
	_ = v4388
	var v4392 int32
	_ = v4392
	var v4401 int32
	_ = v4401
	var v4421 int32
	_ = v4421
	var v4425 int32
	_ = v4425
	var v4426 int32
	_ = v4426
	var v4427 int64
	_ = v4427
	var v4428 int64
	_ = v4428
	var v4431 int64
	_ = v4431
	var v4437 int32
	_ = v4437
	var v4438 int32
	_ = v4438
	var v4439 int32
	_ = v4439
	var v4441 int32
	_ = v4441
	var v4444 int32
	_ = v4444
	var v4447 int32
	_ = v4447
	var v4449 int32
	_ = v4449
	var v4452 int32
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4457 int32
	_ = v4457
	var v4458 int32
	_ = v4458
	var v4465 int32
	_ = v4465
	var v4466 int32
	_ = v4466
	var v4467 int32
	_ = v4467
	var v4468 int32
	_ = v4468
	var v4477 int32
	_ = v4477
	var v4494 int32
	_ = v4494
	var v4513 int32
	_ = v4513
	var v4515 int64
	_ = v4515
	var v4518 int64
	_ = v4518
	var v4521 int64
	_ = v4521
	var v4533 int32
	_ = v4533
	var v4562 int32
	_ = v4562
	var v4566 int32
	_ = v4566
	var v4567 int32
	_ = v4567
	var v4570 int32
	_ = v4570
	var v4572 int32
	_ = v4572
	var v4574 int64
	_ = v4574
	var v4575 int64
	_ = v4575
	var v4578 int64
	_ = v4578
	var v4582 int64
	_ = v4582
	var v4584 int32
	_ = v4584
	var v4586 int32
	_ = v4586
	var v4588 int32
	_ = v4588
	var v4589 int32
	_ = v4589
	var v4591 int32
	_ = v4591
	var v4593 int32
	_ = v4593
	var v4598 int32
	_ = v4598
	var v4599 int32
	_ = v4599
	var v4632 int32
	_ = v4632
	var v4633 int32
	_ = v4633
	var v4634 int32
	_ = v4634
	var v4637 int32
	_ = v4637
	var v4639 int32
	_ = v4639
	var v4641 int32
	_ = v4641
	var v4646 int32
	_ = v4646
	var v4678 int32
	_ = v4678
	var v4680 int32
	_ = v4680
	var v4681 int32
	_ = v4681
	var v4684 int32
	_ = v4684
	var v4685 int32
	_ = v4685
	var v4686 int32
	_ = v4686
	var v4691 int32
	_ = v4691
	var v4700 int32
	_ = v4700
	var v4706 int32
	_ = v4706
	var v4713 int32
	_ = v4713
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4720 int64
	_ = v4720
	var v4725 int32
	_ = v4725
	var v4732 int32
	_ = v4732
	var v4733 int32
	_ = v4733
	var v4735 int32
	_ = v4735
	var v4737 int32
	_ = v4737
	var v4740 int32
	_ = v4740
	var v4742 int32
	_ = v4742
	var v4747 int64
	_ = v4747
	var v4751 int64
	_ = v4751
	var v4763 int32
	_ = v4763
	var v4764 int32
	_ = v4764
	var v4765 int32
	_ = v4765
	var v4766 int32
	_ = v4766
	var v4773 int32
	_ = v4773
	var v4775 int32
	_ = v4775
	var v4777 int32
	_ = v4777
	var v4803 int32
	_ = v4803
	var v4804 int32
	_ = v4804
	var v4813 int32
	_ = v4813
	var v4817 int32
	_ = v4817
	var v4849 int32
	_ = v4849
	var v4851 int32
	_ = v4851
	var v4859 int32
	_ = v4859
	var v4884 int32
	_ = v4884
	var v4887 int32
	_ = v4887
	var v4888 int32
	_ = v4888
	var v4891 int32
	_ = v4891
	var v4895 int32
	_ = v4895
	var v4904 int32
	_ = v4904
	var v4933 int32
	_ = v4933
	var v4939 int32
	_ = v4939
	var v4942 int32
	_ = v4942
	var v4943 int32
	_ = v4943
	var v4944 int32
	_ = v4944
	var v4952 int32
	_ = v4952
	var v4957 int32
	_ = v4957
	var v4963 int32
	_ = v4963
	var v4987 int64
	_ = v4987
	v3 = int32(0)
	v28 = int64(0)
	v32 = m.G0
	v34 = v32 - int32(_a_F_gingetbitmap_0)
	m.G0 = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ginFreeScanKeys(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v41 = m.G0
	v43 = v41 - int32(96)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v47 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+88)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v43)+80)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v43)+72)) = v47
	*(*int64)(unsafe.Add(mBase, uint32(v43)+64)) = v47
	v55 = int32(_a_F_gingetbitmap_1)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v58
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v61 <= v60 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v64 = v60
	goto L5
L4:
	;
	v64 = v61
	goto L5
L5:
	;
	v67 = F_palloc(m, v64*int32(92))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[2]))) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4]))) = v67
	v75 = F_palloc(m, int32(128))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[5]))) = uint8(v77)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[6]))) = v75
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v80 <= v77 {
		v496 = v3
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	if v521 != 0 {
		goto L72
	} else {
		goto L73
	}
L9:
	;
	v93 = v3
	v101 = v3
	goto L10
L10:
	;
	v118 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+60)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+56)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+52)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v118
	v130 = v45 + v101*int32(48)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v131&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v488 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[5]))) = uint8(v488)
	v496 = v93
	goto L8
L12:
	;
	goto L11
L13:
	;
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v130)+4)))
	v136 = v134 - int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(_a_F_gingetbitmap_2)+v136<<(uint(int32(2))%32))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)))
	v156 = F_FunctionCall7Coll(m, v46+int32(1936)+v136*int32(28), v143, v144, v43+int32(16), v147, v43+int32(60), v43+int32(56), v43+int32(52), v43+int32(48))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	if base.Ui32(int32(3)) <= base.Ui32(v158) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v161 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v161
	v164 = v161
	goto L17
L16:
	;
	v164 = v158
	goto L17
L17:
	;
	if v156 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v177 = base.B2i32(v164 != int32(0)) | v93
	v178 = F_palloc0(m, v174)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L24
	}
L19:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if int32(0) < v165 {
		v174 = v165
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v164 == int32(0) {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v171 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v171
	v174 = v171
	goto L18
L24:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	if v181 == int32(0) {
		v239 = v177
		v241 = v180
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+4)))
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+6)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v43)+60))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	v271 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3]))) = v270 + v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	v277 = v274 + v270*int32(92)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v241
	v280 = int32(2)
	v282 = v241 + v271
	v285 = F_palloc(m, v282<<(uint(v280)%32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L34
	}
L26:
	;
	v184 = int32(0)
	if v180 <= v184 {
		v239 = v177
		v241 = v180
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v193 = v177
	v194 = v184
	v195 = v180
	goto L28
L28:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218+v194))))
	if v220 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v239 = v228
	v241 = v229
	goto L25
L30:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v194+v178))) = uint8(v224)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v228 = v224
	v229 = v226
	goto L32
L31:
	;
	v228 = v193
	v229 = v195
	goto L32
L32:
	;
	v231 = v194 + int32(1)
	if v231 < v229 {
		v193 = v228
		v194 = v231
		v195 = v229
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v285
	v288 = F_palloc0(m, v282)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+78)) = uint8(base.B2i32(v266 == int32(2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+76)) = uint16(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+72)) = v266
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+68)) = uint16(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+64)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v277)+60)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v277)+56)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v277)+52)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v277)+28)) = v288
	v301 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v277)+12)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v277)+20)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v277)+80)) = v301
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+88)) = uint8(v307)
	v311 = v46 + int32(4)
	if v266 == int32(3) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v241 != 0 {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+32)) = int32(55)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+36)) = int32(56)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+76)))
	v324 = v311 + v321*int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+44)) = v324 + int32(3696)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+40)) = v324 + int32(2800)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v311+v321<<(uint(int32(2))%32))+uint32(_c_F_gingetbitmap[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+48)) = v336
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v324+int32(2804))))
	if v342 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v343 = int32(57)
	goto L42
L41:
	;
	v343 = int32(58)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+32)) = v343
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v324+int32(3700))))
	if v349 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v350 = int32(59)
	goto L45
L44:
	;
	v350 = int32(60)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+36)) = v350
	goto L36
L46:
	;
	v368 = v307
	goto L49
L47:
	;
	goto L48
L48:
	;
	switch v266 - int32(1) {
	case 0:
		v452 = v280
		goto L60
	default:
		goto L59
	case 2:
		goto L61
	}
L49:
	;
	v389 = v368 << (uint(int32(2)) % 32)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v156+v389)))
	v393 = int32(*(*int8)(unsafe.Add(mBase, uint32(v178+v368))))
	v394 = int32(0)
	if v268 == v394 {
		v405 = v394
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	if v269 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v397 = int32(0)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v264)+uint32(_c_F_gingetbitmap[8]))))
	if v398&int32(1) == v397 {
		v405 = v397
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368+v268))))
	v405 = v404
	goto L51
L54:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v269+v389)))
	v409 = v407
	goto L56
L55:
	;
	v409 = int32(0)
	goto L56
L56:
	;
	v410 = F_ginFillScanEntry(m, v46, v264, v265, v266, v391, v393, v405, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v412+v389))) = v410
	v416 = v368 + int32(1)
	if v416 != v241 {
		v368 = v416
		goto L49
	} else {
		goto L58
	}
L58:
	;
	goto L50
L59:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	if v474 != int32(2) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v453 + int32(1)
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+76)))
	v458 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v277)+72))
	v464 = F_ginFillScanEntry(m, v46, v457, v458, v459, v458, base.I32_extend8_s(v452), v458, v458)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v452 = int32(255)
	goto L60
L62:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v466+v453<<(uint(int32(2))%32)))) = v464
	goto L59
L63:
	;
	v477 = int32(*(*int16)(unsafe.Add(mBase, uint32(v130)+4)))
	v479 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v477)+63)) = uint8(v479)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v482 = v101 + int32(1)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v482 < v483 {
		v93 = v239
		v101 = v482
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v496 = v239
	goto L8
L67:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022)+uint32(_c_F_gingetbitmap[5]))))
	if v1023 != 0 {
		v4963 = v34
		v4987 = v28
		goto L135
	} else {
		goto L136
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L130
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v56
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)+272))
	if v974 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L70:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[5]))))
	if v931 != 0 {
		goto L69
	} else {
		goto L118
	}
L71:
	;
	if v496&int32(1) == int32(0) {
		goto L69
	} else {
		goto L117
	}
L72:
	;
	v525 = v3
	v531 = v521
	v535 = int32(0)
	goto L75
L73:
	;
	goto L74
L74:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[5]))))
	if v769 != 0 {
		goto L71
	} else {
		goto L103
	}
L75:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	v557 = v554 + v525*int32(92)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v557)+72))
	if v558 != int32(2) {
		v592 = v531
		v594 = v535
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if int32(0) < v594 {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	v596 = v525 + int32(1)
	if base.Ui32(v596) < base.Ui32(v592) {
		v525 = v596
		v531 = v592
		v535 = v594
		goto L75
	} else {
		goto L83
	}
L78:
	;
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+76)))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v561)+63)))
	if v563 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v566 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v557)+78)) = uint8(v566)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	*(*int32)(unsafe.Add(mBase, uint32(v557))) = v568 + int32(1)
	v578 = F_ginFillScanEntry(m, v46, v561, v566, int32(2), v566, int32(-1), v566, v566)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v592 = v531
	v594 = v535 + int32(1)
	goto L77
L82:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v557)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v580+v568<<(uint(int32(2))%32)))) = v578
	v585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+76)))
	v587 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v585)+63)) = uint8(v587)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	v592 = v589
	v594 = v535
	goto L77
L83:
	;
	goto L76
L84:
	;
	v602 = F_palloc(m, v592*int32(92))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v737 = v592
	goto L86
L86:
	;
	if v737 != 0 {
		goto L71
	} else {
		goto L102
	}
L87:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	if v604 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v606 = int32(0)
	v610 = v606
	v615 = v606
	v621 = v604 - v594
	goto L91
L89:
	;
	v700 = int32(0)
	goto L90
L90:
	;
	if v700 != 0 {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	v642 = v639 + v615*int32(92)
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+78)))
	if v643 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v700 = v664 * int32(92)
	goto L90
L93:
	;
	v663 = v615 + int32(1)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v663) < base.Ui32(v664) {
		v610 = v660
		v615 = v663
		v621 = v661
		goto L91
	} else {
		goto L97
	}
L94:
	;
	v646 = int32(92)
	base.MemoryCopy(m, v602+v621*v646, v642, v646)
	v660 = v610
	v661 = v621 + int32(1)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v653 = int32(92)
	base.MemoryCopy(m, v602+v610*v653, v642, v653)
	v660 = v610 + int32(1)
	v661 = v621
	goto L93
L97:
	;
	goto L92
L98:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	base.MemoryCopy(m, v701, v602, v700)
	goto L100
L99:
	;
	goto L100
L100:
	;
	F_pfree(m, v602)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	v737 = v705
	goto L86
L102:
	;
	goto L74
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3]))) = int32(1)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v772))) = int64(0)
	v776 = F_palloc(m, int32(4))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772)+8)) = v776
	v780 = F_palloc0(m, int32(1))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v782 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v772)+52)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v772)+28)) = v780
	*(*int64)(unsafe.Add(mBase, uint32(v772)+60)) = v782
	v787 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v772)+68)) = uint16(v787)
	*(*uint8)(unsafe.Add(mBase, uint32(v772)+78)) = uint8(v787)
	v791 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v772)+76)) = uint16(v791)
	*(*int32)(unsafe.Add(mBase, uint32(v772)+72)) = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v772)+12)) = v782
	*(*int64)(unsafe.Add(mBase, uint32(v772)+20)) = v782
	*(*int64)(unsafe.Add(mBase, uint32(v772)+80)) = v782
	*(*uint8)(unsafe.Add(mBase, uint32(v772)+88)) = uint8(v787)
	goto L107
L106:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v847 + int32(1)
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v772)+76)))
	v852 = int32(0)
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v772)+72))
	v858 = F_ginFillScanEntry(m, v46, v851, v852, v853, v852, int32(-1), v852, v852)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L116
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772)+32)) = int32(55)
	*(*int32)(unsafe.Add(mBase, uint32(v772)+36)) = int32(56)
	goto L106
L116:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v772)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v860+v847<<(uint(int32(2))%32)))) = v858
	goto L70
L117:
	;
	goto L70
L118:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ginGetStats(m, v932, v43+int32(16))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
	if v937 <= int32(0) {
		goto L68
	} else {
		goto L120
	}
L120:
	;
	goto L69
L121:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v990 != 0 {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973)+268)))
	if v977 != int32(1) {
		goto L121
	} else {
		goto L125
	}
L123:
	;
	v984 = v974
	goto L124
L124:
	;
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v984)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v984)+16)) = v985 + int64(1)
	goto L121
L125:
	;
	F_pgstat_assoc_relation(m, v973)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v982)+272))
	v984 = v983
	goto L124
L127:
	;
	v991 = *(*int64)(unsafe.Add(mBase, uint32(v990)))
	*(*int64)(unsafe.Add(mBase, uint32(v990))) = v991 + int64(1)
	goto L129
L128:
	;
	goto L129
L129:
	;
	m.G0 = v43 + int32(96)
	goto L67
L130:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_gingetbitmap_3), int32(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v1010 + int32(4)
	F_errhint(m, int32(_a_F_gingetbitmap_4), v43)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_5), int32(482), int32(_a_F_gingetbitmap_6))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
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
	m.G0 = v4963 + int32(_a_F_gingetbitmap_0)
	return v4987
L136:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1026 = F_ReadBuffer(m, v1024, int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v1028, int32(0), v1030)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_LockBuffer(m, v1026, int32(1))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	if v1026 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v1969)+36))
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[2])))
	if v2001 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L141:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+24))
	if v1054 == int32(-1) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1039+(v1026^int32(-1))<<(uint(int32(2))%32))))
	v1053 = v1045
	goto L141
L143:
	;
	goto L144
L144:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v1053 = v1047 + v1026<<(uint(int32(13))%32) + int32(-8192)
	goto L141
L145:
	;
	F_UnlockReleaseBuffer(m, v1026)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1060 = F_ReadBuffer(m, v1059, v1054)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v1969 = l0
	v1970 = l1
	v1974 = v34
	v1998 = v28
	goto L140
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v1060
	F_LockBuffer(m, v1060, int32(1))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v1066 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+40)) = uint16(v1066)
	F_UnlockReleaseBuffer(m, v1026)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+uint32(_c_F_gingetbitmap[3])))
	v1071 = F_palloc(m, v1070)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v1071
	v1076 = F_scanGetCandidate(m, l0, v34+int32(36))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if v1076 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1080 = l0
	v1081 = l1
	v1085 = v34
	v1091 = v1022
	v1095 = v34 + int32(44)
	v1109 = v28
	goto L157
L155:
	;
	v1935 = l0
	v1936 = l1
	v1940 = v34
	v1964 = v28
	goto L156
L156:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1940)+52))
	F_pfree(m, v1966)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L1
	} else {
		goto L278
	}
L157:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+36))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+uint32(_c_F_gingetbitmap[3])))
	if v1113 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1935 = v1194
	v1936 = v1195
	v1940 = v1199
	v1964 = v1929
	goto L156
L159:
	;
	v1118 = int32(0)
	goto L162
L160:
	;
	v1160 = int32(0)
	goto L161
L161:
	;
	if v1160 != 0 {
		goto L168
	} else {
		goto L169
	}
L162:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+uint32(_c_F_gingetbitmap[4])))
	v1149 = v1146 + v1118*int32(92)
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1149)))
	if v1150 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1160 = v1156
	goto L161
L164:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+28))
	base.MemoryFill(m, v1151, int32(0), v1150)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v1155 = v1118 + int32(1)
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v1155) < base.Ui32(v1156) {
		v1118 = v1155
		goto L162
	} else {
		goto L167
	}
L167:
	;
	goto L163
L168:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+52))
	base.MemoryFill(m, v1189, int32(0), v1160)
	goto L170
L169:
	;
	goto L170
L170:
	;
	v1194 = v1080
	v1195 = v1081
	v1199 = v1085
	v1205 = v1091
	v1206 = v1112
	v1209 = v1095
	v1210 = v1112 + int32(4)
	v1223 = v1109
	goto L172
L171:
	;
	if v1643 != 0 {
		goto L253
	} else {
		goto L254
	}
L172:
	;
	v1225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+42)))
	v1226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+40)))
	v1227 = v1225 - v1226
	if v1227 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L1
	} else {
		goto L249
	}
L174:
	;
	base.MemoryFill(m, v1226+v1199+int32(63), int32(0), v1227)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+36))
	if v1233 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+uint32(_c_F_gingetbitmap[3])))
	if v1252 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1237+(v1233^int32(-1))<<(uint(int32(2))%32))))
	v1251 = v1243
	goto L177
L179:
	;
	goto L180
L180:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v1251 = v1245 + v1233<<(uint(int32(13))%32) + int32(-8192)
	goto L177
L181:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1199)+40)) = uint16(v1644)
	v1673 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1251)+16)))
	v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251+v1673)+6)))
	if v1675&int32(32) != 0 {
		goto L171
	} else {
		goto L237
	}
L182:
	;
	v1643 = int32(0)
	v1644 = v1225
	goto L181
L183:
	;
	goto L184
L184:
	;
	v1261 = v1252
	v1272 = int32(0)
	goto L185
L185:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+uint32(_c_F_gingetbitmap[4])))
	v1293 = v1290 + v1272*int32(92)
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1293)))
	if v1294 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v1640 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+42)))
	v1643 = v1608
	v1644 = v1640
	goto L181
L187:
	;
	v1299 = v1294
	v1303 = int32(0)
	goto L190
L188:
	;
	v1608 = v1261
	goto L189
L189:
	;
	v1638 = v1272 + int32(1)
	if base.Ui32(v1638) < base.Ui32(v1608) {
		v1261 = v1608
		v1272 = v1638
		goto L185
	} else {
		goto L236
	}
L190:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+28))
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327+v1303))))
	if v1329 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+uint32(_c_F_gingetbitmap[3])))
	v1608 = v1605
	goto L189
L192:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+8))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1332+v1303<<(uint(int32(2))%32))))
	v1337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+40)))
	v1338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+42)))
	if base.Ui32(v1338) <= base.Ui32(v1337) {
		v1483 = v1338
		goto L196
	} else {
		goto L197
	}
L193:
	;
	v1574 = v1299
	goto L194
L194:
	;
	v1603 = v1303 + int32(1)
	if base.Ui32(v1603) < base.Ui32(v1574) {
		v1299 = v1574
		v1303 = v1603
		goto L190
	} else {
		goto L235
	}
L195:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+52))
	v1561 = v1560 + v1272
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1561))))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+28))
	v1565 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1563+v1303))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1561))) = uint8(base.B2i32(v1562|v1565 != int32(0)))
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1293)))
	v1574 = v1570
	goto L194
L196:
	;
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336)+5)))
	if v1512 != int32(1) {
		goto L195
	} else {
		goto L233
	}
L197:
	;
	v1342 = v1338
	v1343 = v1337
	goto L198
L198:
	;
	v1376 = int32(base.Ui32((v1342-v1343)&int32(_a_F_gingetbitmap_7))>>(uint(int32(1))%32)) + v1343
	v1378 = v1376 & int32(_a_F_gingetbitmap_8)
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1251+int32(20)+v1378<<(uint(int32(2))%32))))
	v1385 = v1251 + v1382&int32(_a_F_gingetbitmap_9)
	v1386 = F_gintuple_get_attrnum(m, v1210, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L1
	} else {
		goto L200
	}
L199:
	;
	v1483 = v1471
	goto L196
L200:
	;
	v1388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1293)+76)))
	if base.Ui32(v1388) < base.Ui32(v1386) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v1476 = int32(_a_F_gingetbitmap_8)
	if base.Ui32(v1472&v1476) < base.Ui32(v1471&v1476) {
		v1342 = v1471
		v1343 = v1472
		goto L198
	} else {
		goto L232
	}
L202:
	;
	v1471 = v1376
	v1472 = v1343
	goto L201
L203:
	;
	goto L204
L204:
	;
	if base.Ui32(v1386) < base.Ui32(v1388) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1471 = v1342
	v1472 = v1376 + int32(1)
	goto L201
L206:
	;
	goto L207
L207:
	;
	v1394 = v1378 - int32(1)
	v1397 = v1394 + (v1199 - int32(-64))
	v1398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1397))))
	if v1398 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1409 = F_gintuple_get_key(m, v1210, v1385, v1199+int32(1088)+v1394)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v1414 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1336)+4)))
	if v1414 == int32(-1) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1199+int32(2112)+v1394<<(uint(int32(2))%32)))) = v1409
	v1412 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1397))) = uint8(v1412)
	goto L210
L212:
	;
	v1468 = base.B2i32(v1464 < int32(0))
	if v1464 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L213:
	;
	v1442 = int32(1)
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336)+5)))
	if v1443 == v1442 {
		goto L221
	} else {
		goto L222
	}
L214:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+16))
	if v1417 != int32(2) {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1336)+20)))
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1336)))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1199+int32(2112)+v1394<<(uint(int32(2))%32))))
	v1438 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1199+int32(1088)+v1394))))
	v1439 = F_ginCompareEntries(m, v1210, v1427, v1428, v1414, v1434, v1438)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199+int32(1088)+v1394))))
	if v1424 != int32(3) {
		goto L213
	} else {
		goto L218
	}
L218:
	;
	v1464 = int32(-1)
	goto L212
L219:
	;
	if v1439 != 0 {
		v1464 = v1439
		goto L212
	} else {
		goto L220
	}
L220:
	;
	goto L213
L221:
	;
	v1446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+42)))
	v1453 = F_matchPartialInPendingList(m, v1210, v1251, v1378, v1446, v1336, v1199+int32(2112), v1199+int32(1088), v1199-int32(-64))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L224
	}
L222:
	;
	v1455 = v1442
	goto L223
L223:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1456+v1303))) = uint8(v1455)
	v1459 = int32(_a_F_gingetbitmap_8)
	if base.Ui32(v1343&v1459) < base.Ui32(v1342&v1459) {
		goto L195
	} else {
		goto L225
	}
L224:
	;
	v1455 = v1453
	goto L223
L225:
	;
	v1483 = v1342
	goto L196
L226:
	;
	v1469 = v1343
	goto L228
L227:
	;
	v1469 = v1376 + int32(1)
	goto L228
L228:
	;
	if v1464 < int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1470 = v1376
	goto L231
L230:
	;
	v1470 = v1342
	goto L231
L231:
	;
	v1471 = v1470
	v1472 = v1469
	goto L201
L232:
	;
	goto L199
L233:
	;
	v1517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+42)))
	v1524 = F_matchPartialInPendingList(m, v1210, v1251, v1483&int32(_a_F_gingetbitmap_8), v1517, v1336, v1199+int32(2112), v1199+int32(1088), v1199-int32(-64))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1526+v1303))) = uint8(v1524)
	goto L195
L235:
	;
	goto L191
L236:
	;
	goto L186
L237:
	;
	v1678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1209)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1199)+60)) = uint16(v1678)
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1209)))
	*(*int32)(unsafe.Add(mBase, uint32(v1199)+56)) = v1680
	v1684 = F_scanGetCandidate(m, v1194, v1199+int32(36))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	if v1684 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1687 = v1199 + int32(56)
	v1688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1209)+2)))
	v1689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1209))))
	v1690 = int32(16)
	v1693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1687)+2)))
	v1694 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1687))))
	if v1688|v1689<<(uint(v1690)%32) == v1693|v1694<<(uint(v1690)%32) {
		goto L244
	} else {
		goto L245
	}
L240:
	;
	goto L241
L241:
	;
	goto L173
L242:
	;
	if v1704 != 0 {
		goto L172
	} else {
		goto L248
	}
L243:
	;
	goto L242
L244:
	;
	v1700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1209)+4)))
	v1701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1687)+4)))
	if v1700 == v1701 {
		v1704 = int32(1)
		goto L243
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v1704 = int32(0)
	goto L243
L247:
	;
	goto L246
L248:
	;
	goto L241
L249:
	;
	F_errmsg_internal(m, int32(_a_F_gingetbitmap_10), int32(0))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_11), int32(1814), int32(_a_F_gingetbitmap_12))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	v1933 = F_scanGetCandidate(m, v1194, v1199+int32(36))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L1
	} else {
		goto L276
	}
L253:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+52))
	v1723 = int32(0)
	goto L256
L254:
	;
	goto L255
L255:
	;
	v1796 = int32(0)
	v1797 = int32(_a_F_gingetbitmap_1)
	v1798 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1800
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+uint32(_c_F_gingetbitmap[3])))
	if v1803 != 0 {
		goto L263
	} else {
		goto L264
	}
L256:
	;
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1723+v1719))))
	if v1752 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	goto L255
L258:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+uint32(_c_F_gingetbitmap[4])))
	v1759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1755+v1723*int32(92))+78)))
	if v1759 != int32(1) {
		v1929 = v1223
		goto L252
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1763 = v1723 + int32(1)
	if v1763 != v1643 {
		v1723 = v1763
		goto L256
	} else {
		goto L262
	}
L261:
	;
	goto L260
L262:
	;
	goto L257
L263:
	;
	v1807 = v1796
	v1810 = v1796
	goto L266
L264:
	;
	v1865 = v1796
	goto L265
L265:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1798
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	F_MemoryContextReset(m, v1892)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L1
	} else {
		goto L274
	}
L266:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+uint32(_c_F_gingetbitmap[4])))
	v1838 = v1835 + v1807*int32(92)
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1838)+32))
	v1840 = m.T0[v1839].(func(*base.Module, int32) int32)(m, v1838)
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L268
	}
L267:
	;
	v1865 = v1854
	goto L265
L268:
	;
	if v1840 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1798
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	F_MemoryContextReset(m, v1846)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1838)+87)))
	v1850 = int32(1)
	v1854 = base.B2i32(v1849|v1810&v1850 != int32(0))
	v1856 = v1807 + v1850
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v1856) < base.Ui32(v1857) {
		v1807 = v1856
		v1810 = v1854
		goto L266
	} else {
		goto L273
	}
L272:
	;
	v1929 = v1223
	goto L252
L273:
	;
	goto L267
L274:
	;
	F_tbm_add_tuples(m, v1195, v1209, int32(1), v1865)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v1929 = v1223 + int64(1)
	goto L252
L276:
	;
	if v1933 != 0 {
		v1080 = v1194
		v1081 = v1195
		v1085 = v1199
		v1091 = v1205
		v1095 = v1209
		v1109 = v1929
		goto L157
	} else {
		goto L277
	}
L277:
	;
	goto L158
L278:
	;
	v1969 = v1935
	v1970 = v1936
	v1974 = v1940
	v1998 = v1964
	goto L140
L279:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4939 = m.ExcPending
	if v4939 != 0 {
		goto L1
	} else {
		goto L688
	}
L280:
	;
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[3])))
	if v3171 != 0 {
		goto L488
	} else {
		goto L489
	}
L281:
	;
	v2005 = v2000 + int32(4)
	v2023 = int32(0)
	goto L282
L282:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[6])))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v2038+v2023<<(uint(int32(2))%32))))
	v2044 = v2042 + int32(648)
	v2046 = v2042 + int32(24)
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1969)+8))
	goto L288
L283:
	;
	if v3040 == int32(0) {
		goto L280
	} else {
		goto L479
	}
L284:
	;
	F_freeGinBtreeStack(m, v2141)
	mBase = m.M
	v3037 = m.ExcPending
	if v3037 != 0 {
		goto L1
	} else {
		goto L477
	}
L285:
	;
	v3001 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	F_LockBuffer(m, v3001, int32(0))
	mBase = m.M
	v3004 = m.ExcPending
	if v3004 != 0 {
		goto L1
	} else {
		goto L476
	}
L286:
	;
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1100))
	v2835 = m.T0[v2834].(func(*base.Module, int32, int32) int32)(m, v1974+int32(1088), v2141)
	mBase = m.M
	v2836 = m.ExcPending
	if v2836 != 0 {
		goto L1
	} else {
		goto L446
	}
L287:
	;
	if v2792 == int32(0) {
		goto L285
	} else {
		goto L442
	}
L288:
	;
	v2079 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2042)+652)) = uint16(v2079)
	*(*uint16)(unsafe.Add(mBase, uint32(v2046)+8)) = uint16(v2079)
	*(*int64)(unsafe.Add(mBase, uint32(v2046))) = int64(0)
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+644))
	if v2085 != 0 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+36))
	v2792 = v2789
	goto L287
L290:
	;
	F_pfree(m, v2085)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L1
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v2088 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+648)) = v2088
	*(*int64)(unsafe.Add(mBase, uint32(v2042)+640)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+36)) = v2088
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+656)) = v2088
	*(*uint8)(unsafe.Add(mBase, uint32(v2042)+655)) = uint8(v2088)
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+44)) = int32(-1)
	v2101 = v1974 + int32(1088)
	v2102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2042)+20)))
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2042)))
	v2104 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2042)+4)))
	base.MemoryFill(m, v2101, v2088, int32(68))
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v2005)))
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+48)) = v2005
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+40)) = v2109
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+32)) = int32(43)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+24)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+20)) = int32(45)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+16)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+12)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+8)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+4)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v2101))) = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(v2101)+60)) = uint8(v2104)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+56)) = v2103
	*(*uint16)(unsafe.Add(mBase, uint32(v2101)+54)) = uint16(v2102)
	*(*uint16)(unsafe.Add(mBase, uint32(v2101)+52)) = uint16(v2088)
	*(*uint8)(unsafe.Add(mBase, uint32(v2101)+36)) = uint8(v2088)
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+28)) = int32(51)
	goto L294
L293:
	;
	goto L292
L294:
	;
	v2141 = F_ginFindLeafPage(m, v2101, int32(1), int32(0))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L1
	} else {
		goto L296
	}
L295:
	;
	v2162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2042)+654)) = uint8(v2162)
	v2164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2042)+5)))
	if v2164 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L296:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2143 < int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v2147+(v2143^int32(-1))<<(uint(int32(2))%32))))
	v2161 = v2153
	goto L295
L298:
	;
	goto L299
L299:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2161 = v2155 + v2143<<(uint(int32(13))%32) + int32(-8192)
	goto L295
L300:
	;
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2042)+4)))
	if v2167 != int32(255) {
		goto L286
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1100))
	v2173 = m.T0[v2172].(func(*base.Module, int32, int32) int32)(m, v1974+int32(1088), v2141)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L304
	}
L303:
	;
	goto L302
L304:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[11]))
	v2180 = F_tbm_create(m, v2176<<(uint(int32(10))%32), int32(0))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+36)) = v2180
	v2183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2042)+5)))
	if v2183 == int32(1) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v2186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2042)+4)))
	if v2186 != 0 {
		v2792 = v2180
		goto L287
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1136))
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(v2187)+8))
	v2189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2042)+20)))
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2191 < int32(0) {
		goto L311
	} else {
		goto L312
	}
L309:
	;
	goto L308
L310:
	;
	F_PredicateLockPage(m, v2190, v2210, v2047)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L1
	} else {
		goto L314
	}
L311:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v2195+(v2191^int32(-1))<<(uint(int32(6))%32))+16))
	v2210 = v2201
	goto L310
L312:
	;
	goto L313
L313:
	;
	v2203 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2203+v2191<<(uint(int32(6))%32)+int32(-64))+16))
	v2210 = v2209
	goto L310
L314:
	;
	v2214 = v2189 - int32(1)
	v2219 = v2188 + v2214<<(uint(int32(4))%32) + int32(20)
	goto L316
L315:
	;
	goto L289
L316:
	;
	v2253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)))
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2254 < int32(0) {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	v2765 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+36))
	if v2765 != 0 {
		goto L432
	} else {
		goto L433
	}
L318:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1136))
	v2341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)))
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2339+v2341<<(uint(int32(2))%32))+20))
	v2348 = v2339 + v2345&int32(_a_F_gingetbitmap_9)
	v2349 = F_gintuple_get_attrnum(m, v2340, v2348)
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L1
	} else {
		goto L339
	}
L319:
	;
	v2273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2272)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2273) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v2258+(v2254^int32(-1))<<(uint(int32(2))%32))))
	v2272 = v2264
	goto L319
L321:
	;
	goto L322
L322:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2272 = v2266 + v2254<<(uint(int32(13))%32) + int32(-8192)
	goto L319
L323:
	;
	v2281 = int32(base.Ui32(v2273+int32(_a_F_gingetbitmap_13)) >> (uint(int32(2)) % 32))
	goto L325
L324:
	;
	v2281 = int32(0)
	goto L325
L325:
	;
	if base.Ui32(v2281&int32(_a_F_gingetbitmap_8)) < base.Ui32(v2253) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2272)+16)))
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2272+v2285)))
	if v2287 == int32(-1) {
		goto L315
	} else {
		goto L329
	}
L327:
	;
	v2321 = v2254
	goto L328
L328:
	;
	if v2321 < int32(0) {
		goto L336
	} else {
		goto L337
	}
L329:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	v2292 = F_ginStepRight(m, v2254, v2290, int32(1))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2141)+4)) = v2292
	if v2292 < int32(0) {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	v2314 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)) = uint16(v2314)
	*(*int32)(unsafe.Add(mBase, uint32(v2141))) = v2313
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	F_PredicateLockPage(m, v2317, v2313, v2047)
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L1
	} else {
		goto L335
	}
L332:
	;
	v2298 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2298+(v2292^int32(-1))<<(uint(int32(6))%32))+16))
	v2313 = v2304
	goto L331
L333:
	;
	goto L334
L334:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2306+v2292<<(uint(int32(6))%32)+int32(-64))+16))
	v2313 = v2312
	goto L331
L335:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	v2321 = v2320
	goto L328
L336:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v2325+(v2321^int32(-1))<<(uint(int32(2))%32))))
	v2339 = v2331
	goto L318
L337:
	;
	goto L338
L338:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2339 = v2333 + v2321<<(uint(int32(13))%32) + int32(-8192)
	goto L318
L339:
	;
	if v2349 != v2189 {
		goto L315
	} else {
		goto L340
	}
L340:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1136))
	v2355 = F_gintuple_get_key(m, v2352, v2348, v1974-int32(-64))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2042)+5)))
	if v2357 == int32(1) {
		goto L345
	} else {
		goto L346
	}
L342:
	;
	goto L317
L343:
	;
	v2761 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)))
	v2763 = v2761 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)) = uint16(v2763)
	goto L316
L344:
	;
	v2387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2348)+4)))
	if v2387 == int32(_a_F_gingetbitmap_8) {
		goto L354
	} else {
		goto L355
	}
L345:
	;
	v2360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1974)+64)))
	if v2360 != 0 {
		goto L315
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+16))
	if v2380 != int32(2) {
		goto L344
	} else {
		goto L352
	}
L348:
	;
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1136))
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v2361+v2214<<(uint(int32(2))%32))+uint32(_c_F_gingetbitmap[14])))
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2042)))
	v2372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2042)+12)))
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+8))
	v2374 = F_FunctionCall4Coll(m, v2361+v2214*int32(28)+int32(_a_F_gingetbitmap_14), v2370, v2371, v2355, v2372, v2373)
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	if int32(0) < v2374 {
		goto L315
	} else {
		goto L350
	}
L350:
	;
	if int32(0) <= v2374 {
		goto L344
	} else {
		goto L351
	}
L351:
	;
	goto L343
L352:
	;
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1974)+64)))
	if v2383 == int32(3) {
		goto L315
	} else {
		goto L353
	}
L353:
	;
	goto L344
L354:
	;
	v2390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2348)+2)))
	v2391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2348))))
	v2394 = v2390 | v2391<<(uint(int32(16))%32)
	v2395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1974)+64)))
	if v2395 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	goto L356
L356:
	;
	v2717 = F_ginReadTuple(m, v2348, v1974+int32(2112))
	mBase = m.M
	v2718 = m.ExcPending
	if v2718 != 0 {
		goto L1
	} else {
		goto L429
	}
L357:
	;
	v2398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2219)+6)))
	v2399 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2219)+4)))
	v2400 = F_datumCopy(m, v2355, v2398, v2399)
	mBase = m.M
	v2401 = m.ExcPending
	if v2401 != 0 {
		goto L1
	} else {
		goto L360
	}
L358:
	;
	v2402 = v2355
	goto L359
L359:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	F_LockBuffer(m, v2403, int32(0))
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L1
	} else {
		goto L361
	}
L360:
	;
	v2402 = v2400
	goto L359
L361:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	F_PredicateLockPage(m, v2407, v2394, v2047)
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	v2413 = F_ginScanBeginPostingTree(m, v1974+int32(2112), v2412, v2394)
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v2415 = *(*int32)(unsafe.Add(mBase, uint32(v2413)+4))
	F_IncrBufferRefCount(m, v2415)
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	F_freeGinBtreeStack(m, v2413)
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v2426 = v2415
	goto L366
L366:
	;
	if v2426 < int32(0) {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	F_UnlockReleaseBuffer(m, v2426)
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L388
	}
L368:
	;
	v2469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2468)+16)))
	v2471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2469+v2468)+6)))
	if v2471&int32(4) != 0 {
		goto L372
	} else {
		goto L373
	}
L369:
	;
	v2454 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2454+(v2426^int32(-1))<<(uint(int32(2))%32))))
	v2468 = v2460
	goto L368
L370:
	;
	goto L371
L371:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2468 = v2462 + v2426<<(uint(int32(13))%32) + int32(-8192)
	goto L368
L372:
	;
	v2522 = v2469
	goto L374
L373:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+36))
	v2475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2468)+16)))
	v2476 = v2468 + v2475
	v2477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476)+6)))
	if v2477&int32(128) != 0 {
		goto L376
	} else {
		goto L377
	}
L374:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v2522+v2468)))
	if v2524 != int32(-1) {
		goto L384
	} else {
		goto L385
	}
L375:
	;
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+656))
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+656)) = v2514 + v2515
	v2518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2468)+16)))
	v2522 = v2518
	goto L374
L376:
	;
	v2480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2468)+12)))
	v2481 = int32(32)
	v2483 = m.G0
	v2485 = v2483 - int32(16)
	m.G0 = v2485
	v2491 = F_ginPostingListDecodeAllSegments(m, v2468+v2481, v2480-v2481, v2485+int32(12))
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L1
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	v2502 = int32(0)
	v2503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2476)+4)))
	if v2503 == v2502 {
		v2514 = v2502
		goto L375
	} else {
		goto L382
	}
L379:
	;
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(v2485)+12))
	F_tbm_add_tuples(m, v2474, v2491, v2493, int32(0))
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	F_pfree(m, v2491)
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	m.G0 = v2485 + int32(16)
	v2514 = v2493
	goto L375
L382:
	;
	F_tbm_add_tuples(m, v2474, v2468+int32(32), v2503, int32(0))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v2514 = v2503
	goto L375
L384:
	;
	v2528 = F_ginStepRight(m, v2426, v2412, int32(1))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	goto L367
L387:
	;
	v2426 = v2528
	goto L366
L388:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	F_LockBuffer(m, v2532, int32(1))
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2536 < int32(0) {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v2555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2554)+16)))
	v2557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2555+v2554)+6)))
	if v2557&int32(2) == int32(0) {
		goto L342
	} else {
		goto L394
	}
L391:
	;
	v2540 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v2540+(v2536^int32(-1))<<(uint(int32(2))%32))))
	v2554 = v2546
	goto L390
L392:
	;
	goto L393
L393:
	;
	v2548 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2554 = v2548 + v2536<<(uint(int32(13))%32) + int32(-8192)
	goto L390
L394:
	;
	v2564 = v2536
	goto L395
L395:
	;
	v2593 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)))
	if v2564 < int32(0) {
		goto L399
	} else {
		goto L400
	}
L396:
	;
	v2709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1974)+64)))
	if v2709 != 0 {
		goto L343
	} else {
		goto L426
	}
L397:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1136))
	v2680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)))
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(v2678+v2680<<(uint(int32(2))%32))+20))
	v2687 = v2678 + v2684&int32(_a_F_gingetbitmap_9)
	v2688 = F_gintuple_get_attrnum(m, v2679, v2687)
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L1
	} else {
		goto L419
	}
L398:
	;
	v2612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2611)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2612) {
		goto L402
	} else {
		goto L403
	}
L399:
	;
	v2597 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(v2597+(v2564^int32(-1))<<(uint(int32(2))%32))))
	v2611 = v2603
	goto L398
L400:
	;
	goto L401
L401:
	;
	v2605 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2611 = v2605 + v2564<<(uint(int32(13))%32) + int32(-8192)
	goto L398
L402:
	;
	v2620 = int32(base.Ui32(v2612+int32(_a_F_gingetbitmap_13)) >> (uint(int32(2)) % 32))
	goto L404
L403:
	;
	v2620 = int32(0)
	goto L404
L404:
	;
	if base.Ui32(v2620&int32(_a_F_gingetbitmap_8)) < base.Ui32(v2593) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v2624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2611)+16)))
	v2626 = *(*int32)(unsafe.Add(mBase, uint32(v2611+v2624)))
	if v2626 == int32(-1) {
		goto L279
	} else {
		goto L408
	}
L406:
	;
	v2660 = v2564
	goto L407
L407:
	;
	if v2660 < int32(0) {
		goto L415
	} else {
		goto L416
	}
L408:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	v2631 = F_ginStepRight(m, v2564, v2629, int32(1))
	mBase = m.M
	v2632 = m.ExcPending
	if v2632 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2141)+4)) = v2631
	if v2631 < int32(0) {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	v2653 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)) = uint16(v2653)
	*(*int32)(unsafe.Add(mBase, uint32(v2141))) = v2652
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	F_PredicateLockPage(m, v2656, v2652, v2047)
	mBase = m.M
	v2658 = m.ExcPending
	if v2658 != 0 {
		goto L1
	} else {
		goto L414
	}
L411:
	;
	v2637 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2637+(v2631^int32(-1))<<(uint(int32(6))%32))+16))
	v2652 = v2643
	goto L410
L412:
	;
	goto L413
L413:
	;
	v2645 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2645+v2631<<(uint(int32(6))%32)+int32(-64))+16))
	v2652 = v2651
	goto L410
L414:
	;
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	v2660 = v2659
	goto L407
L415:
	;
	v2664 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2664+(v2660^int32(-1))<<(uint(int32(2))%32))))
	v2678 = v2670
	goto L397
L416:
	;
	goto L417
L417:
	;
	v2672 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2678 = v2672 + v2660<<(uint(int32(13))%32) + int32(-8192)
	goto L397
L418:
	;
	goto L396
L419:
	;
	if v2688 == v2189 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1136))
	v2694 = F_gintuple_get_key(m, v2691, v2687, v1974+int32(2112))
	mBase = m.M
	v2695 = m.ExcPending
	if v2695 != 0 {
		goto L1
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v2704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)))
	v2706 = v2704 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)) = uint16(v2706)
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	v2564 = v2708
	goto L395
L423:
	;
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1136))
	v2697 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1974)+2112)))
	v2698 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1974)+64)))
	v2699 = F_ginCompareEntries(m, v2696, v2189, v2694, v2697, v2402, v2698)
	mBase = m.M
	v2700 = m.ExcPending
	if v2700 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	if v2699 == int32(0) {
		goto L418
	} else {
		goto L425
	}
L425:
	;
	goto L422
L426:
	;
	v2710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2219)+6)))
	if v2710 != 0 {
		goto L343
	} else {
		goto L427
	}
L427:
	;
	F_pfree(m, v2402)
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	goto L343
L429:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+36))
	v2720 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+2112))
	F_tbm_add_tuples(m, v2719, v2717, v2720, int32(0))
	mBase = m.M
	v2723 = m.ExcPending
	if v2723 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+656))
	v2725 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2348)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+656)) = v2724 + v2725
	F_pfree(m, v2717)
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	goto L343
L432:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+40))
	if v2766 != 0 {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	v2780 = v2536
	goto L434
L434:
	;
	F_LockBuffer(m, v2780, int32(0))
	mBase = m.M
	v2783 = m.ExcPending
	if v2783 != 0 {
		goto L1
	} else {
		goto L440
	}
L435:
	;
	F_pfree(m, v2766)
	mBase = m.M
	v2768 = m.ExcPending
	if v2768 != 0 {
		goto L1
	} else {
		goto L438
	}
L436:
	;
	v2770 = v2765
	goto L437
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+40)) = int32(0)
	F_tbm_free(m, v2770)
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		goto L1
	} else {
		goto L439
	}
L438:
	;
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+36))
	v2770 = v2769
	goto L437
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+36)) = int32(0)
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	v2780 = v2777
	goto L434
L440:
	;
	F_freeGinBtreeStack(m, v2141)
	mBase = m.M
	v2785 = m.ExcPending
	if v2785 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	goto L288
L442:
	;
	v2823 = *(*int32)(unsafe.Add(mBase, uint32(v2792)+16))
	goto L443
L443:
	;
	if v2823 == int32(0) {
		goto L285
	} else {
		goto L444
	}
L444:
	;
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+36))
	v2827 = F_tbm_begin_private_iterate(m, v2826)
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	v2829 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2042)+654)) = uint8(v2829)
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+40)) = v2827
	goto L285
L446:
	;
	if v2835 != 0 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2141)+8)))
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v2161+v2837<<(uint(int32(2))%32))+20))
	v2844 = v2161 + v2841&int32(_a_F_gingetbitmap_9)
	v2845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2844)+4)))
	if v2845 == int32(_a_F_gingetbitmap_8) {
		goto L450
	} else {
		goto L451
	}
L448:
	;
	goto L449
L449:
	;
	v2947 = *(*int32)(unsafe.Add(mBase, uint32(v2005)))
	v2948 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2948 < int32(0) {
		goto L472
	} else {
		goto L473
	}
L450:
	;
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v2005)))
	v2849 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2844)+2)))
	v2850 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2844))))
	v2853 = v2849 | v2850<<(uint(int32(16))%32)
	F_PredicateLockPage(m, v2848, v2853, v2047)
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L1
	} else {
		goto L453
	}
L451:
	;
	goto L452
L452:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v2005)))
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	if v2914 < int32(0) {
		goto L465
	} else {
		goto L466
	}
L453:
	;
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2141)+4))
	F_LockBuffer(m, v2856, int32(0))
	mBase = m.M
	v2859 = m.ExcPending
	if v2859 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	v2862 = *(*int32)(unsafe.Add(mBase, uint32(v2005)))
	v2863 = F_ginScanBeginPostingTree(m, v2042+int32(660), v2862, v2853)
	mBase = m.M
	v2864 = m.ExcPending
	if v2864 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(v2863)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+24)) = v2865
	F_IncrBufferRefCount(m, v2865)
	mBase = m.M
	v2868 = m.ExcPending
	if v2868 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+24))
	if v2869 < int32(0) {
		goto L458
	} else {
		goto L459
	}
L457:
	;
	v2888 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1974)+2116)) = uint16(v2888)
	*(*int32)(unsafe.Add(mBase, uint32(v1974)+2112)) = v2888
	*(*int32)(unsafe.Add(mBase, uint32(v1974)+28)) = v2888
	*(*uint16)(unsafe.Add(mBase, uint32(v1974)+32)) = uint16(v2888)
	v2898 = F_GinDataLeafPageGetItems(m, v2887, v2044, v1974+int32(28))
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L1
	} else {
		goto L461
	}
L458:
	;
	v2873 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2879 = *(*int32)(unsafe.Add(mBase, uint32(v2873+(v2869^int32(-1))<<(uint(int32(2))%32))))
	v2887 = v2879
	goto L457
L459:
	;
	goto L460
L460:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2887 = v2881 + v2869<<(uint(int32(13))%32) + int32(-8192)
	goto L457
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+644)) = v2898
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+648))
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2863)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+656)) = v2901 * v2902
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+24))
	F_LockBuffer(m, v2905, int32(0))
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	F_freeGinBtreeStack(m, v2863)
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L1
	} else {
		goto L463
	}
L463:
	;
	v2911 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2042)+654)) = uint8(v2911)
	goto L284
L464:
	;
	F_PredicateLockPage(m, v2913, v2933, v2047)
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L468
	}
L465:
	;
	v2918 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2918+(v2914^int32(-1))<<(uint(int32(6))%32))+16))
	v2933 = v2924
	goto L464
L466:
	;
	goto L467
L467:
	;
	v2926 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2932 = *(*int32)(unsafe.Add(mBase, uint32(v2926+v2914<<(uint(int32(6))%32)+int32(-64))+16))
	v2933 = v2932
	goto L464
L468:
	;
	v2936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2844)+4)))
	if v2936 == int32(0) {
		goto L285
	} else {
		goto L469
	}
L469:
	;
	v2940 = F_ginReadTuple(m, v2844, v2044)
	mBase = m.M
	v2941 = m.ExcPending
	if v2941 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+644)) = v2940
	v2943 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2042)+654)) = uint8(v2943)
	v2945 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+648))
	*(*int32)(unsafe.Add(mBase, uint32(v2042)+656)) = v2945
	goto L285
L471:
	;
	F_PredicateLockPage(m, v2947, v2967, v2047)
	mBase = m.M
	v2969 = m.ExcPending
	if v2969 != 0 {
		goto L1
	} else {
		goto L475
	}
L472:
	;
	v2952 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v2952+(v2948^int32(-1))<<(uint(int32(6))%32))+16))
	v2967 = v2958
	goto L471
L473:
	;
	goto L474
L474:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2966 = *(*int32)(unsafe.Add(mBase, uint32(v2960+v2948<<(uint(int32(6))%32)+int32(-64))+16))
	v2967 = v2966
	goto L471
L475:
	;
	goto L285
L476:
	;
	goto L284
L477:
	;
	v3039 = v2023 + int32(1)
	v3040 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[2])))
	if base.Ui32(v3039) < base.Ui32(v3040) {
		v2023 = v3039
		goto L282
	} else {
		goto L478
	}
L478:
	;
	goto L283
L479:
	;
	v3045 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[15]))
	if v3045 <= int32(0) {
		goto L280
	} else {
		goto L480
	}
L480:
	;
	v3049 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[6])))
	v3053 = int32(0)
	goto L481
L481:
	;
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v3049+v3053<<(uint(int32(2))%32))))
	v3086 = *(*int32)(unsafe.Add(mBase, uint32(v3085)+656))
	if base.Ui32(v3086) <= base.Ui32(v3045*v3040) {
		goto L280
	} else {
		goto L483
	}
L482:
	;
	v3094 = int32(0)
	v3095 = v3040
	goto L485
L483:
	;
	v3089 = v3053 + int32(1)
	if v3089 != v3040 {
		v3053 = v3089
		goto L481
	} else {
		goto L484
	}
L484:
	;
	goto L482
L485:
	;
	v3124 = v3094 << (uint(int32(2)) % 32)
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[6])))
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(v3124+v3125)))
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v3127)+656))
	v3129 = base.I32_div_u_s(v3128, v3095)
	*(*int32)(unsafe.Add(mBase, uint32(v3127)+656)) = v3129
	v3131 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[6])))
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v3131+v3124)))
	v3134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3133)+655)) = uint8(v3134)
	v3137 = v3094 + v3134
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[2])))
	if base.Ui32(v3137) < base.Ui32(v3138) {
		v3094 = v3137
		v3095 = v3138
		goto L485
	} else {
		goto L487
	}
L486:
	;
	goto L280
L487:
	;
	goto L486
L488:
	;
	v3173 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v3178 = int32(0)
	goto L491
L489:
	;
	goto L490
L490:
	;
	v3805 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1974)+68)) = uint16(v3805)
	*(*int32)(unsafe.Add(mBase, uint32(v1974)+64)) = v3805
	v3809 = v1969
	v3810 = v1970
	v3814 = v1974
	v3838 = v1998
	goto L549
L491:
	;
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[4])))
	v3209 = v3206 + v3178*int32(92)
	v3210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3209)+88)) = uint8(v3210)
	*(*int64)(unsafe.Add(mBase, uint32(v3209)+80)) = int64(0)
	v3214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3209)+78)))
	if v3214 == int32(1) {
		goto L494
	} else {
		goto L495
	}
L492:
	;
	goto L490
L493:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3173
	v3771 = v3178 + int32(1)
	v3772 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v3771) < base.Ui32(v3772) {
		v3178 = v3771
		goto L491
	} else {
		goto L548
	}
L494:
	;
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3218
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+16)) = int32(0)
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+24)) = v3222
	v3226 = F_palloc(m, v3222<<(uint(int32(2))%32))
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L1
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	if base.Ui32(int32(2)) <= base.Ui32(v3276) {
		goto L502
	} else {
		goto L503
	}
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+20)) = v3226
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+24))
	if v3229 <= int32(0) {
		goto L493
	} else {
		goto L498
	}
L498:
	;
	v3235 = int32(0)
	goto L499
L499:
	;
	v3265 = v3235 << (uint(int32(2)) % 32)
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+20))
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+8))
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(v3268+v3265)))
	*(*int32)(unsafe.Add(mBase, uint32(v3265+v3266))) = v3270
	v3273 = v3235 + int32(1)
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+24))
	if v3273 < v3274 {
		v3235 = v3273
		goto L499
	} else {
		goto L501
	}
L500:
	;
	goto L493
L501:
	;
	goto L500
L502:
	;
	v3281 = *(*int32)(unsafe.Add(mBase, uint32(v2000)))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3281
	v3283 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	v3286 = F_palloc(m, v3283<<(uint(int32(2))%32))
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		goto L1
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v3724 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3724
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+16)) = int32(1)
	v3731 = F_palloc(m, int32(4))
	mBase = m.M
	v3732 = m.ExcPending
	if v3732 != 0 {
		goto L1
	} else {
		goto L547
	}
L505:
	;
	v3289 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	if v3289 != 0 {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v3292 = int32(0)
	goto L509
L507:
	;
	v3333 = int32(0)
	goto L508
L508:
	;
	F_qsort_arg(m, v3286, v3333, int32(4), int32(52), v3209)
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
		goto L1
	} else {
		goto L512
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3286+v3292<<(uint(int32(2))%32)))) = v3292
	v3326 = v3292 + int32(1)
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	if base.Ui32(v3326) < base.Ui32(v3327) {
		v3292 = v3326
		goto L509
	} else {
		goto L511
	}
L510:
	;
	v3333 = v3327
	goto L508
L511:
	;
	goto L510
L512:
	;
	v3364 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	if base.Ui32(int32(2)) <= base.Ui32(v3364) {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v3370 = int32(1)
	goto L516
L514:
	;
	v3415 = v3364
	goto L515
L515:
	;
	v3442 = int32(1)
	if v3415 != v3442 {
		goto L519
	} else {
		goto L520
	}
L516:
	;
	v3399 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+28))
	v3400 = int32(2)
	v3403 = *(*int32)(unsafe.Add(mBase, uint32(v3286+v3370<<(uint(v3400)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3399+v3403))) = uint8(v3400)
	v3408 = v3370 + int32(1)
	v3409 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	if base.Ui32(v3408) < base.Ui32(v3409) {
		v3370 = v3408
		goto L516
	} else {
		goto L518
	}
L517:
	;
	v3415 = v3409
	goto L515
L518:
	;
	goto L517
L519:
	;
	v3448 = int32(0)
	goto L522
L520:
	;
	v3511 = v3442
	goto L521
L521:
	;
	v3534 = int32(0)
	v3536 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3536
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+16)) = v3511
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+24)) = v3539 - v3511
	v3544 = F_palloc(m, v3511<<(uint(int32(2))%32))
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
		goto L1
	} else {
		goto L532
	}
L522:
	;
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+28))
	v3481 = *(*int32)(unsafe.Add(mBase, uint32(v3286+v3448<<(uint(int32(2))%32))))
	v3483 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3477+v3481))) = uint8(v3483)
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+36))
	v3486 = m.T0[v3485].(func(*base.Module, int32) int32)(m, v3209)
	mBase = m.M
	v3487 = m.ExcPending
	if v3487 != 0 {
		goto L1
	} else {
		goto L525
	}
L523:
	;
	v3511 = v3500 + int32(1)
	goto L521
L524:
	;
	goto L523
L525:
	;
	if v3486 == int32(0) {
		v3500 = v3448
		goto L524
	} else {
		goto L526
	}
L526:
	;
	v3491 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[16]))
	if v3491 != 0 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3493 = m.ExcPending
	if v3493 != 0 {
		goto L1
	} else {
		goto L530
	}
L528:
	;
	goto L529
L529:
	;
	v3494 = int32(1)
	v3495 = v3448 + v3494
	v3496 = *(*int32)(unsafe.Add(mBase, uint32(v3209)))
	if base.Ui32(v3495) < base.Ui32(v3496-v3494) {
		v3448 = v3495
		goto L522
	} else {
		goto L531
	}
L530:
	;
	goto L529
L531:
	;
	v3500 = v3495
	goto L524
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+12)) = v3544
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+24))
	v3550 = F_palloc(m, v3547<<(uint(int32(2))%32))
	mBase = m.M
	v3551 = m.ExcPending
	if v3551 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+20)) = v3550
	v3553 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+16))
	if int32(0) < v3553 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v3558 = v3534
	goto L537
L535:
	;
	v3605 = v3534
	goto L536
L536:
	;
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+24))
	if int32(0) < v3634 {
		goto L540
	} else {
		goto L541
	}
L537:
	;
	v3587 = int32(2)
	v3588 = v3558 << (uint(v3587) % 32)
	v3589 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+12))
	v3591 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+8))
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v3588+v3286)))
	v3597 = *(*int32)(unsafe.Add(mBase, uint32(v3591+v3593<<(uint(v3587)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3588+v3589))) = v3597
	v3600 = v3558 + int32(1)
	v3601 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+16))
	if v3600 < v3601 {
		v3558 = v3600
		goto L537
	} else {
		goto L539
	}
L538:
	;
	v3605 = v3600
	goto L536
L539:
	;
	goto L538
L540:
	;
	v3640 = v3605
	v3642 = int32(0)
	goto L543
L541:
	;
	goto L542
L542:
	;
	v3720 = *(*int32)(unsafe.Add(mBase, uint32(v2000)))
	F_MemoryContextReset(m, v3720)
	mBase = m.M
	v3722 = m.ExcPending
	if v3722 != 0 {
		goto L1
	} else {
		goto L546
	}
L543:
	;
	v3669 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+20))
	v3670 = int32(2)
	v3673 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+8))
	v3677 = *(*int32)(unsafe.Add(mBase, uint32(v3286+v3640<<(uint(v3670)%32))))
	v3681 = *(*int32)(unsafe.Add(mBase, uint32(v3673+v3677<<(uint(v3670)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3669+v3642<<(uint(v3670)%32)))) = v3681
	v3683 = int32(1)
	v3686 = v3642 + v3683
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+24))
	if v3686 < v3687 {
		v3640 = v3640 + v3683
		v3642 = v3686
		goto L543
	} else {
		goto L545
	}
L544:
	;
	goto L542
L545:
	;
	goto L544
L546:
	;
	goto L493
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3209)+12)) = v3731
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+8))
	v3735 = *(*int32)(unsafe.Add(mBase, uint32(v3734)))
	*(*int32)(unsafe.Add(mBase, uint32(v3731))) = v3735
	goto L493
L548:
	;
	goto L492
L549:
	;
	v3840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+68)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1092)) = uint16(v3840)
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v3814)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v3814)+1088)) = v3842
	v3844 = *(*int32)(unsafe.Add(mBase, uint32(v3809)+36))
	v3846 = v3844 + int32(4)
	goto L552
L551:
	;
	F_tbm_add_tuples(m, v3810, v3814-int32(-64), int32(1), v4904)
	mBase = m.M
	v4933 = m.ExcPending
	if v4933 != 0 {
		goto L1
	} else {
		goto L687
	}
L552:
	;
	v3879 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[16]))
	if v3879 != 0 {
		goto L554
	} else {
		goto L555
	}
L553:
	;
	if v4804 == int32(0) {
		goto L677
	} else {
		goto L678
	}
L554:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3881 = m.ExcPending
	if v3881 != 0 {
		goto L1
	} else {
		goto L557
	}
L555:
	;
	goto L556
L556:
	;
	v3882 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+68)) = uint16(v3882)
	*(*int32)(unsafe.Add(mBase, uint32(v3814)+64)) = v3882
	v3886 = *(*int32)(unsafe.Add(mBase, uint32(v3844)+uint32(_c_F_gingetbitmap[3])))
	if v3886 == v3882 {
		goto L558
	} else {
		goto L559
	}
L557:
	;
	goto L556
L558:
	;
	v4904 = int32(0)
	goto L551
L559:
	;
	goto L560
L560:
	;
	v3891 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1090)))
	v3892 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1088)))
	v3900 = v3891
	v3903 = v3892
	v3910 = int32(0)
	goto L561
L561:
	;
	v3924 = *(*int32)(unsafe.Add(mBase, uint32(v3844)+uint32(_c_F_gingetbitmap[4])))
	v3927 = v3924 + v3910*int32(92)
	v3928 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+68)))
	if v3928 != int32(_a_F_gingetbitmap_8) {
		goto L564
	} else {
		goto L565
	}
L562:
	;
	if v4775 == int32(0) {
		goto L552
	} else {
		goto L675
	}
L563:
	;
	v4803 = v3910 + int32(1)
	v4804 = *(*int32)(unsafe.Add(mBase, uint32(v3844)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v4803) < base.Ui32(v4804) {
		goto L671
	} else {
		goto L672
	}
L564:
	;
	v3941 = *(*int32)(unsafe.Add(mBase, uint32(v3844)))
	v3942 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1092)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2116)) = uint16(v3942)
	v3944 = *(*int32)(unsafe.Add(mBase, uint32(v3814)+1088))
	*(*int32)(unsafe.Add(mBase, uint32(v3814)+2112)) = v3944
	v3947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2114)))
	v3949 = int64(32)
	v3952 = int64(48)
	v3955 = base.I64_extend_i32_u(v3942) | (base.I64_extend_i32_u(v3947)<<(uint(v3949)%64) | base.I64_extend_i32_u(v3944)<<(uint(v3952)%64))
	v3956 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3927)+84)))
	v3957 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3927)+82)))
	v3960 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3927)+80)))
	if base.Ui64(v3955) < base.Ui64(v3956|(v3957<<(uint(v3949)%64)|v3960<<(uint(v3952)%64))) {
		goto L568
	} else {
		goto L569
	}
L565:
	;
	v3931 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+64)))
	v3932 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+66)))
	if v3931&v3932 == int32(_a_F_gingetbitmap_8) {
		goto L564
	} else {
		goto L566
	}
L566:
	;
	v3937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927)+78)))
	if v3937 == int32(0) {
		goto L564
	} else {
		goto L567
	}
L567:
	;
	v4773 = v3900
	v4775 = int32(1)
	v4777 = v3903
	goto L563
L568:
	;
	v4678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927)+88)))
	if v4678 != 0 {
		v4963 = v3814
		v4987 = v3838
		goto L135
	} else {
		goto L649
	}
L569:
	;
	v3966 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+16))
	if v3966 != 0 {
		goto L574
	} else {
		goto L575
	}
L570:
	;
	v4215 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+24))
	if v4215 != 0 {
		goto L595
	} else {
		goto L596
	}
L571:
	;
	v4186 = v4073
	v4187 = v4072
	v4188 = v4073
	v4192 = v4182
	v4197 = v4072
	v4200 = v4183
	goto L570
L572:
	;
	v4186 = v3947
	v4187 = v3944
	v4188 = v3947
	v4192 = v3942 + int32(1)
	v4197 = v3944
	v4200 = v3942
	goto L570
L573:
	;
	v4120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927)+78)))
	if v4120 != 0 {
		goto L572
	} else {
		goto L590
	}
L574:
	;
	v3967 = int32(_a_F_gingetbitmap_8)
	v3977 = v3967
	v3978 = v3967
	v3980 = int32(0)
	v3982 = v3967
	v3995 = int32(1)
	goto L577
L575:
	;
	goto L576
L576:
	;
	v4117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927)+78)))
	if v4117 != 0 {
		goto L572
	} else {
		goto L589
	}
L577:
	;
	v4005 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+12))
	v4009 = *(*int32)(unsafe.Add(mBase, uint32(v4005+v3980<<(uint(int32(2))%32))))
	v4010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4009)+654)))
	if v4010 != 0 {
		v4072 = v3977
		v4073 = v3978
		v4074 = v3982
		v4078 = v3995
		goto L579
	} else {
		goto L580
	}
L578:
	;
	if v4078 == int32(0) {
		goto L573
	} else {
		goto L588
	}
L579:
	;
	v4081 = v3980 + int32(1)
	v4082 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+16))
	if base.Ui32(v4081) < base.Ui32(v4082) {
		v3977 = v4072
		v3978 = v4073
		v3980 = v4081
		v3982 = v4074
		v3995 = v4078
		goto L577
	} else {
		goto L587
	}
L580:
	;
	v4011 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4009)+32)))
	v4013 = int64(65535)
	v4015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4009)+30)))
	v4021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4009)+28)))
	v4026 = base.I64_extend_i32_u(v4011)&v4013 | (base.I64_extend_i32_u(v4015)&v4013<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4021)<<(uint(int64(48))%64))
	if base.Ui64(v4026) <= base.Ui64(v3955) {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v4028 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1092)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+12)) = uint16(v4028)
	v4030 = *(*int32)(unsafe.Add(mBase, uint32(v3814)+1088))
	*(*int32)(unsafe.Add(mBase, uint32(v3814)+8)) = v4030
	F_entryGetItem(m, v3846, v4009, v3814+int32(8))
	mBase = m.M
	v4035 = m.ExcPending
	if v4035 != 0 {
		goto L1
	} else {
		goto L584
	}
L582:
	;
	v4053 = v4015
	v4054 = v4021
	v4055 = v4011
	v4056 = v4026
	goto L583
L583:
	;
	v4057 = int32(0)
	v4059 = int64(65535)
	if base.Ui64(base.I64_extend_i32_u(v3982)&v4059|(base.I64_extend_i32_u(v3977)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v3978)&v4059<<(uint(int64(32))%64))) <= base.Ui64(v4056) {
		v4072 = v3977
		v4073 = v3978
		v4074 = v3982
		v4078 = v4057
		goto L579
	} else {
		goto L586
	}
L584:
	;
	v4036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4009)+654)))
	if v4036 != 0 {
		v4072 = v3977
		v4073 = v3978
		v4074 = v3982
		v4078 = v3995
		goto L579
	} else {
		goto L585
	}
L585:
	;
	v4037 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4009)+32)))
	v4039 = int64(65535)
	v4041 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4009)+30)))
	v4047 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4009)+28)))
	v4053 = v4041
	v4054 = v4047
	v4055 = v4037
	v4056 = base.I64_extend_i32_u(v4037)&v4039 | (base.I64_extend_i32_u(v4041)&v4039<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4047)<<(uint(int64(48))%64))
	goto L583
L586:
	;
	v4072 = v4054
	v4073 = v4053
	v4074 = v4055
	v4078 = v4057
	goto L579
L587:
	;
	goto L578
L588:
	;
	goto L576
L589:
	;
	v4118 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3927)+88)) = uint8(v4118)
	goto L568
L590:
	;
	v4121 = int32(_a_F_gingetbitmap_8)
	if v4074&v4121 != v4121 {
		goto L591
	} else {
		goto L592
	}
L591:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2114)) = uint16(v4073)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2112)) = uint16(v4072)
	v4146 = v4074 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2116)) = uint16(v4146)
	v4182 = v4074
	v4183 = v4146
	goto L571
L592:
	;
	v4129 = v4073&int32(_a_F_gingetbitmap_8) | v4072<<(uint(int32(16))%32)
	if v4129 == int32(-1) {
		goto L591
	} else {
		goto L593
	}
L593:
	;
	v4132 = int32(_a_F_gingetbitmap_8)
	if base.Ui32(v4129) <= base.Ui32(v3944&v3967<<(uint(int32(16))%32)|v3947) {
		v4186 = v3947
		v4187 = v4072
		v4188 = v4073
		v4192 = v4132
		v4197 = v3944
		v4200 = v3942
		goto L570
	} else {
		goto L594
	}
L594:
	;
	v4137 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2116)) = uint16(v4137)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2114)) = uint16(v4073)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2112)) = uint16(v4072)
	v4182 = v4132
	v4183 = v4137
	goto L571
L595:
	;
	v4217 = int64(65535)
	v4233 = v4187
	v4234 = v4188
	v4236 = int32(0)
	v4238 = v4192
	goto L598
L596:
	;
	v4339 = v4187
	v4340 = v4188
	v4344 = v4192
	goto L597
L597:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3927)+84)) = uint16(v4344)
	*(*uint16)(unsafe.Add(mBase, uint32(v3927)+82)) = uint16(v4340)
	*(*uint16)(unsafe.Add(mBase, uint32(v3927)+80)) = uint16(v4339)
	v4370 = *(*int32)(unsafe.Add(mBase, uint32(v3927)))
	if v4370 == int32(0) {
		goto L610
	} else {
		goto L611
	}
L598:
	;
	v4261 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+20))
	v4265 = *(*int32)(unsafe.Add(mBase, uint32(v4261+v4236<<(uint(int32(2))%32))))
	v4266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4265)+654)))
	if v4266 != 0 {
		v4325 = v4233
		v4326 = v4234
		v4327 = v4238
		goto L600
	} else {
		goto L601
	}
L599:
	;
	v4339 = v4325
	v4340 = v4326
	v4344 = v4327
	goto L597
L600:
	;
	v4333 = v4236 + int32(1)
	v4334 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+24))
	if base.Ui32(v4333) < base.Ui32(v4334) {
		v4233 = v4325
		v4234 = v4326
		v4236 = v4333
		v4238 = v4327
		goto L598
	} else {
		goto L608
	}
L601:
	;
	v4267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4265)+32)))
	v4269 = int64(65535)
	v4271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4265)+30)))
	v4277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4265)+28)))
	v4282 = base.I64_extend_i32_u(v4267)&v4269 | (base.I64_extend_i32_u(v4271)&v4269<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4277)<<(uint(int64(48))%64))
	if base.Ui64(v4282) <= base.Ui64(base.I64_extend_i32_u(v4200)&v4217|(base.I64_extend_i32_u(v4197)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v4186)&v4217<<(uint(int64(32))%64))) {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v4284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+2116)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+4)) = uint16(v4284)
	v4286 = *(*int32)(unsafe.Add(mBase, uint32(v3814)+2112))
	*(*int32)(unsafe.Add(mBase, uint32(v3814))) = v4286
	F_entryGetItem(m, v3846, v4265, v3814)
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L1
	} else {
		goto L605
	}
L603:
	;
	v4307 = v4271
	v4308 = v4277
	v4309 = v4267
	v4310 = v4282
	goto L604
L604:
	;
	v4312 = int64(65535)
	if base.Ui64(base.I64_extend_i32_u(v4238)&v4312|(base.I64_extend_i32_u(v4233)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v4234)&v4312<<(uint(int64(32))%64))) <= base.Ui64(v4310) {
		v4325 = v4233
		v4326 = v4234
		v4327 = v4238
		goto L600
	} else {
		goto L607
	}
L605:
	;
	v4290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4265)+654)))
	if v4290 != 0 {
		v4325 = v4233
		v4326 = v4234
		v4327 = v4238
		goto L600
	} else {
		goto L606
	}
L606:
	;
	v4291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4265)+32)))
	v4293 = int64(65535)
	v4295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4265)+30)))
	v4301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4265)+28)))
	v4307 = v4295
	v4308 = v4301
	v4309 = v4291
	v4310 = base.I64_extend_i32_u(v4291)&v4293 | (base.I64_extend_i32_u(v4295)&v4293<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4301)<<(uint(int64(48))%64))
	goto L604
L607:
	;
	v4325 = v4308
	v4326 = v4307
	v4327 = v4309
	goto L600
L608:
	;
	goto L599
L609:
	;
	v4513 = *(*int32)(unsafe.Add(mBase, uint32(v3927)))
	if v4513 != 0 {
		goto L627
	} else {
		goto L628
	}
L610:
	;
	v4373 = int32(_a_F_gingetbitmap_1)
	v4374 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3941
	v4494 = v4374
	goto L609
L611:
	;
	goto L612
L612:
	;
	v4378 = int64(65535)
	v4388 = int32(0)
	v4392 = v4388
	v4401 = v4388
	goto L613
L613:
	;
	v4421 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+8))
	v4425 = *(*int32)(unsafe.Add(mBase, uint32(v4421+v4392<<(uint(int32(2))%32))))
	v4426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4425)+654)))
	if v4426 != 0 {
		goto L616
	} else {
		goto L617
	}
L614:
	;
	v4457 = int32(_a_F_gingetbitmap_1)
	v4458 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3941
	if v4452&int32(1) == int32(0) {
		v4494 = v4458
		goto L609
	} else {
		goto L623
	}
L615:
	;
	v4454 = v4392 + int32(1)
	v4455 = *(*int32)(unsafe.Add(mBase, uint32(v3927)))
	if base.Ui32(v4454) < base.Ui32(v4455) {
		v4392 = v4454
		v4401 = v4452
		goto L613
	} else {
		goto L622
	}
L616:
	;
	v4447 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+28))
	v4449 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4447+v4392))) = uint8(v4449)
	v4452 = v4401
	goto L615
L617:
	;
	v4427 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4425)+32)))
	v4428 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4425)+30)))
	v4431 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4425)+28)))
	if v4427|(v4428<<(uint(int64(32))%64)|v4431<<(uint(int64(48))%64)) != base.I64_extend_i32_u(v4340)&v4378<<(uint(int64(32))%64)|base.I64_extend_i32_u(v4339)<<(uint(int64(48))%64)|v4378 {
		goto L616
	} else {
		goto L618
	}
L618:
	;
	v4437 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+28))
	v4438 = v4437 + v4392
	v4439 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+4))
	if base.Ui32(v4392) < base.Ui32(v4439) {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	v4441 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4438))) = uint8(v4441)
	v4452 = int32(1)
	goto L615
L620:
	;
	goto L621
L621:
	;
	v4444 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4438))) = uint8(v4444)
	v4452 = v4444
	goto L615
L622:
	;
	goto L614
L623:
	;
	v4465 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+36))
	v4466 = m.T0[v4465].(func(*base.Module, int32) int32)(m, v3927)
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	v4468 = int32(1)
	if base.Ui32(v4468) < base.Ui32((v4466-v4468)&int32(255)) {
		v4494 = v4458
		goto L609
	} else {
		goto L625
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4458
	F_MemoryContextReset(m, v3941)
	mBase = m.M
	v4477 = m.ExcPending
	if v4477 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3927)+84)) = int32(16908287)
	*(*uint16)(unsafe.Add(mBase, uint32(v3927)+82)) = uint16(v4340)
	*(*uint16)(unsafe.Add(mBase, uint32(v3927)+80)) = uint16(v4339)
	goto L568
L627:
	;
	v4515 = int64(65535)
	v4518 = base.I64_extend_i32_u(v4340) & v4515 << (uint(int64(32)) % 64)
	v4521 = base.I64_extend_i32_u(v4339) << (uint(int64(48)) % 64)
	v4533 = int32(0)
	goto L630
L628:
	;
	goto L629
L629:
	;
	v4632 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+36))
	v4633 = m.T0[v4632].(func(*base.Module, int32) int32)(m, v3927)
	mBase = m.M
	v4634 = m.ExcPending
	if v4634 != 0 {
		goto L1
	} else {
		goto L647
	}
L630:
	;
	v4562 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+8))
	v4566 = *(*int32)(unsafe.Add(mBase, uint32(v4562+v4533<<(uint(int32(2))%32))))
	v4567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4566)+654)))
	if v4567 == int32(1) {
		goto L633
	} else {
		goto L634
	}
L631:
	;
	goto L629
L632:
	;
	v4598 = v4533 + int32(1)
	v4599 = *(*int32)(unsafe.Add(mBase, uint32(v3927)))
	if base.Ui32(v4598) < base.Ui32(v4599) {
		v4533 = v4598
		goto L630
	} else {
		goto L642
	}
L633:
	;
	v4570 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+28))
	v4572 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4570+v4533))) = uint8(v4572)
	goto L632
L634:
	;
	goto L635
L635:
	;
	v4574 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4566)+32)))
	v4575 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4566)+30)))
	v4578 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4566)+28)))
	v4582 = v4574 | (v4575<<(uint(int64(32))%64) | v4578<<(uint(int64(48))%64))
	if v4518|v4521|v4515 == v4582 {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v4584 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+28))
	v4586 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4584+v4533))) = uint8(v4586)
	goto L632
L637:
	;
	goto L638
L638:
	;
	v4588 = *(*int32)(unsafe.Add(mBase, uint32(v3927)+28))
	v4589 = v4588 + v4533
	if v4582 == v4518|(v4521|base.I64_extend_i32_u(v4344)&v4515) {
		goto L639
	} else {
		goto L640
	}
L639:
	;
	v4591 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4589))) = uint8(v4591)
	goto L632
L640:
	;
	goto L641
L641:
	;
	v4593 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4589))) = uint8(v4593)
	goto L632
L642:
	;
	goto L631
L643:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4494
	F_MemoryContextReset(m, v3941)
	mBase = m.M
	v4646 = m.ExcPending
	if v4646 != 0 {
		goto L1
	} else {
		goto L648
	}
L644:
	;
	v4641 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v3927)+86)) = uint16(v4641)
	goto L643
L645:
	;
	v4639 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3927)+86)) = uint8(v4639)
	goto L643
L646:
	;
	v4637 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3927)+86)) = uint8(v4637)
	goto L643
L647:
	;
	switch v4633 & int32(255) {
	case 0:
		goto L645
	case 1:
		goto L646
	default:
		goto L644
	}
L648:
	;
	goto L568
L649:
	;
	v4680 = v3927 + int32(80)
	v4681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3927)+86)))
	if v4681 == int32(1) {
		goto L652
	} else {
		goto L653
	}
L650:
	;
	v4765 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+66)))
	v4766 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+64)))
	v4773 = v4763
	v4775 = base.B2i32(v4691 == v4765|v4766<<(uint(int32(16))%32))
	v4777 = v4764
	goto L563
L651:
	;
	v4747 = int64(48)
	v4751 = int64(32)
	v4773 = v4684
	v4775 = base.B2i32(base.I64_extend_i32_u(v4686)|base.I64_extend_i32_u(v4685)<<(uint(v4747)%64)|base.I64_extend_i32_u(v4684)<<(uint(v4751)%64) == base.I64_extend_i32_u(v4719)<<(uint(v4747)%64)|v4720|base.I64_extend_i32_u(v4718)<<(uint(v4751)%64))
	v4777 = v4685
	goto L563
L652:
	;
	v4684 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3927)+82)))
	v4685 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3927)+80)))
	v4686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3927)+84)))
	if v4686 == int32(_a_F_gingetbitmap_8) {
		goto L658
	} else {
		goto L659
	}
L653:
	;
	goto L654
L654:
	;
	v4740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4680)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1092)) = uint16(v4740)
	v4742 = *(*int32)(unsafe.Add(mBase, uint32(v4680)))
	*(*int32)(unsafe.Add(mBase, uint32(v3814)+1088)) = v4742
	goto L552
L655:
	;
	v4735 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4680)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+68)) = uint16(v4735)
	v4737 = *(*int32)(unsafe.Add(mBase, uint32(v4680)))
	*(*int32)(unsafe.Add(mBase, uint32(v3814)+64)) = v4737
	v4773 = v4732
	v4775 = int32(1)
	v4777 = v4733
	goto L563
L656:
	;
	if v3910 != 0 {
		v4763 = v3900
		v4764 = v3903
		goto L650
	} else {
		goto L670
	}
L657:
	;
	v4718 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+66)))
	v4719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+64)))
	v4720 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+68)))
	if v4720 != int64(65535) {
		goto L651
	} else {
		goto L668
	}
L658:
	;
	v4691 = v4685<<(uint(int32(16))%32) | v4684
	if v4691 != int32(-1) {
		goto L661
	} else {
		goto L662
	}
L659:
	;
	goto L660
L660:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1090)) = uint16(v4684)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1088)) = uint16(v4685)
	v4713 = v4686 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1092)) = uint16(v4713)
	if v3910 == int32(0) {
		v4732 = v4684
		v4733 = v4685
		goto L655
	} else {
		goto L667
	}
L661:
	;
	if base.Ui32(v4691) <= base.Ui32(v3900&int32(_a_F_gingetbitmap_8)|v3903<<(uint(int32(16))%32)) {
		goto L656
	} else {
		goto L664
	}
L662:
	;
	goto L663
L663:
	;
	v4706 = int32(_a_F_gingetbitmap_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1092)) = uint16(v4706)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1090)) = uint16(v4684)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1088)) = uint16(v4685)
	if v3910 != 0 {
		goto L657
	} else {
		goto L666
	}
L664:
	;
	v4700 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1092)) = uint16(v4700)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1090)) = uint16(v4684)
	*(*uint16)(unsafe.Add(mBase, uint32(v3814)+1088)) = uint16(v4685)
	if v3910 == v4700 {
		v4732 = v4684
		v4733 = v4685
		goto L655
	} else {
		goto L665
	}
L665:
	;
	v4763 = v4684
	v4764 = v4685
	goto L650
L666:
	;
	v4732 = v4684
	v4733 = v4685
	goto L655
L667:
	;
	goto L657
L668:
	;
	v4725 = v4719<<(uint(int32(16))%32) | v4718
	if v4725 == int32(-1) {
		goto L651
	} else {
		goto L669
	}
L669:
	;
	v4773 = v4684
	v4775 = base.B2i32(v4685<<(uint(int32(16))%32)|v4684 == v4725)
	v4777 = v4685
	goto L563
L670:
	;
	v4732 = v3900
	v4733 = v3903
	goto L655
L671:
	;
	if v4775 != 0 {
		v3900 = v4773
		v3903 = v4777
		v3910 = v4803
		goto L561
	} else {
		goto L674
	}
L672:
	;
	goto L673
L673:
	;
	goto L562
L674:
	;
	goto L673
L675:
	;
	goto L553
L676:
	;
	v4884 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+68)))
	if v4884 != int32(_a_F_gingetbitmap_8) {
		v4904 = v4859
		goto L551
	} else {
		goto L684
	}
L677:
	;
	v4859 = int32(0)
	goto L676
L678:
	;
	goto L679
L679:
	;
	v4813 = *(*int32)(unsafe.Add(mBase, uint32(v3844)+uint32(_c_F_gingetbitmap[4])))
	v4817 = int32(0)
	goto L680
L680:
	;
	v4849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4813+v4817*int32(92))+87)))
	if v4849 != 0 {
		v4859 = v4849
		goto L676
	} else {
		goto L682
	}
L681:
	;
	v4859 = v4849
	goto L676
L682:
	;
	v4851 = v4817 + int32(1)
	if v4851 != v4804 {
		v4817 = v4851
		goto L680
	} else {
		goto L683
	}
L683:
	;
	goto L681
L684:
	;
	v4887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+66)))
	v4888 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3814)+64)))
	v4891 = v4887 | v4888<<(uint(int32(16))%32)
	if v4891 == int32(-1) {
		v4904 = v4859
		goto L551
	} else {
		goto L685
	}
L685:
	;
	F_tbm_add_page(m, v3810, v4891)
	mBase = m.M
	v4895 = m.ExcPending
	if v4895 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	v3838 = v3838 + int64(1)
	goto L549
L687:
	;
	v3838 = v3838 + int64(1)
	goto L549
L688:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v4942 = m.ExcPending
	if v4942 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	v4943 = *(*int32)(unsafe.Add(mBase, uint32(v1974)+1128))
	v4944 = *(*int32)(unsafe.Add(mBase, uint32(v4943)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1974)+16)) = v4944 + int32(4)
	F_errmsg(m, int32(_a_F_gingetbitmap_15), v1974+int32(16))
	mBase = m.M
	v4952 = m.ExcPending
	if v4952 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_11), int32(272), int32(_a_F_gingetbitmap_16))
	mBase = m.M
	v4957 = m.ExcPending
	if v4957 != 0 {
		goto L1
	} else {
		goto L691
	}
L691:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ginint4_queryextract(m *base.Module, l0 int32) int32 {
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
	var v19 int32
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v2
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v29 = v24 + int32(8)
		if v19 == int32(20) {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			if v32 <= int32(0) {
				v244 = v2
				m.G0 = v16 + int32(16)
				return v244
			} else {
				v37 = F_query_has_required_values(m, v24)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					if v37 != 0 {
						v39 = int32(0)
					} else {
						v39 = int32(2)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					v44 = F_palloc(m, v41<<(uint(int32(2))%32))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v46
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						if v48 <= v46 {
							v244 = v44
						} else {
							v52 = int32(0)
							v53 = v48
							v58 = v2
							for {
								v67 = v29 + v52<<(uint(int32(3))%32)
								v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67))))
								if v68 == int32(2) {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v44+v58<<(uint(int32(2))%32)))) = v74
									v77 = v58 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v20))) = v77
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
									v80 = v79
									v81 = v77
								} else {
									v80 = v53
									v81 = v58
								}
								v83 = v52 + int32(1)
								if v83 < v80 {
									v52 = v83
									v53 = v80
									v58 = v81
									continue
								} else {
									break
								}
								break
							}
							v244 = v44
						}
						m.G0 = v16 + int32(16)
						return v244
					}
				}
			}
		} else {
			v85 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			if v85 != 0 {
				v86 = F_array_contains_nulls(m, v24)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					if v86 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v261 = m.ExcPending
						if v261 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v264 = m.ExcPending
							if v264 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_ginint4_queryextract_0), int32(0))
								mBase = m.M
								v268 = m.ExcPending
								if v268 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(61), int32(_a_F_ginint4_queryextract_2))
									mBase = m.M
									v273 = m.ExcPending
									if v273 != 0 {
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
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						v91 = F_ArrayGetNItemsSafe(m, v88, v24+int32(16))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20))) = v91
							if v91 <= int32(0) {
								v206 = v2
								v208 = v2
								switch v19 - int32(3) {
								case 0:
									v239 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
									v244 = v208
									m.G0 = v16 + int32(16)
									return v244
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
										F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
										mBase = m.M
										v232 = m.ExcPending
										if v232 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								case 3:
									v239 = v206 ^ int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
									v244 = v208
									m.G0 = v16 + int32(16)
									return v244
								case 4, 10:
									if v206 != 0 {
										v224 = int32(0)
									} else {
										v224 = int32(2)
									}
									v239 = v224
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
									v244 = v208
									m.G0 = v16 + int32(16)
									return v244
								case 5, 11:
									v239 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
									v244 = v208
									m.G0 = v16 + int32(16)
									return v244
								}
							} else {
								v98 = F_palloc(m, v91<<(uint(int32(2))%32))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
									if v100 == int32(0) {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
										v110 = (v103<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									} else {
										v110 = v100
									}
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
									if v111 <= int32(0) {
										v206 = int32(0)
										v208 = v98
									} else {
										v115 = v110 + v24
										v117 = v111 & int32(3)
										v118 = int32(0)
										if base.Ui32(v111) < base.Ui32(int32(4)) {
											v167 = v118
											v180 = v167
											v192 = v2
											for {
												v194 = v180 << (uint(int32(2)) % 32)
												v197 = *(*int32)(unsafe.Add(mBase, uint32(v194+v115)))
												*(*int32)(unsafe.Add(mBase, uint32(v98+v194))) = v197
												v199 = int32(1)
												v203 = v192 + v199
												if v203 != v117 {
													v180 = v180 + v199
													v192 = v203
													continue
												} else {
													break
												}
												break
											}
											v206 = v199
											v208 = v98
										} else {
											v124 = v118
											v129 = int32(0)
											for {
												v138 = v124 << (uint(int32(2)) % 32)
												v141 = *(*int32)(unsafe.Add(mBase, uint32(v138+v115)))
												*(*int32)(unsafe.Add(mBase, uint32(v98+v138))) = v141
												v143 = int32(4)
												v144 = v138 | v143
												v147 = *(*int32)(unsafe.Add(mBase, uint32(v115+v144)))
												*(*int32)(unsafe.Add(mBase, uint32(v98+v144))) = v147
												v150 = v138 | int32(8)
												v153 = *(*int32)(unsafe.Add(mBase, uint32(v115+v150)))
												*(*int32)(unsafe.Add(mBase, uint32(v98+v150))) = v153
												v156 = v138 | int32(12)
												v159 = *(*int32)(unsafe.Add(mBase, uint32(v156+v115)))
												*(*int32)(unsafe.Add(mBase, uint32(v98+v156))) = v159
												v162 = v124 + v143
												v164 = v129 + v143
												if v164 != v111&int32(2147483644) {
													v124 = v162
													v129 = v164
													continue
												} else {
													break
												}
												break
											}
											if v117 != 0 {
												v167 = v162
												v180 = v167
												v192 = v2
												for {
													v194 = v180 << (uint(int32(2)) % 32)
													v197 = *(*int32)(unsafe.Add(mBase, uint32(v194+v115)))
													*(*int32)(unsafe.Add(mBase, uint32(v98+v194))) = v197
													v199 = int32(1)
													v203 = v192 + v199
													if v203 != v117 {
														v180 = v180 + v199
														v192 = v203
														continue
													} else {
														break
													}
													break
												}
												v206 = v199
												v208 = v98
											} else {
												v206 = int32(1)
												v208 = v98
											}
										}
									}
									switch v19 - int32(3) {
									case 0:
										v239 = v2
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
										v244 = v208
										m.G0 = v16 + int32(16)
										return v244
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
											F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									case 3:
										v239 = v206 ^ int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
										v244 = v208
										m.G0 = v16 + int32(16)
										return v244
									case 4, 10:
										if v206 != 0 {
											v224 = int32(0)
										} else {
											v224 = int32(2)
										}
										v239 = v224
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
										v244 = v208
										m.G0 = v16 + int32(16)
										return v244
									case 5, 11:
										v239 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
										v244 = v208
										m.G0 = v16 + int32(16)
										return v244
									}
								}
							}
						}
					}
				}
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v91 = F_ArrayGetNItemsSafe(m, v88, v24+int32(16))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v91
					if v91 <= int32(0) {
						v206 = v2
						v208 = v2
						switch v19 - int32(3) {
						case 0:
							v239 = v2
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
							v244 = v208
							m.G0 = v16 + int32(16)
							return v244
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
								F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
								mBase = m.M
								v232 = m.ExcPending
								if v232 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v239 = v206 ^ int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
							v244 = v208
							m.G0 = v16 + int32(16)
							return v244
						case 4, 10:
							if v206 != 0 {
								v224 = int32(0)
							} else {
								v224 = int32(2)
							}
							v239 = v224
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
							v244 = v208
							m.G0 = v16 + int32(16)
							return v244
						case 5, 11:
							v239 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
							v244 = v208
							m.G0 = v16 + int32(16)
							return v244
						}
					} else {
						v98 = F_palloc(m, v91<<(uint(int32(2))%32))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
							if v100 == int32(0) {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
								v110 = (v103<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v110 = v100
							}
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
							if v111 <= int32(0) {
								v206 = int32(0)
								v208 = v98
							} else {
								v115 = v110 + v24
								v117 = v111 & int32(3)
								v118 = int32(0)
								if base.Ui32(v111) < base.Ui32(int32(4)) {
									v167 = v118
									v180 = v167
									v192 = v2
									for {
										v194 = v180 << (uint(int32(2)) % 32)
										v197 = *(*int32)(unsafe.Add(mBase, uint32(v194+v115)))
										*(*int32)(unsafe.Add(mBase, uint32(v98+v194))) = v197
										v199 = int32(1)
										v203 = v192 + v199
										if v203 != v117 {
											v180 = v180 + v199
											v192 = v203
											continue
										} else {
											break
										}
										break
									}
									v206 = v199
									v208 = v98
								} else {
									v124 = v118
									v129 = int32(0)
									for {
										v138 = v124 << (uint(int32(2)) % 32)
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v138+v115)))
										*(*int32)(unsafe.Add(mBase, uint32(v98+v138))) = v141
										v143 = int32(4)
										v144 = v138 | v143
										v147 = *(*int32)(unsafe.Add(mBase, uint32(v115+v144)))
										*(*int32)(unsafe.Add(mBase, uint32(v98+v144))) = v147
										v150 = v138 | int32(8)
										v153 = *(*int32)(unsafe.Add(mBase, uint32(v115+v150)))
										*(*int32)(unsafe.Add(mBase, uint32(v98+v150))) = v153
										v156 = v138 | int32(12)
										v159 = *(*int32)(unsafe.Add(mBase, uint32(v156+v115)))
										*(*int32)(unsafe.Add(mBase, uint32(v98+v156))) = v159
										v162 = v124 + v143
										v164 = v129 + v143
										if v164 != v111&int32(2147483644) {
											v124 = v162
											v129 = v164
											continue
										} else {
											break
										}
										break
									}
									if v117 != 0 {
										v167 = v162
										v180 = v167
										v192 = v2
										for {
											v194 = v180 << (uint(int32(2)) % 32)
											v197 = *(*int32)(unsafe.Add(mBase, uint32(v194+v115)))
											*(*int32)(unsafe.Add(mBase, uint32(v98+v194))) = v197
											v199 = int32(1)
											v203 = v192 + v199
											if v203 != v117 {
												v180 = v180 + v199
												v192 = v203
												continue
											} else {
												break
											}
											break
										}
										v206 = v199
										v208 = v98
									} else {
										v206 = int32(1)
										v208 = v98
									}
								}
							}
							switch v19 - int32(3) {
							case 0:
								v239 = v2
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
								v244 = v208
								m.G0 = v16 + int32(16)
								return v244
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v228 = m.ExcPending
								if v228 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
									F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
									mBase = m.M
									v232 = m.ExcPending
									if v232 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 3:
								v239 = v206 ^ int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
								v244 = v208
								m.G0 = v16 + int32(16)
								return v244
							case 4, 10:
								if v206 != 0 {
									v224 = int32(0)
								} else {
									v224 = int32(2)
								}
								v239 = v224
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
								v244 = v208
								m.G0 = v16 + int32(16)
								return v244
							case 5, 11:
								v239 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v239
								v244 = v208
								m.G0 = v16 + int32(16)
								return v244
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
							v42 = F_palloc0(m, v37<<(uint(int32(2))%32))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v42
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
								v46 = F_palloc(m, v45)
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
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = int32(1831)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+40)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v145)+60)) = v153
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+36)) = uint8(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+16)) = int32(1832)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = int32(1833)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = int32(1834)
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = int32(1835)
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
	F_errfinish(m, int32(_a_F_gistbuild_3), int32(511), int32(_a_F_gistbuild_4))
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
	v288 = m.T0[v287].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v277, v278, v277, v278, int32(-1), int32(93), v15+int32(32), v278)
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
	v525 = m.T0[v524].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v511, v515, v511, v515, int32(-1), int32(94), v15+int32(32), v515)
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
func F_gistdentryinit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	v7 = l6
	v9 = int32(0)
	if l7 == v9 {
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)) = uint8(v12)
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = l3
		v20 = l0 + l1*int32(28)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(2712))))
		if v23 == v12 {
			return
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+uint32(_c_F_gistdentryinit[0])))
			v34 = F_FunctionCall1Coll(m, v20+int32(2708), v33, l2)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				if l2 == v34 {
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v41
					v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+12)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)) = uint16(v43)
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+14)))
					v53 = v45
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)) = uint8(v53)
				}
				return
			}
		}
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
		v53 = v9
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)) = uint8(v53)
		return
	}
}
func F_gistendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	F_brin_free_desc(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	v5 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(1392)
	m.G0 = v24
	F_gistDeCompressAtt(m, l3, l0, l1, v24+int32(736), v24+int32(192))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_gistDeCompressAtt(m, l3, l0, l2, v24+int32(224), v24+int32(160))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+10)))
	if v41 <= int32(0) {
		v285 = v5
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v24 + int32(1392)
	return v285
L5:
	;
	v47 = l3 + int32(_a_F_gistgetadjusted_0)
	v51 = v24 + int32(1268)
	v53 = v24 + int32(1252)
	v59 = v5
	v61 = v5
	goto L6
L6:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(160)+v59))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(192)+v59))))
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+1376)) = v83
	v85 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+1248)) = v85
	v88 = v59 << (uint(v85) % 32)
	v91 = v88 + (v24 + int32(32))
	v92 = v24 + v59
	if v78&v82&int32(1) == v83 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v169 = int32(0)
	if v155 == v169 {
		v285 = v169
		goto L4
	} else {
		goto L29
	}
L8:
	;
	v165 = v59 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v167 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166)+10)))
	if v165 < v167 {
		v59 = v165
		v61 = v155
		goto L6
	} else {
		goto L28
	}
L9:
	;
	v99 = v59 << (uint(int32(4)) % 32)
	v102 = v99 + (v24 + int32(224))
	v105 = v24 + int32(736) + v99
	v107 = v82 & int32(1)
	if v107 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
	v155 = v61
	goto L8
L12:
	;
	v108 = v102
	goto L14
L13:
	;
	v108 = v105
	goto L14
L14:
	;
	v109 = v78 | v82
	if v109 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v110 = v108
	goto L17
L16:
	;
	v110 = v105
	goto L17
L17:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v113
	if v109 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v115 = v108
	goto L20
L19:
	;
	v115 = v102
	goto L20
L20:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v115)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = v116
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v115)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v118
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v120)
	v123 = v59 * int32(28)
	v125 = v88 + v47
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v128 = v24 + int32(1248)
	v131 = F_FunctionCall2Coll(m, l3+int32(916)+v123, v126, v128, v24+int32(1376))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v131
	if v78|v61 != 0 {
		v155 = v61
		goto L8
	} else {
		goto L22
	}
L22:
	;
	if v107 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1248)) = uint8(v138)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v143 = F_FunctionCall3Coll(m, v123+(l3+int32(_a_F_gistgetadjusted_1)), v142, v137, v131, v128)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v155 = int32(1)
	goto L8
L26:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1248)))
	if v145 != 0 {
		v155 = v138
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
	if int32(0) < v167 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v178 = v166
	v181 = v169
	goto L33
L31:
	;
	goto L32
L32:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v272 = F_index_form_tuple(m, v269, v24+int32(1248), v24)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L44
	}
L33:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v181))))
	if v198 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L32
L35:
	;
	v245 = v181 + int32(1)
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v240)+10)))
	if v245 < v246 {
		v178 = v240
		v181 = v245
		goto L33
	} else {
		goto L43
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(1248)+v181<<(uint(int32(2))%32)))) = int32(0)
	v240 = v178
	goto L35
L37:
	;
	goto L38
L38:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1390)) = uint8(v208)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+1388)) = uint16(v208)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+1384)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v24)+1380)) = l0
	v216 = v181 << (uint(int32(2)) % 32)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216+(v24+int32(32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+1376)) = v220
	v224 = l3 + int32(1812) + v181*int32(28)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v225 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v216+v47)))
	v230 = F_FunctionCall1Coll(m, v224, v227, v24+int32(1376))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v234 = v178
	v235 = v220
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+int32(1248)+v216))) = v235
	v240 = v234
	goto L35
L42:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v234 = v233
	v235 = v232
	goto L41
L43:
	;
	goto L34
L44:
	;
	v274 = int32(_a_F_gistgetadjusted_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v272)+4)) = uint16(v274)
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v272)+4)) = uint16(v276)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v278
	v285 = v272
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v111 int32
	_ = v111
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
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
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
	return v435
L2:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v369 != 0 {
		goto L86
	} else {
		goto L87
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v17 != int32(1) {
		v435 = int32(0)
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
	v358 = m.ExcPending
	if v358 != 0 {
		goto L15
	} else {
		goto L83
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
	v74 = v71
	v75 = v70
	goto L29
L27:
	;
	v272 = v70
	goto L28
L28:
	;
	if v272 == int32(0) {
		v322 = v272
		goto L72
	} else {
		goto L73
	}
L29:
	;
	v82 = int32(_a_F_gistgettuple_0)
	v83 = v75 & v82
	if base.B2i32(v83 != v74&v82)|base.B2i32(v83 == int32(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v272 = v268
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
	v111 = v95
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
	v111 = v105
	goto L36
L38:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+v115<<(uint(int32(4))%32))+44)))
	v120 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v112 + v120
	*(*uint16)(unsafe.Add(mBase, uint32(v111+v112<<(uint(v120)%32)))) = uint16(v119)
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
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	if base.Ui32(v265) <= base.Ui32(v268) {
		v74 = v265
		v75 = v268
		goto L29
	} else {
		goto L71
	}
L41:
	;
	v242 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	if v244 == v242 {
		v435 = v242
		goto L1
	} else {
		goto L61
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
	F_LockBuffer(m, v148, int32(1))
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
	v230 = m.ExcPending
	if v230 != 0 {
		goto L15
	} else {
		goto L60
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
	v187 = int32(0)
	goto L56
L56:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	v196 = int32(1)
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195+v187<<(uint(v196)%32)))))
	v202 = v175 + int32(20) + v199<<(uint(int32(2))%32)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v203 | int32(_a_F_gistgettuple_2)
	v208 = v187 + v196
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v146)+28))
	if v208 < v209 {
		v187 = v208
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+16)))
	v212 = v175 + v211
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+12)))
	v215 = v213 | int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v212)+12)) = uint16(v215)
	F_MarkBufferDirtyHint(m, v148, int32(1))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L15
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	goto L52
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+28)) = int32(0)
	goto L41
L61:
	;
	v247 = F_pairingheap_remove_first(m, v243)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	if v247 == int32(0) {
		v435 = v242
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[6]))
	if v252 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L15
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v255
	v259 = int32(0)
	F_gistScanPage(m, l0, v247, v247+int32(32), v259, v259)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L15
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	F_pfree(m, v247)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L15
	} else {
		goto L69
	}
L69:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[0]))))
	if v265 == int32(0) {
		goto L39
	} else {
		goto L70
	}
L70:
	;
	goto L40
L71:
	;
	goto L30
L72:
	;
	v325 = v16 + int32(48)
	v328 = int32(4)
	v330 = v325 + v322&int32(_a_F_gistgettuple_0)<<(uint(v328)%32)
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v331)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v333
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+v335<<(uint(v328)%32))+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v339)
	v341 = int32(1)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v342 == v341 {
		goto L80
	} else {
		goto L81
	}
L73:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v281&int32(1) == int32(0) {
		v322 = v272
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v286 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v289 = int32(_a_F_gistgettuple_1)
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3]))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v293
	v296 = F_palloc(m, int32(816))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L15
	} else {
		goto L78
	}
L76:
	;
	v302 = v286
	v303 = v272
	goto L77
L77:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if int32(407) < v304 {
		v322 = v303
		goto L72
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v296
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v290
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v302 = v296
	v303 = v301
	goto L77
L79:
	;
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+v303&int32(_a_F_gistgettuple_0)<<(uint(int32(4))%32))+44)))
	v313 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v304 + v313
	*(*uint16)(unsafe.Add(mBase, uint32(v302+v304<<(uint(v313)%32)))) = uint16(v312)
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v322 = v320
	goto L72
L80:
	;
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v325+v345<<(uint(int32(4))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v349
	goto L82
L81:
	;
	goto L82
L82:
	;
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))))
	v353 = v351 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_gistgettuple[2]))) = uint16(v353)
	v435 = v341
	goto L1
L83:
	;
	F_errmsg_internal(m, int32(_a_F_gistgettuple_3), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_gistgettuple_4), int32(617), int32(_a_F_gistgettuple_5))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_pfree(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L15
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+8))
	if v375 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	goto L88
L90:
	;
	v435 = int32(0)
	goto L1
L91:
	;
	goto L92
L92:
	;
	v380 = l0 + int32(60)
	v382 = v374
	goto L93
L93:
	;
	v390 = F_pairingheap_remove_first(m, v382)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L15
	} else {
		goto L95
	}
L94:
	;
	v435 = v422
	goto L1
L95:
	;
	if v390 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v435 = int32(0)
	goto L1
L97:
	;
	goto L98
L98:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	if v395 == int32(-1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v380)+4)) = uint16(v398)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v400
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v402)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+23)))
	F_index_store_float8_orderby_distances(m, l0, v404, v390+int32(32), v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L15
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[6]))
	if v419 != 0 {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v410 = int32(1)
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v411 == v410 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v390)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v414
	goto L105
L104:
	;
	goto L105
L105:
	;
	F_pfree(m, v390)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L15
	} else {
		goto L106
	}
L106:
	;
	v435 = v410
	goto L1
L107:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L15
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v422 = int32(0)
	F_gistScanPage(m, l0, v390, v390+int32(32), v422, v422)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L15
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	F_pfree(m, v390)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L15
	} else {
		goto L112
	}
L112:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	if v432 != 0 {
		v382 = v431
		goto L93
	} else {
		goto L113
	}
L113:
	;
	goto L94
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v11 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 < v11 {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_gistinserttuples[0]))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+(v18^int32(-1))<<(uint(int32(6))%32))+16))
		v37 = v28
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_gistinserttuples[1]))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+v18<<(uint(int32(6))%32)+int32(-64))+16))
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
						F_LockBuffer(m, l6, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
							if v65 != 0 {
								F_gistfinishsplit(m, l0, l1, l2, v65, l8)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
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
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									F_LockBuffer(m, v70, int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										m.G0 = v14 + int32(16)
										return v51
									}
								}
							}
						}
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						if v65 != 0 {
							F_gistfinishsplit(m, l0, l1, l2, v65, l8)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
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
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								F_LockBuffer(m, v70, int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
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
					F_LockBuffer(m, l6, int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						if v65 != 0 {
							F_gistfinishsplit(m, l0, l1, l2, v65, l8)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
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
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								F_LockBuffer(m, v70, int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(16)
									return v51
								}
							}
						}
					}
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					if v65 != 0 {
						F_gistfinishsplit(m, l0, l1, l2, v65, l8)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
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
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							F_LockBuffer(m, v70, int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
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
func F_gistoptions(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = int32(2276)
	v6 = F_get_opfamily_proc(m, l1, v3, v3, int32(12))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v11 = F_OidFunctionCall1Coll(m, v6, int32(0), l0)
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14 & int32(_a_F_gisttranslatecmptype_0)
			}
		} else {
			v14 = int32(0)
			return v14 & int32(_a_F_gisttranslatecmptype_0)
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
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
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
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
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v636 int32
	_ = v636
	var v647 int64
	_ = v647
	var v648 int64
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int64
	_ = v694
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	v18 = m.G0
	v20 = v18 - int32(304)
	m.G0 = v20
	v23 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if int32(0) < v330 {
		goto L73
	} else {
		goto L74
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v29 = v27 + v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v33 = F_get_opfamily_name(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L70
	}
L7:
	;
	v37 = int32(0)
	v39 = F_SearchSysCacheList(m, int32(4), int32(1), v32, v37, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v41 = int32(1)
	v44 = int32(0)
	v46 = F_SearchSysCacheList(m, int32(5), v41, v32, v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v48 <= int32(0) {
		v318 = v41
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = v31
	goto L13
L12:
	;
	v51 = v30
	goto L13
L13:
	;
	v58 = int32(0)
	v59 = v41
	goto L14
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(48)+v58<<(uint(int32(2))%32))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+56))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
	v77 = v75 + v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v78 == v79 {
		v108 = v59
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v318 = v294
	goto L1
L16:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	if v109 != v30 {
		v294 = v108
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v81 = int32(0)
	v84 = F_errstart(m, int32(17), v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v84 == int32(0) {
		v108 = v81
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v92 = F_format_procedure(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+296)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v20)+292)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+288)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_1), v20+int32(288))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(86), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v108 = v81
	goto L16
L24:
	;
	v297 = v58 + int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v297 < v298 {
		v58 = v297
		v59 = v294
		goto L14
	} else {
		goto L69
	}
L25:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+16)))
	switch v111 - int32(1) {
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
	v275 = m.ExcPending
	if v275 != 0 {
		goto L2
	} else {
		goto L65
	}
L27:
	;
	v262 = int32(0)
	v265 = F_errstart(m, int32(17), v262)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L63
	}
L28:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+144)) = int64(9796820404457)
	v256 = int32(2)
	v260 = F_check_amproc_signature(m, v252, v51, int32(0), v256, v256, v20+int32(144))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L61
	}
L29:
	;
	v243 = int32(0)
	v246 = F_errstart(m, int32(17), v243)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L59
	}
L30:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = int32(23)
	v228 = int32(1)
	v233 = F_check_amproc_signature(m, v224, int32(21), v228, v228, v228, v20+int32(272))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L55
	}
L31:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = int32(2281)
	v215 = int32(1)
	v220 = F_check_amproc_signature(m, v211, int32(2278), v215, v215, v215, v20+int32(256))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L53
	}
L32:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v207 = F_check_amoptsproc_signature(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L51
	}
L33:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v189 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v20)+232)) = int64(111669149717)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v189
	v198 = int32(5)
	v202 = F_check_amproc_signature(m, v188, int32(701), int32(0), v198, v198, v20+int32(224))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L49
	}
L34:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v174 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+216)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v51
	v180 = int32(3)
	v184 = F_check_amproc_signature(m, v173, v174, int32(0), v180, v180, v20+int32(208))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L47
	}
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+192)) = int64(9796820404457)
	v165 = int32(2)
	v169 = F_check_amproc_signature(m, v160, int32(2281), int32(1), v165, v165, v20+int32(192))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L45
	}
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v146 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(9796820404457)
	v152 = int32(3)
	v156 = F_check_amproc_signature(m, v145, v146, int32(1), v152, v152, v20+int32(176))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L43
	}
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v133 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v133
	v136 = int32(1)
	v141 = F_check_amproc_signature(m, v132, v133, v136, v136, v136, v20+int32(160))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L41
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v115 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v115
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = int64(111669149717)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v115
	v124 = int32(5)
	v128 = F_check_amproc_signature(m, v114, int32(16), int32(0), v124, v124, v20+int32(112))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	if v128 == int32(0) {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v294 = v108
	goto L24
L41:
	;
	if v141 == int32(0) {
		goto L27
	} else {
		goto L42
	}
L42:
	;
	v294 = v108
	goto L24
L43:
	;
	if v156 == int32(0) {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	v294 = v108
	goto L24
L45:
	;
	if v169 == int32(0) {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v294 = v108
	goto L24
L47:
	;
	if v184 == int32(0) {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	v294 = v108
	goto L24
L49:
	;
	if v202 == int32(0) {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v294 = v108
	goto L24
L51:
	;
	if v207 == int32(0) {
		goto L27
	} else {
		goto L52
	}
L52:
	;
	v294 = v108
	goto L24
L53:
	;
	if v220 == int32(0) {
		goto L27
	} else {
		goto L54
	}
L54:
	;
	v294 = v108
	goto L24
L55:
	;
	if v233 == int32(0) {
		goto L27
	} else {
		goto L56
	}
L56:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	if v237 != int32(2276) {
		goto L27
	} else {
		goto L57
	}
L57:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v240 != int32(2276) {
		goto L27
	} else {
		goto L58
	}
L58:
	;
	v294 = v108
	goto L24
L59:
	;
	if v246 == int32(0) {
		v294 = v243
		goto L24
	} else {
		goto L60
	}
L60:
	;
	v271 = int32(153)
	v272 = int32(_a_F_gistvalidate_4)
	goto L26
L61:
	;
	if v260 != 0 {
		v294 = v108
		goto L24
	} else {
		goto L62
	}
L62:
	;
	goto L27
L63:
	;
	if v265 == int32(0) {
		v294 = v262
		goto L24
	} else {
		goto L64
	}
L64:
	;
	v271 = int32(165)
	v272 = int32(_a_F_gistvalidate_5)
	goto L26
L65:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v277 = F_format_procedure(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v279 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v33
	F_errmsg(m, v272, v20+int32(96))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), v271, int32(_a_F_gistvalidate_3))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v294 = int32(0)
	goto L24
L69:
	;
	goto L15
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_gistvalidate_6), v20)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(52), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
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
	v340 = int32(0)
	v341 = v318
	goto L76
L74:
	;
	v507 = v318
	goto L75
L75:
	;
	v519 = F_identify_opfamily_groups(m, v39, v46)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L2
	} else {
		goto L117
	}
L76:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(48)+v340<<(uint(int32(2))%32))))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+56))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+22)))
	v359 = v357 + v358
	v360 = int32(*(*int16)(unsafe.Add(mBase, uint32(v359)+16)))
	if int32(0) < v360 {
		v393 = v341
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v507 = v497
	goto L75
L78:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+18)))
	if v395 == int32(115) {
		v463 = int32(16)
		v464 = v393
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v363 = int32(0)
	v366 = F_errstart(m, int32(17), v363)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	if v366 == int32(0) {
		v393 = v363
		goto L78
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v374 = F_format_operator(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v359)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_7), v20+int32(80))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(185), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v393 = v363
	goto L78
L86:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v468 = F_check_amop_signature(m, v465, v463, v466, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L2
	} else {
		goto L107
	}
L87:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	v400 = F_get_opfamily_proc(m, v32, v398, v398, int32(8))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L2
	} else {
		goto L89
	}
L88:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v431 = F_get_op_rettype(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L2
	} else {
		goto L97
	}
L89:
	;
	if v400 != 0 {
		v429 = v393
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v402 = int32(0)
	v405 = F_errstart(m, int32(17), v402)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	if v405 == int32(0) {
		v429 = v402
		goto L88
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v413 = F_format_operator(m, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_8), v20-int32(-64))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(202), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	v429 = v402
	goto L88
L97:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v359)+28))
	v434 = F_opfamily_can_sort_type(m, v433, v431)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	if v434 != 0 {
		v463 = v431
		v464 = v429
		goto L86
	} else {
		goto L99
	}
L99:
	;
	v436 = int32(0)
	v439 = F_errstart(m, int32(17), v436)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v439 == int32(0) {
		v463 = v431
		v464 = v436
		goto L86
	} else {
		goto L101
	}
L101:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v447 = F_format_operator(m, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_9), v20+int32(48))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(213), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v463 = v431
	v464 = v436
	goto L86
L106:
	;
	v499 = v340 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if v499 < v500 {
		v340 = v499
		v341 = v497
		goto L76
	} else {
		goto L115
	}
L107:
	;
	if v468 != 0 {
		v497 = v464
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v470 = int32(0)
	v473 = F_errstart(m, int32(17), v470)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	if v473 == int32(0) {
		v497 = v470
		goto L106
	} else {
		goto L110
	}
L110:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v481 = F_format_operator(m, v480)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_10), v20+int32(32))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(232), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v497 = v470
	goto L106
L115:
	;
	goto L77
L116:
	;
	v636 = v507
	v647 = int64(1)
	goto L150
L117:
	;
	if v519 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v611 = int32(0)
	goto L116
L119:
	;
	goto L120
L120:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	if v524 <= int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v611 = int32(0)
	goto L116
L122:
	;
	goto L123
L123:
	;
	v528 = int32(0)
	if v524 != int32(1) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v532 = int32(0)
	if v532 < v524 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v584 = v528
	v588 = v528
	goto L126
L126:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v519)+12))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v601+v588<<(uint(int32(2))%32))))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)))
	if v606 != v30 {
		v611 = v584
		goto L116
	} else {
		goto L146
	}
L127:
	;
	v535 = v524
	goto L129
L128:
	;
	v535 = v532
	goto L129
L129:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v519)+12))
	v541 = int32(0)
	v543 = v541
	v546 = v541
	v547 = v528
	goto L130
L130:
	;
	v562 = v540 + v547<<(uint(int32(2))%32)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	if v30 == v564 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v535&int32(1) == int32(0) {
		v611 = v576
		goto L116
	} else {
		goto L145
	}
L132:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v566 == v30 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v569 = v543
	goto L134
L134:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	if v30 == v571 {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	v568 = v563
	goto L137
L136:
	;
	v568 = v543
	goto L137
L137:
	;
	v569 = v568
	goto L134
L138:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v573 == v30 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v576 = v569
	goto L140
L140:
	;
	v577 = int32(2)
	v578 = v547 + v577
	v580 = v546 + v577
	if v580 != v535&int32(2147483646) {
		v543 = v576
		v546 = v580
		v547 = v578
		goto L130
	} else {
		goto L144
	}
L141:
	;
	v575 = v570
	goto L143
L142:
	;
	v575 = v569
	goto L143
L143:
	;
	v576 = v575
	goto L140
L144:
	;
	goto L131
L145:
	;
	v584 = v576
	v588 = v578
	goto L126
L146:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v608 == v30 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v610 = v605
	goto L149
L148:
	;
	v610 = v584
	goto L149
L149:
	;
	v611 = v610
	goto L116
L150:
	;
	if v611 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L2
	} else {
		goto L167
	}
L152:
	;
	v694 = v647 + int64(1)
	if v694 != int64(13) {
		v636 = v692
		v647 = v694
		goto L150
	} else {
		goto L166
	}
L153:
	;
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v611)+16))
	if base.I32_wrap_i64(int64(base.Ui64(v648)>>(uint(v647)%64)))&int32(1) != 0 {
		v692 = v636
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v653 = base.I32_wrap_i64(v647)
	v654 = int32(12)
	if int32(1)<<(uint(v653)%32)&int32(_a_F_gistvalidate_11) != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L155
L157:
	;
	v665 = base.B2i32(base.Ui32(v653) <= base.Ui32(v654))
	goto L159
L158:
	;
	v665 = int32(0)
	goto L159
L159:
	;
	if base.B2i32(v653&v654 == int32(8))|v665 != 0 {
		v692 = v636
		goto L152
	} else {
		goto L160
	}
L160:
	;
	v667 = int32(0)
	v670 = F_errstart(m, int32(17), v667)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L2
	} else {
		goto L161
	}
L161:
	;
	if v670 == int32(0) {
		v692 = v667
		goto L152
	} else {
		goto L162
	}
L162:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L2
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v29 + int32(8)
	F_errmsg(m, int32(_a_F_gistvalidate_12), v20+int32(16))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(273), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	v692 = v667
	goto L152
L166:
	;
	goto L151
L167:
	;
	F_ReleaseCatCacheList(m, v39)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	F_ReleaseCatCache(m, v23)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	m.G0 = v20 + int32(304)
	return v692 & int32(1)
}
func F_gseg_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v14)+12)))
	if v16&int32(1) != 0 {
		v19 = int32(1)
		v20 = v6 - v19
		v22 = v20 & int32(_a_F_gseg_consistent_0)
		if base.B2i32(base.Ui32(int32(13)) < base.Ui32(v22))|base.B2i32(int32(base.Ui32(int32(_a_F_gseg_consistent_1))>>(uint(v22)%32))&v19 == int32(0)) != 0 {
			v99 = v2
			return v99
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v20&int32(_a_F_gseg_consistent_0)<<(uint(int32(2))%32))+uint32(_c_F_gseg_consistent[0])))
			v40 = F_DirectFunctionCall2Coll(m, v38, int32(0), v12, v7)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				return v40
			}
		}
	} else {
		switch v6&int32(_a_F_gseg_consistent_0) - int32(1) {
		case 0:
			v51 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_2), int32(0), v12, v7)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v51 == int32(0))
			}
		case 1:
			v58 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_3), int32(0), v12, v7)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v58 == int32(0))
			}
		case 2:
			v65 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_4), int32(0), v12, v7)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v65 != int32(0))
			}
		case 3:
			v72 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_5), int32(0), v12, v7)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v72 == int32(0))
			}
		case 4:
			v79 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_6), int32(0), v12, v7)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v79 == int32(0))
			}
		case 5, 6, 12:
			v86 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_7), int32(0), v12, v7)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v86 != int32(0))
			}
		case 7, 13:
			v93 = F_DirectFunctionCall2Coll(m, int32(_a_F_gseg_consistent_4), int32(0), v12, v7)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v99 = base.B2i32(v93 != int32(0))
				return v99
			}
		default:
			v99 = v2
			return v99
		}
	}
}
func F_gtsquery_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3))) = uint8(base.B2i32(v5 == v7))
	return v3
}
func F_gtsvectorout(m *base.Module, l0 int32) int32 {
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
	var v122 int64
	_ = v122
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int64
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int64
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int64
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v259 int64
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v270 int64
	_ = v270
	var v280 int64
	_ = v280
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
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
				return int32(0)
			} else {
				v300 = v34
				v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v301 != v16 {
					F_pfree(m, v16)
					mBase = m.M
					v304 = m.ExcPending
					if v304 != 0 {
						return int32(0)
					} else {
						m.G0 = v13 + int32(32)
						return v300
					}
				} else {
					m.G0 = v13 + int32(32)
					return v300
				}
			}
		} else {
			if v20&int32(4) != 0 {
				v39 = F_pstrdup(m, int32(_a_F_gtsvectorout_1))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v300 = v39
					v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v301 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v304 = m.ExcPending
						if v304 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(32)
							return v300
						}
					} else {
						m.G0 = v13 + int32(32)
						return v300
					}
				}
			} else {
				v41 = int32(8)
				v42 = v16 + v41
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
				v47 = v45 - v41
				if base.Ui32(v43) <= base.Ui32(int32(47)) {
					if v47 == int32(0) {
						v280 = v10
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
								v280 = v83
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
								v280 = v112
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
							v280 = v112
						}
					}
				} else {
					v122 = int64(0)
					if base.B2i32(v42 != (v16+int32(11))&int32(-4))|base.B2i32(v47 < int32(4)) != 0 {
						v201 = v42
						v202 = v47
						v207 = v122
					} else {
						v132 = v47 - int32(4)
						v136 = int32(base.Ui32(v132)>>(uint(int32(2))%32)) + int32(1)
						v138 = v136 & int32(3)
						if base.Ui32(int32(12)) <= base.Ui32(v132) {
							v143 = v42
							v144 = v47
							v147 = int32(0)
							v149 = v122
							for {
								v150 = int32(16)
								v151 = v144 - v150
								v153 = v143 + v150
								v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
								v169 = base.I64_extend_i32_u(base.I32_popcnt(v154)) + (base.I64_extend_i32_u(base.I32_popcnt(v157)) + (base.I64_extend_i32_u(base.I32_popcnt(v160)) + (v149 + base.I64_extend_i32_u(base.I32_popcnt(v163)))))
								v171 = v147 + int32(4)
								if v171 != v136&int32(2147483644) {
									v143 = v153
									v144 = v151
									v147 = v171
									v149 = v169
									continue
								} else {
									break
								}
								break
							}
							if v138 == int32(0) {
								v201 = v153
								v202 = v151
								v207 = v169
							} else {
								v175 = v153
								v176 = v151
								v181 = v169
								v183 = v175
								v184 = v176
								v185 = int32(0)
								v189 = v181
								for {
									v190 = int32(4)
									v191 = v184 - v190
									v193 = v183 + v190
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
									v197 = v189 + base.I64_extend_i32_u(base.I32_popcnt(v194))
									v199 = v185 + int32(1)
									if v199 != v138 {
										v183 = v193
										v184 = v191
										v185 = v199
										v189 = v197
										continue
									} else {
										break
									}
									break
								}
								v201 = v193
								v202 = v191
								v207 = v197
							}
						} else {
							v175 = v42
							v176 = v47
							v181 = v122
							v183 = v175
							v184 = v176
							v185 = int32(0)
							v189 = v181
							for {
								v190 = int32(4)
								v191 = v184 - v190
								v193 = v183 + v190
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
								v197 = v189 + base.I64_extend_i32_u(base.I32_popcnt(v194))
								v199 = v185 + int32(1)
								if v199 != v138 {
									v183 = v193
									v184 = v191
									v185 = v199
									v189 = v197
									continue
								} else {
									break
								}
								break
							}
							v201 = v193
							v202 = v191
							v207 = v197
						}
					}
					if v202 == int32(0) {
						v270 = v207
					} else {
						v211 = v202 & int32(3)
						if v211 == int32(0) {
							v232 = v201
							v234 = v202
							v238 = v207
						} else {
							v215 = v201
							v217 = v202
							v219 = int32(0)
							v221 = v207
							for {
								v222 = int32(1)
								v223 = v215 + v222
								v225 = v217 - v222
								v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
								v227 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v226)+uint32(_c_F_gtsvectorout[0]))))
								v228 = v221 + v227
								v230 = v219 + v222
								if v230 != v211 {
									v215 = v223
									v217 = v225
									v219 = v230
									v221 = v228
									continue
								} else {
									break
								}
								break
							}
							v232 = v223
							v234 = v225
							v238 = v228
						}
						if base.Ui32(v202) < base.Ui32(int32(4)) {
							v270 = v238
						} else {
							v241 = v232
							v243 = v234
							v247 = v238
							for {
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+3)))
								v249 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v248)+uint32(_c_F_gtsvectorout[0]))))
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+2)))
								v251 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v250)+uint32(_c_F_gtsvectorout[0]))))
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
								v253 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v252)+uint32(_c_F_gtsvectorout[0]))))
								v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
								v255 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v254)+uint32(_c_F_gtsvectorout[0]))))
								v259 = v249 + (v251 + (v253 + (v247 + v255)))
								v260 = int32(4)
								v263 = v243 - v260
								if v263 != 0 {
									v241 = v241 + v260
									v243 = v263
									v247 = v259
									continue
								} else {
									break
								}
								break
							}
							v270 = v259
						}
					}
					v280 = v270
				}
				v281 = base.I32_wrap_i64(v280)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v281
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v47<<(uint(int32(3))%32) - v281
				v288 = F_psprintf(m, int32(_a_F_gtsvectorout_2), v13)
				mBase = m.M
				v289 = m.ExcPending
				if v289 != 0 {
					return int32(0)
				} else {
					v300 = v288
					v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v301 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v304 = m.ExcPending
						if v304 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(32)
							return v300
						}
					} else {
						m.G0 = v13 + int32(32)
						return v300
					}
				}
			}
		}
	}
}
