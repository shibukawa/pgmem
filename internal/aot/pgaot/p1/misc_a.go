package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AcquireDeletionLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v3 - int32(1259) {
	case 0:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_LockRelationOid(m, v6, int32(8))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	default:
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_LockDatabaseObject(m, v3, v15, int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	case 2:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_LockSharedObject(m, int32(1261), v11, int32(8))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
func F_AlignedAllocRealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v11 = l0 - int32(8)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v18 = v11 - base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(34))%64)))&int32(1073741822)
	v19 = F_GetMemoryChunkSpace(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = F_GetMemoryChunkContext(m, v18)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v29 = base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(5))%64))) & int32(1073741823)
			v30 = F_MemoryContextAllocAligned(m, v23, l1, v29, l2)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					v34 = F_MemoryContextAllocationFailure(m, v23, l1, l2)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						return v34
					}
				} else {
					v39 = v19 - v29 - int32(8)
					if base.Ui32(l1) < base.Ui32(v39) {
						v41 = l1
					} else {
						v41 = v39
					}
					if v41 != 0 {
						v42 = F__emscripten_memcpy_bulkmem(m, v30, l0, v41)
						mBase = m.M
						v43 = v42
					} else {
						v43 = v30
					}
					F_pfree(m, v18)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						return v43
					}
				}
			}
		}
	}
}
func F_AllocSetGetChunkContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	v5 = l0 - int32(8)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	if v6&int64(16) != int64(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(32))))
		return v13
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v5-base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(34))%64)))&int32(1073741822))))
		return v21
	}
}
func F_AllocSetRealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = l0 - int32(8)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	if v18&int64(16) != int64(0) {
		v24 = l0 - int32(32)
		if v24 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v171 = m.ExcPending
			if v171 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
				F_errmsg_internal(m, int32(238564), v14)
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491560), int32(1198), int32(487180))
					mBase = m.M
					v180 = m.ExcPending
					if v180 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			if v27 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v171 = m.ExcPending
				if v171 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
					F_errmsg_internal(m, int32(238564), v14)
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491560), int32(1198), int32(487180))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				if v30 != int32(474) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
						F_errmsg_internal(m, int32(238564), v14)
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(491560), int32(1198), int32(487180))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(20))))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(16))))
					if v35 != v38 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
							F_errmsg_internal(m, int32(238564), v14)
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(491560), int32(1198), int32(487180))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						if base.Ui32(int32(1073741824)) <= base.Ui32(l1) {
							if l1 < int32(0) {
								F_MemoryContextSizeFailure(m, l1)
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if l2&int32(1) == int32(0) {
									F_MemoryContextSizeFailure(m, l1)
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v53 = (l1+int32(7))&int32(-8) + int32(32)
									v54 = F_emscripten_builtin_realloc(m, v24, v53)
									mBase = m.M
									if v54 == int32(0) {
										v57 = F_MemoryContextAllocationFailure(m, v27, l1, l2)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v156 = v57
											m.G0 = v14 + int32(16)
											return v156
										}
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v61 + (v53 + v24 - v35)
										v66 = v54 + v53
										*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v66
										*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v66
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
										if v69 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v54
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v54
										}
										v73 = v54 + int32(32)
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
										if v74 == int32(0) {
											v156 = v73
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v54
											v156 = v73
										}
										m.G0 = v14 + int32(16)
										return v156
									}
								}
							}
						} else {
							v53 = (l1+int32(7))&int32(-8) + int32(32)
							v54 = F_emscripten_builtin_realloc(m, v24, v53)
							mBase = m.M
							if v54 == int32(0) {
								v57 = F_MemoryContextAllocationFailure(m, v27, l1, l2)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v156 = v57
									m.G0 = v14 + int32(16)
									return v156
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v61 + (v53 + v24 - v35)
								v66 = v54 + v53
								*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v66
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
								if v69 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v54
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v54
								}
								v73 = v54 + int32(32)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
								if v74 == int32(0) {
									v156 = v73
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v54
									v156 = v73
								}
								m.G0 = v14 + int32(16)
								return v156
							}
						}
					}
				}
			}
		}
	} else {
		v82 = int32(8) << (uint(base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(5))%64)))) % 32)
		if base.Ui32(l1) <= base.Ui32(v82) {
			v156 = l0
			m.G0 = v14 + int32(16)
			return v156
		} else {
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v17-base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(34))%64)))&int32(1073741822))))
			v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+104))
			if base.Ui32(v91) < base.Ui32(l1) {
				v93 = F_AllocSetAllocLarge(m, v90, l1, l2)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v142 = v93
					if v142 == int32(0) {
						v148 = F_MemoryContextAllocationFailure(m, v90, l1, l2)
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int32(0)
						} else {
							v156 = v148
							m.G0 = v14 + int32(16)
							return v156
						}
					} else {
						if v82 != 0 {
							v150 = F__emscripten_memcpy_bulkmem(m, v142, l0, v82)
							mBase = m.M
						} else {
						}
						F_AllocSetFree(m, l0)
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int32(0)
						} else {
							v156 = v142
							m.G0 = v14 + int32(16)
							return v156
						}
					}
				}
			} else {
				if base.Ui32(int32(8)) < base.Ui32(l1) {
					v103 = int32(29) - base.I32_clz(l1-int32(1))
				} else {
					v103 = int32(0)
				}
				v108 = v90 + v103<<(uint(int32(2))%32) + int32(48)
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
				if v109 != 0 {
					v111 = v109 + int32(8)
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					*(*int32)(unsafe.Add(mBase, uint32(v108))) = v112
					v142 = v111
					if v142 == int32(0) {
						v148 = F_MemoryContextAllocationFailure(m, v90, l1, l2)
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int32(0)
						} else {
							v156 = v148
							m.G0 = v14 + int32(16)
							return v156
						}
					} else {
						if v82 != 0 {
							v150 = F__emscripten_memcpy_bulkmem(m, v142, l0, v82)
							mBase = m.M
						} else {
						}
						F_AllocSetFree(m, l0)
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int32(0)
						} else {
							v156 = v142
							m.G0 = v14 + int32(16)
							return v156
						}
					}
				} else {
					v114 = int32(8)
					v115 = v114 << (uint(v103) % 32)
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v90)+44))
					v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
					if base.Ui32(v119-v120) < base.Ui32(v115+v114) {
						v123 = F_AllocSetAllocFromNewBlock(m, v90, l1, l2, v103)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							v142 = v123
							if v142 == int32(0) {
								v148 = F_MemoryContextAllocationFailure(m, v90, l1, l2)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									v156 = v148
									m.G0 = v14 + int32(16)
									return v156
								}
							} else {
								if v82 != 0 {
									v150 = F__emscripten_memcpy_bulkmem(m, v142, l0, v82)
									mBase = m.M
								} else {
								}
								F_AllocSetFree(m, l0)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									v156 = v142
									m.G0 = v14 + int32(16)
									return v156
								}
							}
						}
					} else {
						v126 = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v118)+12)) = v120 + v115 + v126
						*(*int64)(unsafe.Add(mBase, uint32(v120))) = base.I64_extend_i32_u(v103<<(uint(int32(5))%32)) | base.I64_extend_i32_u(v120-v118)<<(uint(int64(34))%64) | int64(3)
						v142 = v120 + v126
						if v142 == int32(0) {
							v148 = F_MemoryContextAllocationFailure(m, v90, l1, l2)
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int32(0)
							} else {
								v156 = v148
								m.G0 = v14 + int32(16)
								return v156
							}
						} else {
							if v82 != 0 {
								v150 = F__emscripten_memcpy_bulkmem(m, v142, l0, v82)
								mBase = m.M
							} else {
							}
							F_AllocSetFree(m, l0)
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return int32(0)
							} else {
								v156 = v142
								m.G0 = v14 + int32(16)
								return v156
							}
						}
					}
				}
			}
		}
	}
}
func F_AppendSeconds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v9 = l1 >> (uint(int32(31)) % 32)
	v11 = l1 ^ v9 - v9
	if l3 != 0 {
		v12 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v11) {
			v26 = F_pg_ultoa_n(m, v11, l0)
			mBase = m.M
			if v12 <= v26 {
				v37 = l0 + v26
			} else {
				v29 = l0 + v12
				v31 = F_memmove(m, v29-v26, l0, v26)
				mBase = m.M
				v34 = F___memset(m, l0, int32(48), v12-v26)
				mBase = m.M
				v37 = v29
			}
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v22)
			v37 = l0 + int32(2)
		}
		v40 = v37
	} else {
		v38 = F_pg_ultoa_n(m, v11, l0)
		mBase = m.M
		v40 = v38 + l0
	}
	if l2 == int32(0) {
		return v40
	} else {
		v44 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v44)
		v47 = l2 >> (uint(int32(31)) % 32)
		v49 = l2 ^ v47 - v47
		v51 = base.I32_div_s(v49, int32(10))
		v54 = v51*int32(-10) + v49
		if v54 != 0 {
			v56 = v54 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)) = uint8(v56)
			v62 = v40 + int32(7)
		} else {
			v62 = v40 + int32(6)
		}
		v64 = base.I32_div_s(v49, int32(100))
		v67 = v64*int32(-10) + v51
		v68 = v54 | v67
		if v68 == int32(0) {
			v76 = v40 + int32(5)
		} else {
			v74 = v67 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)) = uint8(v74)
			v76 = v62
		}
		v78 = base.I32_div_s(v49, int32(1000))
		v81 = v78*int32(-10) + v64
		v82 = v68 | v81
		if v82 == int32(0) {
			v90 = v40 + int32(4)
		} else {
			v88 = v81 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)) = uint8(v88)
			v90 = v76
		}
		v92 = base.I32_div_s(v49, int32(10000))
		v95 = v92*int32(-10) + v78
		v96 = v82 | v95
		if v96 == int32(0) {
			v104 = v40 + int32(3)
		} else {
			v102 = v95 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+3)) = uint8(v102)
			v104 = v90
		}
		v106 = base.I32_div_s(v49, int32(100000))
		v109 = v106*int32(-10) + v92
		v110 = v96 | v109
		if v110 == int32(0) {
			v118 = v40 + int32(2)
		} else {
			v116 = v109 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+2)) = uint8(v116)
			v118 = v104
		}
		v120 = base.I32_div_s(v49, int32(1000000))
		v123 = v120*int32(-10) + v106
		if v110|v123 == int32(0) {
			v132 = v40 + int32(1)
		} else {
			v130 = v123 + int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)) = uint8(v130)
			v132 = v118
		}
		if base.Ui32(int32(19)) <= base.Ui32(v106+int32(9)) {
			v138 = v40 + int32(1)
			v139 = F_pg_ultoa_n(m, v49, v138)
			mBase = m.M
			v141 = v139 + v138
		} else {
			v141 = v132
		}
		return v141
	}
}
func F_AtAbort_Portals(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _consts[891]))
	F_hash_seq_init(m, v6+int32(12), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v16
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v6 + int32(32)
	return
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	if v22 != int32(3) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1230])))
	if v26 != int32(1) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = int32(5)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	m.T0[v31].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = int32(0)
	goto L9
L14:
	;
	v79 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L32
	}
L15:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+85)))
	if v42 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	if v43 == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v62 != 0 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = int32(0)
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = int32(5)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v53 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	m.T0[v48].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	m.T0[v53].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L18
L26:
	;
	F_ReleaseCachedPlan(m, v62, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	if v70 == int32(3) {
		goto L14
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(0)
	goto L28
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	F_MemoryContextDeleteChildren(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L14
L32:
	;
	if v79 != 0 {
		v18 = v79
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L8
}
func F_AtProcExit_Buffers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[740]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(228784)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(493364)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v21 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v20 | v21
	if v20&v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v6 + int32(32)
	return
L4:
	;
	goto L7
L5:
	;
	v40 = v20
	goto L6
L6:
	;
	v45 = int32(4104476)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(8))+8))
	if v48 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_perform_spin_delay(m, v6+int32(8))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v40 = v33
	goto L6
L9:
	;
	return
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v34 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v33 | v34
	if v33&v34 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	if v40&int32(536870912) != 0 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v63
	goto L13
L15:
	;
	if int32(999) < v46 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v46 < int32(11) {
		goto L13
	} else {
		goto L22
	}
L18:
	;
	v53 = int32(900)
	if v53 <= v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v56 = v53
	goto L21
L20:
	;
	v56 = v46
	goto L21
L21:
	;
	v63 = v56 + int32(100)
	goto L14
L22:
	;
	v63 = v46 - int32(1)
	goto L14
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v71 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v69 == v71 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v74 = v40
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v74 & int32(-4194305)
	*(*int32)(unsafe.Add(mBase, _consts[740])) = int32(0)
	goto L3
L26:
	;
	v73 = v40 & int32(-536870913)
	goto L28
L27:
	;
	v73 = v40
	goto L28
L28:
	;
	v74 = v73
	goto L25
}
func F_AtProcExit_Twophase(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_AtAbort_Twophase(m)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_AuxiliaryPidGetProc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v2 = int32(0)
	if l0 == v2 {
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v12 = v2
	goto L5
L4:
	;
	return v33
L5:
	;
	v17 = v10 + v12*int32(640)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	if v18 == l0 {
		v33 = v17
		goto L4
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v24 = v10 + (v12|int32(1))*int32(640)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v25 == l0 {
		v33 = v24
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v28 = v12 + int32(2)
	if v28 != int32(38) {
		v12 = v28
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
}
func F_a_cas(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2 != 0 {
		v4 = v2
	} else {
		v4 = int32(1073741823)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
	return v2
}
func F_accumArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v3 = l2
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == int32(0) {
		v17 = F_AllocSetContextCreateInternal(m, l4, int32(97881), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = F_MemoryContextAlloc(m, v17, int32(32))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)) = uint8(v24)
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(64)
				v30 = F_MemoryContextAlloc(m, v17, int32(256))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v30
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					v34 = F_MemoryContextAlloc(m, v17, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v34
						F_get_typlenbyvalalign(m, l3, v22+int32(24), v22+int32(26), v22+int32(27))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = v22
							v50 = int32(4489152)
							v51 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v53
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							if v55 <= v56 {
								*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v55 << (uint(int32(1)) % 32)
								v62 = v55 << (uint(int32(3)) % 32)
								if base.Ui32(int32(1073741824)) <= base.Ui32(v62) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(261))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1073741823)
											F_errmsg(m, int32(662538), v9)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(492766), int32(5384), int32(97881))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
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
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v66 = F_repalloc(m, v65, v62)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v66
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
										v71 = F_repalloc(m, v69, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v71
											if v3 != 0 {
												v84 = l1
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
												m.G0 = v9 + int32(16)
												return v48
											} else {
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
												if v75 != 0 {
													v84 = l1
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
													v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
													*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
													*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
													m.G0 = v9 + int32(16)
													return v48
												} else {
													v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
													if v76 == int32(-1) {
														v79 = F_pg_detoast_datum_copy(m, l1)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v84 = v79
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
															v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
															*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
															m.G0 = v9 + int32(16)
															return v48
														}
													} else {
														v82 = F_datumCopy(m, l1, int32(0), v76)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															v84 = v82
															v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
															v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
															*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
															*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
															m.G0 = v9 + int32(16)
															return v48
														}
													}
												}
											}
										}
									}
								}
							} else {
								if v3 != 0 {
									v84 = l1
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
									*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
									m.G0 = v9 + int32(16)
									return v48
								} else {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
									if v75 != 0 {
										v84 = l1
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
										m.G0 = v9 + int32(16)
										return v48
									} else {
										v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
										if v76 == int32(-1) {
											v79 = F_pg_detoast_datum_copy(m, l1)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v84 = v79
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
												m.G0 = v9 + int32(16)
												return v48
											}
										} else {
											v82 = F_datumCopy(m, l1, int32(0), v76)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = v82
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
												m.G0 = v9 + int32(16)
												return v48
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
	} else {
		v48 = l0
		v50 = int32(4489152)
		v51 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v53
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
		if v55 <= v56 {
			*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v55 << (uint(int32(1)) % 32)
			v62 = v55 << (uint(int32(3)) % 32)
			if base.Ui32(int32(1073741824)) <= base.Ui32(v62) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1073741823)
						F_errmsg(m, int32(662538), v9)
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492766), int32(5384), int32(97881))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
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
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v66 = F_repalloc(m, v65, v62)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v66
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
					v71 = F_repalloc(m, v69, v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v71
						if v3 != 0 {
							v84 = l1
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
							m.G0 = v9 + int32(16)
							return v48
						} else {
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
							if v75 != 0 {
								v84 = l1
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
								m.G0 = v9 + int32(16)
								return v48
							} else {
								v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
								if v76 == int32(-1) {
									v79 = F_pg_detoast_datum_copy(m, l1)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v84 = v79
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
										m.G0 = v9 + int32(16)
										return v48
									}
								} else {
									v82 = F_datumCopy(m, l1, int32(0), v76)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										v84 = v82
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
										m.G0 = v9 + int32(16)
										return v48
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v3 != 0 {
				v84 = l1
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
				m.G0 = v9 + int32(16)
				return v48
			} else {
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+26)))
				if v75 != 0 {
					v84 = l1
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
					m.G0 = v9 + int32(16)
					return v48
				} else {
					v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
					if v76 == int32(-1) {
						v79 = F_pg_detoast_datum_copy(m, l1)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v84 = v79
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
							m.G0 = v9 + int32(16)
							return v48
						}
					} else {
						v82 = F_datumCopy(m, l1, int32(0), v76)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v84 = v82
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v87<<(uint(int32(2))%32)))) = v84
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*uint8)(unsafe.Add(mBase, uint32(v92+v93))) = uint8(v3)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v96 + int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
							m.G0 = v9 + int32(16)
							return v48
						}
					}
				}
			}
		}
	}
}
func F_aclcopy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) <= v10 {
		v16 = v10<<(uint(int32(4))%32) + int32(24)
		v17 = F_palloc0(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v16 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v10
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v33 != 0 {
				v41 = v33
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v41 = (v34<<(uint(int32(3))%32) + int32(23)) & int32(-8)
			}
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v45 = v43 << (uint(int32(4)) % 32)
			if v45 != 0 {
				v46 = F__emscripten_memcpy_bulkmem(m, v17+int32(24), v41+l0, v45)
				mBase = m.M
			} else {
			}
			m.G0 = v8 + int32(16)
			return v17
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
			F_errmsg_internal(m, int32(481885), v8)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495726), int32(433), int32(307550))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
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
func F_aclinsert(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(442651), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495726), int32(1597), int32(81356))
				v19 = m.ExcPending
				if v19 != 0 {
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
func F_aclmembers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	if l0 != 0 {
		v10 = l0 + int32(16)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 != 0 {
			F_check_acl(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v24 = F_palloc(m, v21<<(uint(int32(3))%32))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v26 == int32(0) {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v36 = (v29<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					} else {
						v36 = v26
					}
					v37 = int32(0)
					v39 = l0 + int32(16)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					if v37 < v40 {
						v45 = int32(0)
						v47 = v37
						for {
							v55 = l0 + v36 + v45<<(uint(int32(4))%32)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							if v56 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v24+v47<<(uint(int32(2))%32)))) = v56
								v63 = v47 + int32(1)
							} else {
								v63 = v47
							}
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
							if v64 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v24+v63<<(uint(int32(2))%32)))) = v64
								v71 = v63 + int32(1)
							} else {
								v71 = v63
							}
							v73 = v45 + int32(1)
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
							if v73 < v74 {
								v45 = v73
								v47 = v71
								continue
							} else {
								break
							}
							break
						}
						F_pg_qsort(m, v24, v71, int32(4), int32(471))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
							if base.Ui32(int32(2)) <= base.Ui32(v71) {
								v85 = int32(1)
								v88 = int32(0)
								for {
									v93 = int32(2)
									v95 = v24 + v85<<(uint(v93)%32)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v24+v88<<(uint(v93)%32))))
									if base.B2i32(base.Ui32(v100) < base.Ui32(v99))-base.B2i32(base.Ui32(v99) < base.Ui32(v100)) == int32(0) {
										v114 = v88
									} else {
										v107 = v88 + int32(1)
										if v85 == v107 {
											v114 = v85
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
											*(*int32)(unsafe.Add(mBase, uint32(v24+v107<<(uint(int32(2))%32)))) = v112
											v114 = v107
										}
									}
									v117 = v85 + int32(1)
									if v117 != v71 {
										v85 = v117
										v88 = v114
										continue
									} else {
										break
									}
									break
								}
								v129 = v114 + int32(1)
							} else {
								v129 = v71
							}
							return v129
						}
					} else {
						F_pg_qsort(m, v24, int32(0), int32(4), int32(471))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
							return int32(0)
						}
					}
				}
			}
		} else {
			v13 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13
			return v13
		}
	} else {
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13
		return v13
	}
}
func F_acquire_sample_rows(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v196 int64
	_ = v196
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v209 int64
	_ = v209
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v216 int64
	_ = v216
	var v221 int64
	_ = v221
	var v226 int64
	_ = v226
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v301 float64
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 float64
	_ = v340
	var v342 float64
	_ = v342
	var v351 float64
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
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
	var v383 float64
	_ = v383
	var v385 float64
	_ = v385
	var v401 int32
	_ = v401
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 float64
	_ = v423
	var v424 float64
	_ = v424
	var v429 int32
	_ = v429
	var v463 int64
	_ = v463
	var v464 int64
	_ = v464
	var v465 int64
	_ = v465
	var v486 float64
	_ = v486
	var v490 float64
	_ = v490
	var v492 float64
	_ = v492
	var v502 float64
	_ = v502
	var v503 float64
	_ = v503
	var v506 float64
	_ = v506
	var v527 float64
	_ = v527
	var v528 float64
	_ = v528
	var v530 float64
	_ = v530
	var v533 float64
	_ = v533
	var v535 float64
	_ = v535
	var v536 float64
	_ = v536
	var v537 float64
	_ = v537
	var v539 float64
	_ = v539
	var v540 float64
	_ = v540
	var v542 int32
	_ = v542
	var v543 float64
	_ = v543
	var v550 float64
	_ = v550
	var v577 int64
	_ = v577
	var v578 int64
	_ = v578
	var v579 int64
	_ = v579
	var v600 float64
	_ = v600
	var v605 float64
	_ = v605
	var v606 float64
	_ = v606
	var v607 float64
	_ = v607
	var v611 float64
	_ = v611
	var v613 float64
	_ = v613
	var v615 float64
	_ = v615
	var v618 float64
	_ = v618
	var v621 float64
	_ = v621
	var v627 float64
	_ = v627
	var v628 int32
	_ = v628
	var v629 float64
	_ = v629
	var v632 float64
	_ = v632
	var v639 float64
	_ = v639
	var v640 float64
	_ = v640
	var v642 float64
	_ = v642
	var v665 float64
	_ = v665
	var v666 float64
	_ = v666
	var v669 float64
	_ = v669
	var v680 float64
	_ = v680
	var v735 int64
	_ = v735
	var v736 int64
	_ = v736
	var v737 int64
	_ = v737
	var v758 float64
	_ = v758
	var v761 float64
	_ = v761
	var v763 float64
	_ = v763
	var v764 float64
	_ = v764
	var v767 float64
	_ = v767
	var v778 float64
	_ = v778
	var v814 float64
	_ = v814
	var v841 float64
	_ = v841
	var v901 int64
	_ = v901
	var v902 int64
	_ = v902
	var v903 int64
	_ = v903
	var v924 float64
	_ = v924
	var v927 float64
	_ = v927
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v984 float64
	_ = v984
	var v1002 int32
	_ = v1002
	var v1010 float64
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1026 float64
	_ = v1026
	var v1028 float64
	_ = v1028
	var v1044 int32
	_ = v1044
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1078 int32
	_ = v1078
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 float64
	_ = v1142
	var v1144 float64
	_ = v1144
	var v1145 float64
	_ = v1145
	var v1147 float64
	_ = v1147
	var v1149 float64
	_ = v1149
	var v1152 float64
	_ = v1152
	var v1158 float64
	_ = v1158
	var v1161 float64
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 float64
	_ = v1168
	var v1169 float64
	_ = v1169
	var v1171 float64
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	v20 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(128)
	m.G0 = v34
	v36 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+120)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v34)+112)) = v36
	v41 = F_RelationGetNumberOfBlocksInFork(m, l0, v20)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v45 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = v34 + int32(80)
	v51 = int32(4572664)
	v52 = int32(4572656)
	v53 = *(*int64)(unsafe.Add(mBase, _consts[49]))
	v55 = *(*int64)(unsafe.Add(mBase, _consts[50]))
	v56 = v53 ^ v55
	*(*int64)(unsafe.Add(mBase, _consts[50])) = base.I64_rotl(v56, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[49])) = v56<<(uint(int64(16))%64) ^ base.I64_rotl(v53, int64(24)) ^ v56
	goto L4
L4:
	;
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v48)+8)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v41
	v82 = v34 + int32(96)
	v83 = base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v53*int64(5), int64(7))*int64(9)) >> (uint(int64(32)) % 64))))
	v87 = v83 + int64(4354685564936845354)
	v88 = int64(30)
	v91 = int64(-4658895280553007687)
	v92 = (int64(base.Ui64(v87)>>(uint(v88)%64)) ^ v87) * v91
	v93 = int64(27)
	v96 = int64(-7723592293110705685)
	v97 = (int64(base.Ui64(v92)>>(uint(v93)%64)) ^ v92) * v96
	v98 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = int64(base.Ui64(v97)>>(uint(v98)%64)) ^ v97
	v103 = v83 - int64(7046029254386353131)
	v108 = (int64(base.Ui64(v103)>>(uint(v88)%64)) ^ v103) * v91
	v113 = (int64(base.Ui64(v108)>>(uint(v93)%64)) ^ v108) * v96
	*(*int64)(unsafe.Add(mBase, uint32(v82))) = int64(base.Ui64(v113)>>(uint(v98)%64)) ^ v113
	if v103|v87 == v77 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if base.Ui32(v126) < base.Ui32(v127) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, uint32(v82))) = int64(6364136223846793005)
	goto L8
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	v129 = v126
	goto L11
L10:
	;
	v129 = v127
	goto L11
L11:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v133 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v167 = v34 + int32(64)
	v170 = int32(4572664)
	v171 = int32(4572656)
	v172 = *(*int64)(unsafe.Add(mBase, _consts[49]))
	v174 = *(*int64)(unsafe.Add(mBase, _consts[50]))
	v175 = v172 ^ v174
	*(*int64)(unsafe.Add(mBase, _consts[50])) = base.I64_rotl(v175, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[49])) = v175<<(uint(int64(16))%64) ^ base.I64_rotl(v172, int64(24)) ^ v175
	goto L16
L13:
	;
	goto L12
L14:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v137 != int32(1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v140 = int32(4483812)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v143 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v142 + v143
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v146 + v143
	*(*int64)(unsafe.Add(mBase, uint32(v133+int32(8))+232)) = base.I64_extend_i32_u(v129)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v154 + v143
	v160 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v160 - v143
	goto L13
L16:
	;
	v196 = base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v172*int64(5), int64(7))*int64(9)) >> (uint(int64(32)) % 64))))
	v200 = v196 + int64(4354685564936845354)
	v201 = int64(30)
	v204 = int64(-4658895280553007687)
	v205 = (int64(base.Ui64(v200)>>(uint(v201)%64)) ^ v200) * v204
	v206 = int64(27)
	v209 = int64(-7723592293110705685)
	v210 = (int64(base.Ui64(v205)>>(uint(v206)%64)) ^ v205) * v209
	v211 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = int64(base.Ui64(v210)>>(uint(v211)%64)) ^ v210
	v216 = v196 - int64(7046029254386353131)
	v221 = (int64(base.Ui64(v216)>>(uint(v201)%64)) ^ v216) * v204
	v226 = (int64(base.Ui64(v221)>>(uint(v206)%64)) ^ v221) * v209
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = int64(base.Ui64(v226)>>(uint(v211)%64)) ^ v226
	if v216|v200 == int64(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L21
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = int64(6364136223846793005)
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v167)))
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v167)+8))
	v273 = v271 ^ v272
	*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = base.I64_rotl(v273, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = v273<<(uint(int64(16))%64) ^ base.I64_rotl(v271, int64(24)) ^ v273
	v294 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v271*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L23
L22:
	;
	v297 = F_log(m, v294)
	mBase = m.M
	v301 = F_exp(m, base.F64_div(base.F64_neg(v297), base.F64_convert_i32_s(l3)))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(v34+int32(56)))) = v301
	v303 = int32(0)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v310 = m.T0[v309].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v303, v303, v303, v303, int32(32))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	if base.F64_eq(v294, float64(0)) != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v313 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _consts[352]))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v323 = F_read_stream_begin_relation(m, int32(9), v317, v318, int32(502), v34+int32(80), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+188))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+132))
	v328 = m.T0[v327].(func(*base.Module, int32, int32) int32)(m, v310, v323)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v328 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v334 = l0
	v335 = l1
	v336 = l2
	v337 = l3
	v338 = l4
	v339 = l5
	v340 = float64(-1)
	v342 = float64(0)
	v351 = base.F64_convert_i32_s(l3)
	v353 = v34
	v354 = v34 - int32(-64)
	v356 = v310
	v357 = v313
	v358 = v20
	v360 = v41
	v361 = v45
	v362 = v323
	v364 = v20
	goto L32
L30:
	;
	v1093 = l0
	v1094 = l1
	v1095 = l2
	v1096 = l3
	v1097 = l4
	v1098 = l5
	v1112 = v34
	v1115 = v310
	v1116 = v313
	v1117 = v20
	v1119 = v41
	v1121 = v323
	goto L31
L31:
	;
	F_read_stream_end(m, v1121)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L107
	}
L32:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	v1093 = v334
	v1094 = v335
	v1095 = v336
	v1096 = v337
	v1097 = v338
	v1098 = v339
	v1112 = v353
	v1115 = v356
	v1116 = v357
	v1117 = v1044
	v1119 = v360
	v1121 = v362
	goto L31
L34:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+188))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+136))
	v375 = m.T0[v374].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v356, v361, v353+int32(120), v353+int32(112), v357)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v375 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v383 = v340
	v385 = v342
	v401 = v358
	goto L39
L37:
	;
	v1026 = v340
	v1028 = v342
	v1044 = v358
	goto L38
L38:
	;
	v1053 = v364 + int32(1)
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v1057 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L39:
	;
	if v401 < v337 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v1026 = v984
	v1028 = v1010
	v1044 = v1002
	goto L38
L41:
	;
	v1010 = base.F64_add(v385, float64(1))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+188))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+136))
	v1018 = m.T0[v1017].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v356, v361, v353+int32(120), v353+int32(112), v357)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L99
	}
L42:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v357)+8))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+44))
	v414 = m.T0[v413].(func(*base.Module, int32) int32)(m, v357)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if base.F64_lt(v383, float64(0)) != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336+v401<<(uint(int32(2))%32)))) = v414
	v984 = v383
	v1002 = v401 + int32(1)
	goto L41
L46:
	;
	v422 = v353 + int32(56)
	v423 = float64(0)
	v424 = base.F64_convert_i32_s(v337)
	if base.F64_ge(base.F64_mul(v424, float64(22)), v385) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v841 = v383
	goto L48
L48:
	;
	if base.F64_le(v841, float64(0)) != 0 {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	v841 = v814
	goto L48
L50:
	;
	v429 = v353 + int32(64)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v535 = float64(1)
	v536 = base.F64_add(v385, v535)
	v537 = base.F64_sub(v385, v424)
	v539 = base.F64_add(v537, v535)
	v540 = base.F64_div(v536, v539)
	v542 = v353 + int32(64)
	v543 = *(*float64)(unsafe.Add(mBase, uint32(v422)))
	v550 = v543
	goto L61
L53:
	;
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v429)))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v429)+8))
	v465 = v463 ^ v464
	*(*int64)(unsafe.Add(mBase, uint32(v429)+8)) = base.I64_rotl(v465, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v429))) = v465<<(uint(int64(16))%64) ^ base.I64_rotl(v463, int64(24)) ^ v465
	v486 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v463*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L55
L54:
	;
	v490 = base.F64_add(v385, float64(1))
	v492 = base.F64_div(base.F64_sub(v490, v424), v490)
	if base.F64_gt(v492, v486) == int32(0) {
		v814 = v423
		goto L49
	} else {
		goto L57
	}
L55:
	;
	if base.F64_eq(v486, float64(0)) != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v502 = v490
	v503 = v492
	v506 = v423
	goto L58
L58:
	;
	v527 = float64(1)
	v528 = base.F64_add(v506, v527)
	v530 = base.F64_add(v502, v527)
	v533 = base.F64_mul(v503, base.F64_div(base.F64_sub(v530, v424), v530))
	if base.F64_gt(v533, v486) != 0 {
		v502 = v530
		v503 = v533
		v506 = v528
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v814 = v528
	goto L49
L60:
	;
	goto L59
L61:
	;
	v577 = *(*int64)(unsafe.Add(mBase, uint32(v542)))
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v542)+8))
	v579 = v577 ^ v578
	*(*int64)(unsafe.Add(mBase, uint32(v542)+8)) = base.I64_rotl(v579, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v542))) = v579<<(uint(int64(16))%64) ^ base.I64_rotl(v577, int64(24)) ^ v579
	v600 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v577*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L63
L62:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v422))) = v778
	v814 = v606
	goto L49
L63:
	;
	if base.F64_eq(v600, float64(0)) != 0 {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v605 = base.F64_mul(v385, base.F64_add(v550, float64(-1)))
	v606 = base.F64_floor(v605)
	v607 = base.F64_add(v539, v606)
	v611 = base.F64_add(v385, v605)
	v613 = F_log(m, base.F64_div(base.F64_mul(v607, base.F64_mul(v540, base.F64_mul(v540, v600))), v611))
	mBase = m.M
	v615 = F_exp(m, base.F64_div(v613, v424))
	mBase = m.M
	v618 = base.F64_div(base.F64_mul(v539, base.F64_div(v611, v607)), v385)
	if base.F64_le(v615, v618) != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L62
L66:
	;
	v778 = base.F64_div(v618, v615)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v621 = base.F64_add(v385, v606)
	v627 = base.F64_div(base.F64_mul(base.F64_add(v621, float64(1)), base.F64_div(base.F64_mul(v536, v600), v539)), v611)
	v628 = base.F64_gt(v606, v424)
	if v628 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v629 = v607
	goto L71
L70:
	;
	v629 = v536
	goto L71
L71:
	;
	if base.F64_le(v629, v621) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v628 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v680 = v627
	goto L74
L74:
	;
	goto L81
L75:
	;
	v632 = v385
	goto L77
L76:
	;
	v632 = base.F64_add(v537, v606)
	goto L77
L77:
	;
	v639 = v621
	v640 = v632
	v642 = v627
	goto L78
L78:
	;
	v665 = base.F64_mul(v642, base.F64_div(v639, v640))
	v666 = float64(-1)
	v669 = base.F64_add(v639, v666)
	if base.F64_ge(v669, v629) != 0 {
		v639 = v669
		v640 = base.F64_add(v640, v666)
		v642 = v665
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v680 = v665
	goto L74
L80:
	;
	goto L79
L81:
	;
	v735 = *(*int64)(unsafe.Add(mBase, uint32(v542)))
	v736 = *(*int64)(unsafe.Add(mBase, uint32(v542)+8))
	v737 = v735 ^ v736
	*(*int64)(unsafe.Add(mBase, uint32(v542)+8)) = base.I64_rotl(v737, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v542))) = v737<<(uint(int64(16))%64) ^ base.I64_rotl(v735, int64(24)) ^ v737
	v758 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v735*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L83
L82:
	;
	v761 = F_log(m, v680)
	mBase = m.M
	v763 = F_exp(m, base.F64_div(v761, v424))
	mBase = m.M
	v764 = F_log(m, v758)
	mBase = m.M
	v767 = F_exp(m, base.F64_div(base.F64_neg(v764), v424))
	mBase = m.M
	if base.F64_le(v763, base.F64_div(v611, v385)) == int32(0) {
		v550 = v767
		goto L61
	} else {
		goto L85
	}
L83:
	;
	if base.F64_eq(v758, float64(0)) != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v778 = v767
	goto L65
L86:
	;
	goto L90
L87:
	;
	goto L88
L88:
	;
	v984 = base.F64_add(v841, float64(-1))
	v1002 = v401
	goto L41
L89:
	;
	v936 = v336 + v933<<(uint(int32(2))%32)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)))
	F_pfree(m, v937)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L97
	}
L90:
	;
	v901 = *(*int64)(unsafe.Add(mBase, uint32(v354)))
	v902 = *(*int64)(unsafe.Add(mBase, uint32(v354)+8))
	v903 = v901 ^ v902
	*(*int64)(unsafe.Add(mBase, uint32(v354)+8)) = base.I64_rotl(v903, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v354))) = v903<<(uint(int64(16))%64) ^ base.I64_rotl(v901, int64(24)) ^ v903
	v924 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v901*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L92
L91:
	;
	v927 = base.F64_mul(v924, v351)
	if base.F64_lt(base.F64_abs(v927), float64(2.147483648e+09)) != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	if base.F64_eq(v924, float64(0)) != 0 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v931 = base.I32_trunc_f64_s(v927)
	v933 = v931
	goto L89
L95:
	;
	goto L96
L96:
	;
	v933 = int32(-2147483648)
	goto L89
L97:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v357)+8))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)+44))
	v942 = m.T0[v941].(func(*base.Module, int32) int32)(m, v357)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v936))) = v942
	goto L88
L99:
	;
	if v1018 != 0 {
		v383 = v984
		v385 = v1010
		v401 = v1002
		goto L39
	} else {
		goto L100
	}
L100:
	;
	goto L40
L101:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+188))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+132))
	v1091 = m.T0[v1090].(func(*base.Module, int32, int32) int32)(m, v356, v362)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v1061 != int32(1) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v1064 = int32(4483812)
	v1066 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v1067 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1066 + v1067
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1057)))
	*(*int32)(unsafe.Add(mBase, uint32(v1057))) = v1070 + v1067
	*(*int64)(unsafe.Add(mBase, uint32(v1057+int32(16))+232)) = base.I64_extend_i32_u(v1053)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1057)))
	*(*int32)(unsafe.Add(mBase, uint32(v1057))) = v1078 + v1067
	v1084 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1084 - v1067
	goto L102
L105:
	;
	if v1091 != 0 {
		v340 = v1026
		v342 = v1028
		v358 = v1044
		v364 = v1053
		goto L32
	} else {
		goto L106
	}
L106:
	;
	goto L33
L107:
	;
	F_ExecDropSingleTupleTableSlot(m, v1116)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+188))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+12))
	m.T0[v1130].(func(*base.Module, int32))(m, v1115)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if v1096 == v1117 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_qsort_interruptible(m, v1095, v1096, int32(4), int32(503), int32(0))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+92))
	if v1139 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L112
L114:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1097))) = v1161
	*(*float64)(unsafe.Add(mBase, uint32(v1098))) = v1158
	v1165 = F_errstart(m, v1094, int32(0))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L118
	}
L115:
	;
	v1142 = float64(0)
	v1158 = v1142
	v1161 = v1142
	goto L114
L116:
	;
	goto L117
L117:
	;
	v1144 = *(*float64)(unsafe.Add(mBase, uint32(v1112)+112))
	v1145 = base.F64_convert_i32_u(v1139)
	v1147 = base.F64_convert_i32_u(v1119)
	v1149 = float64(0.5)
	v1152 = *(*float64)(unsafe.Add(mBase, uint32(v1112)+120))
	v1158 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v1144, v1145), v1147), v1149))
	v1161 = base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(v1152, v1145), v1147), v1149))
	goto L114
L118:
	;
	if v1165 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+48))
	v1168 = *(*float64)(unsafe.Add(mBase, uint32(v1097)))
	v1169 = *(*float64)(unsafe.Add(mBase, uint32(v1112)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v1112)+16)) = v1169
	v1171 = *(*float64)(unsafe.Add(mBase, uint32(v1112)+112))
	*(*float64)(unsafe.Add(mBase, uint32(v1112)+24)) = v1171
	*(*int32)(unsafe.Add(mBase, uint32(v1112)+32)) = v1117
	*(*float64)(unsafe.Add(mBase, uint32(v1112)+40)) = v1168
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1112)+4)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v1112)+8)) = v1119
	*(*int32)(unsafe.Add(mBase, uint32(v1112))) = v1167 + int32(4)
	F_errmsg(m, int32(113378), v1112)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	m.G0 = v1112 + int32(128)
	return v1117
L122:
	;
	F_errfinish(m, int32(496562), int32(1352), int32(113305))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L121
}
func F_addCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v7 = m.G0
	v9 = v7 - int32(1024)
	m.G0 = v9
	v12 = l1
	goto L2
L1:
	;
	v46 = v12
	v48 = v17
	v50 = v9
	goto L15
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.Ui32(v17-int32(9)) < base.Ui32(int32(5)) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	v26 = F_pg_mblen_cstr(m, v12)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if v17 == int32(32) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v17 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	return
L10:
	;
	v12 = v26 + v12
	goto L2
L11:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(211687), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(495575), int32(1076), int32(346502))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	switch v48 & int32(255) {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L17
	default:
		goto L18
	}
L16:
	;
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v63 < v62 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	goto L16
L18:
	;
	v53 = F_pg_mblen_cstr(m, v46)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	if v53 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v58 = v46 + v53
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v46 = v58
	v48 = v59
	v50 = v56 + v53
	goto L15
L21:
	;
	v55 = F__emscripten_memcpy_bulkmem(m, v50, v46, v53)
	mBase = m.M
	v56 = v55
	goto L23
L22:
	;
	v56 = v50
	goto L23
L23:
	;
	goto L20
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_setCompoundAffixFlagValue(m, l0, v82+v83*int32(12), v9, l2)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L34
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v82 = v65
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v62 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v80
	v82 = v80
	goto L24
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62 << (uint(int32(1)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v72 = F_repalloc(m, v69, v62*int32(24))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(10)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v78 = F_MemoryContextAlloc(m, v76, int32(120))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L33
	}
L32:
	;
	v80 = v72
	goto L28
L33:
	;
	v80 = v78
	goto L28
L34:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v89)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v91 + v89
	m.G0 = v9 + int32(1024)
	return
}
func F_addFamilyMember(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v30 int32
	_ = v30
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
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v55
	F_errmsg(m, int32(413483), v12)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L14
	} else {
		goto L24
	}
L2:
	;
	v89 = F_lappend(m, v14, l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L23
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = int32(0)
	if v20 < v17 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = v17
	goto L7
L6:
	;
	v23 = v20
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v30 = v3
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25+v30<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v39 != v24 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v78 = v30 + int32(1)
	if v78 != v23 {
		v30 = v78
		goto L8
	} else {
		goto L22
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v41 != v42 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v44 != v45 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v57 = F_format_type_be(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v60 = F_format_type_be(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v47 == int32(1) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	F_errmsg(m, int32(413429), v12+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(492498), int32(1442), int32(228158))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	goto L9
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v89
	m.G0 = v12 + int32(32)
	return
L24:
	;
	F_errfinish(m, int32(492498), int32(1435), int32(228158))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addFkConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+119)))
	switch v27 - int32(112) {
	case 0, 2:
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
		v51 = F_ConstraintNameIsUsed(m, int32(0), v50, l2)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return
		} else {
			if v51 != 0 {
				v53 = int32(0)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
				v58 = F_ChooseConstraintName(m, l2, v53, int32(740129), v56, v53)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					v60 = v58
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v61 == int32(0) {
						v64 = F_pstrdup(m, v60)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v64
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
							v76 = int32(0)
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
							v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
							v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
							v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
							if l7 != 0 {
								v90 = int32(0)
								v91 = int32(1)
							} else {
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
								v90 = base.B2i32(v86 != int32(112))
								v91 = int32(0)
							}
							v92 = F_CreateConstraintEntry(m, v60, v69, int32(102), v71, v72, v73, v74, l7, v75, l10, l8, l8, v76, l6, v77, l9, l11, l12, l13, l8, v78, v79, l15, l14, v80, v76, v76, v76, base.B2i32(l7 == v76), v91, v90, l17, l16)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
								if l7 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l7
									*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(2606)
									if l1 != 0 {
										F_recordDependencyOn(m, l0, v24+int32(4), int32(80))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(1259)
											v113 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
											*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v113
											v120 = int32(83)
											F_recordDependencyOn(m, l0, v24+int32(4), v120)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return
												} else {
													m.G0 = v24 + int32(16)
													return
												}
											}
										}
									} else {
										v120 = int32(105)
										F_recordDependencyOn(m, l0, v24+int32(4), v120)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												m.G0 = v24 + int32(16)
												return
											}
										}
									}
								} else {
									F_CommandCounterIncrement(m)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										m.G0 = v24 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
						v76 = int32(0)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
						v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
						v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
						v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
						if l7 != 0 {
							v90 = int32(0)
							v91 = int32(1)
						} else {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
							v90 = base.B2i32(v86 != int32(112))
							v91 = int32(0)
						}
						v92 = F_CreateConstraintEntry(m, v60, v69, int32(102), v71, v72, v73, v74, l7, v75, l10, l8, l8, v76, l6, v77, l9, l11, l12, l13, l8, v78, v79, l15, l14, v80, v76, v76, v76, base.B2i32(l7 == v76), v91, v90, l17, l16)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
							if l7 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(2606)
								if l1 != 0 {
									F_recordDependencyOn(m, l0, v24+int32(4), int32(80))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(1259)
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
										*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v113
										v120 = int32(83)
										F_recordDependencyOn(m, l0, v24+int32(4), v120)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												m.G0 = v24 + int32(16)
												return
											}
										}
									}
								} else {
									v120 = int32(105)
									F_recordDependencyOn(m, l0, v24+int32(4), v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											m.G0 = v24 + int32(16)
											return
										}
									}
								}
							} else {
								F_CommandCounterIncrement(m)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return
								} else {
									m.G0 = v24 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				v60 = l2
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				if v61 == int32(0) {
					v64 = F_pstrdup(m, v60)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v64
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
						v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
						v76 = int32(0)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
						v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
						v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
						v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
						if l7 != 0 {
							v90 = int32(0)
							v91 = int32(1)
						} else {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
							v90 = base.B2i32(v86 != int32(112))
							v91 = int32(0)
						}
						v92 = F_CreateConstraintEntry(m, v60, v69, int32(102), v71, v72, v73, v74, l7, v75, l10, l8, l8, v76, l6, v77, l9, l11, l12, l13, l8, v78, v79, l15, l14, v80, v76, v76, v76, base.B2i32(l7 == v76), v91, v90, l17, l16)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
							if l7 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l7
								*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(2606)
								if l1 != 0 {
									F_recordDependencyOn(m, l0, v24+int32(4), int32(80))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(1259)
										v113 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
										*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v113
										v120 = int32(83)
										F_recordDependencyOn(m, l0, v24+int32(4), v120)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												m.G0 = v24 + int32(16)
												return
											}
										}
									}
								} else {
									v120 = int32(105)
									F_recordDependencyOn(m, l0, v24+int32(4), v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											m.G0 = v24 + int32(16)
											return
										}
									}
								}
							} else {
								F_CommandCounterIncrement(m)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return
								} else {
									m.G0 = v24 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+12)))
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+13)))
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+14)))
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
					v76 = int32(0)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
					v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+87)))
					v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+88)))
					v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3)+86)))
					if l7 != 0 {
						v90 = int32(0)
						v91 = int32(1)
					} else {
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
						v90 = base.B2i32(v86 != int32(112))
						v91 = int32(0)
					}
					v92 = F_CreateConstraintEntry(m, v60, v69, int32(102), v71, v72, v73, v74, l7, v75, l10, l8, l8, v76, l6, v77, l9, l11, l12, l13, l8, v78, v79, l15, l14, v80, v76, v76, v76, base.B2i32(l7 == v76), v91, v90, l17, l16)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
						if l7 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(2606)
							if l1 != 0 {
								F_recordDependencyOn(m, l0, v24+int32(4), int32(80))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(1259)
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l4)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v113
									v120 = int32(83)
									F_recordDependencyOn(m, l0, v24+int32(4), v120)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											m.G0 = v24 + int32(16)
											return
										}
									}
								}
							} else {
								v120 = int32(105)
								F_recordDependencyOn(m, l0, v24+int32(4), v120)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_CommandCounterIncrement(m)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										m.G0 = v24 + int32(16)
										return
									}
								}
							}
						} else {
							F_CommandCounterIncrement(m)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								m.G0 = v24 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = v37 + int32(4)
				F_errmsg(m, int32(393912), v24)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(492711), int32(10746), int32(90513))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
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
func F_addFkRecurseReferencing(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
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
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v185 int32
	_ = v185
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
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
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v615 int32
	_ = v615
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	v20 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(96)
	m.G0 = v33
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v20
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+119)))
	if v40 != int32(102) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	if v43 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L10
	} else {
		goto L73
	}
L4:
	;
	m.G0 = v33 + int32(96)
	return
L5:
	;
	v218 = F_RelationGetPartitionDesc(m, l2, int32(1))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L30
	}
L6:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	F_createForeignKeyCheckTriggers(m, v46, v47, l1, l5, l4, l16, l17, v33+int32(92), v33+int32(88))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v56 = v40
	goto L9
L9:
	;
	switch v56&int32(255) - int32(112) {
	case 0:
		goto L5
	default:
		goto L4
	case 2:
		goto L6
	}
L10:
	;
	return
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+119)))
	v56 = v55
	goto L9
L12:
	;
	if l14 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+15)))
	if v63 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+14)))
	if v64 != int32(1) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v68 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v199 = F_palloc0(m, int32(32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L27
	}
L17:
	;
	v146 = F_palloc0(m, int32(140))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L10
	} else {
		goto L24
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 <= int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v95 = int32(0)
	goto L20
L20:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v74+v95<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v110 == v67 {
		v185 = v109
		goto L16
	} else {
		goto L22
	}
L21:
	;
	goto L17
L22:
	;
	v113 = v95 + int32(1)
	if v71 != v113 {
		v95 = v113
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v67
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)) = uint8(v152)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v155 = F_CreateTupleDescCopyConstr(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v155
	*(*int64)(unsafe.Add(mBase, uint32(v146)+88)) = int64(0)
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+84)) = uint8(v160)
	v162 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v146)+96)) = uint16(v162)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = F_lappend(m, v164, v146)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v165
	v185 = v146
	goto L16
L27:
	;
	v201 = F_get_constraint_name(m, l5)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v201
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v206
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+16)) = uint8(v210)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v185)+64))
	v214 = F_lappend(m, v213, v199)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+64)) = v214
	goto L4
L30:
	;
	v222 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	if int32(0) < v224 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v33)+92))
	v257 = v20
	goto L35
L33:
	;
	goto L34
L34:
	;
	F_sequence_close(m, v222, int32(3))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L10
	} else {
		goto L72
	}
L35:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263+v257<<(uint(int32(2))%32))))
	v268 = F_table_open(m, v267, l15)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L10
	} else {
		goto L40
	}
L36:
	;
	goto L34
L37:
	;
	v449 = F_RelationGetFKeyList(m, v268)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L10
	} else {
		goto L59
	}
L38:
	;
	if l6&int32(1) == int32(0) {
		goto L37
	} else {
		goto L56
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L10
	} else {
		goto L52
	}
L40:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v268)+48))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+118)))
	if v271 == int32(116) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+24)))
	if v274 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_CheckTableNotInUse(m, v268, int32(538847))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L10
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v268)+52))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v283 = F_build_attrmap_by_name(m, v280, v281, int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	if l6 <= int32(0) {
		goto L37
	} else {
		goto L47
	}
L47:
	;
	v287 = int32(0)
	if l6 == int32(1) {
		v391 = v287
		goto L38
	} else {
		goto L48
	}
L48:
	;
	v310 = v287
	v313 = v287
	goto L49
L49:
	;
	v321 = int32(1)
	v322 = v310 << (uint(v321) % 32)
	v324 = v33 + int32(16)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v328 = int32(*(*int16)(unsafe.Add(mBase, uint32(l8+v322))))
	v332 = int32(2)
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326+v328<<(uint(v321)%32)-v332))))
	*(*uint16)(unsafe.Add(mBase, uint32(v322+v324))) = uint16(v334)
	v337 = v322 | v332
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(l8+v337))))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341+v343<<(uint(v321)%32)-v332))))
	*(*uint16)(unsafe.Add(mBase, uint32(v337+v324))) = uint16(v349)
	v352 = v310 + v332
	v354 = v313 + v332
	if v354 != l6&int32(2147483646) {
		v310 = v352
		v313 = v354
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v391 = v352
	goto L38
L51:
	;
	goto L50
L52:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(143567), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(492711), int32(4460), int32(408449))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v404 = int32(1)
	v405 = v391 << (uint(v404) % 32)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v411 = int32(*(*int16)(unsafe.Add(mBase, uint32(l8+v405))))
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409+v411<<(uint(v404)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v405+(v33+int32(16))))) = uint16(v417)
	goto L37
L57:
	;
	F_sequence_close(m, v268, int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L10
	} else {
		goto L70
	}
L58:
	;
	v534 = int32(1)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_addFkConstraint(m, v33+int32(4), v534, v535, l1, v268, l3, l4, l5, l6, l7, v33+int32(16), l9, l10, l11, l12, l13, v534, l18)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L10
	} else {
		goto L68
	}
L59:
	;
	v451 = F_copyObjectImpl(m, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	if v451 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v455 = int32(0)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	if v456 <= v455 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v478 = v455
	goto L63
L63:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v489+v478<<(uint(int32(2))%32))))
	v496 = F_tryAttachPartitionForeignKey(m, l0, v493, v268, l5, l6, v33+int32(16), l7, l9, v232, v231, v222)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L10
	} else {
		goto L65
	}
L64:
	;
	goto L58
L65:
	;
	if v496 != 0 {
		goto L57
	} else {
		goto L66
	}
L66:
	;
	v499 = v478 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	if v499 < v500 {
		v478 = v499
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_addFkRecurseReferencing(m, l0, l1, v268, l3, l4, v541, l6, l7, v33+int32(16), l9, l10, l11, l12, l13, l14, l15, v232, v231, l18)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	goto L57
L70:
	;
	v580 = v257 + int32(1)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	if v580 < v581 {
		v257 = v580
		goto L35
	} else {
		goto L71
	}
L71:
	;
	goto L36
L72:
	;
	goto L4
L73:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(165211), int32(0))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(492711), int32(11057), int32(335983))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addHLParsedLex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
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
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
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
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v369 int32
	_ = v369
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = l1 + int32(8)
	v19 = l2
	goto L4
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L77
	} else {
		goto L78
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if int32(0) < v31 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 <= v37 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if l3 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v38 << (uint(int32(1)) % 32)
	v45 = F_repalloc(m, v34, v38<<(uint(int32(5))%32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v49 = v34
	v50 = v37
	goto L11
L11:
	;
	v51 = int32(4)
	v53 = v50<<(uint(v51)%32) + v49
	v54 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v58+v59<<(uint(v51)%32))+1)) = uint8(v31)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v65<<(uint(v51)%32))+2)) = uint16(v35)
	v70 = F_palloc(m, v35)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = v45
	v50 = v48
	goto L11
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v74 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(v74)%32))+8)) = v70
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78+v79<<(uint(v74)%32))+8))
	if v35 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v86 + int32(1)
	goto L8
L16:
	;
	v84 = F__emscripten_memcpy_bulkmem(m, v83, v36, v35)
	mBase = m.M
	goto L18
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	F_pfree(m, v19)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L12
	} else {
		goto L75
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = v97
	v107 = l3
	goto L21
L21:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v112 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L22:
	;
	goto L19
L23:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+2)))
	v118 = v106 + v115&int32(1)
	v119 = F_strlen(m, v112)
	mBase = m.M
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v121 <= v122+v123 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v128 = v121
	v130 = v120
	goto L27
L25:
	;
	v157 = v120
	v158 = v122
	goto L26
L26:
	;
	v169 = v157 + v158<<(uint(int32(4))%32)
	v172 = int32(16383)
	if v172 <= v118 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128 << (uint(int32(1)) % 32)
	v145 = F_repalloc(m, v130, v128<<(uint(int32(5))%32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L29
	}
L28:
	;
	v157 = v145
	v158 = v149
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v145
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v148 <= v149+v150 {
		v128 = v148
		v130 = v145
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v175 = v172
	goto L33
L32:
	;
	v175 = v118
	goto L33
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v169-int32(12)))) = uint16(v175)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v177 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v181 = v169 - int32(16)
	v183 = v169 - int32(4)
	v187 = v16
	v189 = int32(0)
	v190 = v177
	goto L37
L35:
	;
	goto L36
L36:
	;
	v289 = v107 + int32(8)
	if v289 != 0 {
		v106 = v118
		v107 = v289
		goto L21
	} else {
		goto L74
	}
L37:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v199 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v271 = v189 + int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v271 < v272 {
		v187 = v187 + int32(12)
		v189 = v271
		v190 = v272
		goto L37
	} else {
		goto L73
	}
L40:
	;
	v202 = int32(12)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	v210 = v205 & int32(4095)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
	if v210 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v236 != 0 {
		goto L39
	} else {
		goto L69
	}
L42:
	;
	if v211 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if v119 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	v236 = int32(0)
	goto L41
L46:
	;
	goto L47
L47:
	;
	v216 = int32(0)
	if v216 < v119 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v219 = int32(-1)
	goto L50
L49:
	;
	v219 = v216
	goto L50
L50:
	;
	v236 = v219
	goto L41
L51:
	;
	v236 = base.B2i32(int32(0) < v210)
	goto L41
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(v210) < base.Ui32(v119) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v225 = v210
	goto L56
L55:
	;
	v225 = v119
	goto L56
L56:
	;
	v226 = F_memcmp(m, v16+v190*v202+int32(base.Ui32(v205)>>(uint(v202)%32)), v112, v225)
	mBase = m.M
	if v211 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v236 = v234
	goto L41
L58:
	;
	if v226 != 0 {
		v234 = v226
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v226 != 0 {
		v234 = v226
		goto L57
	} else {
		goto L62
	}
L61:
	;
	v236 = base.B2i32(v119 < v210)
	goto L41
L62:
	;
	if v210 == v119 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v236 = int32(0)
	goto L41
L64:
	;
	goto L65
L65:
	;
	if v210 < v119 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v233 = int32(-1)
	goto L68
L67:
	;
	v233 = int32(1)
	goto L68
L68:
	;
	v234 = v233
	goto L57
L69:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v237 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v240 = int32(4)
	v242 = v238 + v239<<(uint(v240)%32)
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
	*(*int64)(unsafe.Add(mBase, uint32(v242))) = v243
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v181)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v242)+8)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v247+v248<<(uint(v240)%32))+12)) = v187
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v257 = v253 + v254<<(uint(v240)%32)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v258 | int32(8)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v262 + int32(1)
	goto L39
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v187
	goto L39
L73:
	;
	goto L38
L74:
	;
	goto L22
L75:
	;
	if v304 != 0 {
		v19 = v304
		goto L4
	} else {
		goto L76
	}
L76:
	;
	goto L5
L77:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v321 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	return
L80:
	;
	v325 = l3 + int32(4)
	v326 = l3
	v328 = v321
	goto L83
L81:
	;
	goto L82
L82:
	;
	F_pfree(m, l3)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L12
	} else {
		goto L90
	}
L83:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+2)))
	if v338&int32(1) != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L82
L85:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v341 + int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v346 = v345
	goto L87
L86:
	;
	v346 = v328
	goto L87
L87:
	;
	F_pfree(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L12
	} else {
		goto L88
	}
L88:
	;
	v350 = v326 + int32(12)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	if v353 != 0 {
		v325 = v350
		v326 = v326 + int32(8)
		v328 = v353
		goto L83
	} else {
		goto L89
	}
L89:
	;
	goto L84
L90:
	;
	goto L79
}
func F_addHyperLogLog(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v7 = int32(32) - v6
	v8 = l1 << (uint(v6) % 32)
	if v8 != 0 {
		v15 = int32(32) - (base.I32_clz(v8) ^ int32(31))
		v16 = int32(255)
		if base.Ui32(v7&v16) < base.Ui32(v15&v16) {
			v21 = v7 + int32(1)
		} else {
			v21 = v15
		}
		v25 = v21
	} else {
		v25 = v7 + int32(1)
	}
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = v26 + int32(base.Ui32(l1)>>(uint(v7)%32))
	v30 = v25 & int32(255)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.Ui32(v31) < base.Ui32(v30) {
		v33 = v30
	} else {
		v33 = v31
	}
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v33)
	return
}
func F_add_dummy_return(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if v10 != 0 {
		v14 = F_palloc0(m, int32(32))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
			v19 = int32(1)
			v20 = v18 + v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v23
			v29 = F_list_make1_impl(m, v19, v7+int32(8))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = v14
				v34 = v29
				if v34 != 0 {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32)-int32(4))))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					if v43 == int32(11) {
						m.G0 = v7 + int32(16)
						return
					} else {
						v47 = F_palloc0(m, int32(20))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
							v53 = v51 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
							*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
							v62 = F_lappend(m, v61, v47)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				} else {
					v47 = F_palloc0(m, int32(20))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
						v53 = v51 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
						v62 = F_lappend(m, v61, v47)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		if v11 != 0 {
			v14 = F_palloc0(m, int32(32))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
				v19 = int32(1)
				v20 = v18 + v19
				*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v23
				v29 = F_list_make1_impl(m, v19, v7+int32(8))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = v14
					v34 = v29
					if v34 != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32)-int32(4))))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						if v43 == int32(11) {
							m.G0 = v7 + int32(16)
							return
						} else {
							v47 = F_palloc0(m, int32(20))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
								v53 = v51 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
								*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
								*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
								v62 = F_lappend(m, v61, v47)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
									*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					} else {
						v47 = F_palloc0(m, int32(20))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
							v53 = v51 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
							*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
							v62 = F_lappend(m, v61, v47)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
								*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v34 = v12
			if v34 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32)-int32(4))))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				if v43 == int32(11) {
					m.G0 = v7 + int32(16)
					return
				} else {
					v47 = F_palloc0(m, int32(20))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
						v53 = v51 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
						v62 = F_lappend(m, v61, v47)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
							*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			} else {
				v47 = F_palloc0(m, int32(20))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(11)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
					v53 = v51 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v53
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+472))
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
					v62 = F_lappend(m, v61, v47)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
						*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v62
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_add_outer_joins_to_relids(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
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
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	if l2 == int32(0) {
		v310 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v310
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v11 == int32(0) {
		v310 = l1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v14 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = F_bms_add_member(m, l1, v11)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	v23 = int32(0)
	if v22 == v23 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	return int32(0)
L8:
	;
	return v17
L9:
	;
	if v76 == int32(0) {
		v310 = l1
		goto L1
	} else {
		goto L23
	}
L10:
	;
	v76 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if l1 == int32(0) {
		v67 = v23
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v76 = v67
	goto L9
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v33 < v32 {
		v67 = v23
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v35 = int32(1)
	if v32 <= v35 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v38 = v35
	goto L18
L17:
	;
	v38 = v32
	goto L18
L18:
	;
	v39 = int32(8)
	v44 = int32(0)
	goto L19
L19:
	;
	v51 = v44 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22+v39+v51)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l1+v39))))
	v58 = v53 & (v55 ^ int32(-1))
	v60 = base.B2i32(v58 == int32(0))
	if v58 != 0 {
		v67 = v60
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v67 = v60
	goto L13
L21:
	;
	v62 = v44 + int32(1)
	if v62 != v38 {
		v44 = v62
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v80 = F_bms_add_member(m, l1, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v82 == int32(0) {
		v310 = v80
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v85 = F_bms_copy(m, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v87 == int32(0) {
		v310 = v80
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 <= int32(0) {
		v310 = v80
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v94 = int32(0)
	v95 = v80
	v99 = v85
	goto L29
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v94<<(uint(int32(2))%32))))
	if v106 == l2 {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v310 = v302
	goto L1
L31:
	;
	v306 = v94 + int32(1)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v306 < v307 {
		v94 = v306
		v95 = v302
		v99 = v303
		goto L29
	} else {
		goto L90
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	if v108 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+20))
	if v111 != int32(1) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v114 = F_bms_is_member(m, v108, v99)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	if v114 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	v119 = F_bms_is_member(m, v118, v95)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	if v119 != 0 {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v122 = int32(0)
	if v121 == v122 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v175 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L53
	}
L40:
	;
	v175 = int32(1)
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v95 == int32(0) {
		v166 = v122
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v175 = v166
	goto L39
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v132 < v131 {
		v166 = v122
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v134 = int32(1)
	if v131 <= v134 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v137 = v134
	goto L48
L47:
	;
	v137 = v131
	goto L48
L48:
	;
	v138 = int32(8)
	v143 = int32(0)
	goto L49
L49:
	;
	v150 = v143 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v121+v138+v150)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v95+v138))))
	v157 = v152 & (v154 ^ int32(-1))
	v159 = base.B2i32(v157 == int32(0))
	if v157 != 0 {
		v166 = v159
		goto L43
	} else {
		goto L51
	}
L50:
	;
	v166 = v159
	goto L43
L51:
	;
	v161 = v143 + int32(1)
	if v161 != v137 {
		v143 = v161
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v179 = int32(0)
	if v178 == v179 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v232 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L68
	}
L55:
	;
	v232 = int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	if v95 == int32(0) {
		v223 = v179
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v232 = v223
	goto L54
L59:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v189 < v188 {
		v223 = v179
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v191 = int32(1)
	if v188 <= v191 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v194 = v191
	goto L63
L62:
	;
	v194 = v188
	goto L63
L63:
	;
	v195 = int32(8)
	v200 = int32(0)
	goto L64
L64:
	;
	v207 = v200 << (uint(int32(2)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v178+v195+v207)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+(v95+v195))))
	v214 = v209 & (v211 ^ int32(-1))
	v216 = base.B2i32(v214 == int32(0))
	if v214 != 0 {
		v223 = v216
		goto L58
	} else {
		goto L66
	}
L65:
	;
	v223 = v216
	goto L58
L66:
	;
	v218 = v200 + int32(1)
	if v218 != v194 {
		v200 = v218
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v106)+36))
	v236 = int32(0)
	if v235 == v236 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v289 == int32(0) {
		v302 = v95
		v303 = v99
		goto L31
	} else {
		goto L83
	}
L70:
	;
	v289 = int32(1)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if v95 == int32(0) {
		v280 = v236
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v289 = v280
	goto L69
L74:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v246 < v245 {
		v280 = v236
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v248 = int32(1)
	if v245 <= v248 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v251 = v248
	goto L78
L77:
	;
	v251 = v245
	goto L78
L78:
	;
	v252 = int32(8)
	v257 = int32(0)
	goto L79
L79:
	;
	v264 = v257 << (uint(int32(2)) % 32)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v235+v252+v264)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v264+(v95+v252))))
	v271 = v266 & (v268 ^ int32(-1))
	v273 = base.B2i32(v271 == int32(0))
	if v271 != 0 {
		v280 = v273
		goto L73
	} else {
		goto L81
	}
L80:
	;
	v280 = v273
	goto L73
L81:
	;
	v275 = v257 + int32(1)
	if v275 != v251 {
		v257 = v275
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v106)+24))
	v293 = F_bms_add_member(m, v95, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	if l3 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v296 = F_lappend(m, v295, v106)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v106)+28))
	v300 = F_bms_add_members(m, v99, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L7
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v296
	goto L87
L89:
	;
	v302 = v293
	v303 = v300
	goto L31
L90:
	;
	goto L30
}
func F_add_rtes_to_flat_rtable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v16 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v3
	goto L6
L4:
	;
	goto L5
L5:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+52))
	if v63 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v27<<(uint(int32(2))%32))))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v49 = v27 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v49 < v50 {
		v27 = v49
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	F_add_rte_to_flat_rtable(m, v19, v45, v37)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	switch v40 {
	case 0:
		goto L9
	case 1:
		goto L11
	default:
		goto L8
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v41 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	return
L14:
	;
	goto L8
L15:
	;
	goto L7
L16:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v67 <= v66 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v73 = int32(1)
	v75 = v66
	goto L18
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v75<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L1
L20:
	;
	v149 = int32(1)
	v152 = v75 + v149
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v152 < v153 {
		v73 = v73 + v149
		v75 = v152
		goto L18
	} else {
		goto L45
	}
L21:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+20)))
	if v89 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v90) <= base.Ui32(v73) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v73<<(uint(int32(2))%32))))
	if v96 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+140))
	if v99 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v19
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v103
	v109 = F_query_tree_walker_impl(m, v103, int32(834), v13+int32(8), int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L20
L29:
	;
	v142 = v99
	goto L31
L30:
	;
	v113 = F_fetch_upper_rel(m, v99, int32(7), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L13
	} else {
		goto L32
	}
L31:
	;
	F_add_rtes_to_flat_rtable(m, v142, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L44
	}
L32:
	;
	v115 = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+32))
	if v117 == v115 {
		v137 = v115
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v137 == int32(0) {
		goto L20
	} else {
		goto L43
	}
L34:
	;
	goto L33
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v121 = v120
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if base.Ui32(int32(2)) <= base.Ui32(v125-int32(301)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v137 = int32(1)
	goto L34
L38:
	;
	if v125 != int32(290) {
		v137 = v115
		goto L34
	} else {
		goto L41
	}
L39:
	;
	v121 = v124 + int32(72)
	goto L36
L40:
	;
	goto L37
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v124)+72))
	if v132 != 0 {
		v137 = v115
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v96)+140))
	v142 = v141
	goto L31
L44:
	;
	goto L20
L45:
	;
	goto L19
}
func F_add_unique_group_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 float32
	_ = v24
	var v26 float32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 float64
	_ = v69
	var v72 int32
	_ = v72
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 float64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v5 = int32(0)
	v8 = float64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v5)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v91 = F_remove_nulling_relids(m, l2, v89, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L34
	} else {
		goto L35
	}
L2:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)))
	if v58 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
	v23 = v21 + v22
	v24 = *(*float32)(unsafe.Add(mBase, uint32(v23)+8))
	v26 = *(*float32)(unsafe.Add(mBase, uint32(v23)+16))
	v53 = base.F64_promote_f32(v26)
	v54 = base.F64_promote_f32(v24)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v28 == int32(16) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v53 = float64(2)
	v54 = v8
	goto L2
L7:
	;
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v32 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v39 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
	if v35 != int32(5) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v53 = float64(-1)
	v54 = v8
	goto L2
L12:
	;
	v53 = float64(0)
	v54 = v8
	goto L2
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v42 != int32(6) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
	switch v46 - int32(65530) {
	case 0:
		goto L15
	default:
		goto L12
	case 5:
		v53 = float64(-1)
		v54 = v8
		goto L2
	}
L15:
	;
	v53 = float64(1)
	v54 = v8
	goto L2
L16:
	;
	v59 = base.F64_neg(base.F64_sub(float64(1), v54))
	goto L18
L17:
	;
	v59 = v53
	goto L18
L18:
	;
	if base.F64_gt(v59, float64(0)) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v62 = F_clamp_row_est(m, v59)
	mBase = m.M
	v88 = v62
	goto L1
L20:
	;
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v63 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v66)
	v88 = float64(200)
	goto L1
L23:
	;
	goto L24
L24:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v63)+120))
	if base.F64_le(v69, float64(0)) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v72)
	v88 = float64(200)
	goto L1
L26:
	;
	goto L27
L27:
	;
	if base.F64_lt(v59, float64(0)) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v79 = F_clamp_row_est(m, base.F64_mul(v69, base.F64_neg(v59)))
	mBase = m.M
	v88 = v79
	goto L1
L29:
	;
	goto L30
L30:
	;
	if base.F64_lt(v69, float64(200)) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v82 = F_clamp_row_est(m, v69)
	mBase = m.M
	v88 = v82
	goto L1
L32:
	;
	goto L33
L33:
	;
	v83 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v83)
	v88 = float64(200)
	goto L1
L34:
	;
	return int32(0)
L35:
	;
	if l1 == int32(0) {
		v139 = v5
		goto L37
	} else {
		goto L38
	}
L36:
	;
	m.G0 = v11 + int32(16)
	return v158
L37:
	;
	v144 = F_palloc(m, int32(24))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L34
	} else {
		goto L55
	}
L38:
	;
	v99 = int32(0)
	v102 = l1
	goto L39
L39:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v106 <= v99 {
		v139 = v102
		goto L37
	} else {
		goto L41
	}
L40:
	;
	v139 = v132
	goto L37
L41:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v99<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v114 = F_equal(m, v91, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	if v114 != 0 {
		v158 = v102
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v116 == v117 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v132 != 0 {
		v99 = v131 + int32(1)
		v102 = v132
		goto L39
	} else {
		goto L54
	}
L45:
	;
	v131 = v99
	v132 = v102
	goto L44
L46:
	;
	goto L47
L47:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v121 = F_exprs_known_equal(m, l0, v91, v119, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L34
	} else {
		goto L48
	}
L48:
	;
	if v121 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v131 = v99
	v132 = v102
	goto L44
L50:
	;
	goto L51
L51:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v112)+8))
	if base.F64_le(v125, v88) != 0 {
		v158 = v102
		goto L36
	} else {
		goto L52
	}
L52:
	;
	v129 = F_list_delete_nth_cell(m, v102, v99)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v131 = v99 - int32(1)
	v132 = v129
	goto L44
L54:
	;
	goto L40
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v91
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*float64)(unsafe.Add(mBase, uint32(v144)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v147
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+16)) = uint8(v150)
	v152 = F_lappend(m, v139, v144)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L34
	} else {
		goto L56
	}
L56:
	;
	v158 = v152
	goto L36
}
func F_adjust_view_column_set(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L29
	} else {
		goto L61
	}
L2:
	;
	if int32(0) <= v68 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v68 = base.I32_ctz(v54) | v55<<(uint(int32(5))%32)
	goto L2
L4:
	;
	v68 = int32(-2)
	goto L2
L5:
	;
	v21 = base.I32_div_s(int32(0), int32(32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= v21 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = l0 + int32(8)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v21<<(uint(int32(2))%32))))
	v32 = v29 & int32(-1)
	if v32 != 0 {
		v54 = v32
		v55 = v21
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v34 = v21 + int32(1)
	if v34 == v22 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = v34
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v25+v37<<(uint(int32(2))%32))))
	if v44 != 0 {
		v54 = v44
		v55 = v37
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v46 = v37 + int32(1)
	if v46 != v22 {
		v37 = v46
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v74 = v3
	v76 = v68
	goto L16
L14:
	;
	v238 = v3
	goto L15
L15:
	;
	m.G0 = v10 + int32(16)
	return v238
L16:
	;
	v79 = v76 - int32(7)
	if v79&int32(65535) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v238 = v173
	goto L15
L18:
	;
	if l0 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L19:
	;
	if l1 == int32(0) {
		v173 = v74
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v118 = base.I32_extend16_s(v79)
	if l1 != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v87 <= v86 {
		v173 = v74
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v92 = v86
	v93 = v74
	goto L24
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v92<<(uint(int32(2))%32))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+26)))
	if v102 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v173 = v113
	goto L18
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+8)))
	v109 = F_bms_add_member(m, v93, v106+int32(7))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v113 = v93
	goto L28
L28:
	;
	v115 = v92 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v115 < v116 {
		v92 = v115
		v93 = v113
		goto L24
	} else {
		goto L31
	}
L29:
	;
	return int32(0)
L30:
	;
	v113 = v109
	goto L28
L31:
	;
	goto L25
L32:
	;
	if v156 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L33:
	;
	goto L32
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v122 <= int32(0) {
		v156 = int32(0)
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v156 = int32(0)
	goto L33
L37:
	;
	v125 = int32(0)
	if v125 < v122 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v128 = v122
	goto L40
L39:
	;
	v128 = v125
	goto L40
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v133 = int32(0)
	goto L41
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v129+v133<<(uint(int32(2))%32))))
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+8)))
	if v142 == v118&int32(65535) {
		v156 = v141
		goto L33
	} else {
		goto L43
	}
L42:
	;
	goto L36
L43:
	;
	v145 = v133 + int32(1)
	if v145 != v128 {
		v133 = v145
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+26)))
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v162 != int32(6) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v161)+8)))
	v168 = F_bms_add_member(m, v74, v165+int32(7))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L29
	} else {
		goto L48
	}
L48:
	;
	v173 = v168
	goto L18
L49:
	;
	if int32(0) <= v232 {
		v74 = v173
		v76 = v232
		goto L16
	} else {
		goto L60
	}
L50:
	;
	v232 = base.I32_ctz(v218) | v219<<(uint(int32(5))%32)
	goto L49
L51:
	;
	v232 = int32(-2)
	goto L49
L52:
	;
	v183 = v76 + int32(1)
	v185 = base.I32_div_s(v183, int32(32))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v186 <= v185 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v189 = l0 + int32(8)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v185<<(uint(int32(2))%32))))
	v196 = v193 & (int32(-1) << (uint(v183) % 32))
	if v196 != 0 {
		v218 = v196
		v219 = v185
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v198 = v185 + int32(1)
	if v198 == v186 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v201 = v198
	goto L56
L56:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v189+v201<<(uint(int32(2))%32))))
	if v208 != 0 {
		v218 = v208
		v219 = v201
		goto L50
	} else {
		goto L58
	}
L57:
	;
	goto L51
L58:
	;
	v210 = v201 + int32(1)
	if v210 != v186 {
		v201 = v210
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L17
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v118
	F_errmsg_internal(m, int32(73045), v10)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L29
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(493601), int32(3099), int32(105984))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L29
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_akeys(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_akeys(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_anyarray_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(24047)
			F_errmsg(m, int32(192035), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492364), int32(154), int32(278303))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anycompatiblearray_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(24099)
			F_errmsg(m, int32(192035), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492364), int32(174), int32(278355))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anycompatiblenonarray_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(24077)
			F_errmsg(m, int32(192001), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492364), int32(377), int32(66470))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anycompatiblerange_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(400083)
			F_errmsg(m, int32(192035), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492364), int32(220), int32(279071))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anymultirange_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_multirange_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_anynonarray_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(24065)
			F_errmsg(m, int32(192001), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492364), int32(375), int32(66454))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anytimestamp_typmod_check(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if int32(0) <= l1 {
		if base.Ui32(l1) < base.Ui32(int32(7)) {
			v42 = l1
			m.G0 = v7 + int32(32)
			return v42
		} else {
			v13 = int32(6)
			v16 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v16 == int32(0) {
					v42 = v13
					m.G0 = v7 + int32(32)
					return v42
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(6)
						if l0 != 0 {
							v29 = int32(537825)
						} else {
							v29 = int32(740129)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, int32(486332), v7+int32(16))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494095), int32(138), int32(317275))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = v13
								m.G0 = v7 + int32(32)
								return v42
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if l0 != 0 {
					v56 = int32(537825)
				} else {
					v56 = int32(740129)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(342303), v7)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494095), int32(131), int32(317275))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
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
func F_append_total_cost_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v8 != v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v8 < v10 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = int32(1)
	v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+56))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(v9)+56))
	if base.F64_lt(v18, v19) != 0 {
		v84 = v17
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v15 = int32(1)
	goto L6
L5:
	;
	v15 = int32(-1)
	goto L6
L6:
	;
	return v15
L7:
	;
	return v84
L8:
	;
	if base.F64_gt(v18, v19) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(-1)
L10:
	;
	goto L11
L11:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v7)+48))
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v9)+48))
	if base.F64_lt(v24, v25) != 0 {
		v84 = v17
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if base.F64_gt(v24, v25) != 0 {
		v84 = int32(-1)
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v30 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v84 = v83
	goto L7
L15:
	;
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v32 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v40 = int32(-1)
	goto L20
L19:
	;
	v40 = int32(0)
	goto L20
L20:
	;
	v83 = v40
	goto L14
L21:
	;
	v83 = int32(1)
	goto L14
L22:
	;
	goto L23
L23:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v44 != v45 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v45 < v44 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v51 = int32(8)
	v58 = v44 - int32(1)
	goto L31
L27:
	;
	v50 = int32(1)
	goto L29
L28:
	;
	v50 = int32(-1)
	goto L29
L29:
	;
	v83 = v50
	goto L14
L30:
	;
	if base.Ui32(v67) < base.Ui32(v65) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v63 = v58 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v30+v51+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+(v32+v51))))
	if v65 != v67 {
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v83 = int32(0)
	goto L14
L33:
	;
	v70 = v58 - int32(1)
	if int32(0) <= v70 {
		v58 = v70
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v77 = int32(1)
	goto L37
L36:
	;
	v77 = int32(-1)
	goto L37
L37:
	;
	v83 = v77
	goto L14
}
func F_apply_scanjoin_target_to_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
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
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v449 int32
	_ = v449
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	if v19 == v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l4 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v34 == v32 {
		v54 = v32
		goto L11
	} else {
		goto L12
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if v22 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v25 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if v28 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L3
L8:
	;
	return
L9:
	;
	v63 = int32(0)
	goto L1
L10:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L20
	}
L11:
	;
	goto L10
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v38 = v37
	goto L13
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if base.Ui32(int32(2)) <= base.Ui32(v42-int32(301)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v54 = int32(1)
	goto L11
L15:
	;
	if v42 != int32(290) {
		v54 = v32
		goto L11
	} else {
		goto L18
	}
L16:
	;
	v38 = v41 + int32(72)
	goto L13
L17:
	;
	goto L14
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v49 != 0 {
		v54 = v32
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v54 != 0 {
		v63 = int32(0)
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(0)
	v63 = int32(1)
	goto L1
L22:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v63 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v69
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(0)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v77 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v127 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v80 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v81 <= v80 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v90 = v80
	goto L32
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v101 = v98 + v90<<(uint(int32(2))%32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if l5 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L29
L34:
	;
	v110 = v90 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v110 < v111 {
		v90 = v110
		goto L32
	} else {
		goto L39
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v104
	goto L34
L36:
	;
	goto L37
L37:
	;
	v106 = F_create_projection_path(m, l0, l1, v102, v76)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v106
	goto L34
L39:
	;
	goto L33
L40:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+38)))
	if v178 != 0 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v130 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v131 <= v130 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v140 = v130
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v151 = v148 + v140<<(uint(int32(2))%32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if l5 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L40
L45:
	;
	v160 = v140 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v160 < v161 {
		v140 = v160
		goto L43
	} else {
		goto L50
	}
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v154
	goto L45
L47:
	;
	goto L48
L48:
	;
	v156 = F_create_projection_path(m, l0, l1, v152, v76)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v156
	goto L45
L50:
	;
	goto L44
L51:
	;
	F_adjust_paths_for_srfs(m, l0, l1, l2, l3)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181+v182<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v188
	if v63 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v190 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	goto L57
L57:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v482 != int32(1) {
		goto L123
	} else {
		goto L124
	}
L58:
	;
	if int32(0) <= v247 {
		goto L69
	} else {
		goto L70
	}
L59:
	;
	v247 = base.I32_ctz(v233) | v234<<(uint(int32(5))%32)
	goto L58
L60:
	;
	v247 = int32(-2)
	goto L58
L61:
	;
	v200 = base.I32_div_s(int32(0), int32(32))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v201 <= v200 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v204 = v190 + int32(8)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v200<<(uint(int32(2))%32))))
	v211 = v208 & int32(-1)
	if v211 != 0 {
		v233 = v211
		v234 = v200
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v213 = v200 + int32(1)
	if v213 == v201 {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v216 = v213
	goto L65
L65:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v204+v216<<(uint(int32(2))%32))))
	if v223 != 0 {
		v233 = v223
		v234 = v216
		goto L59
	} else {
		goto L67
	}
L66:
	;
	goto L60
L67:
	;
	v225 = v216 + int32(1)
	if v225 != v201 {
		v216 = v225
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v261 = v247
	v262 = v7
	goto L72
L70:
	;
	v464 = v7
	goto L71
L71:
	;
	F_add_paths_to_append_rel(m, l0, l1, v464)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L8
	} else {
		goto L122
	}
L72:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v264+v261<<(uint(int32(2))%32))))
	v269 = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v268)+32))
	if v271 == v269 {
		v291 = v269
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v464 = v391
	goto L71
L74:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+256))
	if v393 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L75:
	;
	if v291 != 0 {
		v391 = v262
		goto L74
	} else {
		goto L85
	}
L76:
	;
	goto L75
L77:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v275 = v274
	goto L78
L78:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	if base.Ui32(int32(2)) <= base.Ui32(v279-int32(301)) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v291 = int32(1)
	goto L76
L80:
	;
	if v279 != int32(290) {
		v291 = v269
		goto L76
	} else {
		goto L83
	}
L81:
	;
	v275 = v278 + int32(72)
	goto L78
L82:
	;
	goto L79
L83:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v278)+72))
	if v286 != 0 {
		v291 = v269
		goto L76
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
	v296 = F_find_appinfos_by_relids(m, l0, v293, v17+int32(12))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v298 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v298 < v300 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v310 = v298
	v311 = v298
	goto L90
L88:
	;
	v343 = v298
	goto L89
L89:
	;
	F_pfree(m, v296)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L8
	} else {
		goto L96
	}
L90:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v310<<(uint(int32(2))%32))))
	v322 = F_copy_pathtarget(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L8
	} else {
		goto L92
	}
L91:
	;
	v343 = v329
	goto L89
L92:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v326 = F_adjust_appendrel_attrs(m, l0, v324, v325, v296)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v326
	v329 = F_lappend(m, v311, v322)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v332 = v310 + int32(1)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v332 < v333 {
		v310 = v332
		v311 = v329
		goto L90
	} else {
		goto L95
	}
L95:
	;
	goto L91
L96:
	;
	F_apply_scanjoin_target_to_paths(m, l0, v268, v343, l3, l4, l5)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	v353 = int32(0)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v268)+32))
	if v355 == v353 {
		v375 = v353
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v375 != 0 {
		v391 = v262
		goto L74
	} else {
		goto L108
	}
L99:
	;
	goto L98
L100:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	v359 = v358
	goto L101
L101:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	if base.Ui32(int32(2)) <= base.Ui32(v363-int32(301)) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v375 = int32(1)
	goto L99
L103:
	;
	if v363 != int32(290) {
		v375 = v353
		goto L99
	} else {
		goto L106
	}
L104:
	;
	v359 = v362 + int32(72)
	goto L101
L105:
	;
	goto L102
L106:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v362)+72))
	if v370 != 0 {
		v375 = v353
		goto L99
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v377 = F_lappend(m, v262, v268)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	v391 = v377
	goto L74
L110:
	;
	if int32(0) <= v449 {
		v261 = v449
		v262 = v391
		goto L72
	} else {
		goto L121
	}
L111:
	;
	v449 = base.I32_ctz(v435) | v436<<(uint(int32(5))%32)
	goto L110
L112:
	;
	v449 = int32(-2)
	goto L110
L113:
	;
	v400 = v261 + int32(1)
	v402 = base.I32_div_s(v400, int32(32))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v403 <= v402 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v406 = v393 + int32(8)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v406+v402<<(uint(int32(2))%32))))
	v413 = v410 & (int32(-1) << (uint(v400) % 32))
	if v413 != 0 {
		v435 = v413
		v436 = v402
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v415 = v402 + int32(1)
	if v415 == v403 {
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v418 = v415
	goto L117
L117:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v406+v418<<(uint(int32(2))%32))))
	if v425 != 0 {
		v435 = v425
		v436 = v418
		goto L111
	} else {
		goto L119
	}
L118:
	;
	goto L112
L119:
	;
	v427 = v418 + int32(1)
	if v427 != v403 {
		v418 = v427
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	goto L73
L122:
	;
	goto L57
L123:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L8
	} else {
		goto L130
	}
L124:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(1)<<(uint(v485)%32)&int32(44) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v493 = base.B2i32(base.Ui32(v485) <= base.Ui32(int32(5)))
	goto L127
L126:
	;
	v493 = int32(0)
	goto L127
L127:
	;
	if v493 != 0 {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	goto L123
L130:
	;
	m.G0 = v17 + int32(16)
	return
}
func F_arraycontsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int64
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 float64
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 float64
	_ = v162
	var v163 int32
	_ = v163
	var v166 float64
	_ = v166
	var v168 float32
	_ = v168
	var v172 int32
	_ = v172
	var v178 float64
	_ = v178
	var v179 int32
	_ = v179
	var v184 float64
	_ = v184
	var v187 int32
	_ = v187
	var v192 float64
	_ = v192
	var v200 float64
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 float64
	_ = v205
	var v220 float64
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	v9 = float64(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v25 = F_get_restriction_variable(m, v16, v17, v18, v13+int32(8), v13+int32(4), v13+int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		if v25 == int32(0) {
			if v15 == int32(2750) {
				v35 = float64(0.01)
			} else {
				v35 = float64(0.005)
			}
			v220 = v35
			v222 = F_Float8GetDatum(m, v220)
			mBase = m.M
			v223 = m.ExcPending
			if v223 != 0 {
				return int32(0)
			} else {
				m.G0 = v13 + int32(112)
				return v222
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if v37 != int32(7) {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
				if v40 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
					m.T0[v41].(func(*base.Module, int32))(m, v40)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v15 == int32(2750) {
							v48 = float64(0.01)
						} else {
							v48 = float64(0.005)
						}
						v220 = v48
						v222 = F_Float8GetDatum(m, v220)
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(112)
							return v222
						}
					}
				} else {
					if v15 == int32(2750) {
						v48 = float64(0.01)
					} else {
						v48 = float64(0.005)
					}
					v220 = v48
					v222 = F_Float8GetDatum(m, v220)
					mBase = m.M
					v223 = m.ExcPending
					if v223 != 0 {
						return int32(0)
					} else {
						m.G0 = v13 + int32(112)
						return v222
					}
				}
			} else {
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
				if v49 == int32(1) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
					if v52 == int32(0) {
						v220 = v9
						v222 = F_Float8GetDatum(m, v220)
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int32(0)
						} else {
							m.G0 = v13 + int32(112)
							return v222
						}
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
						m.T0[v55].(func(*base.Module, int32))(m, v52)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v220 = v9
							v222 = F_Float8GetDatum(m, v220)
							mBase = m.M
							v223 = m.ExcPending
							if v223 != 0 {
								return int32(0)
							} else {
								m.G0 = v13 + int32(112)
								return v222
							}
						}
					}
				} else {
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
					if v58 != 0 {
						v66 = v15
					} else {
						if v15 == int32(2751) {
							v66 = int32(2752)
						} else {
							if v15 == int32(2752) {
								v65 = int32(2751)
							} else {
								v65 = v15
							}
							v66 = v65
						}
					}
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
					v68 = F_get_base_element_type(m, v67)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						if v68 == int32(0) {
							if v66 == int32(2750) {
								v192 = float64(0.01)
							} else {
								v192 = float64(0.005)
							}
							v200 = v192
							v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
							if v201 != 0 {
								v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
								m.T0[v202].(func(*base.Module, int32))(m, v201)
								mBase = m.M
								v204 = m.ExcPending
								if v204 != 0 {
									return int32(0)
								} else {
									v205 = float64(0)
									if base.F64_lt(v200, v205) != 0 {
										v220 = v205
									} else {
										if base.F64_gt(v200, float64(1)) == int32(0) {
											v220 = v200
										} else {
											v220 = float64(1)
										}
									}
									v222 = F_Float8GetDatum(m, v220)
									mBase = m.M
									v223 = m.ExcPending
									if v223 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(112)
										return v222
									}
								}
							} else {
								v205 = float64(0)
								if base.F64_lt(v200, v205) != 0 {
									v220 = v205
								} else {
									if base.F64_gt(v200, float64(1)) == int32(0) {
										v220 = v200
									} else {
										v220 = float64(1)
									}
								}
								v222 = F_Float8GetDatum(m, v220)
								mBase = m.M
								v223 = m.ExcPending
								if v223 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(112)
									return v222
								}
							}
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							v73 = F_get_base_element_type(m, v72)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v73 != v68 {
									if v66 == int32(2750) {
										v192 = float64(0.01)
									} else {
										v192 = float64(0.005)
									}
									v200 = v192
									v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
									if v201 != 0 {
										v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
										m.T0[v202].(func(*base.Module, int32))(m, v201)
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return int32(0)
										} else {
											v205 = float64(0)
											if base.F64_lt(v200, v205) != 0 {
												v220 = v205
											} else {
												if base.F64_gt(v200, float64(1)) == int32(0) {
													v220 = v200
												} else {
													v220 = float64(1)
												}
											}
											v222 = F_Float8GetDatum(m, v220)
											mBase = m.M
											v223 = m.ExcPending
											if v223 != 0 {
												return int32(0)
											} else {
												m.G0 = v13 + int32(112)
												return v222
											}
										}
									} else {
										v205 = float64(0)
										if base.F64_lt(v200, v205) != 0 {
											v220 = v205
										} else {
											if base.F64_gt(v200, float64(1)) == int32(0) {
												v220 = v200
											} else {
												v220 = float64(1)
											}
										}
										v222 = F_Float8GetDatum(m, v220)
										mBase = m.M
										v223 = m.ExcPending
										if v223 != 0 {
											return int32(0)
										} else {
											m.G0 = v13 + int32(112)
											return v222
										}
									}
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									v79 = F_lookup_type_cache(m, v68, int32(64))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+108))
										if v81 == int32(0) {
											if v66 == int32(2750) {
												v88 = float64(0.01)
											} else {
												v88 = float64(0.005)
											}
											v200 = v88
											v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
											if v201 != 0 {
												v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
												m.T0[v202].(func(*base.Module, int32))(m, v201)
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int32(0)
												} else {
													v205 = float64(0)
													if base.F64_lt(v200, v205) != 0 {
														v220 = v205
													} else {
														if base.F64_gt(v200, float64(1)) == int32(0) {
															v220 = v200
														} else {
															v220 = float64(1)
														}
													}
													v222 = F_Float8GetDatum(m, v220)
													mBase = m.M
													v223 = m.ExcPending
													if v223 != 0 {
														return int32(0)
													} else {
														m.G0 = v13 + int32(112)
														return v222
													}
												}
											} else {
												v205 = float64(0)
												if base.F64_lt(v200, v205) != 0 {
													v220 = v205
												} else {
													if base.F64_gt(v200, float64(1)) == int32(0) {
														v220 = v200
													} else {
														v220 = float64(1)
													}
												}
												v222 = F_Float8GetDatum(m, v220)
												mBase = m.M
												v223 = m.ExcPending
												if v223 != 0 {
													return int32(0)
												} else {
													m.G0 = v13 + int32(112)
													return v222
												}
											}
										} else {
											v89 = F_pg_detoast_datum(m, v77)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
												if v91 == int32(0) {
													v172 = int32(0)
													v178 = F_mcelem_array_selec(m, v89, v79, v172, v172, v172, v172, v172, v172, v66)
													mBase = m.M
													v179 = m.ExcPending
													if v179 != 0 {
														return int32(0)
													} else {
														v184 = v178
														if v89 == v77 {
															v200 = v184
															v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
															if v201 != 0 {
																v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																m.T0[v202].(func(*base.Module, int32))(m, v201)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	v205 = float64(0)
																	if base.F64_lt(v200, v205) != 0 {
																		v220 = v205
																	} else {
																		if base.F64_gt(v200, float64(1)) == int32(0) {
																			v220 = v200
																		} else {
																			v220 = float64(1)
																		}
																	}
																	v222 = F_Float8GetDatum(m, v220)
																	mBase = m.M
																	v223 = m.ExcPending
																	if v223 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v13 + int32(112)
																		return v222
																	}
																}
															} else {
																v205 = float64(0)
																if base.F64_lt(v200, v205) != 0 {
																	v220 = v205
																} else {
																	if base.F64_gt(v200, float64(1)) == int32(0) {
																		v220 = v200
																	} else {
																		v220 = float64(1)
																	}
																}
																v222 = F_Float8GetDatum(m, v220)
																mBase = m.M
																v223 = m.ExcPending
																if v223 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v13 + int32(112)
																	return v222
																}
															}
														} else {
															F_pfree(m, v89)
															mBase = m.M
															v187 = m.ExcPending
															if v187 != 0 {
																return int32(0)
															} else {
																v200 = v184
																v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																if v201 != 0 {
																	v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																	m.T0[v202].(func(*base.Module, int32))(m, v201)
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		v205 = float64(0)
																		if base.F64_lt(v200, v205) != 0 {
																			v220 = v205
																		} else {
																			if base.F64_gt(v200, float64(1)) == int32(0) {
																				v220 = v200
																			} else {
																				v220 = float64(1)
																			}
																		}
																		v222 = F_Float8GetDatum(m, v220)
																		mBase = m.M
																		v223 = m.ExcPending
																		if v223 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v13 + int32(112)
																			return v222
																		}
																	}
																} else {
																	v205 = float64(0)
																	if base.F64_lt(v200, v205) != 0 {
																		v220 = v205
																	} else {
																		if base.F64_gt(v200, float64(1)) == int32(0) {
																			v220 = v200
																		} else {
																			v220 = float64(1)
																		}
																	}
																	v222 = F_Float8GetDatum(m, v220)
																	mBase = m.M
																	v223 = m.ExcPending
																	if v223 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v13 + int32(112)
																		return v222
																	}
																}
															}
														}
													}
												} else {
													v96 = *(*int32)(unsafe.Add(mBase, uint32(v79)+108))
													v97 = F_statistic_proc_security_check(m, v13+int32(8), v96)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int32(0)
													} else {
														if v97 == int32(0) {
															v172 = int32(0)
															v178 = F_mcelem_array_selec(m, v89, v79, v172, v172, v172, v172, v172, v172, v66)
															mBase = m.M
															v179 = m.ExcPending
															if v179 != 0 {
																return int32(0)
															} else {
																v184 = v178
																if v89 == v77 {
																	v200 = v184
																	v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																	if v201 != 0 {
																		v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																		m.T0[v202].(func(*base.Module, int32))(m, v201)
																		mBase = m.M
																		v204 = m.ExcPending
																		if v204 != 0 {
																			return int32(0)
																		} else {
																			v205 = float64(0)
																			if base.F64_lt(v200, v205) != 0 {
																				v220 = v205
																			} else {
																				if base.F64_gt(v200, float64(1)) == int32(0) {
																					v220 = v200
																				} else {
																					v220 = float64(1)
																				}
																			}
																			v222 = F_Float8GetDatum(m, v220)
																			mBase = m.M
																			v223 = m.ExcPending
																			if v223 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v13 + int32(112)
																				return v222
																			}
																		}
																	} else {
																		v205 = float64(0)
																		if base.F64_lt(v200, v205) != 0 {
																			v220 = v205
																		} else {
																			if base.F64_gt(v200, float64(1)) == int32(0) {
																				v220 = v200
																			} else {
																				v220 = float64(1)
																			}
																		}
																		v222 = F_Float8GetDatum(m, v220)
																		mBase = m.M
																		v223 = m.ExcPending
																		if v223 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v13 + int32(112)
																			return v222
																		}
																	}
																} else {
																	F_pfree(m, v89)
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int32(0)
																	} else {
																		v200 = v184
																		v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																		if v201 != 0 {
																			v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																			m.T0[v202].(func(*base.Module, int32))(m, v201)
																			mBase = m.M
																			v204 = m.ExcPending
																			if v204 != 0 {
																				return int32(0)
																			} else {
																				v205 = float64(0)
																				if base.F64_lt(v200, v205) != 0 {
																					v220 = v205
																				} else {
																					if base.F64_gt(v200, float64(1)) == int32(0) {
																						v220 = v200
																					} else {
																						v220 = float64(1)
																					}
																				}
																				v222 = F_Float8GetDatum(m, v220)
																				mBase = m.M
																				v223 = m.ExcPending
																				if v223 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v13 + int32(112)
																					return v222
																				}
																			}
																		} else {
																			v205 = float64(0)
																			if base.F64_lt(v200, v205) != 0 {
																				v220 = v205
																			} else {
																				if base.F64_gt(v200, float64(1)) == int32(0) {
																					v220 = v200
																				} else {
																					v220 = float64(1)
																				}
																			}
																			v222 = F_Float8GetDatum(m, v220)
																			mBase = m.M
																			v223 = m.ExcPending
																			if v223 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v13 + int32(112)
																				return v222
																			}
																		}
																	}
																}
															}
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
															v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+22)))
															v110 = F_get_attstatsslot(m, v13+int32(76), v101, int32(4), int32(0), int32(3))
															mBase = m.M
															v111 = m.ExcPending
															if v111 != 0 {
																return int32(0)
															} else {
																if v110 != 0 {
																	if v66 != int32(2752) {
																		v126 = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v126
																		v131 = int64(0)
																		*(*int64)(unsafe.Add(mBase, uint32(v13-int32(-64)))) = v131
																		*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v131
																		*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v131
																		*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v131
																		v140 = v126
																		v141 = v126
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
																		v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
																		v144 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
																		v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
																		v146 = F_mcelem_array_selec(m, v89, v79, v142, v143, v144, v145, v141, v140, v66)
																		mBase = m.M
																		v147 = m.ExcPending
																		if v147 != 0 {
																			return int32(0)
																		} else {
																			F_free_attstatsslot(m, v13+int32(40))
																			mBase = m.M
																			v151 = m.ExcPending
																			if v151 != 0 {
																				return int32(0)
																			} else {
																				F_free_attstatsslot(m, v13+int32(76))
																				mBase = m.M
																				v155 = m.ExcPending
																				if v155 != 0 {
																					return int32(0)
																				} else {
																					v166 = v146
																					v168 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																					v184 = base.F64_mul(v166, base.F64_sub(float64(1), base.F64_promote_f32(v168)))
																					if v89 == v77 {
																						v200 = v184
																						v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																						if v201 != 0 {
																							v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																							m.T0[v202].(func(*base.Module, int32))(m, v201)
																							mBase = m.M
																							v204 = m.ExcPending
																							if v204 != 0 {
																								return int32(0)
																							} else {
																								v205 = float64(0)
																								if base.F64_lt(v200, v205) != 0 {
																									v220 = v205
																								} else {
																									if base.F64_gt(v200, float64(1)) == int32(0) {
																										v220 = v200
																									} else {
																										v220 = float64(1)
																									}
																								}
																								v222 = F_Float8GetDatum(m, v220)
																								mBase = m.M
																								v223 = m.ExcPending
																								if v223 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v13 + int32(112)
																									return v222
																								}
																							}
																						} else {
																							v205 = float64(0)
																							if base.F64_lt(v200, v205) != 0 {
																								v220 = v205
																							} else {
																								if base.F64_gt(v200, float64(1)) == int32(0) {
																									v220 = v200
																								} else {
																									v220 = float64(1)
																								}
																							}
																							v222 = F_Float8GetDatum(m, v220)
																							mBase = m.M
																							v223 = m.ExcPending
																							if v223 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v13 + int32(112)
																								return v222
																							}
																						}
																					} else {
																						F_pfree(m, v89)
																						mBase = m.M
																						v187 = m.ExcPending
																						if v187 != 0 {
																							return int32(0)
																						} else {
																							v200 = v184
																							v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																							if v201 != 0 {
																								v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																								m.T0[v202].(func(*base.Module, int32))(m, v201)
																								mBase = m.M
																								v204 = m.ExcPending
																								if v204 != 0 {
																									return int32(0)
																								} else {
																									v205 = float64(0)
																									if base.F64_lt(v200, v205) != 0 {
																										v220 = v205
																									} else {
																										if base.F64_gt(v200, float64(1)) == int32(0) {
																											v220 = v200
																										} else {
																											v220 = float64(1)
																										}
																									}
																									v222 = F_Float8GetDatum(m, v220)
																									mBase = m.M
																									v223 = m.ExcPending
																									if v223 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v13 + int32(112)
																										return v222
																									}
																								}
																							} else {
																								v205 = float64(0)
																								if base.F64_lt(v200, v205) != 0 {
																									v220 = v205
																								} else {
																									if base.F64_gt(v200, float64(1)) == int32(0) {
																										v220 = v200
																									} else {
																										v220 = float64(1)
																									}
																								}
																								v222 = F_Float8GetDatum(m, v220)
																								mBase = m.M
																								v223 = m.ExcPending
																								if v223 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v13 + int32(112)
																									return v222
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																		v120 = F_get_attstatsslot(m, v13+int32(40), v116, int32(5), int32(0), int32(2))
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return int32(0)
																		} else {
																			if v120 == int32(0) {
																				v126 = int32(0)
																				*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v126
																				v131 = int64(0)
																				*(*int64)(unsafe.Add(mBase, uint32(v13-int32(-64)))) = v131
																				*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v131
																				*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v131
																				*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v131
																				v140 = v126
																				v141 = v126
																			} else {
																				v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
																				v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
																				v140 = v124
																				v141 = v125
																			}
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
																			v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
																			v144 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
																			v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
																			v146 = F_mcelem_array_selec(m, v89, v79, v142, v143, v144, v145, v141, v140, v66)
																			mBase = m.M
																			v147 = m.ExcPending
																			if v147 != 0 {
																				return int32(0)
																			} else {
																				F_free_attstatsslot(m, v13+int32(40))
																				mBase = m.M
																				v151 = m.ExcPending
																				if v151 != 0 {
																					return int32(0)
																				} else {
																					F_free_attstatsslot(m, v13+int32(76))
																					mBase = m.M
																					v155 = m.ExcPending
																					if v155 != 0 {
																						return int32(0)
																					} else {
																						v166 = v146
																						v168 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																						v184 = base.F64_mul(v166, base.F64_sub(float64(1), base.F64_promote_f32(v168)))
																						if v89 == v77 {
																							v200 = v184
																							v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																							if v201 != 0 {
																								v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																								m.T0[v202].(func(*base.Module, int32))(m, v201)
																								mBase = m.M
																								v204 = m.ExcPending
																								if v204 != 0 {
																									return int32(0)
																								} else {
																									v205 = float64(0)
																									if base.F64_lt(v200, v205) != 0 {
																										v220 = v205
																									} else {
																										if base.F64_gt(v200, float64(1)) == int32(0) {
																											v220 = v200
																										} else {
																											v220 = float64(1)
																										}
																									}
																									v222 = F_Float8GetDatum(m, v220)
																									mBase = m.M
																									v223 = m.ExcPending
																									if v223 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v13 + int32(112)
																										return v222
																									}
																								}
																							} else {
																								v205 = float64(0)
																								if base.F64_lt(v200, v205) != 0 {
																									v220 = v205
																								} else {
																									if base.F64_gt(v200, float64(1)) == int32(0) {
																										v220 = v200
																									} else {
																										v220 = float64(1)
																									}
																								}
																								v222 = F_Float8GetDatum(m, v220)
																								mBase = m.M
																								v223 = m.ExcPending
																								if v223 != 0 {
																									return int32(0)
																								} else {
																									m.G0 = v13 + int32(112)
																									return v222
																								}
																							}
																						} else {
																							F_pfree(m, v89)
																							mBase = m.M
																							v187 = m.ExcPending
																							if v187 != 0 {
																								return int32(0)
																							} else {
																								v200 = v184
																								v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																								if v201 != 0 {
																									v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																									m.T0[v202].(func(*base.Module, int32))(m, v201)
																									mBase = m.M
																									v204 = m.ExcPending
																									if v204 != 0 {
																										return int32(0)
																									} else {
																										v205 = float64(0)
																										if base.F64_lt(v200, v205) != 0 {
																											v220 = v205
																										} else {
																											if base.F64_gt(v200, float64(1)) == int32(0) {
																												v220 = v200
																											} else {
																												v220 = float64(1)
																											}
																										}
																										v222 = F_Float8GetDatum(m, v220)
																										mBase = m.M
																										v223 = m.ExcPending
																										if v223 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v13 + int32(112)
																											return v222
																										}
																									}
																								} else {
																									v205 = float64(0)
																									if base.F64_lt(v200, v205) != 0 {
																										v220 = v205
																									} else {
																										if base.F64_gt(v200, float64(1)) == int32(0) {
																											v220 = v200
																										} else {
																											v220 = float64(1)
																										}
																									}
																									v222 = F_Float8GetDatum(m, v220)
																									mBase = m.M
																									v223 = m.ExcPending
																									if v223 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v13 + int32(112)
																										return v222
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v156 = int32(0)
																	v162 = F_mcelem_array_selec(m, v89, v79, v156, v156, v156, v156, v156, v156, v66)
																	mBase = m.M
																	v163 = m.ExcPending
																	if v163 != 0 {
																		return int32(0)
																	} else {
																		v166 = v162
																		v168 = *(*float32)(unsafe.Add(mBase, uint32(v102+v103)+8))
																		v184 = base.F64_mul(v166, base.F64_sub(float64(1), base.F64_promote_f32(v168)))
																		if v89 == v77 {
																			v200 = v184
																			v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																			if v201 != 0 {
																				v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																				m.T0[v202].(func(*base.Module, int32))(m, v201)
																				mBase = m.M
																				v204 = m.ExcPending
																				if v204 != 0 {
																					return int32(0)
																				} else {
																					v205 = float64(0)
																					if base.F64_lt(v200, v205) != 0 {
																						v220 = v205
																					} else {
																						if base.F64_gt(v200, float64(1)) == int32(0) {
																							v220 = v200
																						} else {
																							v220 = float64(1)
																						}
																					}
																					v222 = F_Float8GetDatum(m, v220)
																					mBase = m.M
																					v223 = m.ExcPending
																					if v223 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v13 + int32(112)
																						return v222
																					}
																				}
																			} else {
																				v205 = float64(0)
																				if base.F64_lt(v200, v205) != 0 {
																					v220 = v205
																				} else {
																					if base.F64_gt(v200, float64(1)) == int32(0) {
																						v220 = v200
																					} else {
																						v220 = float64(1)
																					}
																				}
																				v222 = F_Float8GetDatum(m, v220)
																				mBase = m.M
																				v223 = m.ExcPending
																				if v223 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v13 + int32(112)
																					return v222
																				}
																			}
																		} else {
																			F_pfree(m, v89)
																			mBase = m.M
																			v187 = m.ExcPending
																			if v187 != 0 {
																				return int32(0)
																			} else {
																				v200 = v184
																				v201 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
																				if v201 != 0 {
																					v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
																					m.T0[v202].(func(*base.Module, int32))(m, v201)
																					mBase = m.M
																					v204 = m.ExcPending
																					if v204 != 0 {
																						return int32(0)
																					} else {
																						v205 = float64(0)
																						if base.F64_lt(v200, v205) != 0 {
																							v220 = v205
																						} else {
																							if base.F64_gt(v200, float64(1)) == int32(0) {
																								v220 = v200
																							} else {
																								v220 = float64(1)
																							}
																						}
																						v222 = F_Float8GetDatum(m, v220)
																						mBase = m.M
																						v223 = m.ExcPending
																						if v223 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v13 + int32(112)
																							return v222
																						}
																					}
																				} else {
																					v205 = float64(0)
																					if base.F64_lt(v200, v205) != 0 {
																						v220 = v205
																					} else {
																						if base.F64_gt(v200, float64(1)) == int32(0) {
																							v220 = v200
																						} else {
																							v220 = float64(1)
																						}
																					}
																					v222 = F_Float8GetDatum(m, v220)
																					mBase = m.M
																					v223 = m.ExcPending
																					if v223 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v13 + int32(112)
																						return v222
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
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_asin(m *base.Module, l0 float64) float64 {
	var v7 int64
	_ = v7
	var v12 int32
	_ = v12
	var v36 float64
	_ = v36
	var v71 float64
	_ = v71
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v107 float64
	_ = v107
	var v112 float64
	_ = v112
	var v117 float64
	_ = v117
	var v121 float64
	_ = v121
	var v130 float64
	_ = v130
	var v139 float64
	_ = v139
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	v7 = base.I64_reinterpret_f64(l0)
	v12 = base.I32_wrap_i64(int64(base.Ui64(v7)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v12) {
		if base.I32_wrap_i64(v7)|(v12-int32(1072693248)) == int32(0) {
			return base.F64_add(base.F64_mul(l0, float64(1.5707963267948966)), float64(7.52316384526264e-37))
		} else {
			return base.F64_div(float64(0), base.F64_sub(l0, l0))
		}
	} else {
		if base.Ui32(v12) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v12+int32(-1048576)) < base.Ui32(int32(1044381696)) {
				v144 = l0
				return v144
			} else {
				v36 = base.F64_mul(l0, l0)
				return base.F64_add(base.F64_mul(l0, base.F64_div(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))), l0)
			}
		} else {
			v71 = float64(1)
			v75 = base.F64_mul(base.F64_sub(v71, base.F64_abs(l0)), float64(0.5))
			v76 = base.F64_sqrt(v75)
			v107 = base.F64_div(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, base.F64_add(base.F64_mul(v75, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), v71))
			if base.Ui32(int32(1072640819)) <= base.Ui32(v12) {
				v112 = base.F64_add(base.F64_mul(v76, v107), v76)
				v139 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v112, v112), float64(-6.123233995736766e-17)))
			} else {
				v117 = float64(0.7853981633974483)
				v121 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v76) & int64(-4294967296))
				v130 = base.F64_div(base.F64_sub(v75, base.F64_mul(v121, v121)), base.F64_add(v76, v121))
				v139 = base.F64_add(base.F64_sub(base.F64_sub(v117, base.F64_add(v121, v121)), base.F64_sub(base.F64_mul(base.F64_add(v76, v76), v107), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v130, v130)))), v117)
			}
			if v7 < int64(0) {
				v143 = base.F64_neg(v139)
			} else {
				v143 = v139
			}
			v144 = v143
			return v144
		}
	}
}
func F_assign_checkpoint_completion_target(m *base.Module, l0 float64, l1 int32) {
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
	var v17 float64
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	*(*float64)(unsafe.Add(mBase, _consts[223])) = l0
	v8 = *(*int32)(unsafe.Add(mBase, _consts[222]))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	v12 = base.I32_div_s(v10, int32(1048576))
	v13 = base.I32_div_s(v8, v12)
	v17 = base.F64_div(base.F64_convert_i32_s(v13), base.F64_add(l0, float64(1)))
	if base.F64_lt(base.F64_abs(v17), float64(2.147483648e+09)) != 0 {
		v21 = base.I32_trunc_f64_s(v17)
		v23 = v21
	} else {
		v23 = int32(-2147483648)
	}
	if v23 <= int32(1) {
		v26 = int32(1)
	} else {
		v26 = v23
	}
	*(*int32)(unsafe.Add(mBase, _consts[224])) = v26
	return
}
func F_assign_collations_walker(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
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
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
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
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(96)
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+60)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
	v25 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v28 - int32(1) {
	case 0:
		goto L12
	default:
		goto L6
	case 5, 6, 7, 33, 55, 56, 57:
		goto L11
	case 8:
		goto L10
	case 10:
		goto L9
	case 13:
		goto L7
	case 24:
		goto L19
	case 30:
		goto L20
	case 31:
		goto L8
	case 35:
		goto L18
	case 36:
		goto L17
	case 53, 59, 62, 63, 64, 65, 105:
		goto L14
	case 54:
		goto L16
	case 61:
		goto L15
	case 66:
		goto L13
	}
L3:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	F_merge_collation_state(m, v856, v858, v857, v865, v866, l1)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L21
	} else {
		goto L228
	}
L4:
	;
	F_exprSetCollation(m, l0, int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L21
	} else {
		goto L227
	}
L5:
	;
	v772 = F_exprType(m, l0)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L21
	} else {
		goto L194
	}
L6:
	;
	v758 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L21
	} else {
		goto L192
	}
L7:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v19
	v732 = F_assign_collations_walker(m, v724, v15+int32(72))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L21
	} else {
		goto L188
	}
L8:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v669 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L9:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v656 = F_assign_collations_walker(m, v653, v15+int32(48))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L21
	} else {
		goto L177
	}
L10:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	switch v264 - int32(104) {
	case 0:
		goto L86
	default:
		goto L85
	case 6:
		goto L84
	case 7:
		goto L87
	}
L11:
	;
	v259 = F_exprCollation(m, l0)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L21
	} else {
		goto L82
	}
L12:
	;
	v254 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L21
	} else {
		goto L81
	}
L13:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v239 == int32(0) {
		goto L1
	} else {
		goto L78
	}
L14:
	;
	v237 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L21
	} else {
		goto L77
	}
L15:
	;
	v190 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L21
	} else {
		goto L64
	}
L16:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v162 = F_get_typcollation(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L21
	} else {
		goto L52
	}
L17:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v95 = int32(0)
	v100 = v3
	goto L33
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v52 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v44 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L21
	} else {
		goto L23
	}
L20:
	;
	v34 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v856 = v39
	v857 = v38
	v858 = int32(3)
	goto L3
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v46 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v856 = int32(0)
	v857 = v25
	v858 = v3
	goto L3
L25:
	;
	goto L26
L26:
	;
	v51 = F_exprLocation(m, l0)
	mBase = m.M
	v856 = v46
	v857 = v51
	v858 = int32(1)
	goto L3
L27:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v55 <= int32(0) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v59 = int32(0)
	goto L29
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v59<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v19
	v83 = F_assign_collations_walker(m, v75, v15+int32(72))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L31
	}
L30:
	;
	goto L1
L31:
	;
	v86 = v59 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v86 < v87 {
		v59 = v86
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v104 = int32(0)
	if v90 == v104 {
		v114 = v104
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v89 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v108 <= v95 {
		v114 = int32(0)
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v114 = v110 + v95<<(uint(int32(2))%32)
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L39:
	;
	goto L40
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v119 <= v95 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v131
	v140 = F_list_make2_impl(m, v15+int32(12), v15+int32(8))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L21
	} else {
		goto L46
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v100
	goto L1
L43:
	;
	if v114 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v126 = v123 + v95<<(uint(int32(2))%32)
	if v126 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v129
	v149 = F_assign_collations_walker(m, v140, v15+int32(72))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	if v155 != int32(2) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v158 = v153
	goto L50
L49:
	;
	v158 = int32(0)
	goto L50
L50:
	;
	v159 = F_lappend_oid(m, v100, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	v95 = v95 + int32(1)
	v100 = v159
	goto L33
L52:
	;
	v167 = F_expression_tree_walker_impl(m, l0, int32(483), v15+int32(48))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	if v162 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_exprSetCollation(m, l0, v162)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L21
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v162 != int32(100) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v856 = v162
	v857 = v25
	v858 = v162
	goto L3
L58:
	;
	v176 = F_exprLocation(m, l0)
	mBase = m.M
	F_exprSetCollation(m, l0, v162)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L21
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v179 = int32(2)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v182 == v179 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v856 = v162
	v857 = v176
	v858 = int32(1)
	goto L3
L62:
	;
	F_exprSetCollation(m, l0, v181)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L21
	} else {
		goto L63
	}
L63:
	;
	v856 = v181
	v857 = v180
	v858 = v182
	goto L3
L64:
	;
	v192 = int32(2)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v195 != v192 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v856 = v194
	v857 = v193
	v858 = v195
	goto L3
L66:
	;
	goto L67
L67:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v198 == int32(0) {
		v856 = v194
		v857 = v193
		v858 = v192
		goto L3
	} else {
		goto L68
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v209 = F_get_collation_name(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L21
	} else {
		goto L71
	}
L71:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v212 = F_get_collation_name(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v209
	F_errmsg(m, int32(705642), v15+int32(16))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L21
	} else {
		goto L73
	}
L73:
	;
	F_errhint(m, int32(574723), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	F_parser_errposition(m, v225, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L21
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(496720), int32(480), int32(221024))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	goto L1
L78:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+26)))
	if v244 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v246 = F_exprCollation(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v250 = F_exprLocation(m, v249)
	mBase = m.M
	v856 = v246
	v857 = v250
	v858 = int32(1)
	goto L3
L81:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v856 = v258
	v857 = v256
	v858 = v257
	goto L3
L82:
	;
	v263 = F_exprLocation(m, l0)
	mBase = m.M
	v856 = v259
	v857 = v263
	v858 = base.B2i32(v259 != int32(0))
	goto L3
L83:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v643
	v651 = F_assign_collations_walker(m, v642, v15+int32(72))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L21
	} else {
		goto L176
	}
L84:
	;
	v567 = v15 + int32(48)
	v568 = m.G0
	v570 = v568 - int32(32)
	m.G0 = v570
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v572 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L21
	} else {
		goto L161
	}
L86:
	;
	v345 = v15 + int32(48)
	v347 = m.G0
	v349 = v347 - int32(32)
	m.G0 = v349
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v351 != 0 {
		goto L105
	} else {
		goto L106
	}
L87:
	;
	v268 = v15 + int32(48)
	v269 = m.G0
	v271 = v269 - int32(32)
	m.G0 = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v273 == int32(0) {
		v284 = v3
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v286 = F_assign_collations_walker(m, v285, v268)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L21
	} else {
		goto L92
	}
L89:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v276 != int32(1) {
		v284 = v3
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v280 = F_get_func_variadictype(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L21
	} else {
		goto L91
	}
L91:
	;
	v284 = base.B2i32(v280 == int32(0))
	goto L88
L92:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v288 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	m.G0 = v271 + int32(32)
	goto L83
L94:
	;
	v291 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v292 <= v291 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v300 = v291
	goto L96
L96:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307+v300<<(uint(int32(2))%32))))
	if v284 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L93
L98:
	;
	v326 = v300 + int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v326 < v327 {
		v300 = v326
		goto L96
	} else {
		goto L104
	}
L99:
	;
	v312 = F_assign_collations_walker(m, v311, v268)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L21
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v271)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v271)+8)) = v314
	v322 = F_assign_collations_walker(m, v311, v271+int32(8))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L21
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	goto L98
L104:
	;
	goto L97
L105:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v353 = v352
	goto L107
L106:
	;
	v353 = int32(0)
	goto L107
L107:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v354 == int32(0) {
		v368 = v351
		v369 = v3
		v370 = v3
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if v368 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	if v358 != int32(1) {
		v368 = v351
		v369 = v357
		v370 = v3
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v362 = F_get_func_variadictype(m, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L21
	} else {
		goto L111
	}
L111:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v368 = v364
	v369 = v357
	v370 = base.B2i32(v362 == int32(0))
	goto L108
L112:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	v373 = v371
	goto L114
L113:
	;
	v373 = int32(0)
	goto L114
L114:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v374 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v377 = v375
	goto L117
L116:
	;
	v377 = int32(0)
	goto L117
L117:
	;
	v378 = v373 - v377
	if int32(0) < v378 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v385 = v353
	v386 = v378
	goto L121
L119:
	;
	v415 = v353
	goto L120
L120:
	;
	if v415 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L121:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	v394 = F_assign_collations_walker(m, v393, v345)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L21
	} else {
		goto L123
	}
L122:
	;
	v415 = v406
	goto L120
L123:
	;
	v397 = v385 + int32(4)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	if base.Ui32(v397) < base.Ui32(v400+v401<<(uint(int32(2))%32)) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v406 = v397
	goto L126
L125:
	;
	v406 = int32(0)
	goto L126
L126:
	;
	v407 = int32(1)
	if base.Ui32(v407) < base.Ui32(v386) {
		v385 = v406
		v386 = v386 - v407
		goto L121
	} else {
		goto L127
	}
L127:
	;
	goto L122
L128:
	;
	goto L83
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L21
	} else {
		goto L153
	}
L130:
	;
	m.G0 = v349 + int32(32)
	goto L128
L131:
	;
	if v369 == int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v431 = v415
	v434 = v369
	goto L133
L133:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+28)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v349)+20)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v349)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v441
	v451 = F_assign_collations_walker(m, v440, v349+int32(8))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L21
	} else {
		goto L135
	}
L134:
	;
	goto L130
L135:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v456 = F_assign_collations_walker(m, v453, v349+int32(8))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L21
	} else {
		goto L136
	}
L136:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v349)+16))
	if v458 == int32(2) {
		goto L129
	} else {
		goto L137
	}
L137:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	if v461 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	if v370 != 0 {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v465 = F_exprCollation(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L21
	} else {
		goto L140
	}
L140:
	;
	if v465 == v461 {
		goto L138
	} else {
		goto L141
	}
L141:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v469 = F_exprType(m, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L21
	} else {
		goto L142
	}
L142:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v472 = F_exprTypmod(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L21
	} else {
		goto L143
	}
L143:
	;
	v475 = F_makeRelabelType(m, v468, v469, v472, v461, int32(2))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L21
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+4)) = v475
	goto L138
L145:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v349)+20))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v349)+24))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v349)+28))
	F_merge_collation_state(m, v461, v458, v479, v480, v481, v345)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L21
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v485 = v431 + int32(4)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	if base.Ui32(v487+v488<<(uint(int32(2))%32)) <= base.Ui32(v485) {
		goto L130
	} else {
		goto L149
	}
L148:
	;
	goto L147
L149:
	;
	if v485 == int32(0) {
		goto L130
	} else {
		goto L150
	}
L150:
	;
	v496 = v434 + int32(4)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if base.Ui32(v498+v499<<(uint(int32(2))%32)) <= base.Ui32(v496) {
		goto L130
	} else {
		goto L151
	}
L151:
	;
	if v496 != 0 {
		v431 = v485
		v434 = v496
		goto L133
	} else {
		goto L152
	}
L152:
	;
	goto L134
L153:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L21
	} else {
		goto L154
	}
L154:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v527 = F_get_collation_name(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L21
	} else {
		goto L155
	}
L155:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v349)+24))
	v530 = F_get_collation_name(m, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L21
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+4)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v349))) = v527
	F_errmsg(m, int32(705642), v349)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L21
	} else {
		goto L157
	}
L157:
	;
	F_errhint(m, int32(574723), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L21
	} else {
		goto L158
	}
L158:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v349)+28))
	F_parser_errposition(m, v541, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L21
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(496720), int32(1010), int32(142692))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L21
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	v554 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+50)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v554
	F_errmsg_internal(m, int32(485553), v15+int32(32))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L21
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(496720), int32(616), int32(221024))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L21
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	m.G0 = v570 + int32(32)
	goto L83
L165:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v575 <= int32(0) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v583 = v3
	goto L167
L167:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v590+v583<<(uint(int32(2))%32))))
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+26)))
	if v595 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L164
L169:
	;
	v612 = v583 + int32(1)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v612 < v613 {
		v583 = v612
		goto L167
	} else {
		goto L175
	}
L170:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	*(*int32)(unsafe.Add(mBase, uint32(v570)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v570)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v570)+8)) = v598
	v606 = F_assign_collations_walker(m, v594, v570+int32(8))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L21
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v608 = F_assign_collations_walker(m, v594, v567)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L21
	} else {
		goto L174
	}
L173:
	;
	goto L169
L174:
	;
	goto L169
L175:
	;
	goto L168
L176:
	;
	goto L5
L177:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v659
	v667 = F_assign_collations_walker(m, v658, v15+int32(72))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L21
	} else {
		goto L178
	}
L178:
	;
	goto L5
L179:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v722 = F_assign_collations_walker(m, v719, v15+int32(48))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L21
	} else {
		goto L187
	}
L180:
	;
	v672 = int32(0)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v673 <= v672 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v679 = v672
	goto L182
L182:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v669)+12))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v688+v679<<(uint(int32(2))%32))))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
	v696 = F_assign_collations_walker(m, v693, v15+int32(48))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L21
	} else {
		goto L184
	}
L183:
	;
	goto L179
L184:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v692)+8))
	v701 = F_assign_collations_walker(m, v698, v15+int32(48))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L21
	} else {
		goto L185
	}
L185:
	;
	v704 = v679 + int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v669)+4))
	if v704 < v705 {
		v679 = v704
		goto L182
	} else {
		goto L186
	}
L186:
	;
	goto L183
L187:
	;
	goto L5
L188:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+76)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v735
	v743 = F_assign_collations_walker(m, v734, v15+int32(72))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L21
	} else {
		goto L189
	}
L189:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v748 = F_assign_collations_walker(m, v745, v15+int32(48))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L21
	} else {
		goto L190
	}
L190:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v753 = F_assign_collations_walker(m, v750, v15+int32(48))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L21
	} else {
		goto L191
	}
L191:
	;
	goto L5
L192:
	;
	goto L5
L193:
	;
	F_exprSetCollation(m, l0, v796)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L21
	} else {
		goto L205
	}
L194:
	;
	v774 = F_get_typcollation(m, v772)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L21
	} else {
		goto L195
	}
L195:
	;
	if v774 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v779 = int32(0)
	v793 = v779
	v794 = int32(-1)
	v795 = v779
	v796 = v779
	goto L193
L197:
	;
	goto L198
L198:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v782 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v786 = F_exprLocation(m, l0)
	mBase = m.M
	v793 = v774
	v794 = v786
	v795 = int32(1)
	v796 = v774
	goto L193
L200:
	;
	goto L201
L201:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v782 != int32(2) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v791 = v787
	goto L204
L203:
	;
	v791 = int32(0)
	goto L204
L204:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v793 = v787
	v794 = v792
	v795 = v782
	v796 = v791
	goto L193
L205:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	if v799 == int32(2) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v806 = v804 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v806) {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	goto L208
L208:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v830 = v828 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v830) {
		goto L219
	} else {
		goto L220
	}
L209:
	;
	v856 = v793
	v857 = v794
	v858 = v795
	goto L3
L210:
	;
	goto L209
L211:
	;
	v810 = int32(1) << (uint(v806) % 32)
	if v810&int32(3904) == int32(0) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v822))) = int32(0)
	goto L210
L213:
	;
	if v810&int32(5) != 0 {
		v822 = int32(16)
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v822 = int32(24)
	goto L212
L216:
	;
	if v806 != int32(30) {
		goto L210
	} else {
		goto L217
	}
L217:
	;
	v822 = int32(12)
	goto L212
L218:
	;
	v856 = v793
	v857 = v794
	v858 = v795
	goto L3
L219:
	;
	goto L218
L220:
	;
	v834 = int32(1) << (uint(v830) % 32)
	if v834&int32(3904) == int32(0) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v846))) = v826
	goto L219
L222:
	;
	if v834&int32(5) != 0 {
		v846 = int32(16)
		goto L221
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v846 = int32(24)
	goto L221
L225:
	;
	if v830 != int32(30) {
		goto L219
	} else {
		goto L226
	}
L226:
	;
	v846 = int32(12)
	goto L221
L227:
	;
	v856 = v181
	v857 = v180
	v858 = v179
	goto L3
L228:
	;
	goto L1
}
func F_assign_synchronous_commit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = l0 - int32(2)
	if base.Ui32(int32(3)) <= base.Ui32(v6) {
		v9 = int32(-1)
	} else {
		v9 = v6
	}
	*(*int32)(unsafe.Add(mBase, _consts[671])) = v9
	return
}
func F_atoi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	v5 = l0
	for {
		v10 = v5 + int32(1)
		v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5))))
		if base.B2i32(v11 == int32(32))|base.B2i32(base.Ui32(v11-int32(9)) < base.Ui32(int32(5))) != 0 {
			v5 = v10
			continue
		} else {
			break
		}
		break
	}
	v19 = int32(1)
	switch v11&int32(255) - int32(43) {
	case 0:
		v25 = v19
		v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
		v27 = v26
		v28 = v10
		v29 = v25
	default:
		v27 = v11
		v28 = v5
		v29 = v19
	case 2:
		v25 = int32(0)
		v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
		v27 = v26
		v28 = v10
		v29 = v25
	}
	v30 = int32(0)
	v32 = v27 - int32(48)
	if base.Ui32(v32) <= base.Ui32(int32(9)) {
		v35 = v30
		v36 = v32
		v37 = v28
		for {
			v39 = int32(10)
			v41 = v35*v39 - v36
			v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v37)+1)))
			v46 = v42 - int32(48)
			if base.Ui32(v46) < base.Ui32(v39) {
				v35 = v41
				v36 = v46
				v37 = v37 + int32(1)
				continue
			} else {
				break
			}
			break
		}
		v49 = v41
	} else {
		v49 = v30
	}
	if v29 != 0 {
		v55 = int32(0) - v49
	} else {
		v55 = v49
	}
	return v55
}
func F_attnumAttName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 <= int32(0) {
		v12 = F_SystemAttributeDefinition(m, base.I32_extend16_s(l1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v29 = v12
			m.G0 = v7 + int32(16)
			return v29 + int32(4)
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v17 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(470294), v7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494547), int32(3649), int32(380871))
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
			v29 = v16 + v17<<(uint(int32(4))%32) + l1*int32(100) - int32(80)
			m.G0 = v7 + int32(16)
			return v29 + int32(4)
		}
	}
}
func F_attnumCollationId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 <= int32(0) {
		v24 = int32(0)
		m.G0 = v7 + int32(16)
		return v24
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(470294), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494547), int32(3689), int32(463925))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13<<(uint(int32(4))%32)+l1*int32(100))+16))
			v24 = v21
			m.G0 = v7 + int32(16)
			return v24
		}
	}
}
