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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
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
			v135 = m.ExcPending
			if v135 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(_a_F_GetConflictingVirtualXIDs_0))
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_GetConflictingVirtualXIDs_1), int32(0))
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetConflictingVirtualXIDs_2), int32(3436), int32(_a_F_GetConflictingVirtualXIDs_3))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
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
					v47 = int32(0)
					v50 = v3
					v52 = v43
					for {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v47<<(uint(int32(2))%32))))
						v61 = v52 + v58*int32(640)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
						if v62 == int32(0) {
							v101 = v50
							v102 = v52
						} else {
							if l1 != 0 {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+60))
								if v65 != l1 {
									v101 = v50
									v102 = v52
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
									if l0 != 0 {
										if v67 == int32(0) {
											v101 = v50
											v102 = v52
										} else {
											if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v67)) == int32(0) {
												v81 = base.B2i32(base.Ui32(l0) < base.Ui32(v67))
											} else {
												v81 = base.B2i32(int32(0) < v67-l0)
											}
											v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
											if v81 != 0 {
												v101 = v50
												v102 = v83
											} else {
												v85 = v83
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
												if v86 == int32(0) {
													v101 = v50
													v102 = v85
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
													v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
													v94 = v91 + v50<<(uint(int32(3))%32)
													*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
													*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
													v101 = v50 + int32(1)
													v102 = v85
												}
											}
										}
									} else {
										v85 = v52
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
										if v86 == int32(0) {
											v101 = v50
											v102 = v85
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
											v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
											v94 = v91 + v50<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
											*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
											v101 = v50 + int32(1)
											v102 = v85
										}
									}
								}
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
								if l0 != 0 {
									if v67 == int32(0) {
										v101 = v50
										v102 = v52
									} else {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v67)) == int32(0) {
											v81 = base.B2i32(base.Ui32(l0) < base.Ui32(v67))
										} else {
											v81 = base.B2i32(int32(0) < v67-l0)
										}
										v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
										if v81 != 0 {
											v101 = v50
											v102 = v83
										} else {
											v85 = v83
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
											if v86 == int32(0) {
												v101 = v50
												v102 = v85
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
												v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
												v94 = v91 + v50<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
												*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
												v101 = v50 + int32(1)
												v102 = v85
											}
										}
									}
								} else {
									v85 = v52
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
									if v86 == int32(0) {
										v101 = v50
										v102 = v85
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
										v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
										v94 = v91 + v50<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
										*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
										v101 = v50 + int32(1)
										v102 = v85
									}
								}
							}
						}
						v105 = v47 + int32(1)
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						if v105 < v106 {
							v47 = v105
							v50 = v101
							v52 = v102
							continue
						} else {
							break
						}
						break
					}
					v113 = v101
				} else {
					v113 = v3
				}
				v119 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
				F_LWLockRelease(m, v119+int32(512))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
					*(*int64)(unsafe.Add(mBase, uint32(v125+v113<<(uint(int32(3))%32)))) = int64(4294967295)
					return v125
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
				v47 = int32(0)
				v50 = v3
				v52 = v43
				for {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v47<<(uint(int32(2))%32))))
					v61 = v52 + v58*int32(640)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
					if v62 == int32(0) {
						v101 = v50
						v102 = v52
					} else {
						if l1 != 0 {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+60))
							if v65 != l1 {
								v101 = v50
								v102 = v52
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
								if l0 != 0 {
									if v67 == int32(0) {
										v101 = v50
										v102 = v52
									} else {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v67)) == int32(0) {
											v81 = base.B2i32(base.Ui32(l0) < base.Ui32(v67))
										} else {
											v81 = base.B2i32(int32(0) < v67-l0)
										}
										v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
										if v81 != 0 {
											v101 = v50
											v102 = v83
										} else {
											v85 = v83
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
											if v86 == int32(0) {
												v101 = v50
												v102 = v85
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
												v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
												v94 = v91 + v50<<(uint(int32(3))%32)
												*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
												*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
												v101 = v50 + int32(1)
												v102 = v85
											}
										}
									}
								} else {
									v85 = v52
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
									if v86 == int32(0) {
										v101 = v50
										v102 = v85
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
										v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
										v94 = v91 + v50<<(uint(int32(3))%32)
										*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
										*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
										v101 = v50 + int32(1)
										v102 = v85
									}
								}
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
							if l0 != 0 {
								if v67 == int32(0) {
									v101 = v50
									v102 = v52
								} else {
									if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v67)) == int32(0) {
										v81 = base.B2i32(base.Ui32(l0) < base.Ui32(v67))
									} else {
										v81 = base.B2i32(int32(0) < v67-l0)
									}
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[3]))
									if v81 != 0 {
										v101 = v50
										v102 = v83
									} else {
										v85 = v83
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
										if v86 == int32(0) {
											v101 = v50
											v102 = v85
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
											v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
											v94 = v91 + v50<<(uint(int32(3))%32)
											*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
											*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
											v101 = v50 + int32(1)
											v102 = v85
										}
									}
								}
							} else {
								v85 = v52
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
								if v86 == int32(0) {
									v101 = v50
									v102 = v85
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
									v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
									v94 = v91 + v50<<(uint(int32(3))%32)
									*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v86
									*(*int32)(unsafe.Add(mBase, uint32(v94))) = v89
									v101 = v50 + int32(1)
									v102 = v85
								}
							}
						}
					}
					v105 = v47 + int32(1)
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v105 < v106 {
						v47 = v105
						v50 = v101
						v52 = v102
						continue
					} else {
						break
					}
					break
				}
				v113 = v101
			} else {
				v113 = v3
			}
			v119 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[2]))
			F_LWLockRelease(m, v119+int32(512))
			mBase = m.M
			v123 = m.ExcPending
			if v123 != 0 {
				return int32(0)
			} else {
				v125 = *(*int32)(unsafe.Add(mBase, _c_F_GetConflictingVirtualXIDs[1]))
				*(*int64)(unsafe.Add(mBase, uint32(v125+v113<<(uint(int32(3))%32)))) = int64(4294967295)
				return v125
			}
		}
	}
}
func F_GetIntoRelEFlags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 != 0 {
		v5 = int32(64)
	} else {
		v5 = int32(0)
	}
	return v5
}
func F_GetLockmodeName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = int32(2)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v3)%32))+uint32(_c_F_GetLockmodeName[0])))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+l1<<(uint(v3)%32))))
	return v12
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_gbtreekey_out_0)
			F_errmsg(m, int32(_a_F_gbtreekey_out_1), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gbtreekey_out_2), int32(45), int32(_a_F_gbtreekey_out_3))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
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
	v8 = F_DirectFunctionCall2Coll(m, int32(2443), int32(0), v6, v7)
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v233 int32
	_ = v233
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	v8 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v21 = l4 << (uint(int32(2)) % 32)
	v22 = F_palloc(m, v21)
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
	v24 = F_palloc(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
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
	if v18 < v19 {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v28 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l4) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = v28
	v48 = v8
	goto L9
L7:
	;
	v94 = v28
	goto L8
L8:
	;
	v105 = l4 & int32(3)
	if v105 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L9:
	;
	v51 = v40 << (uint(int32(2)) % 32)
	v53 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v51))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v51+v22))) = v53
	v58 = int32(4)
	v59 = v51 | v58
	*(*int32)(unsafe.Add(mBase, uint32(v24+v59))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v22+v59))) = v53
	v67 = v51 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v67))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v22+v67))) = v53
	v75 = v51 | int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v75))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v75+v22))) = v53
	v83 = v40 + v58
	v85 = v48 + v58
	if v85 != l4&int32(2147483644) {
		v40 = v83
		v48 = v85
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v94 = v83
	goto L8
L11:
	;
	goto L10
L12:
	;
	v115 = v94
	v117 = v8
	goto L13
L13:
	;
	v126 = v115 << (uint(int32(2)) % 32)
	v128 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v126))) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v126+v22))) = v128
	v133 = int32(1)
	v136 = v117 + v133
	if v136 != v105 {
		v115 = v115 + v133
		v117 = v136
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
	v156 = v19
	goto L18
L17:
	;
	v156 = v18
	goto L18
L18:
	;
	if int32(0) < v156 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v167 = int32(0)
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
	if v19 <= v167 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	if v18 <= v167 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v167<<(uint(int32(2))%32))))
	if v182 < int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+v182<<(uint(int32(2))%32)))) = v167
	goto L24
L27:
	;
	v204 = v167 + int32(1)
	if v204 != v156 {
		v167 = v204
		goto L22
	} else {
		goto L30
	}
L28:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v167<<(uint(int32(2))%32))))
	if v195 < int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+v195<<(uint(int32(2))%32)))) = v167
	goto L27
L30:
	;
	goto L23
L31:
	;
	v233 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	F_pfree(m, v22)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L48
	}
L34:
	;
	v244 = v233 << (uint(int32(2)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v24+v244)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v22)))
	if v248&v246 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v252 = int32(0)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v252 <= v248 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v281 = v233 + int32(1)
	if v281 != l4 {
		v233 = v281
		goto L34
	} else {
		goto L47
	}
L39:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257+v248<<(uint(int32(2))%32))))
	v262 = v261
	goto L41
L40:
	;
	v262 = v252
	goto L41
L41:
	;
	v263 = F_lappend(m, v254, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v263
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if int32(0) <= v246 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269+v246<<(uint(int32(2))%32))))
	v274 = v273
	goto L45
L44:
	;
	v274 = v252
	goto L45
L45:
	;
	v275 = F_lappend(m, v266, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v275
	goto L38
L47:
	;
	goto L35
L48:
	;
	F_pfree(m, v24)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
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
	var v11 int32
	_ = v11
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v298 int32
	_ = v298
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	var v304 float64
	_ = v304
	var v305 float64
	_ = v305
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 float64
	_ = v476
	var v477 float64
	_ = v477
	var v481 float64
	_ = v481
	var v482 float64
	_ = v482
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	v11 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v20 = F_find_mergeclauses_for_outer_pathkeys(m, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = int32(0)
	if l4&int32(-2) != int32(8) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
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
		goto L3
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v43 = F_make_inner_pathkeys_for_merge(m, l0, v20, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
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
	v34 = v11
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
	v37 = v36
	goto L16
L15:
	;
	v37 = v11
	goto L16
L16:
	;
	if v37 != v34 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	F_try_mergejoin_path(m, l0, l1, l3, l7, l8, v20, int32(0), v43, v29, l5, l9)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if l4 == int32(9) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l7)+64))
	if v43 == v49 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v43 == int32(0) {
		goto L3
	} else {
		goto L39
	}
L22:
	;
	v102 = int32(1)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v58 = int32(0)
	goto L26
L25:
	;
	v102 = v94
	goto L21
L26:
	;
	v62 = int32(0)
	if v43 == v62 {
		v72 = v62
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v94 = int32(0)
	goto L25
L28:
	;
	if v49 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v66 <= v58 {
		v72 = int32(0)
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v72 = v68 + v58<<(uint(int32(2))%32)
	goto L28
L31:
	;
	v78 = base.B2i32(v72 == int32(0))
	if v72 == int32(0) {
		v94 = v78
		goto L25
	} else {
		goto L36
	}
L32:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v58 < v73 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v102 = base.B2i32(v72 == int32(0))
	goto L21
L35:
	;
	goto L34
L36:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v84 = v81 + v58<<(uint(int32(2))%32)
	if v84 == int32(0) {
		v94 = v78
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v89 == v90 {
		v58 = v58 + int32(1)
		goto L26
	} else {
		goto L38
	}
L38:
	;
	goto L27
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if l6 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v102 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	if v105 <= int32(0) {
		goto L3
	} else {
		goto L45
	}
L42:
	;
	if v105 < int32(2) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v108 = F_list_copy(m, v43)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v112 = v108
	goto L40
L45:
	;
	v112 = v43
	goto L40
L46:
	;
	v114 = l7
	goto L48
L47:
	;
	v114 = int32(0)
	goto L48
L48:
	;
	v119 = v105
	v126 = v112
	v128 = v114
	v131 = v114
	goto L49
L49:
	;
	v132 = int32(0)
	if v126 == v132 {
		v140 = v132
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L3
L51:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v142 = int32(0)
	if v141 != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	goto L51
L53:
	;
	if v119 <= int32(0) {
		v140 = v132
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v119 < v137 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v119
	goto L57
L56:
	;
	goto L57
L57:
	;
	v140 = v126
	goto L52
L58:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v334 = int32(0)
	if v333 != 0 {
		goto L142
	} else {
		goto L143
	}
L59:
	;
	if v261 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L60:
	;
	goto L59
L61:
	;
	v160 = v142
	v163 = v142
	goto L66
L62:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if int32(0) < v151 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v261 = v142
	goto L60
L65:
	;
	goto L64
L66:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166+v163<<(uint(int32(2))%32))))
	if l9 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v261 = v245
	goto L60
L68:
	;
	v252 = v163 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v252 < v253 {
		v160 = v245
		v163 = v252
		goto L66
	} else {
		goto L101
	}
L69:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+21)))
	if v171 != int32(1) {
		v245 = v160
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v160 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	v174 = F_compare_path_costs(m, v160, v170, int32(1))
	mBase = m.M
	if v174 <= int32(0) {
		v245 = v160
		goto L68
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v170)+64))
	if v140 == v177 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v170)+16))
	if v233 != 0 {
		goto L95
	} else {
		goto L96
	}
L78:
	;
	v185 = int32(0)
	goto L79
L79:
	;
	v192 = int32(0)
	if v140 == v192 {
		v202 = v192
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v202 != 0 {
		v245 = v160
		goto L68
	} else {
		goto L94
	}
L81:
	;
	if v177 != 0 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v196 <= v185 {
		v202 = int32(0)
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v202 = v198 + v185<<(uint(int32(2))%32)
	goto L81
L84:
	;
	if v202 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v185 < v203 {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v202 == int32(0) {
		goto L77
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v245 = v160
	goto L68
L90:
	;
	goto L80
L91:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v212 = v209 + v185<<(uint(int32(2))%32)
	if v212 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v217 == v218 {
		v185 = v185 + int32(1)
		goto L79
	} else {
		goto L93
	}
L93:
	;
	v245 = v160
	goto L68
L94:
	;
	goto L77
L95:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	v236 = v234
	goto L97
L96:
	;
	v236 = int32(0)
	goto L97
L97:
	;
	v237 = F_bms_is_subset(m, v236, v142)
	mBase = m.M
	if v237 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v238 = v170
	goto L100
L99:
	;
	v238 = v160
	goto L100
L100:
	;
	v245 = v238
	goto L68
L101:
	;
	goto L67
L102:
	;
	v331 = int32(0)
	v332 = v128
	goto L58
L103:
	;
	goto L104
L104:
	;
	if v128 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v261)+40))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v128)+40))
	if v276 != v277 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L107
L107:
	;
	if v119 < v105 {
		goto L133
	} else {
		goto L134
	}
L108:
	;
	if int32(0) <= v319 {
		v331 = int32(0)
		v332 = v128
		goto L58
	} else {
		goto L132
	}
L109:
	;
	if v276 < v277 {
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
	v282 = int32(-1)
	goto L114
L113:
	;
	v282 = int32(1)
	goto L114
L114:
	;
	v319 = v282
	goto L108
L115:
	;
	v319 = v313
	goto L108
L116:
	;
	v313 = int32(0)
	goto L115
L118:
	;
	goto L119
L119:
	;
	v298 = int32(-1)
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v261)+56))
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v128)+56))
	if base.F64_lt(v299, v300) != 0 {
		v313 = v298
		goto L115
	} else {
		goto L126
	}
L126:
	;
	if base.F64_gt(v299, v300) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v319 = int32(1)
	goto L108
L128:
	;
	goto L129
L129:
	;
	v304 = *(*float64)(unsafe.Add(mBase, uint32(v261)+48))
	v305 = *(*float64)(unsafe.Add(mBase, uint32(v128)+48))
	if base.F64_lt(v304, v305) != 0 {
		v313 = v298
		goto L115
	} else {
		goto L130
	}
L130:
	;
	if base.F64_gt(v304, v305) != 0 {
		v313 = int32(1)
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
	v324 = F_trim_mergeclauses_for_inner_pathkeys(m, v20, v140)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	v326 = v20
	goto L135
L135:
	;
	v327 = int32(0)
	F_try_mergejoin_path(m, l0, l1, l3, v261, l8, v326, v327, v327, v29, l5, l9)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	v326 = v324
	goto L135
L137:
	;
	v331 = v326
	v332 = v261
	goto L58
L138:
	;
	if v119 < int32(2) {
		goto L3
	} else {
		goto L219
	}
L139:
	;
	if v453 == int32(0) {
		v523 = v131
		goto L138
	} else {
		goto L182
	}
L140:
	;
	goto L139
L141:
	;
	v352 = v334
	v355 = v334
	goto L146
L142:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if int32(0) < v343 {
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v453 = v334
	goto L140
L145:
	;
	goto L144
L146:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358+v355<<(uint(int32(2))%32))))
	if l9 != 0 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v453 = v437
	goto L140
L148:
	;
	v444 = v355 + int32(1)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v444 < v445 {
		v352 = v437
		v355 = v444
		goto L146
	} else {
		goto L181
	}
L149:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+21)))
	if v363 != int32(1) {
		v437 = v352
		goto L148
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	if v352 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v366 = F_compare_path_costs(m, v352, v362, v334)
	mBase = m.M
	if v366 <= int32(0) {
		v437 = v352
		goto L148
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v362)+64))
	if v140 == v369 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L155
L157:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v362)+16))
	if v425 != 0 {
		goto L175
	} else {
		goto L176
	}
L158:
	;
	v377 = int32(0)
	goto L159
L159:
	;
	v384 = int32(0)
	if v140 == v384 {
		v394 = v384
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v394 != 0 {
		v437 = v352
		goto L148
	} else {
		goto L174
	}
L161:
	;
	if v369 != 0 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v388 <= v377 {
		v394 = int32(0)
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v394 = v390 + v377<<(uint(int32(2))%32)
	goto L161
L164:
	;
	if v394 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L165:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v377 < v395 {
		goto L164
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if v394 == int32(0) {
		goto L157
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	v437 = v352
	goto L148
L170:
	;
	goto L160
L171:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v404 = v401 + v377<<(uint(int32(2))%32)
	if v404 == int32(0) {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	if v409 == v410 {
		v377 = v377 + int32(1)
		goto L159
	} else {
		goto L173
	}
L173:
	;
	v437 = v352
	goto L148
L174:
	;
	goto L157
L175:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v428 = v426
	goto L177
L176:
	;
	v428 = int32(0)
	goto L177
L177:
	;
	v429 = F_bms_is_subset(m, v428, v334)
	mBase = m.M
	if v429 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v430 = v362
	goto L180
L179:
	;
	v430 = v352
	goto L180
L180:
	;
	v437 = v430
	goto L148
L181:
	;
	goto L147
L182:
	;
	if v131 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v453)+40))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v131)+40))
	if v466 != v467 {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	goto L185
L185:
	;
	if v453 != v332 {
		goto L211
	} else {
		goto L212
	}
L186:
	;
	if int32(0) <= v509 {
		v523 = v131
		goto L138
	} else {
		goto L210
	}
L187:
	;
	if v466 < v467 {
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
	v472 = int32(-1)
	goto L192
L191:
	;
	v472 = int32(1)
	goto L192
L192:
	;
	v509 = v472
	goto L186
L193:
	;
	v509 = v503
	goto L186
L194:
	;
	v503 = int32(0)
	goto L193
L195:
	;
	v475 = int32(-1)
	v476 = *(*float64)(unsafe.Add(mBase, uint32(v453)+48))
	v477 = *(*float64)(unsafe.Add(mBase, uint32(v131)+48))
	if base.F64_lt(v476, v477) != 0 {
		v503 = v475
		goto L193
	} else {
		goto L198
	}
L198:
	;
	if base.F64_gt(v476, v477) != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v509 = int32(1)
	goto L186
L200:
	;
	goto L201
L201:
	;
	v481 = *(*float64)(unsafe.Add(mBase, uint32(v453)+56))
	v482 = *(*float64)(unsafe.Add(mBase, uint32(v131)+56))
	if base.F64_lt(v481, v482) != 0 {
		v503 = v475
		goto L193
	} else {
		goto L202
	}
L202:
	;
	if base.F64_gt(v481, v482) == int32(0) {
		goto L194
	} else {
		goto L203
	}
L203:
	;
	v503 = int32(1)
	goto L193
L210:
	;
	goto L185
L211:
	;
	if v331 != 0 {
		v516 = v331
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	v523 = v453
	goto L138
L214:
	;
	v517 = int32(0)
	F_try_mergejoin_path(m, l0, l1, l3, v453, l8, v516, v517, v517, v29, l5, l9)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L218
	}
L215:
	;
	if v105 <= v119 {
		v516 = v20
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v514 = F_trim_mergeclauses_for_inner_pathkeys(m, v20, v140)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v516 = v514
	goto L214
L218:
	;
	goto L213
L219:
	;
	if l6 == int32(0) {
		v119 = v119 - int32(1)
		v126 = v140
		v128 = v332
		v131 = v523
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
			v18 = F_get_namespace_name_or_temp(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v68
						F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_0), v5+int32(-48))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_generate_qualified_relation_name_1), int32(_a_F_generate_qualified_relation_name_2), int32(_a_F_generate_qualified_relation_name_3))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					F_initStringInfo(m, v5+int32(-16))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_quote_identifier(m, v18)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v26
							F_appendStringInfo(m, v5+int32(-16), int32(_a_F_generate_qualified_relation_name_4), v5+int32(-32))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v40 = F_quote_identifier(m, v16+int32(4))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_appendStringInfoString(m, v5+int32(-16), v40)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
										F_ReleaseCatCache(m, v10)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 - int32(-64)
											return v44
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
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_generate_qualified_relation_name_5), v7)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_generate_qualified_relation_name_1), int32(_a_F_generate_qualified_relation_name_6), int32(_a_F_generate_qualified_relation_name_3))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v58 int32
	_ = v58
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
		v11 = int32(1)
		v14 = l3 << (uint(int32(3)) % 32)
		if l2 != v11 {
			v21 = l1
			v22 = int32(0)
			for {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v29 = base.I32_rem_u_s(v28, v14)
				v30 = int32(3)
				v32 = l0 + int32(base.Ui32(v29)>>(uint(v30)%32))
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
				v34 = int32(1)
				v35 = int32(7)
				v38 = v33 | v34<<(uint(v29&v35)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v38)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v41 = base.I32_rem_u_s(v40, v14)
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
			v58 = v53
		} else {
			v58 = l1
		}
		if l2&v11 == int32(0) {
		} else {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
			v68 = base.I32_rem_u_s(v67, v14)
			v71 = l0 + int32(base.Ui32(v68)>>(uint(int32(3))%32))
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v77 = v72 | int32(1)<<(uint(v68&int32(7))%32)
			*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v77)
		}
	}
	return
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
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
			if v35 == v36 {
				v44 = v36
			} else {
				if v25 <= int32(0) {
					v44 = v36
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					if v25 < v41 {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v25
					} else {
					}
					v44 = v35
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v44
			*(*int32)(unsafe.Add(mBase, _c_F_geqo_eval[0])) = v20
			F_MemoryContextDelete(m, v15)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	v5 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = v16 & int32(268435455)
	if v18 == v5 {
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
	v24 = l0 + int32(4)
	v27 = v24 + v18<<(uint(int32(3))%32)
	v33 = v5
	v37 = v5
	v39 = v18
	goto L4
L4:
	;
	v46 = int32(base.Ui32(v39-v37)>>(uint(int32(1))%32)) + v37
	v51 = v46
	v52 = v33
	goto L6
L5:
	;
	return v381
L6:
	;
	v63 = v51 - int32(1)
	if int32(0) <= v63 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v24+v46<<(uint(int32(2))%32))))
	if v80 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v24+v63<<(uint(int32(2))%32))))
	v72 = v69&int32(268435455) + v52
	if int32(0) <= v69 {
		v51 = v63
		v52 = v72
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v75 = v52
	goto L10
L10:
	;
	goto L7
L11:
	;
	v75 = v72
	goto L10
L12:
	;
	goto L5
L13:
	;
	v368 = int32(0)
	v372 = base.B2i32(v367 < v368)
	if v367 < v368 {
		goto L86
	} else {
		goto L87
	}
L14:
	;
	if v134 != l2 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v88 = v46
	v90 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v134 = v80 & int32(268435455)
	goto L14
L18:
	;
	v100 = v88 - int32(1)
	if int32(0) <= v100 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v134 = v80&int32(268435455) - v112
	goto L14
L20:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v24+v100<<(uint(int32(2))%32))))
	v109 = v106&int32(268435455) + v90
	if int32(0) <= v106 {
		v88 = v100
		v90 = v109
		goto L18
	} else {
		goto L23
	}
L21:
	;
	v112 = v90
	goto L22
L22:
	;
	goto L19
L23:
	;
	v112 = v109
	goto L22
L24:
	;
	if l2 < v134 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v140 = v75 + v27
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v139 = int32(1)
	goto L29
L28:
	;
	v139 = int32(-1)
	goto L29
L29:
	;
	v367 = v139
	goto L13
L30:
	;
	if v202 != 0 {
		v367 = v202
		goto L13
	} else {
		goto L48
	}
L31:
	;
	v202 = int32(0)
	goto L30
L32:
	;
	v176 = v171
	v177 = v172
	v178 = v173
	goto L42
L33:
	;
	if (v140|l1)&int32(3) != 0 {
		v171 = v140
		v172 = l1
		v173 = l2
		goto L32
	} else {
		goto L36
	}
L34:
	;
	v164 = v140
	v165 = l1
	v166 = l2
	goto L35
L35:
	;
	if v166 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L36:
	;
	v148 = v140
	v149 = l1
	v150 = l2
	goto L37
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v153 != v154 {
		v171 = v148
		v172 = v149
		v173 = v150
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v164 = v159
	v165 = v157
	v166 = v161
	goto L35
L39:
	;
	v156 = int32(4)
	v157 = v149 + v156
	v159 = v148 + v156
	v161 = v150 - v156
	if base.Ui32(int32(3)) < base.Ui32(v161) {
		v148 = v159
		v149 = v157
		v150 = v161
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v171 = v164
	v172 = v165
	v173 = v166
	goto L32
L42:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v181 == v182 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v202 = v181 - v182
	goto L30
L44:
	;
	v184 = int32(1)
	v189 = v178 - v184
	if v189 != 0 {
		v176 = v176 + v184
		v177 = v177 + v184
		v178 = v189
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
	v206 = F_palloc(m, int32(20))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v210 = l3
	goto L51
L51:
	;
	v212 = v46 + v18
	v217 = v212
	v218 = int32(0)
	goto L54
L52:
	;
	return int32(0)
L53:
	;
	v210 = v206
	goto L51
L54:
	;
	v229 = v217 - int32(1)
	if int32(0) <= v229 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v246 = l0 + int32(4)
	v249 = v246 + v212<<(uint(int32(2))%32)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	switch int32(base.Ui32(v250)>>(uint(int32(28))%32)) & int32(7) {
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
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v24+v229<<(uint(int32(2))%32))))
	v238 = v235&int32(268435455) + v218
	if int32(0) <= v235 {
		v217 = v229
		v218 = v238
		goto L54
	} else {
		goto L59
	}
L57:
	;
	v241 = v218
	goto L58
L58:
	;
	goto L55
L59:
	;
	v241 = v238
	goto L58
L60:
	;
	v381 = v210
	goto L12
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(18)
	v315 = (v241 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+8)) = v27 + v315
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v318 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)) = uint8(v306)
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(3)
	goto L60
L63:
	;
	v302 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)) = uint8(v302)
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(3)
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v27 + (v241+int32(3))&int32(-4)
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v210)+8)) = v27 + v241
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v261 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(0)
	goto L60
L67:
	;
	v266 = v212
	v267 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v261 & int32(268435455)
	goto L60
L70:
	;
	v273 = v266 - int32(1)
	if int32(0) <= v273 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v261&int32(268435455) - v285
	goto L60
L72:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v246+v273<<(uint(int32(2))%32))))
	v282 = v279&int32(268435455) + v267
	if int32(0) <= v279 {
		v266 = v273
		v267 = v282
		goto L70
	} else {
		goto L75
	}
L73:
	;
	v285 = v267
	goto L74
L74:
	;
	goto L71
L75:
	;
	v285 = v282
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v356 + (v241 - v315)
	goto L60
L77:
	;
	v323 = v212
	v324 = int32(0)
	goto L80
L78:
	;
	goto L79
L79:
	;
	v356 = v318 & int32(268435455)
	goto L76
L80:
	;
	v330 = v323 - int32(1)
	if int32(0) <= v330 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v356 = v318&int32(268435455) - v342
	goto L76
L82:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v246+v330<<(uint(int32(2))%32))))
	v339 = v336&int32(268435455) + v324
	if int32(0) <= v336 {
		v323 = v330
		v324 = v339
		goto L80
	} else {
		goto L85
	}
L83:
	;
	v342 = v324
	goto L84
L84:
	;
	goto L81
L85:
	;
	v342 = v339
	goto L84
L86:
	;
	v373 = v46 + int32(1)
	goto L88
L87:
	;
	v373 = v37
	goto L88
L88:
	;
	if v367 < v368 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v374 = v39
	goto L91
L90:
	;
	v374 = v46
	goto L91
L91:
	;
	if base.Ui32(v373) < base.Ui32(v374) {
		v33 = v368
		v37 = v373
		v39 = v374
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v381 = v368
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	v72 = v69 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v29+v72)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72+v27)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
	v79 = int32(1)
	if l4&v79 != 0 {
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
		v69 = v6
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
		v69 = v6
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
	v69 = int32(1)
	goto L4
L15:
	;
	v69 = int32(1)
	goto L4
L16:
	;
	v61 = int32(0)
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+28)))
	if l2 != v62 {
		v154 = v61
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
		v69 = v46
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
		v69 = v46
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
	v69 = int32(3)
	goto L4
L26:
	;
	v69 = int32(3)
	goto L4
L27:
	;
	v64 = int32(4)
	if l3 == int32(0) {
		v69 = v64
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v67 != l3 {
		v154 = v61
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v69 = v64
	goto L4
L30:
	;
	v85 = F_SysCacheGetAttrNotNull(m, int32(65), l1, v69+int32(27))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
		v154 = v79
		goto L3
	} else {
		goto L45
	}
L33:
	;
	return int32(0)
L34:
	;
	v89 = F_pg_detoast_datum_copy(m, v85)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
	v94 = F_SearchSysCache1(m, int32(82), v91)
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
	F_deconstruct_array(m, v89, v101, v102, v103, l0+int32(12), int32(0), l0+int32(16))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
	goto L39
L41:
	;
	goto L42
L42:
	;
	F_pfree(m, v89)
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
	v130 = F_SysCacheGetAttrNotNull(m, int32(65), l1, v69+int32(22))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v132 = F_pg_detoast_datum_copy(m, v130)
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
	v154 = v79
	goto L3
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v91
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
	F_errfinish(m, int32(_a_F_get_attstatsslot_1), int32(3421), int32(_a_F_get_attstatsslot_2))
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
	F_errfinish(m, int32(_a_F_get_attstatsslot_1), int32(3466), int32(_a_F_get_attstatsslot_2))
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v124 int32
	_ = v124
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
	var v163 int32
	_ = v163
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
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
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v550 int32
	_ = v550
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v968 int32
	_ = v968
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
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
	var v1052 int32
	_ = v1052
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
	var v1120 float64
	_ = v1120
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
	var v1158 int32
	_ = v1158
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_memoize_path[0])))
	if v24 != int32(1) {
		v1158 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(16)
	return v1158
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
	if base.F64_lt(v28, float64(2)) != 0 {
		v1158 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+317)))
	if v31 != int32(1) {
		v390 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v394 != 0 {
		goto L97
	} else {
		goto L98
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v81 == int32(2) {
		v390 = v8
		goto L4
	} else {
		goto L22
	}
L7:
	;
	v81 = int32(0)
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
	v50 = int32(0)
	v52 = v50
	v53 = v50
	goto L13
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(8)+v52<<(uint(int32(2))%32))))
	if v61 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v81 = v74
	goto L6
L15:
	;
	goto L14
L16:
	;
	v62 = int32(2)
	if v53 != 0 {
		v74 = v62
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
	v70 = v52 + int32(1)
	if v70 != v47 {
		v52 = v70
		v53 = v67
		goto L13
	} else {
		goto L21
	}
L19:
	;
	v63 = int32(1)
	if base.Ui32(v63) < base.Ui32(base.I32_popcnt(v61)) {
		v74 = v62
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v67 = v63
	goto L18
L21:
	;
	v74 = v67
	goto L15
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v84 == int32(0) {
		v390 = v8
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v87 <= int32(0) {
		v390 = v8
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v103 = v8
	v104 = v8
	goto L25
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v103<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v113 == int32(0) {
		v368 = v104
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v390 = v368
	goto L4
L27:
	;
	v373 = v103 + int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v373 < v374 {
		v103 = v373
		v104 = v368
		goto L25
	} else {
		goto L95
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v117 = int32(0)
	v124 = base.B2i32(v116|v34 == v117)
	if v116 == v117 {
		v163 = v124
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v163 == int32(0) {
		v368 = v104
		goto L27
	} else {
		goto L41
	}
L30:
	;
	goto L29
L31:
	;
	if v34 == int32(0) {
		v163 = v124
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v130 != v131 {
		v163 = int32(0)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v133 = int32(1)
	if v130 <= v133 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v136 = v133
	goto L36
L35:
	;
	v136 = v130
	goto L36
L36:
	;
	v137 = int32(8)
	v142 = int32(0)
	goto L37
L37:
	;
	v150 = v142 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v116+v137+v150)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v34+v137))))
	v155 = base.B2i32(v152 == v154)
	if v154 != v152 {
		v163 = v155
		goto L30
	} else {
		goto L39
	}
L38:
	;
	v163 = v155
	goto L30
L39:
	;
	v158 = v142 + int32(1)
	if v158 != v136 {
		v142 = v158
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v171 = F_pull_varnos(m, l0, v170)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return int32(0)
L43:
	;
	v175 = int32(0)
	if v171 == v175 {
		v216 = v175
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v216 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L45:
	;
	goto L44
L46:
	;
	if v34 == int32(0) {
		v216 = v175
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v184 < v185 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v187 = v184
	goto L50
L49:
	;
	v187 = v185
	goto L50
L50:
	;
	if v187 <= int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v190 = int32(1)
	goto L53
L52:
	;
	v190 = v187
	goto L53
L53:
	;
	v191 = int32(8)
	v196 = int32(0)
	goto L54
L54:
	;
	v203 = v196 << (uint(int32(2)) % 32)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v34+v191+v203)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v203+(v171+v191))))
	v208 = v205 & v207
	v210 = base.B2i32(v208 != int32(0))
	if v208 != 0 {
		v216 = v210
		goto L45
	} else {
		goto L56
	}
L55:
	;
	v216 = v210
	goto L45
L56:
	;
	v212 = v196 + int32(1)
	if v212 != v190 {
		v196 = v212
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v224 = F_lappend(m, v104, v221)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L42
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v227 = F_pull_vars_of_level(m, v221, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L42
	} else {
		goto L63
	}
L61:
	;
	v368 = v224
	goto L27
L62:
	;
	F_list_free(m, v227)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L42
	} else {
		goto L94
	}
L63:
	;
	if v227 == int32(0) {
		v348 = v104
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v231 = int32(0)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v232 <= v231 {
		v348 = v104
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v245 = v231
	v249 = v104
	goto L66
L66:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253+v245<<(uint(int32(2))%32))))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v258 != int32(319) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v348 = v329
	goto L62
L68:
	;
	v331 = v245 + int32(1)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v331 < v332 {
		v245 = v331
		v249 = v329
		goto L66
	} else {
		goto L93
	}
L69:
	;
	v327 = F_lappend(m, v249, v257)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L42
	} else {
		goto L92
	}
L70:
	;
	if v258 != int32(6) {
		v329 = v249
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v267 = F_find_placeholder_info(m, l0, v257)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L42
	} else {
		goto L76
	}
L73:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	v265 = F_bms_is_member(m, v263, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L42
	} else {
		goto L74
	}
L74:
	;
	if v265 != 0 {
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v329 = v249
	goto L68
L76:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	v271 = int32(0)
	if v269 == v271 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v324 == int32(0) {
		v329 = v249
		goto L68
	} else {
		goto L91
	}
L78:
	;
	v324 = int32(1)
	goto L77
L79:
	;
	goto L80
L80:
	;
	if v270 == int32(0) {
		v315 = v271
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v324 = v315
	goto L77
L82:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v281 < v280 {
		v315 = v271
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v283 = int32(1)
	if v280 <= v283 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v286 = v283
	goto L86
L85:
	;
	v286 = v280
	goto L86
L86:
	;
	v287 = int32(8)
	v292 = int32(0)
	goto L87
L87:
	;
	v299 = v292 << (uint(int32(2)) % 32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v269+v287+v299)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299+(v270+v287))))
	v306 = v301 & (v303 ^ int32(-1))
	v308 = base.B2i32(v306 == int32(0))
	if v306 != 0 {
		v315 = v308
		goto L81
	} else {
		goto L89
	}
L88:
	;
	v315 = v308
	goto L81
L89:
	;
	v310 = v292 + int32(1)
	if v310 != v286 {
		v292 = v310
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L69
L92:
	;
	v329 = v327
	goto L68
L93:
	;
	goto L67
L94:
	;
	v368 = v348
	goto L27
L95:
	;
	goto L26
L96:
	;
	v399 = int32(0)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+8)))
	if base.B2i32(v400&int32(1) == v399)&base.B2i32(l5&int32(-2) == int32(4)) != 0 {
		v1158 = v399
		goto L1
	} else {
		goto L102
	}
L97:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	if v395 != 0 {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v396|v390 != 0 {
		goto L96
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v1158 = int32(0)
	goto L1
L102:
	;
	if v400&int32(1) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v473 = F_contain_volatile_functions(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L42
	} else {
		goto L115
	}
L104:
	;
	if v394 == int32(0) {
		v1158 = v399
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v416 == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v419 = int32(0)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v420 <= v419 {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v394)+20))
	v431 = v419
	goto L108
L108:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v416)+12))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v442+v431<<(uint(int32(2))%32))))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+56))
	v448 = F_bms_is_member(m, v447, v423)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L42
	} else {
		goto L110
	}
L109:
	;
	v1158 = v399
	goto L1
L110:
	;
	if v448 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v451 = v431 + int32(1)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v451 < v452 {
		v431 = v451
		goto L108
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	goto L109
L114:
	;
	goto L103
L115:
	;
	if v473 != 0 {
		v1158 = v399
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v475 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v531 = int32(0)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v533 == v531 {
		v600 = v531
		goto L127
	} else {
		goto L128
	}
L118:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	if v478 <= int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v491 = v399
	goto L120
L120:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v499+v491<<(uint(int32(2))%32))))
	v504 = F_contain_volatile_functions(m, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L42
	} else {
		goto L122
	}
L121:
	;
	v1158 = int32(0)
	goto L1
L122:
	;
	if v504 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v509 = v491 + int32(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	if v509 < v510 {
		v491 = v509
		goto L120
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L121
L126:
	;
	goto L117
L127:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
	if v611 != 0 {
		goto L140
	} else {
		goto L141
	}
L128:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	if v536 == int32(0) {
		v600 = v533
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v539 = int32(0)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v540 <= v539 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v600 = v592
	goto L127
L131:
	;
	v550 = v539
	goto L132
L132:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v536)+12))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v561+v550<<(uint(int32(2))%32))))
	v566 = F_contain_volatile_functions(m, v565)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L42
	} else {
		goto L134
	}
L133:
	;
	v1158 = v531
	goto L1
L134:
	;
	if v566 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v571 = v550 + int32(1)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	if v571 < v572 {
		v550 = v571
		goto L132
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	goto L133
L138:
	;
	goto L130
L139:
	;
	if v1081 == int32(0) {
		v1158 = v531
		goto L1
	} else {
		goto L249
	}
L140:
	;
	v612 = v611
	goto L142
L141:
	;
	v612 = l2
	goto L142
L142:
	;
	v613 = int32(0)
	v615 = v21 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v615))) = v613
	v619 = v21 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = v613
	v623 = v21 + int32(7)
	*(*uint8)(unsafe.Add(mBase, uint32(v623))) = uint8(v613)
	if v600 == v613 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v1081 = v1052
	goto L139
L144:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	F_list_free(m, v1038)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L42
	} else {
		goto L247
	}
L145:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v954 = F_list_concat(m, v390, v953)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L42
	} else {
		goto L229
	}
L146:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v600)+16))
	if v628 == int32(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v631 <= int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v639 = v613
	goto L149
L149:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v652+v639<<(uint(int32(2))%32))))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	if v658 != int32(17) {
		goto L144
	} else {
		goto L151
	}
L150:
	;
	goto L145
L151:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v657)+28))
	if v661 == int32(0) {
		goto L144
	} else {
		goto L152
	}
L152:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v661)+4))
	if v664 != int32(2) {
		goto L144
	} else {
		goto L153
	}
L153:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v656)+44))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v612)+8))
	v670 = int32(0)
	if v668 == v670 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v906+v656)))
	if v908 == int32(0) {
		goto L144
	} else {
		goto L218
	}
L155:
	;
	v900 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v656)+120)) = uint8(v900)
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v657)+28))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)+12))
	v905 = v903
	v906 = int32(160)
	goto L154
L156:
	;
	if v723 != 0 {
		goto L170
	} else {
		goto L171
	}
L157:
	;
	v723 = int32(1)
	goto L156
L158:
	;
	goto L159
L159:
	;
	if v669 == int32(0) {
		v714 = v670
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v723 = v714
	goto L156
L161:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v680 < v679 {
		v714 = v670
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v682 = int32(1)
	if v679 <= v682 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v685 = v682
	goto L165
L164:
	;
	v685 = v679
	goto L165
L165:
	;
	v686 = int32(8)
	v691 = int32(0)
	goto L166
L166:
	;
	v698 = v691 << (uint(int32(2)) % 32)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v668+v686+v698)))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v698+(v669+v686))))
	v705 = v700 & (v702 ^ int32(-1))
	v707 = base.B2i32(v705 == int32(0))
	if v705 != 0 {
		v714 = v707
		goto L160
	} else {
		goto L168
	}
L167:
	;
	v714 = v707
	goto L160
L168:
	;
	v709 = v691 + int32(1)
	if v709 != v685 {
		v691 = v709
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v656)+48))
	v725 = int32(0)
	if v724 == v725 {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	goto L172
L172:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v656)+44))
	v780 = int32(0)
	if v779 == v780 {
		goto L189
	} else {
		goto L190
	}
L173:
	;
	if v778 != 0 {
		goto L155
	} else {
		goto L187
	}
L174:
	;
	v778 = int32(1)
	goto L173
L175:
	;
	goto L176
L176:
	;
	if v667 == int32(0) {
		v769 = v725
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v778 = v769
	goto L173
L178:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v724)+4))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v735 < v734 {
		v769 = v725
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v737 = int32(1)
	if v734 <= v737 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v740 = v737
	goto L182
L181:
	;
	v740 = v734
	goto L182
L182:
	;
	v741 = int32(8)
	v746 = int32(0)
	goto L183
L183:
	;
	v753 = v746 << (uint(int32(2)) % 32)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v724+v741+v753)))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v753+(v667+v741))))
	v760 = v755 & (v757 ^ int32(-1))
	v762 = base.B2i32(v760 == int32(0))
	if v760 != 0 {
		v769 = v762
		goto L177
	} else {
		goto L185
	}
L184:
	;
	v769 = v762
	goto L177
L185:
	;
	v764 = v746 + int32(1)
	if v764 != v740 {
		v746 = v764
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	goto L172
L188:
	;
	if v833 == int32(0) {
		goto L144
	} else {
		goto L202
	}
L189:
	;
	v833 = int32(1)
	goto L188
L190:
	;
	goto L191
L191:
	;
	if v667 == int32(0) {
		v824 = v780
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v833 = v824
	goto L188
L193:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v790 < v789 {
		v824 = v780
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v792 = int32(1)
	if v789 <= v792 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v795 = v792
	goto L197
L196:
	;
	v795 = v789
	goto L197
L197:
	;
	v796 = int32(8)
	v801 = int32(0)
	goto L198
L198:
	;
	v808 = v801 << (uint(int32(2)) % 32)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v779+v796+v808)))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v808+(v667+v796))))
	v815 = v810 & (v812 ^ int32(-1))
	v817 = base.B2i32(v815 == int32(0))
	if v815 != 0 {
		v824 = v817
		goto L192
	} else {
		goto L200
	}
L199:
	;
	v824 = v817
	goto L192
L200:
	;
	v819 = v801 + int32(1)
	if v819 != v795 {
		v801 = v819
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v656)+48))
	v837 = int32(0)
	if v836 == v837 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v890 == int32(0) {
		goto L144
	} else {
		goto L217
	}
L204:
	;
	v890 = int32(1)
	goto L203
L205:
	;
	goto L206
L206:
	;
	if v669 == int32(0) {
		v881 = v837
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v890 = v881
	goto L203
L208:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v836)+4))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v847 < v846 {
		v881 = v837
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v849 = int32(1)
	if v846 <= v849 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v852 = v849
	goto L212
L211:
	;
	v852 = v846
	goto L212
L212:
	;
	v853 = int32(8)
	v858 = int32(0)
	goto L213
L213:
	;
	v865 = v858 << (uint(int32(2)) % 32)
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v836+v853+v865)))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v865+(v669+v853))))
	v872 = v867 & (v869 ^ int32(-1))
	v874 = base.B2i32(v872 == int32(0))
	if v872 != 0 {
		v881 = v874
		goto L207
	} else {
		goto L215
	}
L214:
	;
	v881 = v874
	goto L207
L215:
	;
	v876 = v858 + int32(1)
	if v876 != v852 {
		v858 = v876
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v893 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v656)+120)) = uint8(v893)
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v657)+28))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v895)+12))
	v905 = v896 + int32(4)
	v906 = int32(164)
	goto L154
L218:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v905)))
	v913 = F_list_member(m, v911, v912)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L42
	} else {
		goto L219
	}
L219:
	;
	if v913 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v918 = F_lappend_oid(m, v917, v908)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L42
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v656)+124))
	if v925 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = v918
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	v922 = F_lappend(m, v921, v912)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L42
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615))) = v922
	goto L222
L225:
	;
	v928 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v623))) = uint8(v928)
	goto L227
L226:
	;
	goto L227
L227:
	;
	v931 = v639 + int32(1)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v931 < v932 {
		v639 = v931
		goto L149
	} else {
		goto L228
	}
L228:
	;
	goto L150
L229:
	;
	if v954 == int32(0) {
		v1081 = int32(1)
		goto L139
	} else {
		goto L230
	}
L230:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v954)+4))
	if v959 <= int32(0) {
		v1052 = int32(1)
		goto L143
	} else {
		goto L231
	}
L231:
	;
	v968 = int32(0)
	goto L232
L232:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v954)+12))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v981+v968<<(uint(int32(2))%32))))
	v986 = F_contain_volatile_functions(m, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L42
	} else {
		goto L234
	}
L233:
	;
	v1052 = v1013
	goto L143
L234:
	;
	if v986 != 0 {
		goto L144
	} else {
		goto L235
	}
L235:
	;
	v988 = F_exprType(m, v985)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L42
	} else {
		goto L236
	}
L236:
	;
	v991 = F_lookup_type_cache(m, v988, int32(17))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L42
	} else {
		goto L237
	}
L237:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v991)+68))
	if v993 == int32(0) {
		goto L144
	} else {
		goto L238
	}
L238:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v991)+52))
	if v996 == int32(0) {
		goto L144
	} else {
		goto L239
	}
L239:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	v1000 = F_list_member(m, v999, v985)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L42
	} else {
		goto L240
	}
L240:
	;
	if v1000 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v991)+52))
	v1006 = F_lappend_oid(m, v1004, v1005)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L42
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1013 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v623))) = uint8(v1013)
	v1017 = v968 + v1013
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v954)+4))
	if v1017 < v1018 {
		v968 = v1017
		goto L232
	} else {
		goto L246
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = v1006
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	v1010 = F_lappend(m, v1009, v985)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L42
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615))) = v1010
	goto L243
L246:
	;
	goto L233
L247:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	F_list_free(m, v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L42
	} else {
		goto L248
	}
L248:
	;
	v1052 = int32(0)
	goto L143
L249:
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
		goto L42
	} else {
		goto L250
	}
L250:
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
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+21)))
	v1106 = v1104
	goto L253
L252:
	;
	v1106 = int32(0)
	goto L253
L253:
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
	v1120 = float64(1e+100)
	if base.F64_gt(v1088, v1120) != 0 {
		v1132 = v1120
		goto L255
	} else {
		goto L256
	}
L254:
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
	v1158 = v1090
	goto L1
L255:
	;
	goto L254
L256:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1088)&int64(9223372036854775807)) {
		v1132 = v1120
		goto L255
	} else {
		goto L257
	}
L257:
	;
	v1128 = float64(1)
	if base.F64_le(v1088, v1128) != 0 {
		v1132 = v1128
		goto L255
	} else {
		goto L258
	}
L258:
	;
	v1132 = base.F64_nearest(v1088)
	goto L255
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
	var v22 int32
	_ = v22
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
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
	v22 = v2
	v25 = v2
	goto L6
L4:
	;
	v76 = v2
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(48)+v25<<(uint(int32(2))%32))))
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
	v76 = v69
	goto L5
L8:
	;
	v72 = v25 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v72 < v73 {
		v22 = v69
		v25 = v72
		goto L6
	} else {
		goto L25
	}
L9:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+16)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v61 = F_IndexAmTranslateStrategy(m, v59, v57, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L10:
	;
	v49 = F_GetIndexAmRoutineByAmId(m, v35, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	switch v35 - int32(403) {
	case 0:
		v57 = v35
		goto L9
	case 1:
		goto L10
	case 2:
		v69 = v22
		goto L8
	default:
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v35 == int32(2742) {
		v69 = v22
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
	v69 = v22
	goto L8
L16:
	;
	if v35 == int32(3580) {
		v69 = v22
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v35 == int32(4000) {
		v69 = v22
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+10)))
	F_pfree(m, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v51 != int32(1) {
		v69 = v22
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v57 = v56
	goto L9
L22:
	;
	if v61 != int32(3) {
		v69 = v22
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v66 = F_lappend_oid(m, v22, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v69 = v66
	goto L8
L25:
	;
	goto L7
L26:
	;
	return v76
}
func F_get_negator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+96))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_opname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
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
	v101 = v5
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v20)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
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
	v101 = int32(0)
	goto L5
L8:
	;
	v93 = v30 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v93 < v94 {
		v30 = v93
		goto L6
	} else {
		goto L26
	}
L9:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+16)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v72 = F_IndexAmTranslateStrategy(m, v70, v68, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L22
	}
L10:
	;
	v60 = F_GetIndexAmRoutineByAmId(m, v46, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	if v46 == int32(783) {
		goto L8
	} else {
		goto L18
	}
L12:
	;
	switch v46 - int32(403) {
	case 0:
		v68 = v46
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
	if v46 == int32(2742) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v46 == int32(3580) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v46 == int32(4000) {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	goto L10
L19:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)))
	F_pfree(m, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v62 != int32(1) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	v68 = v67
	goto L9
L22:
	;
	if v72&int32(-5) != int32(1) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v78 != v79 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v83
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v72
	F_ReleaseCatCacheList(m, v20)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	return int32(1)
L26:
	;
	goto L7
L27:
	;
	return v101
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
		goto L13
	} else {
		goto L31
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L28
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v45 = F_get_relids_in_jointree(m, v44, l1, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L13
	} else {
		goto L17
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
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v86 = v4
	goto L1
L9:
	;
	goto L10
L10:
	;
	v23 = v19
	v26 = v4
	goto L11
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v23<<(uint(int32(2))%32))))
	v34 = F_get_relids_in_jointree(m, v33, l1, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v86 = v38
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	v38 = F_bms_join(m, v26, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v41 = v23 + int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v41 < v42 {
		v23 = v41
		v26 = v38
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v48 = F_get_relids_in_jointree(m, v47, l1, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v50 = F_bms_join(m, v45, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v52 == int32(0) {
		v86 = v50
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v55 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if l2 == int32(0) {
		v86 = v50
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if l1 == int32(0) {
		v86 = v50
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v60 = F_bms_add_member(m, v50, v52)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v86 = v60
	goto L1
L26:
	;
	v64 = F_bms_add_member(m, v50, v52)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v86 = v64
	goto L1
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v70
	F_errmsg_internal(m, int32(_a_F_get_relids_in_jointree_0), v9)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_get_relids_in_jointree_1), int32(_a_F_get_relids_in_jointree_2), int32(_a_F_get_relids_in_jointree_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
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
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
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
	v51 = v2
	goto L12
L10:
	;
	v249 = v2
	goto L11
L11:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L57
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v60 = v58 + v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	switch v61 - int32(1247) {
	case 0:
		goto L16
	default:
		v236 = v51
		goto L14
	case 12:
		goto L15
	}
L13:
	;
	v249 = v236
	goto L11
L14:
	;
	v243 = F_systable_getnext(m, v43)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L55
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v77 <= int32(0) {
		v236 = v51
		goto L14
	} else {
		goto L24
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v65 = F_get_typtype(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v65 == int32(100) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = F_get_rels_with_domain(m, v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_find_composite_type_dependencies(m, v67, int32(0), v16)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v72 = F_list_concat(m, v51, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v236 = v72
	goto L14
L23:
	;
	v236 = v51
	goto L14
L24:
	;
	if v51 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+48))
	v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v167)+120)))
	if v168 < v161 {
		v236 = v159
		goto L14
	} else {
		goto L46
	}
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v124 = F_relation_open(m, v122, int32(5))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L27:
	;
	v82 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v82 < v83 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v87 = v83
	goto L30
L29:
	;
	v87 = v82
	goto L30
L30:
	;
	v89 = v82
	goto L31
L31:
	;
	if v89 == v87 {
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v157 = v106
	v159 = v51
	v161 = v77
	goto L25
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32)+v104)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+56))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v108 != v109 {
		v89 = v89 + int32(1)
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v140 = F_palloc(m, int32(12))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L43
	}
L36:
	;
	F_relation_close(m, v124, int32(5))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L42
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+48))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+72))
	if v127 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_find_composite_type_dependencies(m, v127, int32(0), v16)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v132 = v126
	goto L40
L40:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
	switch v133 - int32(109) {
	case 0, 5:
		goto L35
	default:
		goto L36
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v124)+48))
	v132 = v131
	goto L40
L42:
	;
	v236 = v51
	goto L14
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v124
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v124)+48))
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145)+120)))
	v149 = F_palloc(m, v146<<(uint(int32(2))%32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v149
	v152 = F_lappend(m, v51, v140)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v157 = v140
	v159 = v152
	v161 = v154
	goto L25
L46:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+52))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v177 = v170 + v171<<(uint(int32(4))%32) + v161*int32(100)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+11)))
	if v178 != 0 {
		v236 = v159
		goto L14
	} else {
		goto L47
	}
L47:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177-int32(80))+68))
	if v181 != l0 {
		v236 = v159
		goto L14
	} else {
		goto L48
	}
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v183 + int32(1)
	if v183 <= int32(0) {
		v216 = v183
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v226+v216<<(uint(int32(2))%32)))) = v230
	v236 = v159
	goto L14
L50:
	;
	v190 = v183
	goto L51
L51:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v203 = v200 + v190<<(uint(int32(2))%32)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v203-int32(4))))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v206 <= v207 {
		v216 = v190
		goto L49
	} else {
		goto L53
	}
L52:
	;
	v216 = int32(0)
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v206
	v210 = int32(1)
	if v210 < v190 {
		v190 = v190 - v210
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	if v243 != 0 {
		v48 = v243
		v51 = v236
		goto L12
	} else {
		goto L56
	}
L56:
	;
	goto L13
L57:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	m.G0 = v14 + int32(96)
	return v249
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
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 != 0 {
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		if v36 == int32(1) {
			v39 = F_cstring_to_text(m, l1)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v41 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
				return int32(0)
			}
		} else {
			return int32(0)
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v8 != 0 {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
			if v36 == int32(1) {
				v39 = F_cstring_to_text(m, l1)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
					return int32(0)
				}
			} else {
				return int32(0)
			}
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
			if l2 != int32(1) {
				if l2 != int32(11) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
					v29 = F_cstring_to_text_with_len(m, v26, v27-v26)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
						if v36 == int32(1) {
							v39 = F_cstring_to_text(m, l1)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
								return int32(0)
							}
						} else {
							return int32(0)
						}
					}
				} else {
					if v9&int32(1) == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						v29 = F_cstring_to_text_with_len(m, v26, v27-v26)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
							if v36 == int32(1) {
								v39 = F_cstring_to_text(m, l1)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
									return int32(0)
								}
							} else {
								return int32(0)
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
						if v36 == int32(1) {
							v39 = F_cstring_to_text(m, l1)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
								return int32(0)
							}
						} else {
							return int32(0)
						}
					}
				}
			} else {
				if v9&int32(1) == int32(0) {
					if l2 != int32(11) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						v29 = F_cstring_to_text_with_len(m, v26, v27-v26)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
							if v36 == int32(1) {
								v39 = F_cstring_to_text(m, l1)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
									return int32(0)
								}
							} else {
								return int32(0)
							}
						}
					} else {
						if v9&int32(1) == int32(0) {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
							v29 = F_cstring_to_text_with_len(m, v26, v27-v26)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v29
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
								if v36 == int32(1) {
									v39 = F_cstring_to_text(m, l1)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v41 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
										return int32(0)
									}
								} else {
									return int32(0)
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
							if v36 == int32(1) {
								v39 = F_cstring_to_text(m, l1)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
									return int32(0)
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v16 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v16)
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
					if v36 == int32(1) {
						v39 = F_cstring_to_text(m, l1)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v41)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
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
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L63
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
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
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103<<(uint(int32(2))%32))+uint32(_c_F_get_setop_query[1])))
	F_appendStringInfoString(m, v12, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v113 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_get_setop_query_7))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v121 == int32(142) {
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
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	v128 = int32(0)
	goto L51
L51:
	;
	v130 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_4), v128, v130, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	v128 = int32(8)
	goto L51
L53:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)))
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)) = uint8(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_setop_query(m, v137, l1, l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+33)) = uint8(v134)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v141&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v144 - v128
	goto L57
L56:
	;
	goto L57
L57:
	;
	if v121 != int32(142) {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v150 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_setop_query_8), v150, v150, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	goto L9
L60:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v167
	F_errmsg_internal(m, int32(_a_F_get_setop_query_5), v10+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_get_setop_query_1), int32(_a_F_get_setop_query_6), int32(_a_F_get_setop_query_3))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
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
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v183
	F_errmsg_internal(m, int32(_a_F_get_setop_query_0), v10)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_get_setop_query_1), int32(_a_F_get_setop_query_2), int32(_a_F_get_setop_query_3))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
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
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_DeconstructQualifiedName(m, l0, v10+int32(12), v10+int32(8))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v106
L4:
	;
	if l1 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L26
	}
L5:
	;
	v83 = int32(0)
	goto L4
L6:
	;
	v22 = F_LookupExplicitNamespace(m, v20, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = int32(0)
	goto L12
L11:
	;
	v24 = l1
	goto L12
L12:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v27 = int32(0)
	v29 = F_GetSysCacheOid(m, int32(63), v26, v22, v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v83 = v29
	goto L4
L15:
	;
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_get_statistics_object_oid[0]))
	if v35 == v33 {
		v83 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 <= v38 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v46 = v38
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_get_statistics_object_oid[1]))
	if v54 != v56 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L5
L20:
	;
	v59 = int32(0)
	v61 = F_GetSysCacheOid(m, int32(63), v42, v54, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v65 = v46 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v65 < v66 {
		v46 = v65
		goto L18
	} else {
		goto L25
	}
L23:
	;
	if v61 != 0 {
		v106 = v61
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L19
L26:
	;
	if v83 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v91 = F_NameListToString(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v91
	F_errmsg(m, int32(_a_F_get_statistics_object_oid_0), v10)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_get_statistics_object_oid_1), int32(2620), int32(_a_F_get_statistics_object_oid_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_th(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0&int32(3) == int32(0) {
		v34 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v68 = l0 + v67
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68-int32(1)))))
	if base.Ui32((v71-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v67 = v59 - l0
	goto L1
L3:
	;
	v38 = v34
	goto L12
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v23 = l0
	goto L8
L8:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v59 = v27
	goto L2
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v53 = v38
	goto L15
L14:
	;
	goto L13
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v59 = v53
	goto L2
L17:
	;
	goto L16
L18:
	;
	if int32(2) <= v67 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L42
	} else {
		goto L43
	}
L21:
	;
	m.G0 = v9 + int32(16)
	return v107
L22:
	;
	if l1 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68-int32(2)))))
	if v82 == int32(49) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	switch v71 - int32(49) {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	default:
		goto L22
	}
L26:
	;
	goto L25
L27:
	;
	if l1 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	if l1 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	if l1 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v91 = int32(_a_F_get_th_0)
	goto L32
L31:
	;
	v91 = int32(_a_F_get_th_1)
	goto L32
L32:
	;
	v107 = v91
	goto L21
L33:
	;
	v96 = int32(_a_F_get_th_2)
	goto L35
L34:
	;
	v96 = int32(_a_F_get_th_3)
	goto L35
L35:
	;
	v107 = v96
	goto L21
L36:
	;
	v101 = int32(_a_F_get_th_4)
	goto L38
L37:
	;
	v101 = int32(_a_F_get_th_5)
	goto L38
L38:
	;
	v107 = v101
	goto L21
L39:
	;
	v106 = int32(_a_F_get_th_6)
	goto L41
L40:
	;
	v106 = int32(_a_F_get_th_7)
	goto L41
L41:
	;
	v107 = v106
	goto L21
L42:
	;
	return int32(0)
L43:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg(m, int32(_a_F_get_th_8), v9)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_get_th_9), int32(1572), int32(_a_F_get_th_10))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
					v32 = int32(4)
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
					if v34&int32(254) == int32(2) {
						v43 = v32
					} else {
						v43 = base.B2i32(v34 == int32(18)) << (uint(v32) % 32)
					}
					if v34 == int32(1) {
						v46 = v32
					} else {
						v46 = v43
					}
					v57 = v46
				} else {
					v47 = int32(1)
					if v28 != 0 {
						v57 = int32(base.Ui32(v26)>>(uint(v47)%32)) - v47
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
						v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v59 = *(*int32)(unsafe.Add(mBase, _c_F_get_worker[0]))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
				v62 = F_makeJsonLexContextCstringLen(m, int32(0), v29, v57, v60, int32(1))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l3
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v5)
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v62
					v69 = F_palloc0(m, l3)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v69
						v74 = F_palloc(m, l3<<(uint(int32(2))%32))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v74
							if int32(0) < l3 {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
								v80 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v80)
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
								v96 = int32(1351)
								v97 = int32(36)
								*(*int32)(unsafe.Add(mBase, uint32(v97+v12))) = v96
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(1351)
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
								if l3 != 0 {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(1352)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(1353)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(1354)
									v96 = int32(1355)
									v97 = int32(16)
									*(*int32)(unsafe.Add(mBase, uint32(v97+v12))) = v96
								}
							}
							if l1 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1356)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(1357)
							} else {
							}
							if l2 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(1358)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(1359)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(1352)
							} else {
							}
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
							v112 = F_pg_parse_json(m, v111, v12)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								if v112 != 0 {
									F_json_errsave_error(m, v112, v111, int32(0))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
										F_freeJsonLexContext(m, v117)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
											return v120
										}
									}
								} else {
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									F_freeJsonLexContext(m, v117)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
										return v120
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = l0
	goto L3
L1:
	;
	m.G0 = v11 + int32(16)
	return v137
L2:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v112))) = uint8(v117)
	v119 = v108
	goto L38
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.Ui32(v21-int32(9)) < base.Ui32(int32(5)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v30 = v13
	v33 = v21
	v34 = v4
	v36 = v4
	goto L10
L5:
	;
	goto L4
L6:
	;
	v13 = v13 + int32(1)
	goto L3
L7:
	;
	if v21 == int32(32) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v21 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v108 = v13
	v112 = v21
	goto L2
L10:
	;
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v108 = v106
	v112 = v102
	goto L2
L12:
	;
	v106 = v103 + int32(1)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v107 != 0 {
		v30 = v106
		v33 = v107
		v34 = v102
		v36 = v104
		goto L10
	} else {
		goto L37
	}
L13:
	;
	if int32(63) <= v34 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v71 = v30
	goto L13
L15:
	;
	if v33 != int32(34) {
		goto L14
	} else {
		goto L22
	}
L16:
	;
	if v33 == int32(34) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.I32_extend8_s(v33) < int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L19
L19:
	;
	if v33 == int32(95) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if base.B2i32(base.Ui32(v33-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v33|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
		v108 = v30
		v112 = v34
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	if v36 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v102 = v34
	v103 = v30
	v104 = int32(1)
	goto L12
L24:
	;
	goto L25
L25:
	;
	v65 = v30 + int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v66 == int32(34) {
		v71 = v65
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v102 = v34
	v103 = v30
	v104 = int32(0)
	goto L12
L27:
	;
	v74 = int32(0)
	v75 = F_errsave_start(m, l2)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v34))) = uint8(v33)
	v102 = v34 + int32(1)
	v103 = v71
	v104 = v36
	goto L12
L30:
	;
	return int32(0)
L31:
	;
	if v75 == int32(0) {
		v137 = v74
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(34103428))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_getid_0), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(64)
	F_errdetail(m, int32(_a_F_getid_1), v11)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	F_errsave_finish(m, l2, int32(_a_F_getid_2), int32(206), int32(_a_F_getid_3))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v137 = v74
	goto L1
L37:
	;
	goto L11
L38:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v127-int32(9)))&base.B2i32(v127 != int32(32)) != 0 {
		v137 = v119
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v119 = v119 + int32(1)
	goto L38
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
	var v28 int32
	_ = v28
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
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_c_F_getinternalerrposition[1])))
		return v28
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	v3 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v10 - int32(43) {
	case 0:
		goto L4
	default:
		v20 = l0
		v21 = v10
		v23 = v3
		goto L2
	case 2:
		v14 = int32(1)
		goto L3
	}
L1:
	;
	return v155
L2:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v21)-int32(48)) {
		v155 = v3
		goto L1
	} else {
		goto L6
	}
L3:
	;
	v16 = l0 + int32(1)
	if v16 == int32(0) {
		v155 = v3
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v14 = int32(0)
	goto L3
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v20 = v16
	v21 = v19
	v23 = v14
	goto L2
L6:
	;
	v30 = v20
	v32 = v21
	v33 = int32(0)
	goto L7
L7:
	;
	v43 = base.I32_extend8_s(v32) + v33*int32(10) - int32(48)
	if int32(167) < v43 {
		v155 = v3
		goto L1
	} else {
		goto L9
	}
L8:
	;
	if v43 < int32(0) {
		v155 = v3
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v47 = v30 + int32(1)
	v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47))))
	if base.Ui32(v48-int32(48)) < base.Ui32(int32(10)) {
		v30 = v47
		v32 = v48
		v33 = v43
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v56 = v43 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v58 != int32(58) {
		v140 = v47
		v145 = v56
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v23 != 0 {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v62 = v30 + int32(2)
	if v62 == int32(0) {
		v155 = v3
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	if base.Ui32(int32(9)) < base.Ui32(v65-int32(48)) {
		v155 = v3
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v71 = v62
	v73 = int32(0)
	v74 = v65
	goto L16
L16:
	;
	v84 = base.I32_extend8_s(v74) + v73*int32(10) - int32(48)
	if int32(59) < v84 {
		v155 = v3
		goto L1
	} else {
		goto L18
	}
L17:
	;
	if v84 < int32(0) {
		v155 = v3
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v88 = v71 + int32(1)
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88))))
	if base.Ui32(v89-int32(48)) < base.Ui32(int32(10)) {
		v71 = v88
		v73 = v84
		v74 = v89
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v98 = v84*int32(60) + v56
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v100 != int32(58) {
		v140 = v88
		v145 = v98
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v104 = v71 + int32(2)
	if v104 == int32(0) {
		v155 = v3
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v104))))
	if base.Ui32(int32(9)) < base.Ui32(v107-int32(48)) {
		v155 = v3
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v113 = v104
	v115 = int32(0)
	v116 = v107
	goto L24
L24:
	;
	v126 = base.I32_extend8_s(v116) + v115*int32(10) - int32(48)
	if int32(60) < v126 {
		v155 = v3
		goto L1
	} else {
		goto L26
	}
L25:
	;
	if v126 < int32(0) {
		v155 = v3
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v130 = v113 + int32(1)
	v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v130))))
	if base.Ui32(v131-int32(48)) < base.Ui32(int32(10)) {
		v113 = v130
		v115 = v126
		v116 = v131
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v138 = v126 + v98
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v138
	v140 = v130
	v145 = v138
	goto L12
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0) - v145
	goto L31
L30:
	;
	goto L31
L31:
	;
	v155 = v140
	goto L1
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
	var v99 int32
	_ = v99
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
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
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
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int64
	_ = v288
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 < v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
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
		v183 = v23
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
	v183 = v23
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v170 + int32(32)
	v183 = v170
	goto L1
L13:
	;
	v99 = v88
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
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
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+8)))
	if v110&int32(4) == int32(0) {
		v170 = v99
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v116 = v99 + int32(32)
	if base.Ui32(v116) < base.Ui32(v92) {
		v99 = v116
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
	v170 = v132
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
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v193 = v188
	v194 = v189
	goto L44
L42:
	;
	goto L43
L43:
	;
	v225 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v225 < v227 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v193)+24))
	v201 = base.I32_extend16_s(v194)
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v200+v201<<(uint(int32(2))%32)))) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v193)+28))
	v210 = v207 + v201<<(uint(int32(3))%32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v205
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210)+4)))
	if v211 != 0 {
		v193 = v211
		v194 = v214
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
	v235 = v227
	v237 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	if v334&int32(2) == int32(0) {
		v348 = v334
		goto L66
	} else {
		goto L67
	}
L50:
	;
	v242 = v237 << (uint(int32(2)) % 32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v183)+24))
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
	if v246 != v183 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v315 = v235
	goto L54
L54:
	;
	v322 = v237 + int32(1)
	if v322 < v315 {
		v235 = v315
		v237 = v322
		goto L50
	} else {
		goto L65
	}
L55:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v183)+24))
	v302 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v300+v242))) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v183)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v304+v237<<(uint(int32(3))%32)))) = v302
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v315 = v310
	goto L54
L56:
	;
	v256 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+16)))
	v260 = v256
	v261 = v246
	goto L60
L57:
	;
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+16)))
	if v237 != v248 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v183)+28))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v250+v237<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v245)+12)) = v254
	goto L55
L59:
	;
	v281 = int32(3)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v183)+28))
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v284+v237<<(uint(v281)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v280+v279<<(uint(v281)%32)))) = v288
	goto L55
L60:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261)+28))
	v270 = v267 + v260<<(uint(int32(3))%32)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v271 == int32(0) {
		v279 = v260
		v280 = v267
		goto L59
	} else {
		goto L62
	}
L61:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v261)+28))
	v279 = base.I32_extend16_s(v260)
	v280 = v277
	goto L59
L62:
	;
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v270)+4)))
	if v271 != v183 {
		v260 = v274
		v261 = v271
		goto L60
	} else {
		goto L63
	}
L63:
	;
	if v274 != v237 {
		v260 = v274
		v261 = v271
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	goto L51
L66:
	;
	if v348&int32(8) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L67:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v339 == v340 {
		v348 = v334
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(v339) <= base.Ui32(v340) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v344 = v340
	goto L71
L70:
	;
	v344 = int32(0)
	goto L71
L71:
	;
	if v344 != 0 {
		v348 = v334
		goto L66
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v339
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v348 = v346
	goto L66
L73:
	;
	return v183
L74:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v354 == v355 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	if base.Ui32(v354) <= base.Ui32(v355) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v359 = v355
	goto L78
L77:
	;
	v359 = int32(0)
	goto L78
L78:
	;
	if v359 != 0 {
		goto L73
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v354
	goto L73
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
		goto L32
	} else {
		goto L33
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
		goto L25
	}
L5:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v45)
	if v14 <= v45 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v19)
	if v14 <= v19 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v107 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v24 = int32(0)
	goto L10
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v15))))
	if v32 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v107 = int32(0)
	goto L1
L12:
	;
	v35 = int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v12))))
	if v37 != v35 {
		v107 = v35
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v42 = v24 + int32(1)
	if v42 != v14 {
		v24 = v42
		goto L10
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	goto L11
L17:
	;
	v107 = int32(1)
	goto L1
L18:
	;
	goto L19
L19:
	;
	v51 = v45
	goto L20
L20:
	;
	v58 = int32(0)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v15))))
	if v60 != int32(1) {
		v107 = v58
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v107 = v65
	goto L1
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v12))))
	if v64 != 0 {
		v107 = v58
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v65 = int32(1)
	v67 = v51 + v65
	if v14 != v67 {
		v51 = v67
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v75 = v72
	goto L26
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v15))))
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v107 = int32(0)
	goto L1
L28:
	;
	v85 = v75 + int32(1)
	if v14 != v85 {
		v75 = v85
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	v107 = v69
	goto L1
L32:
	;
	return int32(0)
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v16
	F_errmsg_internal(m, int32(_a_F_ginarrayconsistent_0), v10)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_ginarrayconsistent_1), int32(215), int32(_a_F_ginarrayconsistent_2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
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
	var v103 int32
	_ = v103
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
	var v200 int32
	_ = v200
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
	var v246 int32
	_ = v246
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
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v376 int32
	_ = v376
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v749 int32
	_ = v749
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int64
	_ = v794
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int64
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1012 int64
	_ = v1012
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1124 int32
	_ = v1124
	var v1131 int64
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1285 int32
	_ = v1285
	var v1302 int32
	_ = v1302
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1334 int32
	_ = v1334
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1850 int32
	_ = v1850
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1914 int64
	_ = v1914
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1949 int64
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1983 int64
	_ = v1983
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1991 int32
	_ = v1991
	var v2007 int32
	_ = v2007
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2067 int32
	_ = v2067
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2185 int32
	_ = v2185
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2209 int32
	_ = v2209
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2288 int32
	_ = v2288
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2415 int32
	_ = v2415
	var v2444 int32
	_ = v2444
	var v2450 int32
	_ = v2450
	var v2452 int32
	_ = v2452
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2475 int32
	_ = v2475
	var v2477 int32
	_ = v2477
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2503 int32
	_ = v2503
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2516 int32
	_ = v2516
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2532 int32
	_ = v2532
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2556 int32
	_ = v2556
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2595 int32
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2612 int32
	_ = v2612
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2629 int32
	_ = v2629
	var v2635 int32
	_ = v2635
	var v2637 int32
	_ = v2637
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2648 int32
	_ = v2648
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2656 int32
	_ = v2656
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2696 int32
	_ = v2696
	var v2698 int32
	_ = v2698
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2702 int32
	_ = v2702
	var v2704 int32
	_ = v2704
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2721 int32
	_ = v2721
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2766 int32
	_ = v2766
	var v2769 int32
	_ = v2769
	var v2772 int32
	_ = v2772
	var v2775 int32
	_ = v2775
	var v2777 int32
	_ = v2777
	var v2781 int32
	_ = v2781
	var v2784 int32
	_ = v2784
	var v2815 int32
	_ = v2815
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2851 int32
	_ = v2851
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2865 int32
	_ = v2865
	var v2871 int32
	_ = v2871
	var v2873 int32
	_ = v2873
	var v2879 int32
	_ = v2879
	var v2882 int32
	_ = v2882
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2899 int32
	_ = v2899
	var v2902 int32
	_ = v2902
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2912 int32
	_ = v2912
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2937 int32
	_ = v2937
	var v2939 int32
	_ = v2939
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2946 int32
	_ = v2946
	var v2952 int32
	_ = v2952
	var v2954 int32
	_ = v2954
	var v2960 int32
	_ = v2960
	var v2961 int32
	_ = v2961
	var v2963 int32
	_ = v2963
	var v2995 int32
	_ = v2995
	var v2998 int32
	_ = v2998
	var v3031 int32
	_ = v3031
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3039 int32
	_ = v3039
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3181 int32
	_ = v3181
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3208 int32
	_ = v3208
	var v3212 int32
	_ = v3212
	var v3216 int32
	_ = v3216
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3229 int32
	_ = v3229
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3262 int32
	_ = v3262
	var v3264 int32
	_ = v3264
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3270 int32
	_ = v3270
	var v3275 int32
	_ = v3275
	var v3277 int32
	_ = v3277
	var v3280 int32
	_ = v3280
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3286 int32
	_ = v3286
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3331 int32
	_ = v3331
	var v3357 int32
	_ = v3357
	var v3358 int32
	_ = v3358
	var v3364 int32
	_ = v3364
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3397 int32
	_ = v3397
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3413 int32
	_ = v3413
	var v3436 int32
	_ = v3436
	var v3442 int32
	_ = v3442
	var v3471 int32
	_ = v3471
	var v3475 int32
	_ = v3475
	var v3477 int32
	_ = v3477
	var v3479 int32
	_ = v3479
	var v3480 int32
	_ = v3480
	var v3481 int32
	_ = v3481
	var v3485 int32
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3494 int32
	_ = v3494
	var v3503 int32
	_ = v3503
	var v3528 int32
	_ = v3528
	var v3530 int32
	_ = v3530
	var v3533 int32
	_ = v3533
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3541 int32
	_ = v3541
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3547 int32
	_ = v3547
	var v3552 int32
	_ = v3552
	var v3581 int32
	_ = v3581
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3585 int32
	_ = v3585
	var v3587 int32
	_ = v3587
	var v3591 int32
	_ = v3591
	var v3594 int32
	_ = v3594
	var v3595 int32
	_ = v3595
	var v3599 int32
	_ = v3599
	var v3628 int32
	_ = v3628
	var v3634 int32
	_ = v3634
	var v3640 int32
	_ = v3640
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3667 int32
	_ = v3667
	var v3671 int32
	_ = v3671
	var v3675 int32
	_ = v3675
	var v3677 int32
	_ = v3677
	var v3680 int32
	_ = v3680
	var v3681 int32
	_ = v3681
	var v3714 int32
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3718 int32
	_ = v3718
	var v3725 int32
	_ = v3725
	var v3726 int32
	_ = v3726
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3765 int32
	_ = v3765
	var v3766 int32
	_ = v3766
	var v3799 int32
	_ = v3799
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3810 int32
	_ = v3810
	var v3820 int32
	_ = v3820
	var v3822 int32
	_ = v3822
	var v3825 int32
	_ = v3825
	var v3834 int64
	_ = v3834
	var v3837 int32
	_ = v3837
	var v3839 int32
	_ = v3839
	var v3854 int32
	_ = v3854
	var v3856 int32
	_ = v3856
	var v3859 int32
	_ = v3859
	var v3872 int32
	_ = v3872
	var v3874 int32
	_ = v3874
	var v3875 int32
	_ = v3875
	var v3880 int32
	_ = v3880
	var v3893 int32
	_ = v3893
	var v3897 int32
	_ = v3897
	var v3899 int32
	_ = v3899
	var v3902 int32
	_ = v3902
	var v3914 int32
	_ = v3914
	var v3917 int32
	_ = v3917
	var v3918 int32
	_ = v3918
	var v3921 int32
	_ = v3921
	var v3922 int32
	_ = v3922
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3932 int64
	_ = v3932
	var v3933 int64
	_ = v3933
	var v3934 int64
	_ = v3934
	var v3936 int64
	_ = v3936
	var v3937 int64
	_ = v3937
	var v3942 int64
	_ = v3942
	var v3955 int32
	_ = v3955
	var v3958 int32
	_ = v3958
	var v3968 int32
	_ = v3968
	var v3969 int32
	_ = v3969
	var v3971 int32
	_ = v3971
	var v3972 int32
	_ = v3972
	var v3975 int32
	_ = v3975
	var v3994 int32
	_ = v3994
	var v3998 int32
	_ = v3998
	var v3999 int32
	_ = v3999
	var v4000 int32
	_ = v4000
	var v4002 int64
	_ = v4002
	var v4004 int32
	_ = v4004
	var v4008 int64
	_ = v4008
	var v4010 int32
	_ = v4010
	var v4012 int64
	_ = v4012
	var v4015 int64
	_ = v4015
	var v4016 int64
	_ = v4016
	var v4017 int64
	_ = v4017
	var v4020 int64
	_ = v4020
	var v4026 int32
	_ = v4026
	var v4028 int32
	_ = v4028
	var v4033 int32
	_ = v4033
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4037 int64
	_ = v4037
	var v4039 int32
	_ = v4039
	var v4045 int32
	_ = v4045
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4053 int32
	_ = v4053
	var v4054 int64
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4057 int64
	_ = v4057
	var v4070 int32
	_ = v4070
	var v4071 int32
	_ = v4071
	var v4072 int32
	_ = v4072
	var v4073 int32
	_ = v4073
	var v4079 int32
	_ = v4079
	var v4080 int32
	_ = v4080
	var v4084 int32
	_ = v4084
	var v4087 int32
	_ = v4087
	var v4095 int32
	_ = v4095
	var v4098 int32
	_ = v4098
	var v4099 int32
	_ = v4099
	var v4100 int32
	_ = v4100
	var v4105 int32
	_ = v4105
	var v4140 int32
	_ = v4140
	var v4141 int32
	_ = v4141
	var v4147 int32
	_ = v4147
	var v4180 int32
	_ = v4180
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4191 int32
	_ = v4191
	var v4193 int32
	_ = v4193
	var v4197 int32
	_ = v4197
	var v4216 int32
	_ = v4216
	var v4223 int32
	_ = v4223
	var v4224 int32
	_ = v4224
	var v4226 int32
	_ = v4226
	var v4230 int32
	_ = v4230
	var v4249 int32
	_ = v4249
	var v4253 int32
	_ = v4253
	var v4254 int32
	_ = v4254
	var v4255 int32
	_ = v4255
	var v4257 int64
	_ = v4257
	var v4259 int32
	_ = v4259
	var v4263 int64
	_ = v4263
	var v4265 int32
	_ = v4265
	var v4267 int64
	_ = v4267
	var v4270 int64
	_ = v4270
	var v4271 int64
	_ = v4271
	var v4272 int64
	_ = v4272
	var v4275 int64
	_ = v4275
	var v4281 int32
	_ = v4281
	var v4283 int32
	_ = v4283
	var v4286 int32
	_ = v4286
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4290 int64
	_ = v4290
	var v4292 int32
	_ = v4292
	var v4298 int32
	_ = v4298
	var v4304 int32
	_ = v4304
	var v4305 int32
	_ = v4305
	var v4306 int32
	_ = v4306
	var v4307 int64
	_ = v4307
	var v4309 int64
	_ = v4309
	var v4322 int32
	_ = v4322
	var v4323 int32
	_ = v4323
	var v4324 int32
	_ = v4324
	var v4330 int32
	_ = v4330
	var v4331 int32
	_ = v4331
	var v4339 int32
	_ = v4339
	var v4341 int32
	_ = v4341
	var v4345 int32
	_ = v4345
	var v4367 int32
	_ = v4367
	var v4370 int32
	_ = v4370
	var v4371 int32
	_ = v4371
	var v4375 int64
	_ = v4375
	var v4385 int32
	_ = v4385
	var v4389 int32
	_ = v4389
	var v4405 int32
	_ = v4405
	var v4418 int32
	_ = v4418
	var v4422 int32
	_ = v4422
	var v4423 int32
	_ = v4423
	var v4424 int64
	_ = v4424
	var v4425 int64
	_ = v4425
	var v4428 int64
	_ = v4428
	var v4434 int32
	_ = v4434
	var v4435 int32
	_ = v4435
	var v4436 int32
	_ = v4436
	var v4438 int32
	_ = v4438
	var v4441 int32
	_ = v4441
	var v4444 int32
	_ = v4444
	var v4446 int32
	_ = v4446
	var v4449 int32
	_ = v4449
	var v4451 int32
	_ = v4451
	var v4452 int32
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4462 int32
	_ = v4462
	var v4463 int32
	_ = v4463
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4474 int32
	_ = v4474
	var v4494 int32
	_ = v4494
	var v4510 int32
	_ = v4510
	var v4512 int64
	_ = v4512
	var v4515 int64
	_ = v4515
	var v4518 int64
	_ = v4518
	var v4530 int32
	_ = v4530
	var v4559 int32
	_ = v4559
	var v4563 int32
	_ = v4563
	var v4564 int32
	_ = v4564
	var v4567 int32
	_ = v4567
	var v4569 int32
	_ = v4569
	var v4571 int64
	_ = v4571
	var v4572 int64
	_ = v4572
	var v4575 int64
	_ = v4575
	var v4579 int64
	_ = v4579
	var v4581 int32
	_ = v4581
	var v4583 int32
	_ = v4583
	var v4585 int32
	_ = v4585
	var v4586 int32
	_ = v4586
	var v4588 int32
	_ = v4588
	var v4590 int32
	_ = v4590
	var v4595 int32
	_ = v4595
	var v4596 int32
	_ = v4596
	var v4629 int32
	_ = v4629
	var v4630 int32
	_ = v4630
	var v4631 int32
	_ = v4631
	var v4634 int32
	_ = v4634
	var v4636 int32
	_ = v4636
	var v4638 int32
	_ = v4638
	var v4643 int32
	_ = v4643
	var v4675 int32
	_ = v4675
	var v4676 int32
	_ = v4676
	var v4679 int32
	_ = v4679
	var v4680 int32
	_ = v4680
	var v4681 int32
	_ = v4681
	var v4686 int32
	_ = v4686
	var v4696 int32
	_ = v4696
	var v4697 int32
	_ = v4697
	var v4698 int32
	_ = v4698
	var v4703 int32
	_ = v4703
	var v4705 int32
	_ = v4705
	var v4706 int32
	_ = v4706
	var v4707 int32
	_ = v4707
	var v4713 int32
	_ = v4713
	var v4714 int32
	_ = v4714
	var v4716 int32
	_ = v4716
	var v4721 int32
	_ = v4721
	var v4722 int32
	_ = v4722
	var v4723 int32
	_ = v4723
	var v4724 int32
	_ = v4724
	var v4727 int32
	_ = v4727
	var v4734 int32
	_ = v4734
	var v4736 int32
	_ = v4736
	var v4737 int64
	_ = v4737
	var v4744 int32
	_ = v4744
	var v4754 int64
	_ = v4754
	var v4758 int64
	_ = v4758
	var v4786 int32
	_ = v4786
	var v4788 int32
	_ = v4788
	var v4791 int32
	_ = v4791
	var v4803 int32
	_ = v4803
	var v4805 int32
	_ = v4805
	var v4806 int32
	_ = v4806
	var v4809 int32
	_ = v4809
	var v4815 int32
	_ = v4815
	var v4819 int32
	_ = v4819
	var v4851 int32
	_ = v4851
	var v4853 int32
	_ = v4853
	var v4865 int32
	_ = v4865
	var v4886 int32
	_ = v4886
	var v4889 int32
	_ = v4889
	var v4890 int32
	_ = v4890
	var v4893 int32
	_ = v4893
	var v4897 int32
	_ = v4897
	var v4908 int32
	_ = v4908
	var v4933 int32
	_ = v4933
	var v4934 int32
	_ = v4934
	var v4935 int32
	_ = v4935
	var v4938 int32
	_ = v4938
	var v4967 int32
	_ = v4967
	var v4968 int32
	_ = v4968
	var v4969 int32
	_ = v4969
	var v4973 int32
	_ = v4973
	var v4976 int32
	_ = v4976
	var v4977 int32
	_ = v4977
	var v4978 int32
	_ = v4978
	var v4986 int32
	_ = v4986
	var v4991 int32
	_ = v4991
	var v4996 int32
	_ = v4996
	var v5020 int64
	_ = v5020
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
		v505 = v3
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	if v530 != 0 {
		goto L72
	} else {
		goto L73
	}
L9:
	;
	v93 = v3
	v103 = v3
	goto L10
L10:
	;
	v118 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+60)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+56)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+52)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v118
	v130 = v45 + v103*int32(48)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v131&int32(1) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v497 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[5]))) = uint8(v497)
	v505 = v93
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
		v246 = v180
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
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v246
	v280 = int32(2)
	v282 = v246 + v271
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
		v246 = v180
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v193 = v177
	v194 = v184
	v200 = v180
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
	v246 = v229
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
	v229 = v200
	goto L32
L32:
	;
	v231 = v194 + int32(1)
	if v231 < v229 {
		v193 = v228
		v194 = v231
		v200 = v229
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
	if v266 == int32(3) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v246 != 0 {
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
	v324 = v46 + int32(3728)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+76)))
	v326 = int32(28)
	v327 = v325 * v326
	v329 = v327 - v326
	*(*int32)(unsafe.Add(mBase, uint32(v277)+44)) = v324 + v329
	v333 = v46 + int32(2832)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+40)) = v329 + v333
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v325<<(uint(int32(2))%32)+(v46+int32(4)))+uint32(_c_F_gingetbitmap[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+48)) = v341
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v333+v327-int32(24))))
	if v348 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v349 = int32(57)
	goto L42
L41:
	;
	v349 = int32(58)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+32)) = v349
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v327+v324-int32(24))))
	if v356 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v357 = int32(59)
	goto L45
L44:
	;
	v357 = int32(60)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+36)) = v357
	goto L36
L46:
	;
	v376 = v307
	goto L49
L47:
	;
	goto L48
L48:
	;
	switch v266 - int32(1) {
	case 0:
		v461 = v280
		goto L60
	default:
		goto L59
	case 2:
		goto L61
	}
L49:
	;
	v398 = v376 << (uint(int32(2)) % 32)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v156+v398)))
	v402 = int32(*(*int8)(unsafe.Add(mBase, uint32(v376+v178))))
	v403 = int32(0)
	if v268 == v403 {
		v414 = v403
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
	v406 = int32(0)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v264)+uint32(_c_F_gingetbitmap[8]))))
	if v407&int32(1) == v406 {
		v414 = v406
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376+v268))))
	v414 = v413
	goto L51
L54:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v269+v398)))
	v418 = v416
	goto L56
L55:
	;
	v418 = int32(0)
	goto L56
L56:
	;
	v419 = F_ginFillScanEntry(m, v46, v264, v265, v266, v400, v402, v414, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v421+v398))) = v419
	v425 = v376 + int32(1)
	if v425 != v246 {
		v376 = v425
		goto L49
	} else {
		goto L58
	}
L58:
	;
	goto L50
L59:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	if v483 != int32(2) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v462 + int32(1)
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+76)))
	v467 = int32(0)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v277)+72))
	v473 = F_ginFillScanEntry(m, v46, v466, v467, v468, v467, base.I32_extend8_s(v461), v467, v467)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v461 = int32(255)
	goto L60
L62:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v277)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v475+v462<<(uint(int32(2))%32)))) = v473
	goto L59
L63:
	;
	v486 = int32(*(*int16)(unsafe.Add(mBase, uint32(v130)+4)))
	v488 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v486)+63)) = uint8(v488)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v491 = v103 + int32(1)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v491 < v492 {
		v93 = v239
		v103 = v491
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v505 = v239
	goto L8
L67:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043)+uint32(_c_F_gingetbitmap[5]))))
	if v1044 != 0 {
		v4996 = v34
		v5020 = v28
		goto L144
	} else {
		goto L145
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L139
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v56
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v994)+272))
	if v995 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L70:
	;
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[5]))))
	if v952 != 0 {
		goto L69
	} else {
		goto L127
	}
L71:
	;
	if v505&int32(1) == int32(0) {
		goto L69
	} else {
		goto L126
	}
L72:
	;
	v534 = v3
	v537 = int32(0)
	v545 = v530
	goto L75
L73:
	;
	goto L74
L74:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[5]))))
	if v781 != 0 {
		goto L71
	} else {
		goto L112
	}
L75:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	v566 = v563 + v534*int32(92)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)+72))
	if v567 != int32(2) {
		v601 = v537
		v602 = v545
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if int32(0) < v601 {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	v605 = v534 + int32(1)
	if base.Ui32(v605) < base.Ui32(v602) {
		v534 = v605
		v537 = v601
		v545 = v602
		goto L75
	} else {
		goto L83
	}
L78:
	;
	v570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v566)+76)))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v570)+63)))
	if v572 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v575 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v566)+78)) = uint8(v575)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	*(*int32)(unsafe.Add(mBase, uint32(v566))) = v577 + int32(1)
	v587 = F_ginFillScanEntry(m, v46, v570, v575, int32(2), v575, int32(-1), v575, v575)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v601 = v537 + int32(1)
	v602 = v545
	goto L77
L82:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v589+v577<<(uint(int32(2))%32)))) = v587
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v566)+76)))
	v596 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v594)+63)) = uint8(v596)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	v601 = v537
	v602 = v598
	goto L77
L83:
	;
	goto L76
L84:
	;
	v611 = F_palloc(m, v602*int32(92))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	v749 = v602
	goto L86
L86:
	;
	if v749 != 0 {
		goto L71
	} else {
		goto L111
	}
L87:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	if v613 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v615 = int32(0)
	v619 = v615
	v624 = v615
	v631 = v613 - v601
	goto L91
L89:
	;
	v711 = int32(0)
	goto L90
L90:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	if v711 != 0 {
		goto L107
	} else {
		goto L108
	}
L91:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	v651 = v648 + v624*int32(92)
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+78)))
	if v652 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v711 = v675 * int32(92)
	goto L90
L93:
	;
	v674 = v624 + int32(1)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v674) < base.Ui32(v675) {
		v619 = v671
		v624 = v674
		v631 = v672
		goto L91
	} else {
		goto L105
	}
L94:
	;
	v655 = int32(92)
	goto L98
L95:
	;
	goto L96
L96:
	;
	v663 = int32(92)
	goto L102
L97:
	;
	v671 = v619
	v672 = v631 + int32(1)
	goto L93
L98:
	;
	v659 = F__emscripten_memcpy_bulkmem(m, v611+v631*v655, v651, v655)
	mBase = m.M
	goto L100
L100:
	;
	goto L97
L101:
	;
	v671 = v619 + int32(1)
	v672 = v631
	goto L93
L102:
	;
	v667 = F__emscripten_memcpy_bulkmem(m, v611+v619*v663, v651, v663)
	mBase = m.M
	goto L104
L104:
	;
	goto L101
L105:
	;
	goto L92
L106:
	;
	F_pfree(m, v611)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L110
	}
L107:
	;
	v713 = F__emscripten_memcpy_bulkmem(m, v712, v611, v711)
	mBase = m.M
	goto L109
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3])))
	v749 = v717
	goto L86
L111:
	;
	goto L74
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[3]))) = int32(1)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_gingetbitmap[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v784))) = int64(0)
	v788 = F_palloc(m, int32(4))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v784)+8)) = v788
	v792 = F_palloc0(m, int32(1))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v794 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v784)+52)) = v794
	*(*int32)(unsafe.Add(mBase, uint32(v784)+28)) = v792
	*(*int64)(unsafe.Add(mBase, uint32(v784)+60)) = v794
	v799 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v784)+68)) = uint16(v799)
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+78)) = uint8(v799)
	v803 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v784)+76)) = uint16(v803)
	*(*int32)(unsafe.Add(mBase, uint32(v784)+72)) = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v784)+12)) = v794
	*(*int64)(unsafe.Add(mBase, uint32(v784)+20)) = v794
	*(*int64)(unsafe.Add(mBase, uint32(v784)+80)) = v794
	*(*uint8)(unsafe.Add(mBase, uint32(v784)+88)) = uint8(v799)
	goto L116
L115:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	*(*int32)(unsafe.Add(mBase, uint32(v784))) = v868 + int32(1)
	v872 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v784)+76)))
	v873 = int32(0)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v784)+72))
	v879 = F_ginFillScanEntry(m, v46, v872, v873, v874, v873, int32(-1), v873, v873)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L1
	} else {
		goto L125
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v784)+32)) = int32(55)
	*(*int32)(unsafe.Add(mBase, uint32(v784)+36)) = int32(56)
	goto L115
L125:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v784)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v881+v868<<(uint(int32(2))%32)))) = v879
	goto L70
L126:
	;
	goto L70
L127:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ginGetStats(m, v953, v43+int32(16))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
	if v958 <= int32(0) {
		goto L68
	} else {
		goto L129
	}
L129:
	;
	goto L69
L130:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1011 != 0 {
		goto L136
	} else {
		goto L137
	}
L131:
	;
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994)+268)))
	if v998 != int32(1) {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	v1005 = v995
	goto L133
L133:
	;
	v1006 = *(*int64)(unsafe.Add(mBase, uint32(v1005)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1005)+16)) = v1006 + int64(1)
	goto L130
L134:
	;
	F_pgstat_assoc_relation(m, v994)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+272))
	v1005 = v1004
	goto L133
L136:
	;
	v1012 = *(*int64)(unsafe.Add(mBase, uint32(v1011)))
	*(*int64)(unsafe.Add(mBase, uint32(v1011))) = v1012 + int64(1)
	goto L138
L137:
	;
	goto L138
L138:
	;
	m.G0 = v43 + int32(96)
	goto L67
L139:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_gingetbitmap_3), int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v1031 + int32(4)
	F_errhint(m, int32(_a_F_gingetbitmap_4), v43)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_5), int32(482), int32(_a_F_gingetbitmap_6))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	m.G0 = v4996 + int32(_a_F_gingetbitmap_0)
	return v5020
L145:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1047 = F_ReadBuffer(m, v1045, int32(0))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v1049, int32(0), v1051)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_LockBuffer(m, v1047, int32(1))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	if v1047 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+36))
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[2])))
	if v1987 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L150:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+24))
	if v1075 == int32(-1) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1060+(v1047^int32(-1))<<(uint(int32(2))%32))))
	v1074 = v1066
	goto L150
L152:
	;
	goto L153
L153:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v1074 = v1068 + v1047<<(uint(int32(13))%32) + int32(-8192)
	goto L150
L154:
	;
	F_UnlockReleaseBuffer(m, v1047)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1081 = F_ReadBuffer(m, v1080, v1075)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L158
	}
L157:
	;
	v1955 = l0
	v1956 = l1
	v1959 = v34
	v1983 = v28
	goto L149
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v1081
	F_LockBuffer(m, v1081, int32(1))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v1087 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+40)) = uint16(v1087)
	F_UnlockReleaseBuffer(m, v1047)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+uint32(_c_F_gingetbitmap[3])))
	v1092 = F_palloc(m, v1091)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v1092
	v1097 = F_scanGetCandidate(m, l0, v34+int32(36))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	if v1097 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1103 = l0
	v1104 = l1
	v1107 = v34
	v1110 = v34 + int32(44)
	v1114 = v34 + int32(63)
	v1124 = v1043
	v1131 = v28
	goto L166
L164:
	;
	v1921 = l0
	v1922 = l1
	v1925 = v34
	v1949 = v28
	goto L165
L165:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+52))
	F_pfree(m, v1952)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L1
	} else {
		goto L281
	}
L166:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1103)+36))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+uint32(_c_F_gingetbitmap[3])))
	if v1136 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v1921 = v1103
	v1922 = v1104
	v1925 = v1107
	v1949 = v1914
	goto L165
L168:
	;
	v1141 = int32(0)
	goto L171
L169:
	;
	v1184 = int32(0)
	goto L170
L170:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+52))
	v1216 = F__emscripten_memset_bulkmem(m, v1213, base.I32_extend8_s(int32(0)), v1184)
	mBase = m.M
	goto L175
L171:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+uint32(_c_F_gingetbitmap[4])))
	v1172 = v1169 + v1141*int32(92)
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+28))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1172)))
	v1177 = F__emscripten_memset_bulkmem(m, v1173, base.I32_extend8_s(int32(0)), v1175)
	mBase = m.M
	goto L173
L172:
	;
	v1184 = v1180
	goto L170
L173:
	;
	v1179 = v1141 + int32(1)
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v1179) < base.Ui32(v1180) {
		v1141 = v1179
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v1218 = v1135 + int32(4)
	goto L177
L176:
	;
	if v1629 != 0 {
		goto L256
	} else {
		goto L257
	}
L177:
	;
	v1250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+40)))
	v1253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+42)))
	v1256 = F__emscripten_memset_bulkmem(m, v1114+v1250, base.I32_extend8_s(int32(0)), v1253-v1250)
	mBase = m.M
	goto L179
L178:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L1
	} else {
		goto L252
	}
L179:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+36))
	if v1257 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+uint32(_c_F_gingetbitmap[3])))
	if v1276 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L181:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1261+(v1257^int32(-1))<<(uint(int32(2))%32))))
	v1275 = v1267
	goto L180
L182:
	;
	goto L183
L183:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v1275 = v1269 + v1257<<(uint(int32(13))%32) + int32(-8192)
	goto L180
L184:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1107)+40)) = uint16(v1630)
	v1659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1275)+16)))
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275+v1659)+6)))
	if v1661&int32(32) != 0 {
		goto L176
	} else {
		goto L240
	}
L185:
	;
	v1629 = int32(0)
	v1630 = v1253
	goto L184
L186:
	;
	goto L187
L187:
	;
	v1285 = v1276
	v1302 = int32(0)
	goto L188
L188:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+uint32(_c_F_gingetbitmap[4])))
	v1317 = v1314 + v1302*int32(92)
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1317)))
	if v1318 != 0 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v1626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+42)))
	v1629 = v1594
	v1630 = v1626
	goto L184
L190:
	;
	v1323 = v1318
	v1334 = int32(0)
	goto L193
L191:
	;
	v1594 = v1285
	goto L192
L192:
	;
	v1624 = v1302 + int32(1)
	if base.Ui32(v1624) < base.Ui32(v1594) {
		v1285 = v1594
		v1302 = v1624
		goto L188
	} else {
		goto L239
	}
L193:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+28))
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1351+v1334))))
	if v1353 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+uint32(_c_F_gingetbitmap[3])))
	v1594 = v1591
	goto L192
L195:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+8))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1356+v1334<<(uint(int32(2))%32))))
	v1361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+40)))
	v1362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+42)))
	if base.Ui32(v1362) <= base.Ui32(v1361) {
		v1496 = v1362
		v1497 = v1361
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v1560 = v1323
	goto L197
L197:
	;
	v1589 = v1334 + int32(1)
	if base.Ui32(v1589) < base.Ui32(v1560) {
		v1323 = v1560
		v1334 = v1589
		goto L193
	} else {
		goto L238
	}
L198:
	;
	v1525 = int32(_a_F_gingetbitmap_7)
	v1528 = v1496 & v1525
	if base.Ui32(v1497&v1525) < base.Ui32(v1528) {
		goto L234
	} else {
		goto L235
	}
L199:
	;
	v1366 = v1362
	v1367 = v1361
	goto L200
L200:
	;
	v1398 = int32(1)
	v1400 = int32(base.Ui32((v1366-v1367)&int32(_a_F_gingetbitmap_8))>>(uint(v1398)%32)) + v1367
	v1402 = v1400 & int32(_a_F_gingetbitmap_7)
	v1404 = v1402 - v1398
	v1406 = v1404 << (uint(int32(2)) % 32)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1275+int32(24)+v1406)))
	v1411 = v1275 + v1408&int32(_a_F_gingetbitmap_9)
	v1412 = F_gintuple_get_attrnum(m, v1218, v1411)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L1
	} else {
		goto L202
	}
L201:
	;
	v1496 = v1486
	v1497 = v1487
	goto L198
L202:
	;
	v1414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1317)+76)))
	if base.Ui32(v1414) < base.Ui32(v1412) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v1489 = int32(_a_F_gingetbitmap_7)
	if base.Ui32(v1487&v1489) < base.Ui32(v1486&v1489) {
		v1366 = v1486
		v1367 = v1487
		goto L200
	} else {
		goto L233
	}
L204:
	;
	v1486 = v1400
	v1487 = v1367
	goto L203
L205:
	;
	goto L206
L206:
	;
	if base.Ui32(v1412) < base.Ui32(v1414) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1486 = v1366
	v1487 = v1400 + int32(1)
	goto L203
L208:
	;
	goto L209
L209:
	;
	v1421 = v1107 - int32(-64) + v1404
	v1422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1421))))
	if v1422 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1431 = F_gintuple_get_key(m, v1218, v1411, v1107+int32(1088)+v1404)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v1436 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1360)+4)))
	if v1436 == int32(-1) {
		goto L216
	} else {
		goto L217
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1107+int32(2112)+v1406))) = v1431
	v1434 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1421))) = uint8(v1434)
	goto L212
L214:
	;
	v1483 = base.B2i32(v1479 < int32(0))
	if v1479 < int32(0) {
		goto L227
	} else {
		goto L228
	}
L215:
	;
	v1462 = int32(1)
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+5)))
	if v1463 == v1462 {
		goto L223
	} else {
		goto L224
	}
L216:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+16))
	if v1439 != int32(2) {
		goto L215
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v1449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1360)+20)))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1360)))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1107+int32(2112)+v1406)))
	v1458 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1107+int32(1088)+v1404))))
	v1459 = F_ginCompareEntries(m, v1218, v1449, v1450, v1436, v1454, v1458)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107+int32(1088)+v1404))))
	if v1446 != int32(3) {
		goto L215
	} else {
		goto L220
	}
L220:
	;
	v1479 = int32(-1)
	goto L214
L221:
	;
	if v1459 != 0 {
		v1479 = v1459
		goto L214
	} else {
		goto L222
	}
L222:
	;
	goto L215
L223:
	;
	v1466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+42)))
	v1473 = F_matchPartialInPendingList(m, v1218, v1275, v1402, v1466, v1360, v1107+int32(2112), v1107+int32(1088), v1107-int32(-64))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	v1475 = v1462
	goto L225
L225:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1476+v1334))) = uint8(v1475)
	v1496 = v1366
	v1497 = v1367
	goto L198
L226:
	;
	v1475 = v1473
	goto L225
L227:
	;
	v1484 = v1367
	goto L229
L228:
	;
	v1484 = v1400 + int32(1)
	goto L229
L229:
	;
	if v1479 < int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1485 = v1400
	goto L232
L231:
	;
	v1485 = v1366
	goto L232
L232:
	;
	v1486 = v1485
	v1487 = v1484
	goto L203
L233:
	;
	goto L201
L234:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+52))
	v1547 = v1546 + v1302
	v1548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1547))))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+28))
	v1551 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1549+v1334))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1547))) = uint8(base.B2i32(v1548|v1551 != int32(0)))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1317)))
	v1560 = v1556
	goto L197
L235:
	;
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360)+5)))
	if v1530 != int32(1) {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1107)+42)))
	v1540 = F_matchPartialInPendingList(m, v1218, v1275, v1528, v1533, v1360, v1107+int32(2112), v1107+int32(1088), v1107-int32(-64))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1542+v1334))) = uint8(v1540)
	goto L234
L238:
	;
	goto L194
L239:
	;
	goto L189
L240:
	;
	v1664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1107)+60)) = uint16(v1664)
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1110)))
	*(*int32)(unsafe.Add(mBase, uint32(v1107)+56)) = v1666
	v1670 = F_scanGetCandidate(m, v1103, v1107+int32(36))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	if v1670 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1673 = v1107 + int32(56)
	v1674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+2)))
	v1675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110))))
	v1676 = int32(16)
	v1679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1673)+2)))
	v1680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1673))))
	if v1674|v1675<<(uint(v1676)%32) == v1679|v1680<<(uint(v1676)%32) {
		goto L247
	} else {
		goto L248
	}
L243:
	;
	goto L244
L244:
	;
	goto L178
L245:
	;
	if v1690 != 0 {
		goto L177
	} else {
		goto L251
	}
L246:
	;
	goto L245
L247:
	;
	v1686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+4)))
	v1687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1673)+4)))
	if v1686 == v1687 {
		v1690 = int32(1)
		goto L246
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1690 = int32(0)
	goto L246
L250:
	;
	goto L249
L251:
	;
	goto L244
L252:
	;
	F_errmsg_internal(m, int32(_a_F_gingetbitmap_10), int32(0))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_11), int32(1814), int32(_a_F_gingetbitmap_12))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	v1919 = F_scanGetCandidate(m, v1103, v1107+int32(36))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L1
	} else {
		goto L279
	}
L256:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+52))
	v1709 = int32(0)
	goto L259
L257:
	;
	goto L258
L258:
	;
	v1782 = int32(0)
	v1783 = int32(_a_F_gingetbitmap_1)
	v1784 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1786
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+uint32(_c_F_gingetbitmap[3])))
	if v1789 != 0 {
		goto L266
	} else {
		goto L267
	}
L259:
	;
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1709+v1705))))
	if v1738 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	goto L258
L261:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+uint32(_c_F_gingetbitmap[4])))
	v1745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1741+v1709*int32(92))+78)))
	if v1745 != int32(1) {
		v1914 = v1131
		goto L255
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1749 = v1709 + int32(1)
	if v1749 != v1629 {
		v1709 = v1749
		goto L259
	} else {
		goto L265
	}
L264:
	;
	goto L263
L265:
	;
	goto L260
L266:
	;
	v1793 = v1782
	v1795 = v1782
	goto L269
L267:
	;
	v1850 = v1782
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1784
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	F_MemoryContextReset(m, v1878)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L277
	}
L269:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+uint32(_c_F_gingetbitmap[4])))
	v1824 = v1821 + v1793*int32(92)
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+32))
	v1826 = m.T0[v1825].(func(*base.Module, int32) int32)(m, v1824)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L271
	}
L270:
	;
	v1850 = v1840
	goto L268
L271:
	;
	if v1826 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v1784
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	F_MemoryContextReset(m, v1832)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1824)+87)))
	v1836 = int32(1)
	v1840 = base.B2i32(v1835|v1795&v1836 != int32(0))
	v1842 = v1793 + v1836
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v1842) < base.Ui32(v1843) {
		v1793 = v1842
		v1795 = v1840
		goto L269
	} else {
		goto L276
	}
L275:
	;
	v1914 = v1131
	goto L255
L276:
	;
	goto L270
L277:
	;
	F_tbm_add_tuples(m, v1104, v1110, int32(1), v1850)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v1914 = v1131 + int64(1)
	goto L255
L279:
	;
	if v1919 != 0 {
		v1131 = v1914
		goto L166
	} else {
		goto L280
	}
L280:
	;
	goto L167
L281:
	;
	v1955 = v1921
	v1956 = v1922
	v1959 = v1925
	v1983 = v1949
	goto L149
L282:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4973 = m.ExcPending
	if v4973 != 0 {
		goto L1
	} else {
		goto L686
	}
L283:
	;
	v3165 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[3])))
	if v3165 != 0 {
		goto L491
	} else {
		goto L492
	}
L284:
	;
	v1991 = v1986 + int32(4)
	v2007 = int32(0)
	goto L285
L285:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[6])))
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2024+v2007<<(uint(int32(2))%32))))
	v2030 = v2028 + int32(648)
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1955)+8))
	goto L291
L286:
	;
	if v3034 == int32(0) {
		goto L283
	} else {
		goto L482
	}
L287:
	;
	F_freeGinBtreeStack(m, v2131)
	mBase = m.M
	v3031 = m.ExcPending
	if v3031 != 0 {
		goto L1
	} else {
		goto L480
	}
L288:
	;
	v2995 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	F_LockBuffer(m, v2995, int32(0))
	mBase = m.M
	v2998 = m.ExcPending
	if v2998 != 0 {
		goto L1
	} else {
		goto L479
	}
L289:
	;
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1100))
	v2827 = m.T0[v2826].(func(*base.Module, int32, int32) int32)(m, v1959+int32(1088), v2131)
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L1
	} else {
		goto L449
	}
L290:
	;
	if v2784 == int32(0) {
		goto L288
	} else {
		goto L445
	}
L291:
	;
	v2067 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2028)+652)) = uint16(v2067)
	*(*uint16)(unsafe.Add(mBase, uint32(v2028+int32(32)))) = uint16(v2067)
	*(*int64)(unsafe.Add(mBase, uint32(v2028+int32(24)))) = int64(0)
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+644))
	if v2073 != 0 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v2781 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+36))
	v2784 = v2781
	goto L290
L293:
	;
	F_pfree(m, v2073)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L1
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v2076 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+648)) = v2076
	*(*int64)(unsafe.Add(mBase, uint32(v2028)+640)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+36)) = v2076
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+656)) = v2076
	*(*uint8)(unsafe.Add(mBase, uint32(v2028)+655)) = uint8(v2076)
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+44)) = int32(-1)
	v2090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2028)+20)))
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2028)))
	v2092 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2028)+4)))
	v2096 = F___memset(m, v1959+int32(1088), v2076, int32(68))
	mBase = m.M
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v1991)))
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+48)) = v1991
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+40)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+32)) = int32(43)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+24)) = int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+20)) = int32(45)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+16)) = int32(46)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+12)) = int32(47)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+8)) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+4)) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v2096))) = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(v2096)+60)) = uint8(v2092)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+56)) = v2091
	*(*uint16)(unsafe.Add(mBase, uint32(v2096)+54)) = uint16(v2090)
	*(*uint16)(unsafe.Add(mBase, uint32(v2096)+52)) = uint16(v2076)
	*(*uint8)(unsafe.Add(mBase, uint32(v2096)+36)) = uint8(v2076)
	*(*int32)(unsafe.Add(mBase, uint32(v2096)+28)) = int32(51)
	goto L297
L296:
	;
	goto L295
L297:
	;
	v2131 = F_ginFindLeafPage(m, v1959+int32(1088), int32(1), int32(0))
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L1
	} else {
		goto L299
	}
L298:
	;
	v2152 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2028)+654)) = uint8(v2152)
	v2154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2028)+5)))
	if v2154 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L299:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	if v2133 < int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v2137+(v2133^int32(-1))<<(uint(int32(2))%32))))
	v2151 = v2143
	goto L298
L301:
	;
	goto L302
L302:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2151 = v2145 + v2133<<(uint(int32(13))%32) + int32(-8192)
	goto L298
L303:
	;
	v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2028)+4)))
	if v2157 != int32(255) {
		goto L289
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1100))
	v2163 = m.T0[v2162].(func(*base.Module, int32, int32) int32)(m, v1959+int32(1088), v2131)
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L1
	} else {
		goto L307
	}
L306:
	;
	goto L305
L307:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[11]))
	v2170 = F_tbm_create(m, v2166<<(uint(int32(10))%32), int32(0))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+36)) = v2170
	v2173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2028)+5)))
	if v2173 == int32(1) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2028)+4)))
	if v2176 != 0 {
		v2784 = v2170
		goto L290
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1136))
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v2177)+8))
	v2179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2028)+20)))
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	if v2181 < int32(0) {
		goto L314
	} else {
		goto L315
	}
L312:
	;
	goto L311
L313:
	;
	F_PredicateLockPage(m, v2180, v2200, v2031)
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L1
	} else {
		goto L317
	}
L314:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v2185+(v2181^int32(-1))<<(uint(int32(6))%32))+16))
	v2200 = v2191
	goto L313
L315:
	;
	goto L316
L316:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2193+v2181<<(uint(int32(6))%32)+int32(-64))+16))
	v2200 = v2199
	goto L313
L317:
	;
	v2204 = v2179 - int32(1)
	v2209 = v2178 + v2204<<(uint(int32(4))%32) + int32(20)
	goto L319
L318:
	;
	goto L292
L319:
	;
	v2243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	if v2244 < int32(0) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+36))
	if v2757 != 0 {
		goto L435
	} else {
		goto L436
	}
L321:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1136))
	v2331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)))
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(v2331<<(uint(int32(2))%32)+v2329)+20))
	v2338 = v2329 + v2335&int32(_a_F_gingetbitmap_9)
	v2339 = F_gintuple_get_attrnum(m, v2330, v2338)
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L1
	} else {
		goto L342
	}
L322:
	;
	v2263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2262)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2263) {
		goto L326
	} else {
		goto L327
	}
L323:
	;
	v2248 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v2248+(v2244^int32(-1))<<(uint(int32(2))%32))))
	v2262 = v2254
	goto L322
L324:
	;
	goto L325
L325:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2262 = v2256 + v2244<<(uint(int32(13))%32) + int32(-8192)
	goto L322
L326:
	;
	v2271 = int32(base.Ui32(v2263+int32(_a_F_gingetbitmap_13)) >> (uint(int32(2)) % 32))
	goto L328
L327:
	;
	v2271 = int32(0)
	goto L328
L328:
	;
	if base.Ui32(v2271&int32(_a_F_gingetbitmap_7)) < base.Ui32(v2243) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v2275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2262)+16)))
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v2262+v2275)))
	if v2277 == int32(-1) {
		goto L318
	} else {
		goto L332
	}
L330:
	;
	v2311 = v2244
	goto L331
L331:
	;
	if v2311 < int32(0) {
		goto L339
	} else {
		goto L340
	}
L332:
	;
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	v2282 = F_ginStepRight(m, v2244, v2280, int32(1))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2131)+4)) = v2282
	if v2282 < int32(0) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v2304 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)) = uint16(v2304)
	*(*int32)(unsafe.Add(mBase, uint32(v2131))) = v2303
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	F_PredicateLockPage(m, v2307, v2303, v2031)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L1
	} else {
		goto L338
	}
L335:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v2288+(v2282^int32(-1))<<(uint(int32(6))%32))+16))
	v2303 = v2294
	goto L334
L336:
	;
	goto L337
L337:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v2282<<(uint(int32(6))%32)+int32(-64))+16))
	v2303 = v2302
	goto L334
L338:
	;
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	v2311 = v2310
	goto L331
L339:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v2315+(v2311^int32(-1))<<(uint(int32(2))%32))))
	v2329 = v2321
	goto L321
L340:
	;
	goto L341
L341:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2329 = v2323 + v2311<<(uint(int32(13))%32) + int32(-8192)
	goto L321
L342:
	;
	if v2339 != v2179 {
		goto L318
	} else {
		goto L343
	}
L343:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1136))
	v2345 = F_gintuple_get_key(m, v2342, v2338, v1959-int32(-64))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	v2347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2028)+5)))
	if v2347 == int32(1) {
		goto L348
	} else {
		goto L349
	}
L345:
	;
	goto L320
L346:
	;
	v2753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)))
	v2755 = v2753 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)) = uint16(v2755)
	goto L319
L347:
	;
	v2377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2338)+4)))
	if v2377 == int32(_a_F_gingetbitmap_7) {
		goto L357
	} else {
		goto L358
	}
L348:
	;
	v2350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959)+64)))
	if v2350 != 0 {
		goto L318
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+16))
	if v2370 != int32(2) {
		goto L347
	} else {
		goto L355
	}
L351:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1136))
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2351+v2204<<(uint(int32(2))%32))+uint32(_c_F_gingetbitmap[14])))
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2028)))
	v2362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2028)+12)))
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+8))
	v2364 = F_FunctionCall4Coll(m, v2351+v2204*int32(28)+int32(_a_F_gingetbitmap_14), v2360, v2361, v2345, v2362, v2363)
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	if int32(0) < v2364 {
		goto L318
	} else {
		goto L353
	}
L353:
	;
	if int32(0) <= v2364 {
		goto L347
	} else {
		goto L354
	}
L354:
	;
	goto L346
L355:
	;
	v2373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959)+64)))
	if v2373 == int32(3) {
		goto L318
	} else {
		goto L356
	}
L356:
	;
	goto L347
L357:
	;
	v2380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2338)+2)))
	v2381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2338))))
	v2384 = v2380 | v2381<<(uint(int32(16))%32)
	v2385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959)+64)))
	if v2385 == int32(0) {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	goto L359
L359:
	;
	v2709 = F_ginReadTuple(m, v2338, v1959+int32(2112))
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		goto L1
	} else {
		goto L432
	}
L360:
	;
	v2388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2209)+6)))
	v2389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2209)+4)))
	v2390 = F_datumCopy(m, v2345, v2388, v2389)
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L1
	} else {
		goto L363
	}
L361:
	;
	v2392 = v2345
	goto L362
L362:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	F_LockBuffer(m, v2393, int32(0))
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L1
	} else {
		goto L364
	}
L363:
	;
	v2392 = v2390
	goto L362
L364:
	;
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	F_PredicateLockPage(m, v2397, v2384, v2031)
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	v2403 = F_ginScanBeginPostingTree(m, v1959+int32(2112), v2402, v2384)
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v2403)+4))
	F_IncrBufferRefCount(m, v2405)
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	F_freeGinBtreeStack(m, v2403)
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	v2415 = v2405
	goto L369
L369:
	;
	if v2415 < int32(0) {
		goto L372
	} else {
		goto L373
	}
L370:
	;
	F_UnlockReleaseBuffer(m, v2415)
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L1
	} else {
		goto L391
	}
L371:
	;
	v2459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2458)+16)))
	v2461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2459+v2458)+6)))
	if v2461&int32(4) == int32(0) {
		goto L375
	} else {
		goto L376
	}
L372:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v2444+(v2415^int32(-1))<<(uint(int32(2))%32))))
	v2458 = v2450
	goto L371
L373:
	;
	goto L374
L374:
	;
	v2452 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2458 = v2452 + v2415<<(uint(int32(13))%32) + int32(-8192)
	goto L371
L375:
	;
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+36))
	v2467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2458)+16)))
	v2468 = v2458 + v2467
	v2469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2468)+6)))
	if v2469&int32(128) != 0 {
		goto L379
	} else {
		goto L380
	}
L376:
	;
	v2512 = v2459
	goto L377
L377:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v2458+v2512)))
	if v2516 != int32(-1) {
		goto L387
	} else {
		goto L388
	}
L378:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+656))
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+656)) = v2507 + v2508
	v2511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2458)+16)))
	v2512 = v2511
	goto L377
L379:
	;
	v2472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2458)+12)))
	v2473 = int32(32)
	v2475 = m.G0
	v2477 = v2475 - int32(16)
	m.G0 = v2477
	v2483 = F_ginPostingListDecodeAllSegments(m, v2458+v2473, v2472-v2473, v2477+int32(12))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L1
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	v2495 = int32(0)
	v2496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2468)+4)))
	if v2496 == v2495 {
		v2507 = v2495
		goto L378
	} else {
		goto L385
	}
L382:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v2477)+12))
	F_tbm_add_tuples(m, v2466, v2483, v2485, int32(0))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	F_pfree(m, v2483)
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(v2477)+12))
	m.G0 = v2477 + int32(16)
	v2507 = v2491
	goto L378
L385:
	;
	F_tbm_add_tuples(m, v2466, v2458+int32(32), v2496, int32(0))
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v2507 = v2496
	goto L378
L387:
	;
	v2520 = F_ginStepRight(m, v2415, v2402, int32(1))
	mBase = m.M
	v2521 = m.ExcPending
	if v2521 != 0 {
		goto L1
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	goto L370
L390:
	;
	v2415 = v2520
	goto L369
L391:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	F_LockBuffer(m, v2524, int32(1))
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	if v2528 < int32(0) {
		goto L394
	} else {
		goto L395
	}
L393:
	;
	v2547 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2546)+16)))
	v2549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547+v2546)+6)))
	if v2549&int32(2) == int32(0) {
		goto L345
	} else {
		goto L397
	}
L394:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2532+(v2528^int32(-1))<<(uint(int32(2))%32))))
	v2546 = v2538
	goto L393
L395:
	;
	goto L396
L396:
	;
	v2540 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2546 = v2540 + v2528<<(uint(int32(13))%32) + int32(-8192)
	goto L393
L397:
	;
	v2556 = v2528
	goto L398
L398:
	;
	v2585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)))
	if v2556 < int32(0) {
		goto L402
	} else {
		goto L403
	}
L399:
	;
	v2701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959)+64)))
	if v2701 != 0 {
		goto L346
	} else {
		goto L429
	}
L400:
	;
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1136))
	v2672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)))
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(v2672<<(uint(int32(2))%32)+v2670)+20))
	v2679 = v2670 + v2676&int32(_a_F_gingetbitmap_9)
	v2680 = F_gintuple_get_attrnum(m, v2671, v2679)
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L1
	} else {
		goto L422
	}
L401:
	;
	v2604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2603)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v2604) {
		goto L405
	} else {
		goto L406
	}
L402:
	;
	v2589 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(v2589+(v2556^int32(-1))<<(uint(int32(2))%32))))
	v2603 = v2595
	goto L401
L403:
	;
	goto L404
L404:
	;
	v2597 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2603 = v2597 + v2556<<(uint(int32(13))%32) + int32(-8192)
	goto L401
L405:
	;
	v2612 = int32(base.Ui32(v2604+int32(_a_F_gingetbitmap_13)) >> (uint(int32(2)) % 32))
	goto L407
L406:
	;
	v2612 = int32(0)
	goto L407
L407:
	;
	if base.Ui32(v2612&int32(_a_F_gingetbitmap_7)) < base.Ui32(v2585) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v2616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2603)+16)))
	v2618 = *(*int32)(unsafe.Add(mBase, uint32(v2603+v2616)))
	if v2618 == int32(-1) {
		goto L282
	} else {
		goto L411
	}
L409:
	;
	v2652 = v2556
	goto L410
L410:
	;
	if v2652 < int32(0) {
		goto L418
	} else {
		goto L419
	}
L411:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	v2623 = F_ginStepRight(m, v2556, v2621, int32(1))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2131)+4)) = v2623
	if v2623 < int32(0) {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	v2645 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)) = uint16(v2645)
	*(*int32)(unsafe.Add(mBase, uint32(v2131))) = v2644
	v2648 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	F_PredicateLockPage(m, v2648, v2644, v2031)
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L1
	} else {
		goto L417
	}
L414:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2635 = *(*int32)(unsafe.Add(mBase, uint32(v2629+(v2623^int32(-1))<<(uint(int32(6))%32))+16))
	v2644 = v2635
	goto L413
L415:
	;
	goto L416
L416:
	;
	v2637 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2637+v2623<<(uint(int32(6))%32)+int32(-64))+16))
	v2644 = v2643
	goto L413
L417:
	;
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	v2652 = v2651
	goto L410
L418:
	;
	v2656 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2656+(v2652^int32(-1))<<(uint(int32(2))%32))))
	v2670 = v2662
	goto L400
L419:
	;
	goto L420
L420:
	;
	v2664 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2670 = v2664 + v2652<<(uint(int32(13))%32) + int32(-8192)
	goto L400
L421:
	;
	goto L399
L422:
	;
	if v2680 == v2179 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1136))
	v2686 = F_gintuple_get_key(m, v2683, v2679, v1959+int32(2112))
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L1
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	v2696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)))
	v2698 = v2696 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)) = uint16(v2698)
	v2700 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	v2556 = v2700
	goto L398
L426:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1136))
	v2689 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1959)+2112)))
	v2690 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1959)+64)))
	v2691 = F_ginCompareEntries(m, v2688, v2179, v2686, v2689, v2392, v2690)
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	if v2691 == int32(0) {
		goto L421
	} else {
		goto L428
	}
L428:
	;
	goto L425
L429:
	;
	v2702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2209)+6)))
	if v2702 != 0 {
		goto L346
	} else {
		goto L430
	}
L430:
	;
	F_pfree(m, v2392)
	mBase = m.M
	v2704 = m.ExcPending
	if v2704 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	goto L346
L432:
	;
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+36))
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+2112))
	F_tbm_add_tuples(m, v2711, v2709, v2712, int32(0))
	mBase = m.M
	v2715 = m.ExcPending
	if v2715 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+656))
	v2717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2338)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+656)) = v2716 + v2717
	F_pfree(m, v2709)
	mBase = m.M
	v2721 = m.ExcPending
	if v2721 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	goto L346
L435:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+40))
	if v2758 != 0 {
		goto L438
	} else {
		goto L439
	}
L436:
	;
	v2772 = v2528
	goto L437
L437:
	;
	F_LockBuffer(m, v2772, int32(0))
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L1
	} else {
		goto L443
	}
L438:
	;
	F_pfree(m, v2758)
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		goto L1
	} else {
		goto L441
	}
L439:
	;
	v2762 = v2757
	goto L440
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+40)) = int32(0)
	F_tbm_free(m, v2762)
	mBase = m.M
	v2766 = m.ExcPending
	if v2766 != 0 {
		goto L1
	} else {
		goto L442
	}
L441:
	;
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+36))
	v2762 = v2761
	goto L440
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+36)) = int32(0)
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	v2772 = v2769
	goto L437
L443:
	;
	F_freeGinBtreeStack(m, v2131)
	mBase = m.M
	v2777 = m.ExcPending
	if v2777 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	goto L291
L445:
	;
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(v2784)+16))
	goto L446
L446:
	;
	if v2815 == int32(0) {
		goto L288
	} else {
		goto L447
	}
L447:
	;
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+36))
	v2819 = F_tbm_begin_private_iterate(m, v2818)
	mBase = m.M
	v2820 = m.ExcPending
	if v2820 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	v2821 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2028)+654)) = uint8(v2821)
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+40)) = v2819
	goto L288
L449:
	;
	if v2827 != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+8)))
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v2829<<(uint(int32(2))%32)+v2151)+20))
	v2836 = v2151 + v2833&int32(_a_F_gingetbitmap_9)
	v2837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2836)+4)))
	if v2837 == int32(_a_F_gingetbitmap_7) {
		goto L453
	} else {
		goto L454
	}
L451:
	;
	goto L452
L452:
	;
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(v1991)))
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	if v2942 < int32(0) {
		goto L475
	} else {
		goto L476
	}
L453:
	;
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v1991)))
	v2841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2836)+2)))
	v2842 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2836))))
	v2845 = v2841 | v2842<<(uint(int32(16))%32)
	F_PredicateLockPage(m, v2840, v2845, v2031)
	mBase = m.M
	v2847 = m.ExcPending
	if v2847 != 0 {
		goto L1
	} else {
		goto L456
	}
L454:
	;
	goto L455
L455:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v1991)))
	v2908 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	if v2908 < int32(0) {
		goto L468
	} else {
		goto L469
	}
L456:
	;
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+4))
	F_LockBuffer(m, v2848, int32(0))
	mBase = m.M
	v2851 = m.ExcPending
	if v2851 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v1991)))
	v2855 = F_ginScanBeginPostingTree(m, v2028+int32(660), v2854, v2845)
	mBase = m.M
	v2856 = m.ExcPending
	if v2856 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v2857 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+24)) = v2857
	F_IncrBufferRefCount(m, v2857)
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L1
	} else {
		goto L459
	}
L459:
	;
	v2861 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+24))
	if v2861 < int32(0) {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	v2882 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1959+int32(2116)))) = uint16(v2882)
	*(*uint16)(unsafe.Add(mBase, uint32(v1959)+32)) = uint16(v2882)
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+2112)) = v2882
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+28)) = v2882
	v2892 = F_GinDataLeafPageGetItems(m, v2879, v2030, v1959+int32(28))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L464
	}
L461:
	;
	v2865 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[9]))
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v2865+(v2861^int32(-1))<<(uint(int32(2))%32))))
	v2879 = v2871
	goto L460
L462:
	;
	goto L463
L463:
	;
	v2873 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[10]))
	v2879 = v2873 + v2861<<(uint(int32(13))%32) + int32(-8192)
	goto L460
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+644)) = v2892
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+648))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+656)) = v2895 * v2896
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+24))
	F_LockBuffer(m, v2899, int32(0))
	mBase = m.M
	v2902 = m.ExcPending
	if v2902 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_freeGinBtreeStack(m, v2855)
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	v2905 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2028)+654)) = uint8(v2905)
	goto L287
L467:
	;
	F_PredicateLockPage(m, v2907, v2927, v2031)
	mBase = m.M
	v2929 = m.ExcPending
	if v2929 != 0 {
		goto L1
	} else {
		goto L471
	}
L468:
	;
	v2912 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v2912+(v2908^int32(-1))<<(uint(int32(6))%32))+16))
	v2927 = v2918
	goto L467
L469:
	;
	goto L470
L470:
	;
	v2920 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2920+v2908<<(uint(int32(6))%32)+int32(-64))+16))
	v2927 = v2926
	goto L467
L471:
	;
	v2930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2836)+4)))
	if v2930 == int32(0) {
		goto L288
	} else {
		goto L472
	}
L472:
	;
	v2934 = F_ginReadTuple(m, v2836, v2030)
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+644)) = v2934
	v2937 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2028)+654)) = uint8(v2937)
	v2939 = *(*int32)(unsafe.Add(mBase, uint32(v2028)+648))
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+656)) = v2939
	goto L288
L474:
	;
	F_PredicateLockPage(m, v2941, v2961, v2031)
	mBase = m.M
	v2963 = m.ExcPending
	if v2963 != 0 {
		goto L1
	} else {
		goto L478
	}
L475:
	;
	v2946 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[12]))
	v2952 = *(*int32)(unsafe.Add(mBase, uint32(v2946+(v2942^int32(-1))<<(uint(int32(6))%32))+16))
	v2961 = v2952
	goto L474
L476:
	;
	goto L477
L477:
	;
	v2954 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[13]))
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(v2954+v2942<<(uint(int32(6))%32)+int32(-64))+16))
	v2961 = v2960
	goto L474
L478:
	;
	goto L288
L479:
	;
	goto L287
L480:
	;
	v3033 = v2007 + int32(1)
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[2])))
	if base.Ui32(v3033) < base.Ui32(v3034) {
		v2007 = v3033
		goto L285
	} else {
		goto L481
	}
L481:
	;
	goto L286
L482:
	;
	v3039 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[15]))
	if v3039 <= int32(0) {
		goto L283
	} else {
		goto L483
	}
L483:
	;
	v3043 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[6])))
	v3047 = int32(0)
	goto L484
L484:
	;
	v3079 = *(*int32)(unsafe.Add(mBase, uint32(v3043+v3047<<(uint(int32(2))%32))))
	v3080 = *(*int32)(unsafe.Add(mBase, uint32(v3079)+656))
	if base.Ui32(v3080) <= base.Ui32(v3039*v3034) {
		goto L283
	} else {
		goto L486
	}
L485:
	;
	v3088 = int32(0)
	v3089 = v3034
	goto L488
L486:
	;
	v3083 = v3047 + int32(1)
	if v3083 != v3034 {
		v3047 = v3083
		goto L484
	} else {
		goto L487
	}
L487:
	;
	goto L485
L488:
	;
	v3118 = v3088 << (uint(int32(2)) % 32)
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[6])))
	v3121 = *(*int32)(unsafe.Add(mBase, uint32(v3118+v3119)))
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v3121)+656))
	v3123 = base.I32_div_u_s(v3122, v3089)
	*(*int32)(unsafe.Add(mBase, uint32(v3121)+656)) = v3123
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[6])))
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(v3125+v3118)))
	v3128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3127)+655)) = uint8(v3128)
	v3131 = v3088 + v3128
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[2])))
	if base.Ui32(v3131) < base.Ui32(v3132) {
		v3088 = v3131
		v3089 = v3132
		goto L488
	} else {
		goto L490
	}
L489:
	;
	goto L283
L490:
	;
	goto L489
L491:
	;
	v3167 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	v3181 = int32(0)
	goto L494
L492:
	;
	goto L493
L493:
	;
	v3799 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1959)+1092)) = uint16(v3799)
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+1088)) = v3799
	v3806 = v1955
	v3807 = v1956
	v3810 = v1959
	v3820 = v3799
	v3822 = v3799
	v3825 = v3799
	v3834 = v1983
	goto L552
L494:
	;
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[4])))
	v3203 = v3200 + v3181*int32(92)
	*(*int64)(unsafe.Add(mBase, uint32(v3203)+80)) = int64(0)
	v3206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3203)+88)) = uint8(v3206)
	v3208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3203)+78)))
	if v3208 == int32(1) {
		goto L497
	} else {
		goto L498
	}
L495:
	;
	goto L493
L496:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3167
	v3765 = v3181 + int32(1)
	v3766 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[3])))
	if base.Ui32(v3765) < base.Ui32(v3766) {
		v3181 = v3765
		goto L494
	} else {
		goto L551
	}
L497:
	;
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3212
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+16)) = int32(0)
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+24)) = v3216
	v3220 = F_palloc(m, v3216<<(uint(int32(2))%32))
	mBase = m.M
	v3221 = m.ExcPending
	if v3221 != 0 {
		goto L1
	} else {
		goto L500
	}
L498:
	;
	goto L499
L499:
	;
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	if base.Ui32(int32(2)) <= base.Ui32(v3270) {
		goto L505
	} else {
		goto L506
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+20)) = v3220
	v3223 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+24))
	if v3223 <= int32(0) {
		goto L496
	} else {
		goto L501
	}
L501:
	;
	v3229 = int32(0)
	goto L502
L502:
	;
	v3259 = v3229 << (uint(int32(2)) % 32)
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+20))
	v3262 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+8))
	v3264 = *(*int32)(unsafe.Add(mBase, uint32(v3262+v3259)))
	*(*int32)(unsafe.Add(mBase, uint32(v3259+v3260))) = v3264
	v3267 = v3229 + int32(1)
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+24))
	if v3267 < v3268 {
		v3229 = v3267
		goto L502
	} else {
		goto L504
	}
L503:
	;
	goto L496
L504:
	;
	goto L503
L505:
	;
	v3275 = *(*int32)(unsafe.Add(mBase, uint32(v1986)))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3275
	v3277 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	v3280 = F_palloc(m, v3277<<(uint(int32(2))%32))
	mBase = m.M
	v3281 = m.ExcPending
	if v3281 != 0 {
		goto L1
	} else {
		goto L508
	}
L506:
	;
	goto L507
L507:
	;
	v3718 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3718
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+16)) = int32(1)
	v3725 = F_palloc(m, int32(4))
	mBase = m.M
	v3726 = m.ExcPending
	if v3726 != 0 {
		goto L1
	} else {
		goto L550
	}
L508:
	;
	v3283 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	if v3283 != 0 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v3286 = int32(0)
	goto L512
L510:
	;
	v3331 = int32(0)
	goto L511
L511:
	;
	F_qsort_arg(m, v3280, v3331, int32(4), int32(52), v3203)
	mBase = m.M
	v3357 = m.ExcPending
	if v3357 != 0 {
		goto L1
	} else {
		goto L515
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3280+v3286<<(uint(int32(2))%32)))) = v3286
	v3320 = v3286 + int32(1)
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	if base.Ui32(v3320) < base.Ui32(v3321) {
		v3286 = v3320
		goto L512
	} else {
		goto L514
	}
L513:
	;
	v3331 = v3321
	goto L511
L514:
	;
	goto L513
L515:
	;
	v3358 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	if base.Ui32(int32(2)) <= base.Ui32(v3358) {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v3364 = int32(1)
	goto L519
L517:
	;
	v3413 = v3358
	goto L518
L518:
	;
	v3436 = int32(1)
	if v3413 != v3436 {
		goto L522
	} else {
		goto L523
	}
L519:
	;
	v3393 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+28))
	v3394 = int32(2)
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(v3280+v3364<<(uint(v3394)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3393+v3397))) = uint8(v3394)
	v3402 = v3364 + int32(1)
	v3403 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	if base.Ui32(v3402) < base.Ui32(v3403) {
		v3364 = v3402
		goto L519
	} else {
		goto L521
	}
L520:
	;
	v3413 = v3403
	goto L518
L521:
	;
	goto L520
L522:
	;
	v3442 = int32(0)
	goto L525
L523:
	;
	v3503 = v3436
	goto L524
L524:
	;
	v3528 = int32(0)
	v3530 = *(*int32)(unsafe.Add(mBase, uint32(v1986)+uint32(_c_F_gingetbitmap[1])))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3530
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+16)) = v3503
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+24)) = v3533 - v3503
	v3538 = F_palloc(m, v3503<<(uint(int32(2))%32))
	mBase = m.M
	v3539 = m.ExcPending
	if v3539 != 0 {
		goto L1
	} else {
		goto L535
	}
L525:
	;
	v3471 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+28))
	v3475 = *(*int32)(unsafe.Add(mBase, uint32(v3280+v3442<<(uint(int32(2))%32))))
	v3477 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3471+v3475))) = uint8(v3477)
	v3479 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+36))
	v3480 = m.T0[v3479].(func(*base.Module, int32) int32)(m, v3203)
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L1
	} else {
		goto L528
	}
L526:
	;
	v3503 = v3494 + int32(1)
	goto L524
L527:
	;
	goto L526
L528:
	;
	if v3480 == int32(0) {
		v3494 = v3442
		goto L527
	} else {
		goto L529
	}
L529:
	;
	v3485 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[16]))
	if v3485 != 0 {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3487 = m.ExcPending
	if v3487 != 0 {
		goto L1
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	v3488 = int32(1)
	v3489 = v3442 + v3488
	v3490 = *(*int32)(unsafe.Add(mBase, uint32(v3203)))
	if base.Ui32(v3489) < base.Ui32(v3490-v3488) {
		v3442 = v3489
		goto L525
	} else {
		goto L534
	}
L533:
	;
	goto L532
L534:
	;
	v3494 = v3489
	goto L527
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+12)) = v3538
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+24))
	v3544 = F_palloc(m, v3541<<(uint(int32(2))%32))
	mBase = m.M
	v3545 = m.ExcPending
	if v3545 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+20)) = v3544
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+16))
	if int32(0) < v3547 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v3552 = v3528
	goto L540
L538:
	;
	v3599 = v3528
	goto L539
L539:
	;
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+24))
	if int32(0) < v3628 {
		goto L543
	} else {
		goto L544
	}
L540:
	;
	v3581 = int32(2)
	v3582 = v3552 << (uint(v3581) % 32)
	v3583 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+12))
	v3585 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+8))
	v3587 = *(*int32)(unsafe.Add(mBase, uint32(v3280+v3582)))
	v3591 = *(*int32)(unsafe.Add(mBase, uint32(v3585+v3587<<(uint(v3581)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3582+v3583))) = v3591
	v3594 = v3552 + int32(1)
	v3595 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+16))
	if v3594 < v3595 {
		v3552 = v3594
		goto L540
	} else {
		goto L542
	}
L541:
	;
	v3599 = v3594
	goto L539
L542:
	;
	goto L541
L543:
	;
	v3634 = v3599
	v3640 = int32(0)
	goto L546
L544:
	;
	goto L545
L545:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v1986)))
	F_MemoryContextReset(m, v3714)
	mBase = m.M
	v3716 = m.ExcPending
	if v3716 != 0 {
		goto L1
	} else {
		goto L549
	}
L546:
	;
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+20))
	v3664 = int32(2)
	v3667 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+8))
	v3671 = *(*int32)(unsafe.Add(mBase, uint32(v3280+v3634<<(uint(v3664)%32))))
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(v3667+v3671<<(uint(v3664)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3663+v3640<<(uint(v3664)%32)))) = v3675
	v3677 = int32(1)
	v3680 = v3640 + v3677
	v3681 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+24))
	if v3680 < v3681 {
		v3634 = v3634 + v3677
		v3640 = v3680
		goto L546
	} else {
		goto L548
	}
L547:
	;
	goto L545
L548:
	;
	goto L547
L549:
	;
	goto L496
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3203)+12)) = v3725
	v3728 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+8))
	v3729 = *(*int32)(unsafe.Add(mBase, uint32(v3728)))
	*(*int32)(unsafe.Add(mBase, uint32(v3725))) = v3729
	goto L496
L551:
	;
	goto L495
L552:
	;
	v3837 = *(*int32)(unsafe.Add(mBase, uint32(v3806)+36))
	v3839 = v3837 + int32(4)
	v3854 = v3820
	v3856 = v3822
	v3859 = v3825
	goto L556
L554:
	;
	v4967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4938)+1092)))
	v4968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4938)+1090)))
	v4969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4938)+1088)))
	v3806 = v4934
	v3807 = v4935
	v3810 = v4938
	v3820 = v4968
	v3822 = v4969
	v3825 = v4967
	v3834 = v3834 + int64(1)
	goto L552
L555:
	;
	F_tbm_add_tuples(m, v3807, v3810+int32(1088), int32(1), v4908)
	mBase = m.M
	v4933 = m.ExcPending
	if v4933 != 0 {
		goto L1
	} else {
		goto L685
	}
L556:
	;
	v3872 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[16]))
	if v3872 != 0 {
		goto L558
	} else {
		goto L559
	}
L557:
	;
	if v4806 == int32(0) {
		goto L675
	} else {
		goto L676
	}
L558:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3874 = m.ExcPending
	if v3874 != 0 {
		goto L1
	} else {
		goto L561
	}
L559:
	;
	goto L560
L560:
	;
	v3875 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1092)) = uint16(v3875)
	*(*int32)(unsafe.Add(mBase, uint32(v3810)+1088)) = v3875
	v3880 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+uint32(_c_F_gingetbitmap[3])))
	if v3880 == v3875 {
		v4908 = v3875
		goto L555
	} else {
		goto L562
	}
L561:
	;
	goto L560
L562:
	;
	v3893 = v3875
	v3897 = v3854
	v3899 = v3856
	v3902 = v3859
	goto L563
L563:
	;
	v3914 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+uint32(_c_F_gingetbitmap[4])))
	v3917 = v3914 + v3893*int32(92)
	v3918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1092)))
	if v3918 != int32(_a_F_gingetbitmap_7) {
		goto L566
	} else {
		goto L567
	}
L564:
	;
	if v4803 == int32(0) {
		v3854 = v4786
		v3856 = v4788
		v3859 = v4791
		goto L556
	} else {
		goto L673
	}
L565:
	;
	v4805 = v3893 + int32(1)
	v4806 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+uint32(_c_F_gingetbitmap[3])))
	if v4803 != 0 {
		goto L669
	} else {
		goto L670
	}
L566:
	;
	v3928 = *(*int32)(unsafe.Add(mBase, uint32(v3837)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)) = uint16(v3902)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2114)) = uint16(v3897)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2112)) = uint16(v3899)
	v3932 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+84)))
	v3933 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+82)))
	v3934 = int64(32)
	v3936 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+80)))
	v3937 = int64(48)
	v3942 = int64(65535)
	if base.Ui64(base.I64_extend_i32_u(v3902)&v3942|(base.I64_extend_i32_u(v3899)<<(uint(v3937)%64)|base.I64_extend_i32_u(v3897)&v3942<<(uint(v3934)%64))) < base.Ui64(v3932|(v3933<<(uint(v3934)%64)|v3936<<(uint(v3937)%64))) {
		goto L570
	} else {
		goto L571
	}
L567:
	;
	v3921 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1088)))
	v3922 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1090)))
	if v3921&v3922 == int32(_a_F_gingetbitmap_7) {
		goto L566
	} else {
		goto L568
	}
L568:
	;
	v3927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3917)+78)))
	if v3927 != 0 {
		v4786 = v3897
		v4788 = v3899
		v4791 = v3902
		v4803 = int32(1)
		goto L565
	} else {
		goto L569
	}
L569:
	;
	goto L566
L570:
	;
	v4675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3917)+88)))
	if v4675 != 0 {
		v4996 = v3810
		v5020 = v3834
		goto L144
	} else {
		goto L648
	}
L571:
	;
	v3955 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+16))
	if v3955 == int32(0) {
		goto L575
	} else {
		goto L576
	}
L572:
	;
	v4216 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+24))
	if v4216 != 0 {
		goto L594
	} else {
		goto L595
	}
L573:
	;
	v4180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)))
	v4183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2114)))
	v4184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2112)))
	v4191 = v4180 + int32(1)
	v4193 = v4183
	v4197 = v4184
	goto L572
L574:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2114)) = uint16(v4071)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2112)) = uint16(v4073)
	v4147 = v4070 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)) = uint16(v4147)
	v4191 = v4070
	v4193 = v4071
	v4197 = v4073
	goto L572
L575:
	;
	v4140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3917)+78)))
	if v4140 != 0 {
		goto L573
	} else {
		goto L593
	}
L576:
	;
	v3958 = int32(_a_F_gingetbitmap_7)
	v3968 = int32(0)
	v3969 = v3958
	v3971 = v3958
	v3972 = int32(1)
	v3975 = v3958
	goto L577
L577:
	;
	v3994 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+12))
	v3998 = *(*int32)(unsafe.Add(mBase, uint32(v3994+v3968<<(uint(int32(2))%32))))
	v3999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3998)+654)))
	if v3999 != 0 {
		v4070 = v3969
		v4071 = v3971
		v4072 = v3972
		v4073 = v3975
		goto L579
	} else {
		goto L580
	}
L578:
	;
	if v4072&int32(1) != 0 {
		goto L575
	} else {
		goto L588
	}
L579:
	;
	v4079 = v3968 + int32(1)
	v4080 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+16))
	if base.Ui32(v4079) < base.Ui32(v4080) {
		v3968 = v4079
		v3969 = v4070
		v3971 = v4071
		v3972 = v4072
		v3975 = v4073
		goto L577
	} else {
		goto L587
	}
L580:
	;
	v4000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3998)+32)))
	v4002 = int64(65535)
	v4004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3998)+30)))
	v4008 = int64(32)
	v4010 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3998)+28)))
	v4012 = int64(48)
	v4015 = base.I64_extend_i32_u(v4000)&v4002 | (base.I64_extend_i32_u(v4004)&v4002<<(uint(v4008)%64) | base.I64_extend_i32_u(v4010)<<(uint(v4012)%64))
	v4016 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)))
	v4017 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2114)))
	v4020 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2112)))
	if base.Ui64(v4015) <= base.Ui64(v4016|(v4017<<(uint(v4008)%64)|v4020<<(uint(v4012)%64))) {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v4026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+12)) = uint16(v4026)
	v4028 = *(*int32)(unsafe.Add(mBase, uint32(v3810)+2112))
	*(*int32)(unsafe.Add(mBase, uint32(v3810)+8)) = v4028
	F_entryGetItem(m, v3839, v3998, v3810+int32(8))
	mBase = m.M
	v4033 = m.ExcPending
	if v4033 != 0 {
		goto L1
	} else {
		goto L584
	}
L582:
	;
	v4051 = v4010
	v4052 = v4004
	v4053 = v4000
	v4054 = v4015
	goto L583
L583:
	;
	v4055 = int32(0)
	v4057 = int64(65535)
	if base.Ui64(base.I64_extend_i32_u(v3969)&v4057|(base.I64_extend_i32_u(v3975)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v3971)&v4057<<(uint(int64(32))%64))) <= base.Ui64(v4054) {
		v4070 = v3969
		v4071 = v3971
		v4072 = v4055
		v4073 = v3975
		goto L579
	} else {
		goto L586
	}
L584:
	;
	v4034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3998)+654)))
	if v4034 != 0 {
		v4070 = v3969
		v4071 = v3971
		v4072 = v3972
		v4073 = v3975
		goto L579
	} else {
		goto L585
	}
L585:
	;
	v4035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3998)+32)))
	v4037 = int64(65535)
	v4039 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3998)+30)))
	v4045 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3998)+28)))
	v4051 = v4045
	v4052 = v4039
	v4053 = v4035
	v4054 = base.I64_extend_i32_u(v4035)&v4037 | (base.I64_extend_i32_u(v4039)&v4037<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4045)<<(uint(int64(48))%64))
	goto L583
L586:
	;
	v4070 = v4053
	v4071 = v4052
	v4072 = v4055
	v4073 = v4051
	goto L579
L587:
	;
	goto L578
L588:
	;
	v4084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3917)+78)))
	if v4084&int32(1) != 0 {
		goto L573
	} else {
		goto L589
	}
L589:
	;
	v4087 = int32(_a_F_gingetbitmap_7)
	if v4070&v4087 != v4087 {
		goto L574
	} else {
		goto L590
	}
L590:
	;
	v4095 = v4071&int32(_a_F_gingetbitmap_7) | v4073<<(uint(int32(16))%32)
	if v4095 == int32(-1) {
		goto L574
	} else {
		goto L591
	}
L591:
	;
	v4098 = int32(_a_F_gingetbitmap_7)
	v4099 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2114)))
	v4100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2112)))
	if base.Ui32(v4095) <= base.Ui32(v4099|v4100<<(uint(int32(16))%32)) {
		v4191 = v4098
		v4193 = v4071
		v4197 = v4073
		goto L572
	} else {
		goto L592
	}
L592:
	;
	v4105 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)) = uint16(v4105)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2114)) = uint16(v4071)
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2112)) = uint16(v4073)
	v4191 = v4098
	v4193 = v4071
	v4197 = v4073
	goto L572
L593:
	;
	v4141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3917)+88)) = uint8(v4141)
	goto L570
L594:
	;
	v4223 = int32(0)
	v4224 = v4191
	v4226 = v4193
	v4230 = v4197
	goto L597
L595:
	;
	v4339 = v4191
	v4341 = v4193
	v4345 = v4197
	goto L596
L596:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3917)+84)) = uint16(v4339)
	*(*uint16)(unsafe.Add(mBase, uint32(v3917)+82)) = uint16(v4341)
	*(*uint16)(unsafe.Add(mBase, uint32(v3917)+80)) = uint16(v4345)
	v4367 = *(*int32)(unsafe.Add(mBase, uint32(v3917)))
	if v4367 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L597:
	;
	v4249 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+20))
	v4253 = *(*int32)(unsafe.Add(mBase, uint32(v4249+v4223<<(uint(int32(2))%32))))
	v4254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4253)+654)))
	if v4254 != 0 {
		v4322 = v4224
		v4323 = v4226
		v4324 = v4230
		goto L599
	} else {
		goto L600
	}
L598:
	;
	v4339 = v4322
	v4341 = v4323
	v4345 = v4324
	goto L596
L599:
	;
	v4330 = v4223 + int32(1)
	v4331 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+24))
	if base.Ui32(v4330) < base.Ui32(v4331) {
		v4223 = v4330
		v4224 = v4322
		v4226 = v4323
		v4230 = v4324
		goto L597
	} else {
		goto L607
	}
L600:
	;
	v4255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4253)+32)))
	v4257 = int64(65535)
	v4259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4253)+30)))
	v4263 = int64(32)
	v4265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4253)+28)))
	v4267 = int64(48)
	v4270 = base.I64_extend_i32_u(v4255)&v4257 | (base.I64_extend_i32_u(v4259)&v4257<<(uint(v4263)%64) | base.I64_extend_i32_u(v4265)<<(uint(v4267)%64))
	v4271 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)))
	v4272 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2114)))
	v4275 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2112)))
	if base.Ui64(v4270) <= base.Ui64(v4271|(v4272<<(uint(v4263)%64)|v4275<<(uint(v4267)%64))) {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v4281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+2116)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3810)+4)) = uint16(v4281)
	v4283 = *(*int32)(unsafe.Add(mBase, uint32(v3810)+2112))
	*(*int32)(unsafe.Add(mBase, uint32(v3810))) = v4283
	F_entryGetItem(m, v3839, v4253, v3810)
	mBase = m.M
	v4286 = m.ExcPending
	if v4286 != 0 {
		goto L1
	} else {
		goto L604
	}
L602:
	;
	v4304 = v4265
	v4305 = v4259
	v4306 = v4255
	v4307 = v4270
	goto L603
L603:
	;
	v4309 = int64(65535)
	if base.Ui64(base.I64_extend_i32_u(v4224)&v4309|(base.I64_extend_i32_u(v4230)<<(uint(int64(48))%64)|base.I64_extend_i32_u(v4226)&v4309<<(uint(int64(32))%64))) <= base.Ui64(v4307) {
		v4322 = v4224
		v4323 = v4226
		v4324 = v4230
		goto L599
	} else {
		goto L606
	}
L604:
	;
	v4287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4253)+654)))
	if v4287 != 0 {
		v4322 = v4224
		v4323 = v4226
		v4324 = v4230
		goto L599
	} else {
		goto L605
	}
L605:
	;
	v4288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4253)+32)))
	v4290 = int64(65535)
	v4292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4253)+30)))
	v4298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4253)+28)))
	v4304 = v4298
	v4305 = v4292
	v4306 = v4288
	v4307 = base.I64_extend_i32_u(v4288)&v4290 | (base.I64_extend_i32_u(v4292)&v4290<<(uint(int64(32))%64) | base.I64_extend_i32_u(v4298)<<(uint(int64(48))%64))
	goto L603
L606:
	;
	v4322 = v4306
	v4323 = v4305
	v4324 = v4304
	goto L599
L607:
	;
	goto L598
L608:
	;
	v4510 = *(*int32)(unsafe.Add(mBase, uint32(v3917)))
	if v4510 != 0 {
		goto L626
	} else {
		goto L627
	}
L609:
	;
	v4370 = int32(_a_F_gingetbitmap_1)
	v4371 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3928
	v4494 = v4371
	goto L608
L610:
	;
	goto L611
L611:
	;
	v4375 = int64(65535)
	v4385 = int32(0)
	v4389 = v4385
	v4405 = v4385
	goto L612
L612:
	;
	v4418 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+8))
	v4422 = *(*int32)(unsafe.Add(mBase, uint32(v4418+v4389<<(uint(int32(2))%32))))
	v4423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4422)+654)))
	if v4423 != 0 {
		goto L615
	} else {
		goto L616
	}
L613:
	;
	v4454 = int32(_a_F_gingetbitmap_1)
	v4455 = *(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v3928
	if v4449&int32(1) == int32(0) {
		v4494 = v4455
		goto L608
	} else {
		goto L622
	}
L614:
	;
	v4451 = v4389 + int32(1)
	v4452 = *(*int32)(unsafe.Add(mBase, uint32(v3917)))
	if base.Ui32(v4451) < base.Ui32(v4452) {
		v4389 = v4451
		v4405 = v4449
		goto L612
	} else {
		goto L621
	}
L615:
	;
	v4444 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+28))
	v4446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4444+v4389))) = uint8(v4446)
	v4449 = v4405
	goto L614
L616:
	;
	v4424 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4422)+32)))
	v4425 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4422)+30)))
	v4428 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4422)+28)))
	if v4424|(v4425<<(uint(int64(32))%64)|v4428<<(uint(int64(48))%64)) != base.I64_extend_i32_u(v4341)&v4375<<(uint(int64(32))%64)|base.I64_extend_i32_u(v4345)<<(uint(int64(48))%64)|v4375 {
		goto L615
	} else {
		goto L617
	}
L617:
	;
	v4434 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+28))
	v4435 = v4434 + v4389
	v4436 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+4))
	if base.Ui32(v4389) < base.Ui32(v4436) {
		goto L618
	} else {
		goto L619
	}
L618:
	;
	v4438 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4435))) = uint8(v4438)
	v4449 = int32(1)
	goto L614
L619:
	;
	goto L620
L620:
	;
	v4441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4435))) = uint8(v4441)
	v4449 = v4441
	goto L614
L621:
	;
	goto L613
L622:
	;
	v4462 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+36))
	v4463 = m.T0[v4462].(func(*base.Module, int32) int32)(m, v3917)
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L1
	} else {
		goto L623
	}
L623:
	;
	v4465 = int32(1)
	if base.Ui32(v4465) < base.Ui32((v4463-v4465)&int32(255)) {
		v4494 = v4455
		goto L608
	} else {
		goto L624
	}
L624:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4455
	F_MemoryContextReset(m, v3928)
	mBase = m.M
	v4474 = m.ExcPending
	if v4474 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3917)+84)) = int32(16908287)
	*(*uint16)(unsafe.Add(mBase, uint32(v3917)+82)) = uint16(v4341)
	*(*uint16)(unsafe.Add(mBase, uint32(v3917)+80)) = uint16(v4345)
	goto L570
L626:
	;
	v4512 = int64(65535)
	v4515 = base.I64_extend_i32_u(v4341) & v4512 << (uint(int64(32)) % 64)
	v4518 = base.I64_extend_i32_u(v4345) << (uint(int64(48)) % 64)
	v4530 = int32(0)
	goto L629
L627:
	;
	goto L628
L628:
	;
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+36))
	v4630 = m.T0[v4629].(func(*base.Module, int32) int32)(m, v3917)
	mBase = m.M
	v4631 = m.ExcPending
	if v4631 != 0 {
		goto L1
	} else {
		goto L646
	}
L629:
	;
	v4559 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+8))
	v4563 = *(*int32)(unsafe.Add(mBase, uint32(v4559+v4530<<(uint(int32(2))%32))))
	v4564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4563)+654)))
	if v4564 == int32(1) {
		goto L632
	} else {
		goto L633
	}
L630:
	;
	goto L628
L631:
	;
	v4595 = v4530 + int32(1)
	v4596 = *(*int32)(unsafe.Add(mBase, uint32(v3917)))
	if base.Ui32(v4595) < base.Ui32(v4596) {
		v4530 = v4595
		goto L629
	} else {
		goto L641
	}
L632:
	;
	v4567 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+28))
	v4569 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4567+v4530))) = uint8(v4569)
	goto L631
L633:
	;
	goto L634
L634:
	;
	v4571 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4563)+32)))
	v4572 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4563)+30)))
	v4575 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v4563)+28)))
	v4579 = v4571 | (v4572<<(uint(int64(32))%64) | v4575<<(uint(int64(48))%64))
	if v4518|v4515|v4512 == v4579 {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v4581 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+28))
	v4583 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4581+v4530))) = uint8(v4583)
	goto L631
L636:
	;
	goto L637
L637:
	;
	v4585 = *(*int32)(unsafe.Add(mBase, uint32(v3917)+28))
	v4586 = v4585 + v4530
	if v4579 == v4515|(v4518|base.I64_extend_i32_u(v4339)&v4512) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v4588 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4586))) = uint8(v4588)
	goto L631
L639:
	;
	goto L640
L640:
	;
	v4590 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4586))) = uint8(v4590)
	goto L631
L641:
	;
	goto L630
L642:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gingetbitmap[0])) = v4494
	F_MemoryContextReset(m, v3928)
	mBase = m.M
	v4643 = m.ExcPending
	if v4643 != 0 {
		goto L1
	} else {
		goto L647
	}
L643:
	;
	v4638 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v3917)+86)) = uint16(v4638)
	goto L642
L644:
	;
	v4636 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3917)+86)) = uint8(v4636)
	goto L642
L645:
	;
	v4634 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v3917)+86)) = uint8(v4634)
	goto L642
L646:
	;
	switch v4630 & int32(255) {
	case 0:
		goto L644
	case 1:
		goto L645
	default:
		goto L643
	}
L647:
	;
	goto L570
L648:
	;
	v4676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3917)+86)))
	if v4676 == int32(1) {
		goto L652
	} else {
		goto L653
	}
L649:
	;
	if v3893 == int32(0) {
		goto L658
	} else {
		goto L659
	}
L650:
	;
	v4705 = v4679
	v4706 = v4680
	v4707 = v4703
	goto L649
L651:
	;
	v4703 = v4681 - int32(1)
	goto L650
L652:
	;
	v4679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+82)))
	v4680 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+80)))
	v4681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+84)))
	if v4681 != int32(_a_F_gingetbitmap_7) {
		goto L651
	} else {
		goto L655
	}
L653:
	;
	goto L654
L654:
	;
	v4696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+84)))
	v4697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+82)))
	v4698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3917)+80)))
	v3854 = v4697
	v3856 = v4698
	v3859 = v4696
	goto L556
L655:
	;
	v4686 = v4680<<(uint(int32(16))%32) | v4679
	if v4686 == int32(-1) {
		goto L651
	} else {
		goto L656
	}
L656:
	;
	if base.Ui32(v4686) <= base.Ui32(v3897&int32(_a_F_gingetbitmap_7)|v3899<<(uint(int32(16))%32)) {
		v4705 = v3897
		v4706 = v3899
		v4707 = v3902
		goto L649
	} else {
		goto L657
	}
L657:
	;
	v4703 = int32(0)
	goto L650
L658:
	;
	v4713 = v3917 + int32(80)
	v4714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4713)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v3810+int32(1092)))) = uint16(v4714)
	v4716 = *(*int32)(unsafe.Add(mBase, uint32(v4713)))
	*(*int32)(unsafe.Add(mBase, uint32(v3810)+1088)) = v4716
	v4786 = v4705
	v4788 = v4706
	v4791 = v4707
	v4803 = int32(1)
	goto L565
L659:
	;
	goto L660
L660:
	;
	if v4681 != int32(_a_F_gingetbitmap_7) {
		goto L662
	} else {
		goto L663
	}
L661:
	;
	v4737 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1092)))
	if v4737 != int64(65535) {
		goto L666
	} else {
		goto L667
	}
L662:
	;
	v4721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1090)))
	v4722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1088)))
	v4734 = v4721
	v4736 = v4722
	goto L661
L663:
	;
	goto L664
L664:
	;
	v4723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1090)))
	v4724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1088)))
	v4727 = v4680<<(uint(int32(16))%32) | v4679
	if v4727 == int32(-1) {
		v4734 = v4723
		v4736 = v4724
		goto L661
	} else {
		goto L665
	}
L665:
	;
	v4786 = v4705
	v4788 = v4706
	v4791 = v4707
	v4803 = base.B2i32(v4727 == v4724<<(uint(int32(16))%32)|v4723)
	goto L565
L666:
	;
	v4754 = int64(48)
	v4758 = int64(32)
	v4786 = v4705
	v4788 = v4706
	v4791 = v4707
	v4803 = base.B2i32(base.I64_extend_i32_u(v4681)|base.I64_extend_i32_u(v4680)<<(uint(v4754)%64)|base.I64_extend_i32_u(v4679)<<(uint(v4758)%64) == base.I64_extend_i32_u(v4736)<<(uint(v4754)%64)|v4737|base.I64_extend_i32_u(v4734)&int64(65535)<<(uint(v4758)%64))
	goto L565
L667:
	;
	v4744 = v4734&int32(_a_F_gingetbitmap_7) | v4736<<(uint(int32(16))%32)
	if v4744 == int32(-1) {
		goto L666
	} else {
		goto L668
	}
L668:
	;
	v4786 = v4705
	v4788 = v4706
	v4791 = v4707
	v4803 = base.B2i32(v4680<<(uint(int32(16))%32)|v4679 == v4744)
	goto L565
L669:
	;
	v4809 = base.B2i32(base.Ui32(v4805) < base.Ui32(v4806))
	goto L671
L670:
	;
	v4809 = int32(0)
	goto L671
L671:
	;
	if v4809 != 0 {
		v3893 = v4805
		v3897 = v4786
		v3899 = v4788
		v3902 = v4791
		goto L563
	} else {
		goto L672
	}
L672:
	;
	goto L564
L673:
	;
	goto L557
L674:
	;
	v4886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1092)))
	if v4886 != int32(_a_F_gingetbitmap_7) {
		v4908 = v4865
		goto L555
	} else {
		goto L682
	}
L675:
	;
	v4865 = int32(0)
	goto L674
L676:
	;
	goto L677
L677:
	;
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(v3837)+uint32(_c_F_gingetbitmap[4])))
	v4819 = int32(0)
	goto L678
L678:
	;
	v4851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4815+v4819*int32(92))+87)))
	if v4851 != 0 {
		v4865 = v4851
		goto L674
	} else {
		goto L680
	}
L679:
	;
	v4865 = v4851
	goto L674
L680:
	;
	v4853 = v4819 + int32(1)
	if v4853 != v4806 {
		v4819 = v4853
		goto L678
	} else {
		goto L681
	}
L681:
	;
	goto L679
L682:
	;
	v4889 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1090)))
	v4890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3810)+1088)))
	v4893 = v4889 | v4890<<(uint(int32(16))%32)
	if v4893 == int32(-1) {
		v4908 = v4865
		goto L555
	} else {
		goto L683
	}
L683:
	;
	F_tbm_add_page(m, v3807, v4893)
	mBase = m.M
	v4897 = m.ExcPending
	if v4897 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
	;
	v4934 = v3806
	v4935 = v3807
	v4938 = v3810
	goto L554
L685:
	;
	v4934 = v3806
	v4935 = v3807
	v4938 = v3810
	goto L554
L686:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v4976 = m.ExcPending
	if v4976 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	v4977 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+1128))
	v4978 = *(*int32)(unsafe.Add(mBase, uint32(v4977)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1959)+16)) = v4978 + int32(4)
	F_errmsg(m, int32(_a_F_gingetbitmap_15), v1959+int32(16))
	mBase = m.M
	v4986 = m.ExcPending
	if v4986 != 0 {
		goto L1
	} else {
		goto L688
	}
L688:
	;
	F_errfinish(m, int32(_a_F_gingetbitmap_11), int32(272), int32(_a_F_gingetbitmap_16))
	mBase = m.M
	v4991 = m.ExcPending
	if v4991 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
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
	var v56 int32
	_ = v56
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
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
				v250 = v2
				m.G0 = v16 + int32(16)
				return v250
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
							v250 = v44
						} else {
							v52 = int32(0)
							v53 = v48
							v56 = v2
							for {
								v67 = v29 + v52<<(uint(int32(3))%32)
								v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67))))
								if v68 == int32(2) {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v44+v56<<(uint(int32(2))%32)))) = v74
									v77 = v56 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v20))) = v77
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
									v80 = v79
									v81 = v77
								} else {
									v80 = v53
									v81 = v56
								}
								v83 = v52 + int32(1)
								if v83 < v80 {
									v52 = v83
									v53 = v80
									v56 = v81
									continue
								} else {
									break
								}
								break
							}
							v250 = v44
						}
						m.G0 = v16 + int32(16)
						return v250
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
						v268 = m.ExcPending
						if v268 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v271 = m.ExcPending
							if v271 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_ginint4_queryextract_0), int32(0))
								mBase = m.M
								v277 = m.ExcPending
								if v277 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(61), int32(_a_F_ginint4_queryextract_2))
									mBase = m.M
									v284 = m.ExcPending
									if v284 != 0 {
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
						v91 = F_ArrayGetNItems(m, v88, v24+int32(16))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20))) = v91
							if v91 <= int32(0) {
								v209 = v2
								v210 = int32(0)
								switch v19 - int32(3) {
								case 0:
									v246 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
									v250 = v210
									m.G0 = v16 + int32(16)
									return v250
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v231 = m.ExcPending
									if v231 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
										F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
											mBase = m.M
											v244 = m.ExcPending
											if v244 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								case 3:
									v246 = v209 ^ int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
									v250 = v210
									m.G0 = v16 + int32(16)
									return v250
								case 4, 10:
									if v209 != 0 {
										v227 = int32(0)
									} else {
										v227 = int32(2)
									}
									v246 = v227
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
									v250 = v210
									m.G0 = v16 + int32(16)
									return v250
								case 5, 11:
									v246 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
									v250 = v210
									m.G0 = v16 + int32(16)
									return v250
								}
							} else {
								v99 = F_palloc(m, v91<<(uint(int32(2))%32))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
									if v101 == int32(0) {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
										v111 = (v104<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									} else {
										v111 = v101
									}
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
									if v112 <= int32(0) {
										v209 = int32(0)
										v210 = v99
									} else {
										v116 = v111 + v24
										v118 = v112 & int32(3)
										v119 = int32(0)
										if base.Ui32(int32(4)) <= base.Ui32(v112) {
											v125 = v119
											v131 = int32(0)
											for {
												v139 = v125 << (uint(int32(2)) % 32)
												v142 = *(*int32)(unsafe.Add(mBase, uint32(v139+v116)))
												*(*int32)(unsafe.Add(mBase, uint32(v99+v139))) = v142
												v144 = int32(4)
												v145 = v139 | v144
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v145+v116)))
												*(*int32)(unsafe.Add(mBase, uint32(v99+v145))) = v148
												v151 = v139 | int32(8)
												v154 = *(*int32)(unsafe.Add(mBase, uint32(v151+v116)))
												*(*int32)(unsafe.Add(mBase, uint32(v99+v151))) = v154
												v157 = v139 | int32(12)
												v160 = *(*int32)(unsafe.Add(mBase, uint32(v157+v116)))
												*(*int32)(unsafe.Add(mBase, uint32(v99+v157))) = v160
												v163 = v125 + v144
												v165 = v131 + v144
												if v165 != v112&int32(2147483644) {
													v125 = v163
													v131 = v165
													continue
												} else {
													break
												}
												break
											}
											v167 = v163
										} else {
											v167 = v119
										}
										if v118 == int32(0) {
											v209 = int32(1)
											v210 = v99
										} else {
											v183 = v167
											v190 = v2
											for {
												v197 = v183 << (uint(int32(2)) % 32)
												v200 = *(*int32)(unsafe.Add(mBase, uint32(v197+v116)))
												*(*int32)(unsafe.Add(mBase, uint32(v99+v197))) = v200
												v202 = int32(1)
												v206 = v190 + v202
												if v206 != v118 {
													v183 = v183 + v202
													v190 = v206
													continue
												} else {
													break
												}
												break
											}
											v209 = v202
											v210 = v99
										}
									}
									switch v19 - int32(3) {
									case 0:
										v246 = v2
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
										v250 = v210
										m.G0 = v16 + int32(16)
										return v250
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v231 = m.ExcPending
										if v231 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
											F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
												mBase = m.M
												v244 = m.ExcPending
												if v244 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									case 3:
										v246 = v209 ^ int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
										v250 = v210
										m.G0 = v16 + int32(16)
										return v250
									case 4, 10:
										if v209 != 0 {
											v227 = int32(0)
										} else {
											v227 = int32(2)
										}
										v246 = v227
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
										v250 = v210
										m.G0 = v16 + int32(16)
										return v250
									case 5, 11:
										v246 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
										v250 = v210
										m.G0 = v16 + int32(16)
										return v250
									}
								}
							}
						}
					}
				}
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v91 = F_ArrayGetNItems(m, v88, v24+int32(16))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v91
					if v91 <= int32(0) {
						v209 = v2
						v210 = int32(0)
						switch v19 - int32(3) {
						case 0:
							v246 = v2
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
							v250 = v210
							m.G0 = v16 + int32(16)
							return v250
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v231 = m.ExcPending
							if v231 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
								F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
								mBase = m.M
								v237 = m.ExcPending
								if v237 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
									mBase = m.M
									v244 = m.ExcPending
									if v244 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v246 = v209 ^ int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
							v250 = v210
							m.G0 = v16 + int32(16)
							return v250
						case 4, 10:
							if v209 != 0 {
								v227 = int32(0)
							} else {
								v227 = int32(2)
							}
							v246 = v227
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
							v250 = v210
							m.G0 = v16 + int32(16)
							return v250
						case 5, 11:
							v246 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
							v250 = v210
							m.G0 = v16 + int32(16)
							return v250
						}
					} else {
						v99 = F_palloc(m, v91<<(uint(int32(2))%32))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
							if v101 == int32(0) {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
								v111 = (v104<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v111 = v101
							}
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
							if v112 <= int32(0) {
								v209 = int32(0)
								v210 = v99
							} else {
								v116 = v111 + v24
								v118 = v112 & int32(3)
								v119 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v112) {
									v125 = v119
									v131 = int32(0)
									for {
										v139 = v125 << (uint(int32(2)) % 32)
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v139+v116)))
										*(*int32)(unsafe.Add(mBase, uint32(v99+v139))) = v142
										v144 = int32(4)
										v145 = v139 | v144
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v145+v116)))
										*(*int32)(unsafe.Add(mBase, uint32(v99+v145))) = v148
										v151 = v139 | int32(8)
										v154 = *(*int32)(unsafe.Add(mBase, uint32(v151+v116)))
										*(*int32)(unsafe.Add(mBase, uint32(v99+v151))) = v154
										v157 = v139 | int32(12)
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v157+v116)))
										*(*int32)(unsafe.Add(mBase, uint32(v99+v157))) = v160
										v163 = v125 + v144
										v165 = v131 + v144
										if v165 != v112&int32(2147483644) {
											v125 = v163
											v131 = v165
											continue
										} else {
											break
										}
										break
									}
									v167 = v163
								} else {
									v167 = v119
								}
								if v118 == int32(0) {
									v209 = int32(1)
									v210 = v99
								} else {
									v183 = v167
									v190 = v2
									for {
										v197 = v183 << (uint(int32(2)) % 32)
										v200 = *(*int32)(unsafe.Add(mBase, uint32(v197+v116)))
										*(*int32)(unsafe.Add(mBase, uint32(v99+v197))) = v200
										v202 = int32(1)
										v206 = v190 + v202
										if v206 != v118 {
											v183 = v183 + v202
											v190 = v206
											continue
										} else {
											break
										}
										break
									}
									v209 = v202
									v210 = v99
								}
							}
							switch v19 - int32(3) {
							case 0:
								v246 = v2
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
								v250 = v210
								m.G0 = v16 + int32(16)
								return v250
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v231 = m.ExcPending
								if v231 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
									F_errmsg_internal(m, int32(_a_F_ginint4_queryextract_3), v16)
									mBase = m.M
									v237 = m.ExcPending
									if v237 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ginint4_queryextract_1), int32(100), int32(_a_F_ginint4_queryextract_2))
										mBase = m.M
										v244 = m.ExcPending
										if v244 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 3:
								v246 = v209 ^ int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
								v250 = v210
								m.G0 = v16 + int32(16)
								return v250
							case 4, 10:
								if v209 != 0 {
									v227 = int32(0)
								} else {
									v227 = int32(2)
								}
								v246 = v227
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
								v250 = v210
								m.G0 = v16 + int32(16)
								return v250
							case 5, 11:
								v246 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v18))) = v246
								v250 = v210
								m.G0 = v16 + int32(16)
								return v250
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	v5 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v10 = F_initGISTstate(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(_a_F_gistbeginscan_0)
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			*(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0])) = v15
			v18 = F_palloc0(m, int32(_a_F_gistbeginscan_1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v10
				v21 = F_createTempGistContext(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
					v31 = F_palloc(m, v28<<(uint(int32(4))%32))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)) = uint8(v33)
						*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v31
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
						if int32(0) < v36 {
							v41 = F_palloc0(m, v36<<(uint(int32(2))%32))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5)+76)) = v41
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
								v45 = F_palloc(m, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v5)+80)) = v45
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
									v51 = F__emscripten_memset_bulkmem(m, v45, base.I32_extend8_s(int32(1)), v49)
									mBase = m.M
									v53 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v53
									*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(-1)
									*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v53
									*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v18
									*(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0])) = v13
									return v5
								}
							}
						} else {
							v53 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v18
							*(*int32)(unsafe.Add(mBase, _c_F_gistbeginscan[0])) = v13
							return v5
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
	var v109 int64
	_ = v109
	var v118 int32
	_ = v118
	var v122 int64
	_ = v122
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 float64
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 float64
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v708 float64
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int64
	_ = v722
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
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	F_MemoryContextDelete(m, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L3
	} else {
		goto L163
	}
L2:
	;
	v452 = F_gistNewBuffer(m, l1, l0)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L3
	} else {
		goto L105
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
	v425 = m.ExcPending
	if v425 != 0 {
		goto L3
	} else {
		goto L101
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
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[1]))
	v143 = m.G0
	v145 = v143 - int32(16)
	m.G0 = v145
	v147 = int32(0)
	v149 = F_tuplesort_begin_common(m, v142, v147, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L34
	}
L11:
	;
	v122 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v122
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = int32(819)
	if v58 != v54 {
		goto L2
	} else {
		goto L33
	}
L12:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v109 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v109
	v118 = base.I32_div_s(int32(_a_F_gistbuild_0)-v106<<(uint(int32(13))%32), int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v118
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
	v151 = int32(_a_F_gistbuild_1)
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v154
	v157 = F_palloc(m, int32(12))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gistbuild[2])))
	if v160 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v180)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = int32(1847)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+40)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v149)+60)) = v157
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+36)) = uint8(v186)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+16)) = int32(1848)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = int32(1849)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = int32(1850)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(1851)
	v196 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v157)+8)) = uint16(v196)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = l0
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v149)+40))
	v203 = F_palloc0(m, v200*int32(36))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L42
	}
L37:
	;
	v165 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if v165 == int32(0) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = int32(102)
	F_errmsg_internal(m, int32(_a_F_gistbuild_2), v145)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_gistbuild_3), int32(511), int32(_a_F_gistbuild_4))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L36
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+44)) = v203
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v149)+40))
	if v206 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v152
	m.G0 = v145 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v149
	v281 = int32(1)
	v282 = int32(0)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+140))
	v292 = m.T0[v291].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v281, v282, v281, v282, int32(-1), int32(93), v15+int32(32), v282)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L3
	} else {
		goto L51
	}
L44:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	v214 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v203)+10)) = uint16(v214)
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+9)) = uint8(v217)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v213
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+20)) = uint8(v220)
	F_PrepareSortSupportFromGistIndexRel(m, l1, v203)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v149)+40))
	if v224 < int32(2) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v233 = v214
	goto L47
L47:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v149)+44))
	v242 = v239 + v233*int32(36)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246+v233<<(uint(int32(2))%32))))
	v251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+20)) = uint8(v251)
	v254 = v233 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+10)) = uint16(v254)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+9)) = uint8(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v250
	F_PrepareSortSupportFromGistIndexRel(m, l1, v242)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L49
	}
L48:
	;
	goto L43
L49:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v149)+40))
	if v254 < v261 {
		v233 = v254
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	F_tuplesort_performsort(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v301 = F_smgr_bulk_start_rel(m, v299, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v301
	v305 = F_palloc0(m, int32(28))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v308 = F_palloc(m, int32(_a_F_gistbuild_5))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+12)) = v308
	v313 = int32(1)
	F_PageInit(m, v308, int32(_a_F_gistbuild_5), int32(16))
	mBase = m.M
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308)+16)))
	v318 = v308 + v317
	v319 = int32(_a_F_gistbuild_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v318)+14)) = uint16(v319)
	*(*uint16)(unsafe.Add(mBase, uint32(v318)+12)) = uint16(v313)
	*(*int32)(unsafe.Add(mBase, uint32(v318)+8)) = int32(-1)
	goto L56
L56:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v325 = F_tuplesort_getheaptuple(m, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	if v325 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v332 = v325
	goto L61
L59:
	;
	goto L60
L60:
	;
	v365 = v305
	goto L67
L61:
	;
	F_gist_indexsortbuild_levelstate_add(m, v15+int32(32), v305, v332)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L3
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	F_MemoryContextReset(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v348 = F_tuplesort_getheaptuple(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	if v348 != 0 {
		v332 = v348
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	if v374 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v399))) = int64(4294967296)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	v403 = F_smgr_bulk_get_buf(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L92
	}
L69:
	;
	goto L68
L70:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	if v377 == int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_gist_indexsortbuild_levelstate_flush(m, v15+int32(32), v365)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L3
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	if v385 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_pfree(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L3
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
	if v388 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	F_pfree(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	if v391 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v365)+24))
	if v394 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	F_pfree(m, v365)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v365 = v384
	goto L67
L92:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	goto L94
L93:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	F_smgr_bulk_write(m, v409, int32(0), v407, int32(1))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L3
	} else {
		goto L97
	}
L94:
	;
	v407 = F__emscripten_memcpy_bulkmem(m, v403, v405, int32(_a_F_gistbuild_5))
	mBase = m.M
	goto L96
L96:
	;
	goto L93
L97:
	;
	F_pfree(m, v365)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
	F_smgr_bulk_finish(m, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	F_tuplesort_end(m, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	v708 = v292
	goto L1
L101:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v426 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistbuild_7), v15+int32(16))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_gistbuild_8), int32(195), int32(_a_F_gistbuild_9))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v472 = int32(_a_F_gistbuild_10)
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3]))
	v475 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3])) = v474 + v475
	if v452 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	if v452 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[4]))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v457+(v452^int32(-1))<<(uint(int32(2))%32))))
	v471 = v463
	goto L104
L107:
	;
	goto L108
L108:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[5]))
	v471 = v465 + v452<<(uint(int32(13))%32) + int32(-8192)
	goto L104
L109:
	;
	F_MarkBufferDirty(m, v452)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L3
	} else {
		goto L114
	}
L110:
	;
	F_PageInit(m, v496, int32(_a_F_gistbuild_5), int32(16))
	mBase = m.M
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496)+16)))
	v501 = v496 + v500
	v502 = int32(_a_F_gistbuild_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v501)+14)) = uint16(v502)
	*(*uint16)(unsafe.Add(mBase, uint32(v501)+12)) = uint16(v475)
	*(*int32)(unsafe.Add(mBase, uint32(v501)+8)) = int32(-1)
	goto L109
L111:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[4]))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v482+(v452^int32(-1))<<(uint(int32(2))%32))))
	v496 = v488
	goto L110
L112:
	;
	goto L113
L113:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[5]))
	v496 = v490 + v452<<(uint(int32(13))%32) + int32(-8192)
	goto L110
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v471))) = int64(4294967296)
	F_UnlockReleaseBuffer(m, v452)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v513 = int32(_a_F_gistbuild_10)
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3]))
	v516 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[3])) = v515 - v516
	v520 = int32(0)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+140))
	v530 = m.T0[v529].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v516, v520, v516, v520, int32(-1), int32(94), v15+int32(32), v520)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	if v532 == int32(4) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v537 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L3
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+118)))
	if v681 != int32(112) {
		v708 = v530
		goto L1
	} else {
		goto L155
	}
L120:
	;
	if v537 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_errmsg_internal(m, int32(_a_F_gistbuild_11), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L3
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v548 = int32(_a_F_gistbuild_1)
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0]))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v552
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+44))
	v557 = v555 - int32(1)
	if int32(0) <= v557 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	F_errfinish(m, int32(_a_F_gistbuild_8), int32(323), int32(_a_F_gistbuild_9))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v560 = v557
	goto L129
L127:
	;
	v656 = v554
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v549
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	F_BufFileClose(m, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L3
	} else {
		goto L154
	}
L129:
	;
	v573 = v560 << (uint(int32(2)) % 32)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v554)+40))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v573+v574)))
	if v576 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v656 = v650
	goto L128
L131:
	;
	v580 = v576
	goto L134
L132:
	;
	goto L133
L133:
	;
	v635 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L3
	} else {
		goto L147
	}
L134:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v580)+12))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	if v591 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L133
L136:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v554)+40))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v618+v573)))
	if v620 != 0 {
		v580 = v620
		goto L134
	} else {
		goto L146
	}
L137:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+16)))
	if v592 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v612 = F_list_delete_first(m, v580)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L3
	} else {
		goto L145
	}
L140:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v554)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v596
	v598 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v590)+16)) = uint8(v598)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v554)+28))
	v601 = F_lcons(m, v590, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L3
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	F_gistProcessEmptyingQueue(m, v15+int32(32))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L3
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v554)+28)) = v601
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistbuild[0])) = v606
	goto L142
L144:
	;
	goto L136
L145:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v554)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v614+v573))) = v612
	goto L136
L146:
	;
	goto L135
L147:
	;
	if v635 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v560
	F_errmsg_internal(m, int32(_a_F_gistbuild_12), v15)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L3
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	if int32(0) < v560 {
		v560 = v560 - int32(1)
		goto L129
	} else {
		goto L153
	}
L151:
	;
	F_errfinish(m, int32(_a_F_gistbuild_8), int32(1418), int32(_a_F_gistbuild_13))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L3
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	goto L130
L154:
	;
	goto L119
L155:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_gistbuild[6]))
	if v685 <= int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v688 != 0 {
		v708 = v530
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v690 = int32(0)
	v692 = F_RelationGetNumberOfBlocksInFork(m, l1, v690)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L3
	} else {
		goto L161
	}
L159:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v689 != 0 {
		v708 = v530
		goto L1
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	F_log_newpage_range(m, l1, v690, v692, int32(1))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	v708 = v530
	goto L1
L163:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	F_freeGISTstate(m, v715)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	v719 = F_palloc(m, int32(16))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v719))) = v708
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v719)+8)) = base.F64_convert_i64_s(v722)
	m.G0 = v15 + int32(96)
	return v719
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
	var v51 int32
	_ = v51
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
					v51 = v45
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)) = uint8(v51)
				}
				return
			}
		}
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
		v51 = v9
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)) = uint8(v51)
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
	F_freeGISTstate(m, v3)
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	v3 = int32(0)
	if v3 < l1 {
		v10 = int32(1)
		if l1 == v10 {
			v14 = int32(0)
			v47 = v14
			v48 = v14
		} else {
			v18 = int32(0)
			v21 = v18
			v22 = v18
			v23 = v3
			for {
				v26 = int32(2)
				v28 = l0 + v21<<(uint(v26)%32)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+6)))
				v31 = int32(_a_F_gistfitpage_0)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+6)))
				v40 = v22 + v30&v31 + v35&v31 + int32(8)
				v42 = v21 + v26
				v44 = v23 + v26
				if v44 != l1&int32(2147483646) {
					v21 = v42
					v22 = v40
					v23 = v44
					continue
				} else {
					break
				}
				break
			}
			v47 = v42
			v48 = v40
		}
		if l1&v10 != 0 {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0+v47<<(uint(int32(2))%32))))
			v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)))
			v62 = v56&int32(_a_F_gistfitpage_0) + v48 + int32(4)
		} else {
			v62 = v48
		}
		v71 = base.B2i32(base.Ui32(v62) < base.Ui32(int32(_a_F_gistfitpage_1)))
	} else {
		v71 = int32(1)
	}
	return v71
}
func F_gistgetadjusted(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(1392)
	m.G0 = v23
	F_gistDeCompressAtt(m, l3, l0, l1, v23+int32(736), v23+int32(192))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_gistDeCompressAtt(m, l3, l0, l2, v23+int32(224), v23+int32(160))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+10)))
	if v40 <= int32(0) {
		v290 = v5
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v23 + int32(1392)
	return v290
L5:
	;
	v46 = l3 + int32(_a_F_gistgetadjusted_0)
	v50 = v23 + int32(1268)
	v52 = v23 + int32(1252)
	v58 = v5
	v60 = v5
	goto L6
L6:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(160)+v58))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(192)+v58))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+1376)) = int32(0)
	v84 = v58 << (uint(int32(4)) % 32)
	v87 = v84 + (v23 + int32(736))
	v88 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+1248)) = v88
	v90 = v23 + v58
	v92 = v58 << (uint(v88) % 32)
	if v80 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v172 = int32(0)
	if v166 == v172 {
		v290 = v172
		goto L4
	} else {
		goto L30
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92+(v23+int32(32))))) = v141
	if v60&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v108 = v23 + int32(224) + v84
	if v80 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v76&int32(1) == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v102)
	v137 = v102
	v141 = int32(0)
	goto L8
L12:
	;
	v109 = v108
	goto L14
L13:
	;
	v109 = v87
	goto L14
L14:
	;
	v112 = (v80 | v76) & int32(1)
	if v112 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v113 = v109
	goto L17
L16:
	;
	v113 = v87
	goto L17
L17:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)))
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v113)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+8)) = v116
	if v112 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v118 = v109
	goto L20
L19:
	;
	v118 = v108
	goto L20
L20:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v118)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = v119
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v121
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v123)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v46+v92)))
	v135 = F_FunctionCall2Coll(m, l3+int32(916)+v58*int32(28), v130, v23+int32(1248), v23+int32(1376))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v137 = v123
	v141 = v135
	goto L8
L22:
	;
	v168 = v58 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v169)+10)))
	if v168 < v170 {
		v58 = v168
		v60 = v166
		goto L6
	} else {
		goto L29
	}
L23:
	;
	v166 = int32(1)
	goto L22
L24:
	;
	v145 = int32(0)
	if (v137|v76)&int32(1) != 0 {
		v166 = v145
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if v80 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1248)) = uint8(v150)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v46+v92)))
	v159 = F_FunctionCall3Coll(m, l3+int32(_a_F_gistgetadjusted_1)+v58*int32(28), v156, v149, v141, v23+int32(1248))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1248)))
	if v161 != 0 {
		v166 = v145
		goto L22
	} else {
		goto L28
	}
L28:
	;
	goto L23
L29:
	;
	goto L7
L30:
	;
	if int32(0) < v170 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v183 = v169
	v186 = v172
	goto L34
L32:
	;
	goto L33
L33:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v275 = F_index_form_tuple(m, v272, v23+int32(1248), v23)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L45
	}
L34:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v186))))
	if v202 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L33
L36:
	;
	v249 = v186 + int32(1)
	v250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244)+10)))
	if v249 < v250 {
		v183 = v244
		v186 = v249
		goto L34
	} else {
		goto L44
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(1248)+v186<<(uint(int32(2))%32)))) = int32(0)
	v244 = v183
	goto L36
L38:
	;
	goto L39
L39:
	;
	v212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1390)) = uint8(v212)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+1388)) = uint16(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+1384)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v23)+1380)) = l0
	v220 = v186 << (uint(int32(2)) % 32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220+(v23+int32(32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+1376)) = v224
	v228 = l3 + int32(1812) + v186*int32(28)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v229 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3+int32(_a_F_gistgetadjusted_0)+v220)))
	v234 = F_FunctionCall1Coll(m, v228, v231, v23+int32(1376))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v238 = v183
	v239 = v224
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(1248)+v220))) = v239
	v244 = v238
	goto L36
L43:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v238 = v237
	v239 = v236
	goto L42
L44:
	;
	goto L35
L45:
	;
	v278 = v275 + int32(4)
	v279 = int32(_a_F_gistgetadjusted_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v278))) = uint16(v279)
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v278))) = uint16(v281)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v283
	v290 = v275
	goto L4
}
func F_gistgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v450
L2:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v382 != 0 {
		goto L89
	} else {
		goto L90
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)))
	if v18 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L17
	} else {
		goto L86
	}
L6:
	;
	v450 = int32(0)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+17)))
	if v21 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+272))
	if v25 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v68 {
		goto L2
	} else {
		goto L27
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+268)))
	if v28 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v37 = v25
	goto L15
L15:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = v38 + int64(1)
	goto L12
L16:
	;
	F_pgstat_assoc_relation(m, v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
	v37 = v36
	goto L15
L19:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v43))) = v44 + int64(1)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[0]))) = v48
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+17)) = uint8(v48)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v48
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[1])))
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_MemoryContextReset(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v59
	F_gistScanPage(m, l0, v13, v59, v59, v59)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	goto L11
L27:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[0]))))
	if base.Ui32(v72) <= base.Ui32(v71) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v77 = v72
	v79 = v71
	goto L31
L29:
	;
	v282 = v71
	goto L30
L30:
	;
	if v282 == int32(0) {
		v333 = v282
		goto L75
	} else {
		goto L76
	}
L31:
	;
	v86 = int32(_a_F_gistgettuple_0)
	v87 = v79 & v86
	if v87 != v77&v86 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v282 = v277
	goto L30
L33:
	;
	goto L42
L34:
	;
	if v87 == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v93&int32(1) == int32(0) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v98 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = int32(_a_F_gistgettuple_1)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v105
	v108 = F_palloc(m, int32(816))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L40
	}
L38:
	;
	v114 = v98
	goto L39
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if int32(407) < v115 {
		goto L33
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v108
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v102
	v114 = v108
	goto L39
L41:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(44)+v118<<(uint(int32(4))%32)))))
	v123 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v115 + v123
	*(*uint16)(unsafe.Add(mBase, uint32(v114+v115<<(uint(v123)%32)))) = uint16(v122)
	goto L33
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v143 == int32(-1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	if base.Ui32(v274) <= base.Ui32(v277) {
		v77 = v274
		v79 = v277
		goto L31
	} else {
		goto L74
	}
L44:
	;
	v251 = int32(0)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	if v253 == v251 {
		v450 = v251
		goto L1
	} else {
		goto L64
	}
L45:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v146 <= int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+32))
	v152 = F_ReadBuffer(m, v149, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	if v152 == int32(0) {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	F_LockBuffer(m, v152, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_gistcheckpage(m, v159, v152)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	if v152 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v180 = F_BufferGetLSNAtomic(m, v152)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L17
	} else {
		goto L56
	}
L52:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[4]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165+(v152^int32(-1))<<(uint(int32(2))%32))))
	v179 = v171
	goto L51
L53:
	;
	goto L54
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[5]))
	v179 = v173 + v152<<(uint(int32(13))%32) + int32(-8192)
	goto L51
L55:
	;
	F_UnlockReleaseBuffer(m, v152)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L17
	} else {
		goto L63
	}
L56:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v150)+40))
	if v180 != v182 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v150)+28))
	if v184 <= int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v191 = int32(0)
	goto L59
L59:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
	v201 = int32(1)
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200+v191<<(uint(v201)%32)))))
	v209 = v204<<(uint(int32(2))%32) + (v179 + int32(24)) - int32(4)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v210 | int32(_a_F_gistgettuple_2)
	v215 = v191 + v201
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v150)+28))
	if v215 < v216 {
		v191 = v215
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179)+16)))
	v219 = v179 + v218
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+12)))
	v222 = v220 | int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+12)) = uint16(v222)
	F_MarkBufferDirtyHint(m, v152, int32(1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L17
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	goto L55
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+28)) = int32(0)
	goto L44
L64:
	;
	v256 = F_pairingheap_remove_first(m, v252)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L17
	} else {
		goto L65
	}
L65:
	;
	if v256 == int32(0) {
		v450 = v251
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[6]))
	if v261 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L17
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v264
	v268 = int32(0)
	F_gistScanPage(m, l0, v256, v256+int32(32), v268, v268)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L17
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v256)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[0]))))
	if v274 == int32(0) {
		goto L42
	} else {
		goto L73
	}
L73:
	;
	goto L43
L74:
	;
	goto L32
L75:
	;
	v336 = v17 + int32(48)
	v339 = int32(4)
	v341 = v336 + v333&int32(_a_F_gistgettuple_0)<<(uint(v339)%32)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v342
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v346)
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336+v348<<(uint(v339)%32))+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v352)
	v354 = int32(1)
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v355 == v354 {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v291&int32(1) == int32(0) {
		v333 = v282
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v296 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v299 = int32(_a_F_gistgettuple_1)
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3]))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v303
	v306 = F_palloc(m, int32(816))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L17
	} else {
		goto L81
	}
L79:
	;
	v313 = v282
	v314 = v296
	goto L80
L80:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if int32(407) < v315 {
		v333 = v313
		goto L75
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v306
	*(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[3])) = v300
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	v313 = v311
	v314 = v306
	goto L80
L82:
	;
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v313&int32(_a_F_gistgettuple_0)<<(uint(int32(4))%32))+44)))
	v324 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v315 + v324
	*(*uint16)(unsafe.Add(mBase, uint32(v314+v315<<(uint(v324)%32)))) = uint16(v323)
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	v333 = v331
	goto L75
L83:
	;
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v336+v358<<(uint(int32(4))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v362
	goto L85
L84:
	;
	goto L85
L85:
	;
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))))
	v366 = v364 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_gistgettuple[2]))) = uint16(v366)
	v450 = v354
	goto L1
L86:
	;
	F_errmsg_internal(m, int32(_a_F_gistgettuple_3), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L17
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_gistgettuple_4), int32(617), int32(_a_F_gistgettuple_5))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L17
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_pfree(m, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L17
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v381)+8))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+8))
	if v388 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	goto L91
L93:
	;
	v450 = int32(0)
	goto L1
L94:
	;
	goto L95
L95:
	;
	v393 = l0 + int32(60)
	v395 = v387
	goto L96
L96:
	;
	v404 = F_pairingheap_remove_first(m, v395)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L17
	} else {
		goto L98
	}
L97:
	;
	v450 = v436
	goto L1
L98:
	;
	if v404 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v450 = int32(0)
	goto L1
L100:
	;
	goto L101
L101:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v404)+12))
	if v409 == int32(-1) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v404)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = v412
	v414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v393)+4)) = uint16(v414)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v416)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+23)))
	F_index_store_float8_orderby_distances(m, l0, v418, v404+int32(32), v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L17
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_gistgettuple[6]))
	if v433 != 0 {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v424 = int32(1)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v425 == v424 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v404)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v428
	goto L108
L107:
	;
	goto L108
L108:
	;
	F_pfree(m, v404)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L17
	} else {
		goto L109
	}
L109:
	;
	v450 = v424
	goto L1
L110:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L17
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v436 = int32(0)
	F_gistScanPage(m, l0, v404, v404+int32(32), v436, v436)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L17
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	F_pfree(m, v404)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L17
	} else {
		goto L115
	}
L115:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v381)+8))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+8))
	if v446 != 0 {
		v395 = v445
		goto L96
	} else {
		goto L116
	}
L116:
	;
	goto L97
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
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
					if l6 == int32(0) {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						if v62 != 0 {
							F_gistfinishsplit(m, l0, l1, l2, v62, l8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								F_LockBuffer(m, v67, int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(16)
									return v51
								}
							}
						}
					} else {
						if l9 == int32(0) {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
							if v62 != 0 {
								F_gistfinishsplit(m, l0, l1, l2, v62, l8)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
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
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									F_LockBuffer(m, v67, int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										m.G0 = v14 + int32(16)
										return v51
									}
								}
							}
						} else {
							F_LockBuffer(m, l6, int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
								if v62 != 0 {
									F_gistfinishsplit(m, l0, l1, l2, v62, l8)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
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
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										F_LockBuffer(m, v67, int32(0))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
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
			} else {
				if l6 == int32(0) {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
					if v62 != 0 {
						F_gistfinishsplit(m, l0, l1, l2, v62, l8)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
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
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							F_LockBuffer(m, v67, int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v14 + int32(16)
								return v51
							}
						}
					}
				} else {
					if l9 == int32(0) {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
						if v62 != 0 {
							F_gistfinishsplit(m, l0, l1, l2, v62, l8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								F_LockBuffer(m, v67, int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(16)
									return v51
								}
							}
						}
					} else {
						F_LockBuffer(m, l6, int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
							if v62 != 0 {
								F_gistfinishsplit(m, l0, l1, l2, v62, l8)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
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
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									F_LockBuffer(m, v67, int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
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
	var v57 int32
	_ = v57
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
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
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
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
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
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v642 int32
	_ = v642
	var v653 int64
	_ = v653
	var v654 int64
	_ = v654
	var v659 int32
	_ = v659
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int64
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
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
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if int32(0) < v335 {
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
	v308 = m.ExcPending
	if v308 != 0 {
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
		v323 = v41
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
	v57 = int32(0)
	v59 = v41
	goto L14
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(48)+v57<<(uint(int32(2))%32))))
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
	v323 = v299
	goto L1
L16:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	if v109 != v30 {
		v299 = v108
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
	v302 = v57 + int32(1)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v302 < v303 {
		v57 = v302
		v59 = v299
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
		goto L37
	case 2, 3, 8:
		goto L36
	case 4:
		goto L35
	case 5:
		goto L34
	case 6:
		goto L33
	case 7:
		goto L32
	case 9:
		goto L31
	case 10:
		goto L30
	case 11:
		goto L29
	default:
		goto L27
	}
L26:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L65
	}
L27:
	;
	v265 = int32(0)
	v268 = F_errstart(m, int32(17), v265)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L2
	} else {
		goto L63
	}
L28:
	;
	v256 = int32(0)
	v259 = F_errstart(m, int32(17), v256)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L2
	} else {
		goto L61
	}
L29:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = int32(23)
	v240 = int32(1)
	v245 = F_check_amproc_signature(m, v236, int32(21), v240, v240, v240, v20+int32(272))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L57
	}
L30:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = int32(2281)
	v227 = int32(1)
	v232 = F_check_amproc_signature(m, v223, int32(2278), v227, v227, v227, v20+int32(256))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L55
	}
L31:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v219 = F_check_amoptsproc_signature(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L53
	}
L32:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v201 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v201
	*(*int64)(unsafe.Add(mBase, uint32(v20)+232)) = int64(111669149717)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v201
	v210 = int32(5)
	v214 = F_check_amproc_signature(m, v200, int32(701), int32(0), v210, v210, v20+int32(224))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L51
	}
L33:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v186 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+216)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v51
	v192 = int32(3)
	v196 = F_check_amproc_signature(m, v185, v186, int32(0), v192, v192, v20+int32(208))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L49
	}
L34:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+192)) = int64(9796820404457)
	v177 = int32(2)
	v181 = F_check_amproc_signature(m, v172, int32(2281), int32(1), v177, v177, v20+int32(192))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L47
	}
L35:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v158 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(9796820404457)
	v164 = int32(3)
	v168 = F_check_amproc_signature(m, v157, v158, int32(1), v164, v164, v20+int32(176))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L45
	}
L36:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v145 = int32(2281)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v145
	v148 = int32(1)
	v153 = F_check_amproc_signature(m, v144, v145, v148, v148, v148, v20+int32(160))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L43
	}
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+144)) = int64(9796820404457)
	v136 = int32(2)
	v140 = F_check_amproc_signature(m, v132, v51, int32(0), v136, v136, v20+int32(144))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
		goto L28
	} else {
		goto L40
	}
L40:
	;
	v299 = v108
	goto L24
L41:
	;
	if v140 == int32(0) {
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v299 = v108
	goto L24
L43:
	;
	if v153 == int32(0) {
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v299 = v108
	goto L24
L45:
	;
	if v168 == int32(0) {
		goto L28
	} else {
		goto L46
	}
L46:
	;
	v299 = v108
	goto L24
L47:
	;
	if v181 == int32(0) {
		goto L28
	} else {
		goto L48
	}
L48:
	;
	v299 = v108
	goto L24
L49:
	;
	if v196 == int32(0) {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	v299 = v108
	goto L24
L51:
	;
	if v214 == int32(0) {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	v299 = v108
	goto L24
L53:
	;
	if v219 == int32(0) {
		goto L28
	} else {
		goto L54
	}
L54:
	;
	v299 = v108
	goto L24
L55:
	;
	if v232 == int32(0) {
		goto L28
	} else {
		goto L56
	}
L56:
	;
	v299 = v108
	goto L24
L57:
	;
	if v245 == int32(0) {
		goto L28
	} else {
		goto L58
	}
L58:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	if v249 != int32(2276) {
		goto L28
	} else {
		goto L59
	}
L59:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v252 == int32(2276) {
		v299 = v108
		goto L24
	} else {
		goto L60
	}
L60:
	;
	goto L28
L61:
	;
	if v259 == int32(0) {
		v299 = v256
		goto L24
	} else {
		goto L62
	}
L62:
	;
	v275 = int32(165)
	v276 = int32(_a_F_gistvalidate_4)
	goto L26
L63:
	;
	if v268 == int32(0) {
		v299 = v265
		goto L24
	} else {
		goto L64
	}
L64:
	;
	v275 = int32(153)
	v276 = int32(_a_F_gistvalidate_5)
	goto L26
L65:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v281 = F_format_procedure(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v33
	F_errmsg(m, v276, v20+int32(96))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), v275, int32(_a_F_gistvalidate_3))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v299 = int32(0)
	goto L24
L69:
	;
	goto L15
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(_a_F_gistvalidate_6), v20)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(52), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	v344 = int32(0)
	v346 = v323
	goto L76
L74:
	;
	v512 = v323
	goto L75
L75:
	;
	v524 = F_identify_opfamily_groups(m, v39, v46)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L2
	} else {
		goto L117
	}
L76:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(48)+v344<<(uint(int32(2))%32))))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+56))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+22)))
	v364 = v362 + v363
	v365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v364)+16)))
	if int32(0) < v365 {
		v398 = v346
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v512 = v502
	goto L75
L78:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+18)))
	if v400 == int32(115) {
		v468 = int32(16)
		v469 = v398
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v368 = int32(0)
	v371 = F_errstart(m, int32(17), v368)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	if v371 == int32(0) {
		v398 = v368
		goto L78
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v364)+20))
	v379 = F_format_operator(m, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v381 = int32(*(*int16)(unsafe.Add(mBase, uint32(v364)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v381
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_7), v20+int32(80))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(185), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v398 = v368
	goto L78
L86:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v364)+20))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v364)+8))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	v473 = F_check_amop_signature(m, v470, v468, v471, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L2
	} else {
		goto L107
	}
L87:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v364)+8))
	v405 = F_get_opfamily_proc(m, v32, v403, v403, int32(8))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L2
	} else {
		goto L89
	}
L88:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v364)+20))
	v436 = F_get_op_rettype(m, v435)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L2
	} else {
		goto L97
	}
L89:
	;
	if v405 != 0 {
		v434 = v398
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v407 = int32(0)
	v410 = F_errstart(m, int32(17), v407)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	if v410 == int32(0) {
		v434 = v407
		goto L88
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v364)+20))
	v418 = F_format_operator(m, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_8), v20-int32(-64))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(202), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	v434 = v407
	goto L88
L97:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v364)+28))
	v439 = F_opfamily_can_sort_type(m, v438, v436)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	if v439 != 0 {
		v468 = v436
		v469 = v434
		goto L86
	} else {
		goto L99
	}
L99:
	;
	v441 = int32(0)
	v444 = F_errstart(m, int32(17), v441)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v444 == int32(0) {
		v468 = v436
		v469 = v441
		goto L86
	} else {
		goto L101
	}
L101:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v364)+20))
	v452 = F_format_operator(m, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_9), v20+int32(48))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(213), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v468 = v436
	v469 = v441
	goto L86
L106:
	;
	v504 = v344 + int32(1)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if v504 < v505 {
		v344 = v504
		v346 = v502
		goto L76
	} else {
		goto L115
	}
L107:
	;
	if v473 != 0 {
		v502 = v469
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v475 = int32(0)
	v478 = F_errstart(m, int32(17), v475)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	if v478 == int32(0) {
		v502 = v475
		goto L106
	} else {
		goto L110
	}
L110:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v364)+20))
	v486 = F_format_operator(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v33
	F_errmsg(m, int32(_a_F_gistvalidate_10), v20+int32(32))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(232), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v502 = v475
	goto L106
L115:
	;
	goto L77
L116:
	;
	v642 = v512
	v653 = int64(1)
	goto L149
L117:
	;
	if v524 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v617 = int32(0)
	goto L116
L119:
	;
	goto L120
L120:
	;
	v529 = int32(0)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if v530 <= v529 {
		v617 = v529
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v533 = int32(0)
	if v533 < v530 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v536 = v530
	goto L124
L123:
	;
	v536 = v533
	goto L124
L124:
	;
	v537 = int32(1)
	if v530 == v537 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v536&v537 == int32(0) {
		v617 = v588
		goto L116
	} else {
		goto L144
	}
L126:
	;
	v541 = int32(0)
	v588 = v541
	v591 = v541
	goto L125
L127:
	;
	goto L128
L128:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	v546 = int32(0)
	v549 = v546
	v552 = v546
	v553 = v546
	goto L129
L129:
	;
	v568 = v545 + v552<<(uint(int32(2))%32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	if v30 == v570 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v588 = v582
	v591 = v584
	goto L125
L131:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	if v572 == v30 {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v575 = v549
	goto L133
L133:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	if v30 == v577 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v574 = v569
	goto L136
L135:
	;
	v574 = v549
	goto L136
L136:
	;
	v575 = v574
	goto L133
L137:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v579 == v30 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v582 = v575
	goto L139
L139:
	;
	v583 = int32(2)
	v584 = v552 + v583
	v586 = v553 + v583
	if v586 != v536&int32(2147483646) {
		v549 = v582
		v552 = v584
		v553 = v586
		goto L129
	} else {
		goto L143
	}
L140:
	;
	v581 = v576
	goto L142
L141:
	;
	v581 = v575
	goto L142
L142:
	;
	v582 = v581
	goto L139
L143:
	;
	goto L130
L144:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v524)+12))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v607+v591<<(uint(int32(2))%32))))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	if v612 != v30 {
		v617 = v588
		goto L116
	} else {
		goto L145
	}
L145:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if v614 == v30 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v616 = v611
	goto L148
L147:
	;
	v616 = v588
	goto L148
L148:
	;
	v617 = v616
	goto L116
L149:
	;
	if v617 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L2
	} else {
		goto L167
	}
L151:
	;
	v699 = v653 + int64(1)
	if v699 != int64(13) {
		v642 = v697
		v653 = v699
		goto L149
	} else {
		goto L166
	}
L152:
	;
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v617)+16))
	if base.I32_wrap_i64(int64(base.Ui64(v654)>>(uint(v653)%64)))&int32(1) != 0 {
		v697 = v642
		goto L151
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v659 = base.I32_wrap_i64(v653)
	if v659&int32(12) == int32(8) {
		v697 = v642
		goto L151
	} else {
		goto L156
	}
L155:
	;
	goto L154
L156:
	;
	if int32(1)<<(uint(v659)%32)&int32(_a_F_gistvalidate_11) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v671 = base.B2i32(base.Ui32(v659) <= base.Ui32(int32(12)))
	goto L159
L158:
	;
	v671 = int32(0)
	goto L159
L159:
	;
	if v671 != 0 {
		v697 = v642
		goto L151
	} else {
		goto L160
	}
L160:
	;
	v672 = int32(0)
	v675 = F_errstart(m, int32(17), v672)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L2
	} else {
		goto L161
	}
L161:
	;
	if v675 == int32(0) {
		v697 = v672
		goto L151
	} else {
		goto L162
	}
L162:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L2
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v659
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_gistvalidate_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v29 + int32(8)
	F_errmsg(m, int32(_a_F_gistvalidate_12), v20+int32(16))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_gistvalidate_2), int32(273), int32(_a_F_gistvalidate_3))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	v697 = v672
	goto L151
L166:
	;
	goto L150
L167:
	;
	F_ReleaseCatCacheList(m, v39)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	F_ReleaseCatCache(m, v23)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	m.G0 = v20 + int32(304)
	return v697 & int32(1)
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
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int64
	_ = v129
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int64
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int64
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v264 int64
	_ = v264
	var v265 int32
	_ = v265
	var v268 int64
	_ = v268
	var v269 int32
	_ = v269
	var v272 int64
	_ = v272
	var v276 int64
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int64
	_ = v287
	var v296 int64
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		if v19&int32(1) != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v23 = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(base.Ui32(int32(base.Ui32(v22)>>(uint(v23)%32))-int32(8)) >> (uint(v23) % 32))
			v33 = F_psprintf(m, int32(_a_F_gtsvectorout_0), v12+int32(16))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v315 = v33
				v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v316 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v319 = m.ExcPending
					if v319 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(32)
						return v315
					}
				} else {
					m.G0 = v12 + int32(32)
					return v315
				}
			}
		} else {
			if v19&int32(4) != 0 {
				v38 = F_pstrdup(m, int32(_a_F_gtsvectorout_1))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v315 = v38
					v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v316 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v319 = m.ExcPending
						if v319 != 0 {
							return int32(0)
						} else {
							m.G0 = v12 + int32(32)
							return v315
						}
					} else {
						m.G0 = v12 + int32(32)
						return v315
					}
				}
			} else {
				v40 = int32(8)
				v41 = v15 + v40
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v44 = int32(base.Ui32(v42) >> (uint(int32(2)) % 32))
				v46 = v44 - v40
				if base.Ui32(v42) <= base.Ui32(int32(47)) {
					if v46 == int32(0) {
						v296 = v9
					} else {
						v51 = int32(3)
						v52 = v44 & v51
						if base.Ui32(v44-int32(9)) < base.Ui32(v51) {
							v95 = v41
							v102 = v9
						} else {
							v61 = v41
							v62 = int32(0)
							v68 = v9
							for {
								v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
								v72 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_gtsvectorout[0]))))
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
								v76 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_gtsvectorout[0]))))
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
								v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_gtsvectorout[0]))))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								v84 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_c_F_gtsvectorout[0]))))
								v88 = v72 + (v76 + (v80 + (v68 + v84)))
								v89 = int32(4)
								v90 = v61 + v89
								v92 = v62 + v89
								if v92 != v46&int32(-4) {
									v61 = v90
									v62 = v92
									v68 = v88
									continue
								} else {
									break
								}
								break
							}
							v95 = v90
							v102 = v88
						}
						if v52 == int32(0) {
							v296 = v102
						} else {
							v107 = v95
							v108 = int32(0)
							v114 = v102
							for {
								v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
								v118 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_gtsvectorout[0]))))
								v119 = v114 + v118
								v120 = int32(1)
								v123 = v108 + v120
								if v123 != v52 {
									v107 = v107 + v120
									v108 = v123
									v114 = v119
									continue
								} else {
									break
								}
								break
							}
							v296 = v119
						}
					}
				} else {
					v129 = int64(0)
					if v46 < int32(4) {
						v208 = v41
						v209 = v46
						v214 = v129
					} else {
						if v41 != (v15+int32(11))&int32(-4) {
							v208 = v41
							v209 = v46
							v214 = v129
						} else {
							v138 = v46 - int32(4)
							v142 = int32(base.Ui32(v138)>>(uint(int32(2))%32)) + int32(1)
							v144 = v142 & int32(3)
							if base.Ui32(v138) < base.Ui32(int32(12)) {
								v180 = v41
								v181 = v46
								v186 = v129
							} else {
								v150 = v41
								v151 = v46
								v152 = int32(0)
								v156 = v129
								for {
									v157 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
									v172 = base.I64_extend_i32_u(base.I32_popcnt(v157)) + (base.I64_extend_i32_u(base.I32_popcnt(v160)) + (base.I64_extend_i32_u(base.I32_popcnt(v163)) + (v156 + base.I64_extend_i32_u(base.I32_popcnt(v166)))))
									v173 = int32(16)
									v174 = v151 - v173
									v176 = v150 + v173
									v178 = v152 + int32(4)
									if v178 != v142&int32(2147483644) {
										v150 = v176
										v151 = v174
										v152 = v178
										v156 = v172
										continue
									} else {
										break
									}
									break
								}
								v180 = v176
								v181 = v174
								v186 = v172
							}
							if v144 == int32(0) {
								v208 = v180
								v209 = v181
								v214 = v186
							} else {
								v191 = v181
								v192 = v180
								v193 = int32(0)
								v196 = v186
								for {
									v197 = int32(4)
									v198 = v191 - v197
									v199 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
									v202 = v196 + base.I64_extend_i32_u(base.I32_popcnt(v199))
									v204 = v192 + v197
									v206 = v193 + int32(1)
									if v206 != v144 {
										v191 = v198
										v192 = v204
										v193 = v206
										v196 = v202
										continue
									} else {
										break
									}
									break
								}
								v208 = v204
								v209 = v198
								v214 = v202
							}
						}
					}
					if v209 == int32(0) {
						v287 = v214
					} else {
						v218 = v209 & int32(3)
						if v218 == int32(0) {
							v241 = v208
							v243 = v209
							v247 = v214
						} else {
							v224 = v209
							v225 = v208
							v226 = int32(0)
							v228 = v214
							for {
								v229 = int32(1)
								v230 = v224 - v229
								v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
								v234 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v231)+uint32(_c_F_gtsvectorout[0]))))
								v235 = v228 + v234
								v237 = v225 + v229
								v239 = v226 + v229
								if v239 != v218 {
									v224 = v230
									v225 = v237
									v226 = v239
									v228 = v235
									continue
								} else {
									break
								}
								break
							}
							v241 = v237
							v243 = v230
							v247 = v235
						}
						if base.Ui32(v209) < base.Ui32(int32(4)) {
							v287 = v247
						} else {
							v250 = v241
							v252 = v243
							v256 = v247
							for {
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+3)))
								v260 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v257)+uint32(_c_F_gtsvectorout[0]))))
								v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+2)))
								v264 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v261)+uint32(_c_F_gtsvectorout[0]))))
								v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
								v268 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_gtsvectorout[0]))))
								v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
								v272 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v269)+uint32(_c_F_gtsvectorout[0]))))
								v276 = v260 + (v264 + (v268 + (v256 + v272)))
								v277 = int32(4)
								v280 = v252 - v277
								if v280 != 0 {
									v250 = v250 + v277
									v252 = v280
									v256 = v276
									continue
								} else {
									break
								}
								break
							}
							v287 = v276
						}
					}
					v296 = v287
				}
				v297 = base.I32_wrap_i64(v296)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v297
				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v46<<(uint(int32(3))%32) - v297
				v304 = F_psprintf(m, int32(_a_F_gtsvectorout_2), v12)
				mBase = m.M
				v305 = m.ExcPending
				if v305 != 0 {
					return int32(0)
				} else {
					v315 = v304
					v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v316 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v319 = m.ExcPending
						if v319 != 0 {
							return int32(0)
						} else {
							m.G0 = v12 + int32(32)
							return v315
						}
					} else {
						m.G0 = v12 + int32(32)
						return v315
					}
				}
			}
		}
	}
}
