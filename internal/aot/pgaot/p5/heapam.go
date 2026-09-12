package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_fetch_row_version(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)) = uint16(v12)
	v15 = l3 + int32(48)
	v19 = F_heap_fetch(m, l0, l2, v15, v8+int32(12), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			F_ExecStorePinnedBufferHeapTuple(m, v15, l3, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v26
				m.G0 = v8 + int32(16)
				return v19
			}
		} else {
			m.G0 = v8 + int32(16)
			return v19
		}
	}
}
func F_heapam_relation_nontransactional_truncate(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_RelationTruncate(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_heapam_scan_sample_next_block(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v5 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if v11 != 0 {
			F_ReleaseBuffer(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v20 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, l1, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v46 = v20
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
						if v46 == int32(-1) {
							v50 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
							return v50
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
							if v55 != 0 {
								F_ProcessInterrupts(m)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = int32(0)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
									v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
										v65 = int32(1)
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
										if v66&v65 != 0 {
											F_heap_prepare_pagescan(m, l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
												return v65
											}
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
											return v65
										}
									}
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
									v65 = int32(1)
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v66&v65 != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
											return v65
										}
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
										return v65
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					if v22 == int32(-1) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v46 = v25
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
						if v46 == int32(-1) {
							v50 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
							return v50
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
							if v55 != 0 {
								F_ProcessInterrupts(m)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = int32(0)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
									v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
										v65 = int32(1)
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
										if v66&v65 != 0 {
											F_heap_prepare_pagescan(m, l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
												return v65
											}
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
											return v65
										}
									}
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
									v65 = int32(1)
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v66&v65 != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
											return v65
										}
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
										return v65
									}
								}
							}
						}
					} else {
						v27 = v22 + int32(1)
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if base.Ui32(v27) < base.Ui32(v29) {
							v31 = v27
						} else {
							v31 = int32(0)
						}
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v32&int32(128) != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_ss_report_location(m, v35, v31)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v31 != v38 {
									v46 = v31
									*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
									if v46 == int32(-1) {
										v50 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
										return v50
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
										if v55 != 0 {
											F_ProcessInterrupts(m)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v59 = int32(0)
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
												v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
													v65 = int32(1)
													v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
													if v66&v65 != 0 {
														F_heap_prepare_pagescan(m, l0)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
															return v65
														}
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
														return v65
													}
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v59 = int32(0)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
											v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
												v65 = int32(1)
												v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
												if v66&v65 != 0 {
													F_heap_prepare_pagescan(m, l0)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
														return v65
													}
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
													return v65
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(-1)
									v42 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v42)
									return v42
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v31 != v38 {
								v46 = v31
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
								if v46 == int32(-1) {
									v50 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
									return v50
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
									if v55 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v59 = int32(0)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
											v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
												v65 = int32(1)
												v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
												if v66&v65 != 0 {
													F_heap_prepare_pagescan(m, l0)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
														return v65
													}
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
													return v65
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v59 = int32(0)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
										v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
											v65 = int32(1)
											v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
											if v66&v65 != 0 {
												F_heap_prepare_pagescan(m, l0)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
													return v65
												}
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
												return v65
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(-1)
								v42 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v42)
								return v42
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			if v18 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v20 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, l1, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v46 = v20
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
					if v46 == int32(-1) {
						v50 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
						return v50
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
						if v55 != 0 {
							F_ProcessInterrupts(m)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
									v65 = int32(1)
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v66&v65 != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
											return v65
										}
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
										return v65
									}
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v59 = int32(0)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
								v65 = int32(1)
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
								if v66&v65 != 0 {
									F_heap_prepare_pagescan(m, l0)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
										return v65
									}
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
									return v65
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				if v22 == int32(-1) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v46 = v25
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
					if v46 == int32(-1) {
						v50 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
						return v50
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
						if v55 != 0 {
							F_ProcessInterrupts(m)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
									v65 = int32(1)
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v66&v65 != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
											return v65
										}
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
										return v65
									}
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v59 = int32(0)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
								v65 = int32(1)
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
								if v66&v65 != 0 {
									F_heap_prepare_pagescan(m, l0)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
										return v65
									}
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
									return v65
								}
							}
						}
					}
				} else {
					v27 = v22 + int32(1)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if base.Ui32(v27) < base.Ui32(v29) {
						v31 = v27
					} else {
						v31 = int32(0)
					}
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v32&int32(128) != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_ss_report_location(m, v35, v31)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v31 != v38 {
								v46 = v31
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
								if v46 == int32(-1) {
									v50 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
									return v50
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
									if v55 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v59 = int32(0)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
											v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
												v65 = int32(1)
												v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
												if v66&v65 != 0 {
													F_heap_prepare_pagescan(m, l0)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
														return v65
													}
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
													return v65
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v59 = int32(0)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
										v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
											v65 = int32(1)
											v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
											if v66&v65 != 0 {
												F_heap_prepare_pagescan(m, l0)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
													return v65
												}
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
												return v65
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(-1)
								v42 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v42)
								return v42
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v31 != v38 {
							v46 = v31
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
							if v46 == int32(-1) {
								v50 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v50)
								return v50
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _consts[45]))
								if v55 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v59 = int32(0)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
										v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
											v65 = int32(1)
											v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
											if v66&v65 != 0 {
												F_heap_prepare_pagescan(m, l0)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
													return v65
												}
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
												return v65
											}
										}
									}
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = int32(0)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
									v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v62
										v65 = int32(1)
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
										if v66&v65 != 0 {
											F_heap_prepare_pagescan(m, l0)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
												return v65
											}
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v65)
											return v65
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(-1)
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v42)
							return v42
						}
					}
				}
			}
		}
	}
}
func F_heapam_tuple_complete_speculative(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v13)
	v16 = l1 + int32(28)
	v20 = F_ExecFetchSlotHeapTuple(m, l1, v13, v11+int32(15))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		if l3 != 0 {
			v22 = m.G0
			v23 = int32(16)
			v24 = v22 - v23
			m.G0 = v24
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+2)))
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16))))
			v31 = F_ReadBuffer(m, l0, v26|v27<<(uint(v23)%32))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_LockBuffer(m, v31, int32(2))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					if v31 < int32(0) {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[12]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v31^int32(-1))<<(uint(int32(2))%32))))
						v53 = v45
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _consts[13]))
						v53 = v47 + v31<<(uint(int32(13))%32) + int32(-8192)
					}
					v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
					if base.Ui32(int32(25)) <= base.Ui32(v54) {
						v62 = int32(base.Ui32(v54+int32(262120)) >> (uint(int32(2)) % 32))
					} else {
						v62 = int32(0)
					}
					v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
					if base.Ui32(v62&int32(65535)) < base.Ui32(v65) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(242536), int32(0))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								F_errfinish(m, int32(509219), int32(6116), int32(350985))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v65<<(uint(int32(2))%32)+v53)+20))
						if v70&int32(98304) != int32(32768) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(242536), int32(0))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									F_errfinish(m, int32(509219), int32(6116), int32(350985))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v75 = int32(4530932)
							v77 = *(*int32)(unsafe.Add(mBase, _consts[14]))
							*(*int32)(unsafe.Add(mBase, _consts[14])) = v77 + int32(1)
							F_MarkBufferDirty(m, v31)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v85 = v53 + v70&int32(32767)
								v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v85)+16)) = uint16(v86)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v88
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+118)))
								if v91 != int32(112) {
									v126 = int32(4530932)
									v128 = *(*int32)(unsafe.Add(mBase, _consts[14]))
									*(*int32)(unsafe.Add(mBase, _consts[14])) = v128 - int32(1)
									F_UnlockReleaseBuffer(m, v31)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										m.G0 = v24 + int32(16)
										v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
										if v157 == int32(1) {
											F_pfree(m, v20)
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return
											} else {
												m.G0 = v11 + int32(16)
												return
											}
										} else {
											m.G0 = v11 + int32(16)
											return
										}
									}
								} else {
									v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
									if v95 <= int32(0) {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v98 != 0 {
											v126 = int32(4530932)
											v128 = *(*int32)(unsafe.Add(mBase, _consts[14]))
											*(*int32)(unsafe.Add(mBase, _consts[14])) = v128 - int32(1)
											F_UnlockReleaseBuffer(m, v31)
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return
											} else {
												m.G0 = v24 + int32(16)
												v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
												if v157 == int32(1) {
													F_pfree(m, v20)
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v99 != 0 {
												v126 = int32(4530932)
												v128 = *(*int32)(unsafe.Add(mBase, _consts[14]))
												*(*int32)(unsafe.Add(mBase, _consts[14])) = v128 - int32(1)
												F_UnlockReleaseBuffer(m, v31)
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return
												} else {
													m.G0 = v24 + int32(16)
													v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
													if v157 == int32(1) {
														F_pfree(m, v20)
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16)
															return
														}
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												}
											} else {
												v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
												*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)) = uint16(v100)
												F_XLogBeginInsert(m)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return
												} else {
													v105 = int32(4432116)
													v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
													v108 = v107 | int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[44])) = uint8(v108)
													F_XLogRegisterData(m, v24+int32(14), int32(2))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return
													} else {
														F_XLogRegisterBuffer(m, int32(0), v31, int32(8))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return
														} else {
															v121 = F_XLogInsert(m, int32(10), int32(80))
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v53))) = base.I64_rotr(v121, int64(32))
																v126 = int32(4530932)
																v128 = *(*int32)(unsafe.Add(mBase, _consts[14]))
																*(*int32)(unsafe.Add(mBase, _consts[14])) = v128 - int32(1)
																F_UnlockReleaseBuffer(m, v31)
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return
																} else {
																	m.G0 = v24 + int32(16)
																	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
																	if v157 == int32(1) {
																		F_pfree(m, v20)
																		mBase = m.M
																		v161 = m.ExcPending
																		if v161 != 0 {
																			return
																		} else {
																			m.G0 = v11 + int32(16)
																			return
																		}
																	} else {
																		m.G0 = v11 + int32(16)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v24)+14)) = uint16(v100)
										F_XLogBeginInsert(m)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											v105 = int32(4432116)
											v107 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
											v108 = v107 | int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[44])) = uint8(v108)
											F_XLogRegisterData(m, v24+int32(14), int32(2))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return
											} else {
												F_XLogRegisterBuffer(m, int32(0), v31, int32(8))
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return
												} else {
													v121 = F_XLogInsert(m, int32(10), int32(80))
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v53))) = base.I64_rotr(v121, int64(32))
														v126 = int32(4530932)
														v128 = *(*int32)(unsafe.Add(mBase, _consts[14]))
														*(*int32)(unsafe.Add(mBase, _consts[14])) = v128 - int32(1)
														F_UnlockReleaseBuffer(m, v31)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return
														} else {
															m.G0 = v24 + int32(16)
															v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
															if v157 == int32(1) {
																F_pfree(m, v20)
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return
																} else {
																	m.G0 = v11 + int32(16)
																	return
																}
															} else {
																m.G0 = v11 + int32(16)
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
					}
				}
			}
		} else {
			F_heap_abort_speculative(m, l0, v16)
			mBase = m.M
			v152 = m.ExcPending
			if v152 != 0 {
				return
			} else {
				v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
				if v157 == int32(1) {
					F_pfree(m, v20)
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			}
		}
	}
}
func F_heapam_tuple_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)) = uint8(v17)
	v22 = F_ExecFetchSlotHeapTuple(m, l2, v17, v15+int32(15))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v26
		v29 = F_heap_update(m, l0, l1, v22, l3, l5, l6, l7, l8, l9)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l2)+32)) = uint16(v31)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v33
			if v29 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l9))) = int32(0)
			} else {
			}
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
			if v37 == int32(1) {
				F_pfree(m, v22)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					m.G0 = v15 + int32(16)
					return v29
				}
			} else {
				m.G0 = v15 + int32(16)
				return v29
			}
		}
	}
}
