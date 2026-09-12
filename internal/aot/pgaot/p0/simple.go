package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSimpleRelationDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_CheckCmdReplicaIdentity(m, v7, int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = l3 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		if v13 == int32(0) {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v29 = m.G0
			v31 = v29 - int32(32)
			m.G0 = v31
			v34 = F_GetCurrentCommandId(m, int32(1))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = int32(0)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+188))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
				v43 = m.T0[v42].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v7, v12, v34, v28, v36, int32(1), v31+int32(12), v36)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					if v43 != 0 {
						switch v43 - int32(2) {
						case 0:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(331277), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									F_errfinish(m, int32(486092), int32(306), int32(343401))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 1:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(438824), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									F_errfinish(m, int32(486092), int32(314), int32(343401))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 2:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(437061), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									F_errfinish(m, int32(486092), int32(318), int32(343401))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v31))) = v43
								F_errmsg_internal(m, int32(56680), v31)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									F_errfinish(m, int32(486092), int32(322), int32(343401))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v31 + int32(32)
						v102 = int32(0)
						F_ExecARDeleteTriggers(m, l1, l0, v12, v102, v102, v102)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+18)))
			if v16 != int32(1) {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v29 = m.G0
				v31 = v29 - int32(32)
				m.G0 = v31
				v34 = F_GetCurrentCommandId(m, int32(1))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(0)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+188))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
					v43 = m.T0[v42].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v7, v12, v34, v28, v36, int32(1), v31+int32(12), v36)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						if v43 != 0 {
							switch v43 - int32(2) {
							case 0:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(331277), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										F_errfinish(m, int32(486092), int32(306), int32(343401))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 1:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(438824), int32(0))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errfinish(m, int32(486092), int32(314), int32(343401))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 2:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(437061), int32(0))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_errfinish(m, int32(486092), int32(318), int32(343401))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v31))) = v43
									F_errmsg_internal(m, int32(56680), v31)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_errfinish(m, int32(486092), int32(322), int32(343401))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							m.G0 = v31 + int32(32)
							v102 = int32(0)
							F_ExecARDeleteTriggers(m, l1, l0, v12, v102, v102, v102)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v19 = int32(0)
				v24 = F_ExecBRDeleteTriggers(m, l1, l2, l0, v12, v19, v19, v19, v19, v19)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					if v24 == int32(0) {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v29 = m.G0
						v31 = v29 - int32(32)
						m.G0 = v31
						v34 = F_GetCurrentCommandId(m, int32(1))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v36 = int32(0)
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+188))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
							v43 = m.T0[v42].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v7, v12, v34, v28, v36, int32(1), v31+int32(12), v36)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								if v43 != 0 {
									switch v43 - int32(2) {
									case 0:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(331277), int32(0))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												F_errfinish(m, int32(486092), int32(306), int32(343401))
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									case 1:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(438824), int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												F_errfinish(m, int32(486092), int32(314), int32(343401))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									case 2:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											F_errmsg_internal(m, int32(437061), int32(0))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												F_errfinish(m, int32(486092), int32(318), int32(343401))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v31))) = v43
											F_errmsg_internal(m, int32(56680), v31)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												F_errfinish(m, int32(486092), int32(322), int32(343401))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									m.G0 = v31 + int32(32)
									v102 = int32(0)
									F_ExecARDeleteTriggers(m, l1, l0, v12, v102, v102, v102)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										return
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
func F_SimpleLruWaitIO(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v8 = l1 << (uint(int32(3)) % 32) & int32(-128)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_LWLockRelease(m, v8+v10)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = l1 << (uint(int32(7)) % 32)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
		v19 = F_LWLockAcquire(m, v15+v16, int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
			F_LWLockRelease(m, v21+v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
				v28 = F_LWLockAcquire(m, v25+v8, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+l1<<(uint(int32(2))%32))))
					switch v34 - int32(1) {
					case 0, 2:
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
						v40 = F_LWLockConditionalAcquire(m, v37+v15, int32(1))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if v40 == int32(0) {
								return
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v47 = v44 + l1<<(uint(int32(2))%32)
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
								if v48 == int32(1) {
									*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(2)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									v57 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v55+l1))) = uint8(v57)
								}
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
								F_LWLockRelease(m, v59+l1<<(uint(int32(7))%32))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									return
								}
							}
						}
					default:
						return
					}
				}
			}
		}
	}
}
func F_SimpleLruZeroPage(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = F_SlruSelectLRUPage(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v14+v10<<(uint(int32(3))%32)))) = l1
		v19 = int32(2)
		v20 = v10 << (uint(v19) % 32)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = v19
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		v27 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10+v25))) = uint8(v27)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
		v34 = v29 + v10>>(uint(int32(4))%32)<<(uint(v19)%32)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v36+v20)))
		if v35 != v38 {
			v41 = v35 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v43+v20))) = v41
		} else {
		}
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v47+v20)))
		v53 = F__emscripten_memset_bulkmem(m, v49, base.I32_extend8_s(int32(0)), int32(8192))
		mBase = m.M
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
		if v55 <= int32(0) {
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+36))
			v59 = v55 * v10
			v60 = int32(3)
			v62 = v58 + v59<<(uint(v60)%32)
			v64 = v55 << (uint(v60) % 32)
			if base.Ui32(int32(1024)) < base.Ui32(v64) {
				v93 = v64
				v96 = F__emscripten_memset_bulkmem(m, v62, base.I32_extend8_s(int32(0)), v93)
				mBase = m.M
			} else {
				if v62&int32(3) != 0 {
					v93 = v64
					v96 = F__emscripten_memset_bulkmem(m, v62, base.I32_extend8_s(int32(0)), v93)
					mBase = m.M
				} else {
					if base.Ui32(v62+v64) <= base.Ui32(v62) {
					} else {
						v71 = int32(3)
						v72 = v59 << (uint(v71) % 32)
						v78 = v72 + v58 + int32(4)
						v84 = v55*(v10<<(uint(v71)%32)+int32(8)) + v58
						if base.Ui32(v84) < base.Ui32(v78) {
							v86 = v78
						} else {
							v86 = v84
						}
						v93 = (v72^int32(-1)-v58+v86)&int32(-4) + int32(4)
						v96 = F__emscripten_memset_bulkmem(m, v62, base.I32_extend8_s(int32(0)), v93)
						mBase = m.M
					}
				}
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = l1
		v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
		v105 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[137])) = uint8(v105)
		*(*uint8)(unsafe.Add(mBase, _consts[138])) = uint8(v105)
		v111 = v103 << (uint(int32(6)) % 32)
		v114 = *(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[139])))
		*(*int64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[139]))) = v114 + int64(1)
		return v10
	}
}
